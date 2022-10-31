package routes

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/cuno-1000/panic-product/api-gateway/handler"
	"github.com/cuno-1000/panic-product/api-gateway/token/middleware"
	"github.com/cuno-1000/panic-product/api-gateway/token/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"net/http"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default()

	//router.Use(TracerWrapper).Use(Cors())
	router.Use(Cors())

	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
	authRouter := router.Group("/user").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.GET("/GetUserInfo", handler.GetUserInfo)
		authRouter.GET("/logout", handler.Logout)

		authRouter.GET("/event/qualify/:event_uuid", handler.CheckQualify)
		authRouter.GET("/event/apply/:apply_link/:event_uuid", handler.ApplyPurchase)
	}
	router.GET("/event/list", handler.GetEventsList)
	router.GET("/testing", handler.Test)
	//router.GET("/fetch/adult/career", handler.FetchNormalAdult)

	return router
}

func TracerWrapper(c *gin.Context) {
	md := make(map[string]string)
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
	defer sp.Finish()

	if err := opentracing.GlobalTracer().Inject(sp.Context(),
		opentracing.TextMap,
		opentracing.TextMapCarrier(md)); err != nil {
		logger.Fatal(err)
	}

	ctx := context.TODO()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	//c.Set(contextTracerKey, ctx)

	c.Next()

	statusCode := c.Writer.Status()
	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
	ext.HTTPMethod.Set(sp, c.Request.Method)
	ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
	if statusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
	}
}

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)
}
