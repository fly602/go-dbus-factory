// Code generated by "./generator ./session/com.deepin.daemon.EventLog"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package EventLog

import (
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type EventLog interface {
	eventLog // interface com.deepin.daemon.EventLog
	proxy.Object
}

type objectEventLog struct {
	interfaceEventLog // interface com.deepin.daemon.EventLog
	proxy.ImplObject
}

func NewEventLog(conn *dbus.Conn) EventLog {
	obj := new(objectEventLog)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.EventLog", "/com/deepin/daemon/EventLog")
	return obj
}

type eventLog interface {
	GoEnable(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call
	Enable(flags dbus.Flags, enable bool) error
	GoReportLog(flags dbus.Flags, ch chan *dbus.Call, log string) *dbus.Call
	ReportLog(flags dbus.Flags, log string) error
	Enabled() proxy.PropBool
}

type interfaceEventLog struct{}

func (v *interfaceEventLog) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceEventLog) GetInterfaceName_() string {
	return "com.deepin.daemon.EventLog"
}

// method Enable

func (v *interfaceEventLog) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enable)
}

func (v *interfaceEventLog) Enable(flags dbus.Flags, enable bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

// method ReportLog

func (v *interfaceEventLog) GoReportLog(flags dbus.Flags, ch chan *dbus.Call, log string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReportLog", flags, ch, log)
}

func (v *interfaceEventLog) ReportLog(flags dbus.Flags, log string) error {
	return (<-v.GoReportLog(flags, make(chan *dbus.Call, 1), log).Done).Err
}

// property Enabled b

func (v *interfaceEventLog) Enabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Enabled",
	}
}
