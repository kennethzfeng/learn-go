package main

import (
	"fmt"
	"time"

	snmp "github.com/soniah/gosnmp"
	"github.com/yext/glog"
)

type Config struct {
	Coummunity string
	Host       string
	OIDs       []string
}

var config Config

func main() {
	config.Coummunity = "public"
	config.Host = "192.168.1.1"
	config.OIDs = []string{
		".1.3.6.1.2.1.2.2.1.16.8",
	}

	forever := time.Tick(10 * time.Second)
	getStat(config)
	for _ = range forever {
		getStat(config)
	}
}

func getStat(c Config) {
	snmp.Default.Target = c.Host
	err := snmp.Default.Connect()
	if err != nil {
		glog.FatalIf(err, "Connection error: ")
	}
	defer snmp.Default.Conn.Close()

	result, err := snmp.Default.Get(c.OIDs)
	if err != nil {
		glog.FatalIf(err, "SNMP Get error: ")
	}

	for i, variable := range result.Variables {
		var line string
		line = fmt.Sprintf("%d: oid: %s ", i, variable.Name)
		switch variable.Type {
		case snmp.OctetString:
			line += fmt.Sprintf("string: %s\n", string(variable.Value.([]byte)))
		default:
			line += fmt.Sprintf("number: %d\n", snmp.ToBigInt(variable.Value))
		}

		glog.Infof(line)
	}
}
