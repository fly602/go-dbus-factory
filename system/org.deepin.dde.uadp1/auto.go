// Code generated by "./generator ./system/org.deepin.dde.uadp1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package uadp1

import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Uadp interface {
	uadp // interface org.deepin.dde.Uadp1
	proxy.Object
}

type objectUadp struct {
	interfaceUadp // interface org.deepin.dde.Uadp1
	proxy.ImplObject
}

func NewUadp(conn *dbus.Conn) Uadp {
	obj := new(objectUadp)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Uadp1", "/org/deepin/dde/Uadp1")
	return obj
}

type uadp interface {
	GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, dataKey string, keyringKey string) *dbus.Call
	SetDataKey(flags dbus.Flags, exePath string, keyName string, dataKey string, keyringKey string) error
	GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, keyringKey string) *dbus.Call
	GetDataKey(flags dbus.Flags, exePath string, keyName string, keyringKey string) (string, error)
	GoAvailable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Available(flags dbus.Flags) (bool, error)
	GoDelete(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	Delete(flags dbus.Flags, name string) error
	GoGet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	Get(flags dbus.Flags, name string) ([]uint8, error)
	GoListName(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ListName(flags dbus.Flags) ([]string, error)
	GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Release(flags dbus.Flags) error
	GoSet(flags dbus.Flags, ch chan *dbus.Call, name string, data []uint8) *dbus.Call
	Set(flags dbus.Flags, name string, data []uint8) error
}

type interfaceUadp struct{}

func (v *interfaceUadp) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceUadp) GetInterfaceName_() string {
	return "org.deepin.dde.Uadp1"
}

// method SetDataKey

func (v *interfaceUadp) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, dataKey string, keyringKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDataKey", flags, ch, exePath, keyName, dataKey, keyringKey)
}

func (v *interfaceUadp) SetDataKey(flags dbus.Flags, exePath string, keyName string, dataKey string, keyringKey string) error {
	return (<-v.GoSetDataKey(flags, make(chan *dbus.Call, 1), exePath, keyName, dataKey, keyringKey).Done).Err
}

// method GetDataKey

func (v *interfaceUadp) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, keyringKey string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDataKey", flags, ch, exePath, keyName, keyringKey)
}

func (*interfaceUadp) StoreGetDataKey(call *dbus.Call) (dataKey string, err error) {
	err = call.Store(&dataKey)
	return
}

func (v *interfaceUadp) GetDataKey(flags dbus.Flags, exePath string, keyName string, keyringKey string) (string, error) {
	return v.StoreGetDataKey(
		<-v.GoGetDataKey(flags, make(chan *dbus.Call, 1), exePath, keyName, keyringKey).Done)
}

// method Available

func (v *interfaceUadp) GoAvailable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Available", flags, ch)
}

func (*interfaceUadp) StoreAvailable(call *dbus.Call) (outArg0 bool, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceUadp) Available(flags dbus.Flags) (bool, error) {
	return v.StoreAvailable(
		<-v.GoAvailable(flags, make(chan *dbus.Call, 1)).Done)
}

// method Delete

func (v *interfaceUadp) GoDelete(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, name)
}

func (v *interfaceUadp) Delete(flags dbus.Flags, name string) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method Get

func (v *interfaceUadp) GoGet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Get", flags, ch, name)
}

func (*interfaceUadp) StoreGet(call *dbus.Call) (outArg0 []uint8, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceUadp) Get(flags dbus.Flags, name string) ([]uint8, error) {
	return v.StoreGet(
		<-v.GoGet(flags, make(chan *dbus.Call, 1), name).Done)
}

// method ListName

func (v *interfaceUadp) GoListName(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListName", flags, ch)
}

func (*interfaceUadp) StoreListName(call *dbus.Call) (outArg0 []string, err error) {
	err = call.Store(&outArg0)
	return
}

func (v *interfaceUadp) ListName(flags dbus.Flags) ([]string, error) {
	return v.StoreListName(
		<-v.GoListName(flags, make(chan *dbus.Call, 1)).Done)
}

// method Release

func (v *interfaceUadp) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Release", flags, ch)
}

func (v *interfaceUadp) Release(flags dbus.Flags) error {
	return (<-v.GoRelease(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Set

func (v *interfaceUadp) GoSet(flags dbus.Flags, ch chan *dbus.Call, name string, data []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Set", flags, ch, name, data)
}

func (v *interfaceUadp) Set(flags dbus.Flags, name string, data []uint8) error {
	return (<-v.GoSet(flags, make(chan *dbus.Call, 1), name, data).Done).Err
}