// Code generated by "./generator system/com.deepin.system.inputdevices"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package inputdevices

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockInputDevices struct {
	MockInterfaceInputDevices // interface com.deepin.system.InputDevices
	proxy.MockObject
}

type MockInterfaceInputDevices struct {
	mock.Mock
}

// signal TouchscreenAdded

func (v *MockInterfaceInputDevices) ConnectTouchscreenAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchscreenRemoved

func (v *MockInterfaceInputDevices) ConnectTouchscreenRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Touchscreens ao

func (v *MockInterfaceInputDevices) Touchscreens() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockTouchscreen struct {
	MockInterfaceTouchscreen // interface com.deepin.system.InputDevices.Touchscreen
	proxy.MockObject
}

type MockInterfaceTouchscreen struct {
	mock.Mock
}

// property DevNode s

func (v *MockInterfaceTouchscreen) DevNode() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property BusType s

func (v *MockInterfaceTouchscreen) BusType() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property UUID s

func (v *MockInterfaceTouchscreen) UUID() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Phys s

func (v *MockInterfaceTouchscreen) Phys() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property OutputName s

func (v *MockInterfaceTouchscreen) OutputName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Width d

func (v *MockInterfaceTouchscreen) Width() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Height d

func (v *MockInterfaceTouchscreen) Height() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *MockInterfaceTouchscreen) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Serial s

func (v *MockInterfaceTouchscreen) Serial() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockTouchpad struct {
	MockInterfaceTouchpad // interface com.deepin.system.InputDevices.Touchpad
	proxy.MockObject
}

type MockInterfaceTouchpad struct {
	mock.Mock
}

// method SetTouchpadEnable

func (v *MockInterfaceTouchpad) GoSetTouchpadEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTouchpad) SetTouchpadEnable(flags dbus.Flags, enabled bool) error {
	mockArgs := v.Called(flags, enabled)

	return mockArgs.Error(0)
}

// property isExist b

func (v *MockInterfaceTouchpad) IsExist() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Enable b

func (v *MockInterfaceTouchpad) Enable() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
