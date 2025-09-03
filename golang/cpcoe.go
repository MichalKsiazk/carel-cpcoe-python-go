package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/goburrow/modbus"
)

type CPCOE_CHDIR uint16

const (
	INPUT  = 0
	OUTPUT = 1
)

type CPCOE_CHTYPE_OUT uint16

const (
	OUT_0_10    = 0
	OUT_PWM_100 = 3
	OUT_PWM_2   = 5
)

type CPCOE_CHTYPE_IN uint16

const (
	NTC            = 0
	PT1000         = 1
	ANALOG_0_1     = 2
	ANALOG_0_10    = 3
	ANALOG_4_20    = 4
	DIN            = 5
	ANALOG_0_5     = 6
	NTC_0_150      = 7
	NTC_TEMP_PRESS = 11
	PTC_R          = 34
	PT500          = 35
	PT100          = 36
	ANALOG_0_20    = 42
	NOT_USED       = 255
)

type ModbusDataChunk struct {
	ConfigPath string
	Start      uint
	Stop       uint
	Count      uint
}

type CPCOE_Device struct {
	Address  uint16
	Baudrate int
	Bytesize int
	Parity   string
	StopBits int
	Port     string
	Handler  modbus.RTUClientHandler
	Client   modbus.Client
}

type CPCOE_DeviceOption func(*CPCOE_Device)

func NewCPCOE_Device(address uint16, port string) *CPCOE_Device {
	dev := &CPCOE_Device{
		Address:  address,
		Port:     port,
		Baudrate: 19200,
		Bytesize: 8,
		Parity:   "N",
		StopBits: 1,
	}

	dev.Handler = *modbus.NewRTUClientHandler(dev.Port)
	dev.Handler.BaudRate = dev.Baudrate
	dev.Handler.DataBits = dev.Bytesize
	dev.Handler.Parity = dev.Parity
	dev.Handler.StopBits = dev.StopBits
	dev.Handler.SlaveId = byte(address)
	dev.Handler.Timeout = 5 * time.Second

	err := dev.Handler.Connect()
	if err != nil {
		fmt.Println("Error during connecting with CPCOE module")
	}
	defer dev.Handler.Close()
	dev.Client = modbus.NewClient(&dev.Handler)
	for {
		results, err := dev.Client.ReadHoldingRegisters(67, 2)
		if err != nil {
			fmt.Println("BIG CHUJ")
		}
		bin := binary.BigEndian.Uint32(results)
		temp := math.Float32frombits(bin)
		fmt.Printf("%f\n", temp)
	}

}
