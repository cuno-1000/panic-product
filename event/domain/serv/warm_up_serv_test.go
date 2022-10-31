package serv

import (
	"github.com/cuno-1000/panic-product/event/domain/model"
	"github.com/cuno-1000/panic-product/event/domain/repo"
	"testing"
)

func TestEventDataService_WarmUpBlacklist(t *testing.T) {
	eventDataService := NewEventDataService(repo.NewEventRepository(RedisDb))
	eventModel := eventDataService.EventByUuid("9af8da7e-f8d3-4b69-7a91-c6315dc0c80a")
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		event *model.Event
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantIsDone bool
		wantErr    bool
	}{
		{
			"good",
			fields{
				EventRepository: repo.NewEventRepository(RedisDb),
			},
			args{
				eventModel,
			},
			true,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			gotIsDone, err := e.WarmUpBlacklist(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("WarmUpBlacklist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsDone != tt.wantIsDone {
				t.Errorf("WarmUpBlacklist() gotIsDone = %v, want %v", gotIsDone, tt.wantIsDone)
			}
		})
	}
}

func TestEventDataService_WarmUpEventRoutines(t *testing.T) {
	type fields struct {
		EventRepository repo.IEventRepository
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			e.WarmUpEventRoutines()
		})
	}
}

func TestEventDataService_WarmUpNormalAdult(t *testing.T) {
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		event *model.Event
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantIsDone bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			gotIsDone, err := e.WarmUpNormalAdult(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("WarmUpNormalAdult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsDone != tt.wantIsDone {
				t.Errorf("WarmUpNormalAdult() gotIsDone = %v, want %v", gotIsDone, tt.wantIsDone)
			}
		})
	}
}

func TestEventDataService_WarmUpRepaymentOverDue(t *testing.T) {
	type fields struct {
		EventRepository repo.IEventRepository
	}
	type args struct {
		event *model.Event
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantIsDone bool
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventDataService{
				EventRepository: tt.fields.EventRepository,
			}
			gotIsDone, err := e.WarmUpRepaymentOverDue(tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("WarmUpRepaymentOverDue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsDone != tt.wantIsDone {
				t.Errorf("WarmUpRepaymentOverDue() gotIsDone = %v, want %v", gotIsDone, tt.wantIsDone)
			}
		})
	}
}
