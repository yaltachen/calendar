package main

import (
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"

	"github.com/yaltachen/calendar"
	"github.com/yaltachen/calendar/service"
)

func TestServer(t *testing.T) {
	var (
		client *rpc.Client
		lDate  calendar.LunarDate
		sDate  calendar.SolarDate
		err    error
	)
	if client, err = jsonrpc.Dial("tcp", "127.0.0.1:8000"); err != nil {
		t.Errorf("dial failed, err: %v\r\n", err)
	}
	if err = client.Call("DateTransServer.Solar2Lunar",
		service.Date{Year: 2020, Month: 2, Date: 15}, &lDate); err != nil {
		t.Errorf("call solar2lunar failed, err: %v\r\n", err)
	} else {
		log.Printf("%v", lDate)
	}

	if client, err = jsonrpc.Dial("tcp", "127.0.0.1:8000"); err != nil {
		t.Errorf("dial failed, err: %v\r\n", err)
	}
	if err = client.Call("DateTransServer.Lunar2Solar",
		service.Date{Year: 2020, Month: 1, Date: 22, Leap: false}, &sDate); err != nil {
		t.Errorf("call solar2lunar failed, err: %v\r\n", err)
	} else {
		log.Printf("%v", sDate)
	}

}
