package tbus

import (
	prot "github.com/evo-bots/tbus/go/tbus/protocol"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// DeviceBase implements basic device operations
type DeviceBase struct {
	Info DeviceInfo

	busPort BusPort
}

// DeviceInfo returns device information
func (d *DeviceBase) DeviceInfo() DeviceInfo {
	return d.Info
}

// AttachTo implements Device
func (d *DeviceBase) AttachTo(busPort BusPort, addr uint8) {
	d.busPort = busPort
	d.Info.Address = uint32(addr)
}

// BusPort implements Device
func (d *DeviceBase) BusPort() BusPort {
	return d.busPort
}

// Reply write reply to bus
func (d *DeviceBase) Reply(msgID uint32, reply proto.Message) error {
	if reply == nil {
		reply = &empty.Empty{}
	}

	encoded, err := proto.Marshal(reply)
	if err != nil {
		return err
	}
	msg, err := prot.EncodeAsMsg(nil, msgID, 0, encoded)
	if err != nil {
		return err
	}
	return d.busPort.SendMsg(msg)
}

// LogicBase implements DeviceLogic
type LogicBase struct {
	Device Device
}

// SetDevice implements DeviceLogic
func (l *LogicBase) SetDevice(dev Device) {
	l.Device = dev
}

// DeviceAddress is a helper to construct device address
func DeviceAddress(devs ...Device) []uint8 {
	addrs := make([]uint8, len(devs))
	for n, dev := range devs {
		addrs[n] = uint8(dev.DeviceInfo().Address)
	}
	return addrs
}

// AddLabel adds a single label to device info
func (d *DeviceInfo) AddLabel(name, value string) *DeviceInfo {
	m := d.Labels
	if m == nil {
		m = make(map[string]string)
		d.Labels = m
	}
	m[name] = value
	return d
}

// AddLabels adds multiple labels
func (d *DeviceInfo) AddLabels(labels map[string]string) *DeviceInfo {
	if d.Labels == nil {
		d.Labels = labels
	} else {
		for k, v := range labels {
			d.Labels[k] = v
		}
	}
	return d
}
