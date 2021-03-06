package tbus

import (
	proto "github.com/golang/protobuf/proto"
)

// Controller is the base class of device controller
type Controller struct {
	LogicBase
	Master  Master
	Address RouteAddr
}

// Invoke invokes a method on a device
func (c *Controller) Invoke(methodIndex uint8, params proto.Message) Invocation {
	return c.Master.Invoke(methodIndex, params, c.Address)
}

// Subscribe subscribes an event channel
func (c *Controller) Subscribe(channel uint8, handler EventHandler) EventSubscription {
	return c.Master.Subscribe(channel, c.Address, handler)
}

// DeviceInfo retrieves device information
func (c *Controller) DeviceInfo() (info DeviceInfo, err error) {
	err = c.Invoke(0, nil).Result(&info)
	return
}

// MethodInvocation provides partial Invocation implementations for
// generated controller code
type MethodInvocation struct {
	Invocation Invocation
}

// Recv implements Invocation
func (i *MethodInvocation) Recv() (MsgReceiver, error) {
	return i.Invocation.Recv()
}

// MsgID implements Invocation
func (i *MethodInvocation) MsgID() MsgID {
	return i.Invocation.MsgID()
}

// Result implements Invocation
func (i *MethodInvocation) Result(reply proto.Message) error {
	return i.Invocation.Result(reply)
}

// Ignore implements Invocation
func (i *MethodInvocation) Ignore() {
	i.Invocation.Ignore()
}
