package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchya1/graylog-go-api-client/graylog/api"
)

func main() {
	c := api.NewClient(os.Getenv("GRAYLOG_HOSTNAME"), os.Getenv("GRAYLOG_API_TOKEN"), os.Getenv("GRAYLOG_USERNAME"), true)

	i := &api.GraylogSyslogUDPGlobalInput{
		Title:  "Test syslog",
		Type:   "org.graylog2.inputs.syslog.udp.SyslogUDPInput",
		Global: true,
		Configuration: api.GraylogSyslogUDPGlobalInputConfiguration{
			BindAddress:         "0.0.0.0",
			NumberWorkerThreads: 1,
			Port:                43454,
			ForceRDNS:           false,
			AllowDateOverride:   true,
			StoreFullMessage:    true,
			RecvBufferSize:      262144,
		},
	}

	id, e := c.CreateGraylogSyslogUDPGlobalInput(i)

	if e != nil {
		panic(e)
	}

	time.Sleep(3 * time.Second)

	i = &api.GraylogSyslogUDPGlobalInput{
		Title:  "Test syslog updated",
		Type:   "org.graylog2.inputs.syslog.udp.SyslogUDPInput",
		Global: true,
		Configuration: api.GraylogSyslogUDPGlobalInputConfiguration{
			BindAddress:         "0.0.0.0",
			NumberWorkerThreads: 1,
			Port:                43454,
			ForceRDNS:           false,
			AllowDateOverride:   true,
			StoreFullMessage:    true,
			RecvBufferSize:      262144,
		},
	}

	e = c.UpdateGraylogSyslogUDPGlobalInput(i, id)

	if e != nil {
		panic(e)
	}

	r, e := c.GetGraylogSyslogUDPGlobalInput(id)

	if e != nil {
		panic(e)
	}

	fmt.Println(r)

	time.Sleep(3 * time.Second)

	e = c.DeleteGraylogSyslogUDPGlobalInput(id)

	if e != nil {
		panic(e)
	}
}
