package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	pb "github.com/cuno-1000/panic-product/api-gateway/proto"
	"github.com/cuno-1000/panic-product/api-gateway/token/middleware"
	"github.com/gin-gonic/gin"
	cmap "github.com/orcaman/concurrent-map"
	"google.golang.org/grpc"
	"log"
	"time"
)

var EventServiceConn *grpc.ClientConn

func GetEventsList(c *gin.Context) {
	client := pb.NewEventEngineClient(EventServiceConn)
	if rsp, err := client.FetchEvents(context.Background(), &pb.FetchEventsRequest{}); err != nil {
		Fail(c, err)
		return
	} else {
		Success(c, rsp.Events)
		return
	}
}

func CheckQualify(c *gin.Context) {
	id, err := middleware.IdFromJWT(c)
	if err != nil {
		Fail(c, "Cannot read token")
		return
	}
	uuid := c.Param("event_uuid")

	client := pb.NewEventEngineClient(EventServiceConn)

	var rsp *pb.CheckBlacklistResponse
	if err = hystrix.Do("CheckQualify", func() error {

		if rsp, err = client.CheckBlacklist(context.Background(), &pb.CheckBlacklistRequest{
			Uuid:   uuid,
			UserId: id,
		}); err != nil {
			return err
		}
		return nil
	}, func(err error) error {

		if rsp, err = client.CheckBlacklistDowngrade(context.Background(), &pb.CheckBlacklistRequest{
			Uuid:   uuid,
			UserId: id,
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		BznFail(c, -1)
	}

	if rsp.Link == "" {
		BznFail(c, int(rsp.Situation))
	} else {
		Success(c, rsp.Link)
	}
}

//type MapWithLock struct {
//	mp    map[string]chan int
//	mutex *sync.RWMutex
//}
//
//var EventApplyRequests = &MapWithLock{mp: make(map[string]chan int), mutex: new(sync.RWMutex)}
var EventApplyRequests = cmap.New()

//
//var EventApplyRequests sync.Map

func ApplyPurchase(c *gin.Context) {
	id, err := middleware.IdFromJWT(c)
	if err != nil {
		Fail(c, "Cannot read token")
		return
	}
	uuid := c.Param("event_uuid")
	link := c.Param("apply_link")
	//if len(link) != 36 {
	//	BznFail(c, 64)
	//	return
	//}

	requestId := CreateUUID()
	EventApplyRequests.Set(requestId, make(chan int))

	var receive int

	req := &ApplyRequest{
		UserId:    id,
		EventUuid: uuid,
		Link:      link,
		RequestId: requestId,
	}
	ApplyPool.JobQueue <- req

	//PushToMq(req)
	//EventApplyRequests.mutex.Lock()
	var responseChannel chan int
	tmp, _ := EventApplyRequests.Get(requestId)
	responseChannel = tmp.(chan int)
	receive = <-responseChannel
	//receive = <-EventApplyRequests[requestId]
	// concurrent map read and map write
	close(responseChannel)
	//_, ok := EventApplyRequests[requestId]
	//if ok {
	EventApplyRequests.Remove(requestId)
	//delete(EventApplyRequests.mp, requestId)
	//defer EventApplyRequests.mutex.Unlock()
	//}

	if receive != 0 {
		BznFail(c, receive)
		return
	} else {
		Success(c, receive)
		return
	}
}

func MqClientProcess(id uint64, uuid, link, requestId string) {
	tmp, _ := EventApplyRequests.Get(requestId)
	client := pb.NewEventEngineClient(EventServiceConn)
	rsp, err := client.ApplyPurchase(context.Background(), &pb.ApplyRequest{
		Link:   link,
		Uuid:   uuid,
		UserId: id,
	})
	if err != nil {
		tmp.(chan int) <- -1
		return
	}
	tmp.(chan int) <- int(rsp.Situation)
}

func init() {
	hystrix.ConfigureCommand("CheckQualify", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second),
		MaxConcurrentRequests:  1,
		SleepWindow:            5000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  30,
	})
}

// CreateUUID create a random UUID with from RFC 4122
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
