// Code generated by "./generator system/org.deepin.dde.accounts1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package accounts1

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Accounts interface {
	accounts // interface org.deepin.dde.Accounts1
	proxy.Object
}

type objectAccounts struct {
	interfaceAccounts // interface org.deepin.dde.Accounts1
	proxy.ImplObject
}

func NewAccounts(conn *dbus.Conn) Accounts {
	obj := new(objectAccounts)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Accounts1", "/org/deepin/dde/Accounts1")
	return obj
}

type accounts interface {
	GoAllowGuestAccount(flags dbus.Flags, ch chan *dbus.Call, allow bool) *dbus.Call
	AllowGuestAccount(flags dbus.Flags, allow bool) error
	GoCreateGuestAccount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CreateGuestAccount(flags dbus.Flags) (string, error)
	GoCreateUser(flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call
	CreateUser(flags dbus.Flags, name string, fullName string, type0 int32) (dbus.ObjectPath, error)
	GoDeleteUser(flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call
	DeleteUser(flags dbus.Flags, name string, rmFiles bool) error
	GoFindUserById(flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call
	FindUserById(flags dbus.Flags, uid string) (string, error)
	GoFindUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	FindUserByName(flags dbus.Flags, name string) (string, error)
	GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call
	IsPasswordValid(flags dbus.Flags, password string) (bool, string, int32, error)
	GoIsUsernameValid(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	IsUsernameValid(flags dbus.Flags, name string) (bool, string, int32, error)
	GoSetTerminalLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call
	SetTerminalLocked(flags dbus.Flags, locked bool) error
	GoRandUserIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RandUserIcon(flags dbus.Flags) (string, error)
	ConnectUserAdded(cb func(objPath string)) (dbusutil.SignalHandlerId, error)
	ConnectUserDeleted(cb func(objPath string)) (dbusutil.SignalHandlerId, error)
	UserList() proxy.PropStringArray
	GuestIcon() proxy.PropString
	AllowGuest() proxy.PropBool
	IsTerminalLocked() proxy.PropBool
}

type interfaceAccounts struct{}

func (v *interfaceAccounts) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceAccounts) GetInterfaceName_() string {
	return "org.deepin.dde.Accounts1"
}

// method AllowGuestAccount

func (v *interfaceAccounts) GoAllowGuestAccount(flags dbus.Flags, ch chan *dbus.Call, allow bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AllowGuestAccount", flags, ch, allow)
}

func (v *interfaceAccounts) AllowGuestAccount(flags dbus.Flags, allow bool) error {
	return (<-v.GoAllowGuestAccount(flags, make(chan *dbus.Call, 1), allow).Done).Err
}

// method CreateGuestAccount

func (v *interfaceAccounts) GoCreateGuestAccount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateGuestAccount", flags, ch)
}

func (*interfaceAccounts) StoreCreateGuestAccount(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *interfaceAccounts) CreateGuestAccount(flags dbus.Flags) (string, error) {
	return v.StoreCreateGuestAccount(
		<-v.GoCreateGuestAccount(flags, make(chan *dbus.Call, 1)).Done)
}

// method CreateUser

func (v *interfaceAccounts) GoCreateUser(flags dbus.Flags, ch chan *dbus.Call, name string, fullName string, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateUser", flags, ch, name, fullName, type0)
}

func (*interfaceAccounts) StoreCreateUser(call *dbus.Call) (user dbus.ObjectPath, err error) {
	err = call.Store(&user)
	return
}

func (v *interfaceAccounts) CreateUser(flags dbus.Flags, name string, fullName string, type0 int32) (dbus.ObjectPath, error) {
	return v.StoreCreateUser(
		<-v.GoCreateUser(flags, make(chan *dbus.Call, 1), name, fullName, type0).Done)
}

// method DeleteUser

func (v *interfaceAccounts) GoDeleteUser(flags dbus.Flags, ch chan *dbus.Call, name string, rmFiles bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteUser", flags, ch, name, rmFiles)
}

func (v *interfaceAccounts) DeleteUser(flags dbus.Flags, name string, rmFiles bool) error {
	return (<-v.GoDeleteUser(flags, make(chan *dbus.Call, 1), name, rmFiles).Done).Err
}

// method FindUserById

func (v *interfaceAccounts) GoFindUserById(flags dbus.Flags, ch chan *dbus.Call, uid string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserById", flags, ch, uid)
}

func (*interfaceAccounts) StoreFindUserById(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *interfaceAccounts) FindUserById(flags dbus.Flags, uid string) (string, error) {
	return v.StoreFindUserById(
		<-v.GoFindUserById(flags, make(chan *dbus.Call, 1), uid).Done)
}

// method FindUserByName

func (v *interfaceAccounts) GoFindUserByName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FindUserByName", flags, ch, name)
}

func (*interfaceAccounts) StoreFindUserByName(call *dbus.Call) (user string, err error) {
	err = call.Store(&user)
	return
}

func (v *interfaceAccounts) FindUserByName(flags dbus.Flags, name string) (string, error) {
	return v.StoreFindUserByName(
		<-v.GoFindUserByName(flags, make(chan *dbus.Call, 1), name).Done)
}

// method IsPasswordValid

func (v *interfaceAccounts) GoIsPasswordValid(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPasswordValid", flags, ch, password)
}

func (*interfaceAccounts) StoreIsPasswordValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *interfaceAccounts) IsPasswordValid(flags dbus.Flags, password string) (bool, string, int32, error) {
	return v.StoreIsPasswordValid(
		<-v.GoIsPasswordValid(flags, make(chan *dbus.Call, 1), password).Done)
}

// method IsUsernameValid

func (v *interfaceAccounts) GoIsUsernameValid(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsUsernameValid", flags, ch, name)
}

func (*interfaceAccounts) StoreIsUsernameValid(call *dbus.Call) (ok bool, errReason string, errCode int32, err error) {
	err = call.Store(&ok, &errReason, &errCode)
	return
}

func (v *interfaceAccounts) IsUsernameValid(flags dbus.Flags, name string) (bool, string, int32, error) {
	return v.StoreIsUsernameValid(
		<-v.GoIsUsernameValid(flags, make(chan *dbus.Call, 1), name).Done)
}

// method SetTerminalLocked

func (v *interfaceAccounts) GoSetTerminalLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTerminalLocked", flags, ch, locked)
}

func (v *interfaceAccounts) SetTerminalLocked(flags dbus.Flags, locked bool) error {
	return (<-v.GoSetTerminalLocked(flags, make(chan *dbus.Call, 1), locked).Done).Err
}

// method RandUserIcon

func (v *interfaceAccounts) GoRandUserIcon(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RandUserIcon", flags, ch)
}

func (*interfaceAccounts) StoreRandUserIcon(call *dbus.Call) (iconFile string, err error) {
	err = call.Store(&iconFile)
	return
}

func (v *interfaceAccounts) RandUserIcon(flags dbus.Flags) (string, error) {
	return v.StoreRandUserIcon(
		<-v.GoRandUserIcon(flags, make(chan *dbus.Call, 1)).Done)
}

// signal UserAdded

func (v *interfaceAccounts) ConnectUserAdded(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var objPath string
		err := dbus.Store(sig.Body, &objPath)
		if err == nil {
			cb(objPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal UserDeleted

func (v *interfaceAccounts) ConnectUserDeleted(cb func(objPath string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "UserDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".UserDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var objPath string
		err := dbus.Store(sig.Body, &objPath)
		if err == nil {
			cb(objPath)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property UserList as

func (v *interfaceAccounts) UserList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "UserList",
	}
}

// property GuestIcon s

func (v *interfaceAccounts) GuestIcon() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "GuestIcon",
	}
}

// property AllowGuest b

func (v *interfaceAccounts) AllowGuest() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AllowGuest",
	}
}

// property IsTerminalLocked b

func (v *interfaceAccounts) IsTerminalLocked() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "IsTerminalLocked",
	}
}

type User interface {
	user // interface org.deepin.dde.Accounts1.User
	proxy.Object
}

type objectUser struct {
	interfaceUser // interface org.deepin.dde.Accounts1.User
	proxy.ImplObject
}

func NewUser(conn *dbus.Conn, path dbus.ObjectPath) (User, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectUser)
	obj.ImplObject.Init_(conn, "org.deepin.dde.Accounts1", path)
	return obj, nil
}

type user interface {
	GoAddGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call
	AddGroup(flags dbus.Flags, group string) error
	GoDeleteGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call
	DeleteGroup(flags dbus.Flags, group string) error
	GoDeleteIconFile(flags dbus.Flags, ch chan *dbus.Call, icon string) *dbus.Call
	DeleteIconFile(flags dbus.Flags, icon string) error
	GoEnableNoPasswdLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	EnableNoPasswdLogin(flags dbus.Flags, enabled bool) error
	GoGetReminderInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetReminderInfo(flags dbus.Flags) (LoginReminderInfo, error)
	GoIsPasswordExpired(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	IsPasswordExpired(flags dbus.Flags) (bool, error)
	GoPasswordExpiredInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PasswordExpiredInfo(flags dbus.Flags) (int32, int64, error)
	GoSetAutomaticLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	SetAutomaticLogin(flags dbus.Flags, enabled bool) error
	GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, currentWorkspace int32) *dbus.Call
	SetCurrentWorkspace(flags dbus.Flags, currentWorkspace int32) error
	GoSetDesktopBackgrounds(flags dbus.Flags, ch chan *dbus.Call, val []string) *dbus.Call
	SetDesktopBackgrounds(flags dbus.Flags, val []string) error
	GoSetFullName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	SetFullName(flags dbus.Flags, name string) error
	GoSetGreeterBackground(flags dbus.Flags, ch chan *dbus.Call, bg string) *dbus.Call
	SetGreeterBackground(flags dbus.Flags, bg string) error
	GoSetGroups(flags dbus.Flags, ch chan *dbus.Call, groups []string) *dbus.Call
	SetGroups(flags dbus.Flags, groups []string) error
	GoSetHistoryLayout(flags dbus.Flags, ch chan *dbus.Call, list []string) *dbus.Call
	SetHistoryLayout(flags dbus.Flags, list []string) error
	GoSetHomeDir(flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call
	SetHomeDir(flags dbus.Flags, home string) error
	GoSetIconFile(flags dbus.Flags, ch chan *dbus.Call, iconURI string) *dbus.Call
	SetIconFile(flags dbus.Flags, iconURI string) error
	GoSetLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call
	SetLayout(flags dbus.Flags, layout string) error
	GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call
	SetLocale(flags dbus.Flags, locale string) error
	GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call
	SetLocked(flags dbus.Flags, locked bool) error
	GoSetLongDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetLongDateFormat(flags dbus.Flags, value int32) error
	GoSetLongTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetLongTimeFormat(flags dbus.Flags, value int32) error
	GoSetMaxPasswordAge(flags dbus.Flags, ch chan *dbus.Call, nDays int32) *dbus.Call
	SetMaxPasswordAge(flags dbus.Flags, nDays int32) error
	GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call
	SetPassword(flags dbus.Flags, password string) error
	GoSetShell(flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call
	SetShell(flags dbus.Flags, shell string) error
	GoSetShortDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetShortDateFormat(flags dbus.Flags, value int32) error
	GoSetShortTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetShortTimeFormat(flags dbus.Flags, value int32) error
	GoSetUse24HourFormat(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call
	SetUse24HourFormat(flags dbus.Flags, value bool) error
	GoSetWeekBegins(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetWeekBegins(flags dbus.Flags, value int32) error
	GoSetWeekdayFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call
	SetWeekdayFormat(flags dbus.Flags, value int32) error
	ShortDateFormat() proxy.PropInt32
	DesktopBackgrounds() proxy.PropStringArray
	Groups() proxy.PropStringArray
	GreeterBackground() proxy.PropString
	HistoryLayout() proxy.PropStringArray
	UserName() proxy.PropString
	FullName() proxy.PropString
	Use24HourFormat() proxy.PropBool
	UUID() proxy.PropString
	IconFile() proxy.PropString
	LongDateFormat() proxy.PropInt32
	PasswordLastChange() proxy.PropInt32
	LoginTime() proxy.PropUint64
	Gid() proxy.PropString
	PasswordStatus() proxy.PropString
	MaxPasswordAge() proxy.PropInt32
	AutomaticLogin() proxy.PropBool
	Workspace() proxy.PropInt32
	NoPasswdLogin() proxy.PropBool
	Shell() proxy.PropString
	ShortTimeFormat() proxy.PropInt32
	LongTimeFormat() proxy.PropInt32
	Layout() proxy.PropString
	WeekdayFormat() proxy.PropInt32
	CreatedTime() proxy.PropUint64
	XSession() proxy.PropString
	Locked() proxy.PropBool
	HomeDir() proxy.PropString
	Locale() proxy.PropString
	WeekBegins() proxy.PropInt32
	IconList() proxy.PropStringArray
	Uid() proxy.PropString
	SystemAccount() proxy.PropBool
	AccountType() proxy.PropInt32
}

type interfaceUser struct{}

func (v *interfaceUser) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceUser) GetInterfaceName_() string {
	return "org.deepin.dde.Accounts1.User"
}

// method AddGroup

func (v *interfaceUser) GoAddGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddGroup", flags, ch, group)
}

func (v *interfaceUser) AddGroup(flags dbus.Flags, group string) error {
	return (<-v.GoAddGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

// method DeleteGroup

func (v *interfaceUser) GoDeleteGroup(flags dbus.Flags, ch chan *dbus.Call, group string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteGroup", flags, ch, group)
}

func (v *interfaceUser) DeleteGroup(flags dbus.Flags, group string) error {
	return (<-v.GoDeleteGroup(flags, make(chan *dbus.Call, 1), group).Done).Err
}

// method DeleteIconFile

func (v *interfaceUser) GoDeleteIconFile(flags dbus.Flags, ch chan *dbus.Call, icon string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteIconFile", flags, ch, icon)
}

func (v *interfaceUser) DeleteIconFile(flags dbus.Flags, icon string) error {
	return (<-v.GoDeleteIconFile(flags, make(chan *dbus.Call, 1), icon).Done).Err
}

// method EnableNoPasswdLogin

func (v *interfaceUser) GoEnableNoPasswdLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableNoPasswdLogin", flags, ch, enabled)
}

func (v *interfaceUser) EnableNoPasswdLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableNoPasswdLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method GetReminderInfo

func (v *interfaceUser) GoGetReminderInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetReminderInfo", flags, ch)
}

func (*interfaceUser) StoreGetReminderInfo(call *dbus.Call) (info LoginReminderInfo, err error) {
	err = call.Store(&info)
	return
}

func (v *interfaceUser) GetReminderInfo(flags dbus.Flags) (LoginReminderInfo, error) {
	return v.StoreGetReminderInfo(
		<-v.GoGetReminderInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method IsPasswordExpired

func (v *interfaceUser) GoIsPasswordExpired(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IsPasswordExpired", flags, ch)
}

func (*interfaceUser) StoreIsPasswordExpired(call *dbus.Call) (expired bool, err error) {
	err = call.Store(&expired)
	return
}

func (v *interfaceUser) IsPasswordExpired(flags dbus.Flags) (bool, error) {
	return v.StoreIsPasswordExpired(
		<-v.GoIsPasswordExpired(flags, make(chan *dbus.Call, 1)).Done)
}

// method PasswordExpiredInfo

func (v *interfaceUser) GoPasswordExpiredInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PasswordExpiredInfo", flags, ch)
}

func (*interfaceUser) StorePasswordExpiredInfo(call *dbus.Call) (expiredStatus int32, dayLeft int64, err error) {
	err = call.Store(&expiredStatus, &dayLeft)
	return
}

func (v *interfaceUser) PasswordExpiredInfo(flags dbus.Flags) (int32, int64, error) {
	return v.StorePasswordExpiredInfo(
		<-v.GoPasswordExpiredInfo(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetAutomaticLogin

func (v *interfaceUser) GoSetAutomaticLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutomaticLogin", flags, ch, enabled)
}

func (v *interfaceUser) SetAutomaticLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoSetAutomaticLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method SetCurrentWorkspace

func (v *interfaceUser) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, currentWorkspace int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspace", flags, ch, currentWorkspace)
}

func (v *interfaceUser) SetCurrentWorkspace(flags dbus.Flags, currentWorkspace int32) error {
	return (<-v.GoSetCurrentWorkspace(flags, make(chan *dbus.Call, 1), currentWorkspace).Done).Err
}

// method SetDesktopBackgrounds

func (v *interfaceUser) GoSetDesktopBackgrounds(flags dbus.Flags, ch chan *dbus.Call, val []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDesktopBackgrounds", flags, ch, val)
}

func (v *interfaceUser) SetDesktopBackgrounds(flags dbus.Flags, val []string) error {
	return (<-v.GoSetDesktopBackgrounds(flags, make(chan *dbus.Call, 1), val).Done).Err
}

// method SetFullName

func (v *interfaceUser) GoSetFullName(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFullName", flags, ch, name)
}

func (v *interfaceUser) SetFullName(flags dbus.Flags, name string) error {
	return (<-v.GoSetFullName(flags, make(chan *dbus.Call, 1), name).Done).Err
}

// method SetGreeterBackground

func (v *interfaceUser) GoSetGreeterBackground(flags dbus.Flags, ch chan *dbus.Call, bg string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetGreeterBackground", flags, ch, bg)
}

func (v *interfaceUser) SetGreeterBackground(flags dbus.Flags, bg string) error {
	return (<-v.GoSetGreeterBackground(flags, make(chan *dbus.Call, 1), bg).Done).Err
}

// method SetGroups

func (v *interfaceUser) GoSetGroups(flags dbus.Flags, ch chan *dbus.Call, groups []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetGroups", flags, ch, groups)
}

func (v *interfaceUser) SetGroups(flags dbus.Flags, groups []string) error {
	return (<-v.GoSetGroups(flags, make(chan *dbus.Call, 1), groups).Done).Err
}

// method SetHistoryLayout

func (v *interfaceUser) GoSetHistoryLayout(flags dbus.Flags, ch chan *dbus.Call, list []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHistoryLayout", flags, ch, list)
}

func (v *interfaceUser) SetHistoryLayout(flags dbus.Flags, list []string) error {
	return (<-v.GoSetHistoryLayout(flags, make(chan *dbus.Call, 1), list).Done).Err
}

// method SetHomeDir

func (v *interfaceUser) GoSetHomeDir(flags dbus.Flags, ch chan *dbus.Call, home string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetHomeDir", flags, ch, home)
}

func (v *interfaceUser) SetHomeDir(flags dbus.Flags, home string) error {
	return (<-v.GoSetHomeDir(flags, make(chan *dbus.Call, 1), home).Done).Err
}

// method SetIconFile

func (v *interfaceUser) GoSetIconFile(flags dbus.Flags, ch chan *dbus.Call, iconURI string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetIconFile", flags, ch, iconURI)
}

func (v *interfaceUser) SetIconFile(flags dbus.Flags, iconURI string) error {
	return (<-v.GoSetIconFile(flags, make(chan *dbus.Call, 1), iconURI).Done).Err
}

// method SetLayout

func (v *interfaceUser) GoSetLayout(flags dbus.Flags, ch chan *dbus.Call, layout string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLayout", flags, ch, layout)
}

func (v *interfaceUser) SetLayout(flags dbus.Flags, layout string) error {
	return (<-v.GoSetLayout(flags, make(chan *dbus.Call, 1), layout).Done).Err
}

// method SetLocale

func (v *interfaceUser) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *interfaceUser) SetLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoSetLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// method SetLocked

func (v *interfaceUser) GoSetLocked(flags dbus.Flags, ch chan *dbus.Call, locked bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocked", flags, ch, locked)
}

func (v *interfaceUser) SetLocked(flags dbus.Flags, locked bool) error {
	return (<-v.GoSetLocked(flags, make(chan *dbus.Call, 1), locked).Done).Err
}

// method SetLongDateFormat

func (v *interfaceUser) GoSetLongDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongDateFormat", flags, ch, value)
}

func (v *interfaceUser) SetLongDateFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetLongDateFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetLongTimeFormat

func (v *interfaceUser) GoSetLongTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLongTimeFormat", flags, ch, value)
}

func (v *interfaceUser) SetLongTimeFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetLongTimeFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetMaxPasswordAge

func (v *interfaceUser) GoSetMaxPasswordAge(flags dbus.Flags, ch chan *dbus.Call, nDays int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMaxPasswordAge", flags, ch, nDays)
}

func (v *interfaceUser) SetMaxPasswordAge(flags dbus.Flags, nDays int32) error {
	return (<-v.GoSetMaxPasswordAge(flags, make(chan *dbus.Call, 1), nDays).Done).Err
}

// method SetPassword

func (v *interfaceUser) GoSetPassword(flags dbus.Flags, ch chan *dbus.Call, password string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetPassword", flags, ch, password)
}

func (v *interfaceUser) SetPassword(flags dbus.Flags, password string) error {
	return (<-v.GoSetPassword(flags, make(chan *dbus.Call, 1), password).Done).Err
}

// method SetShell

func (v *interfaceUser) GoSetShell(flags dbus.Flags, ch chan *dbus.Call, shell string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShell", flags, ch, shell)
}

func (v *interfaceUser) SetShell(flags dbus.Flags, shell string) error {
	return (<-v.GoSetShell(flags, make(chan *dbus.Call, 1), shell).Done).Err
}

// method SetShortDateFormat

func (v *interfaceUser) GoSetShortDateFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortDateFormat", flags, ch, value)
}

func (v *interfaceUser) SetShortDateFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetShortDateFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetShortTimeFormat

func (v *interfaceUser) GoSetShortTimeFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetShortTimeFormat", flags, ch, value)
}

func (v *interfaceUser) SetShortTimeFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetShortTimeFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetUse24HourFormat

func (v *interfaceUser) GoSetUse24HourFormat(flags dbus.Flags, ch chan *dbus.Call, value bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetUse24HourFormat", flags, ch, value)
}

func (v *interfaceUser) SetUse24HourFormat(flags dbus.Flags, value bool) error {
	return (<-v.GoSetUse24HourFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetWeekBegins

func (v *interfaceUser) GoSetWeekBegins(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWeekBegins", flags, ch, value)
}

func (v *interfaceUser) SetWeekBegins(flags dbus.Flags, value int32) error {
	return (<-v.GoSetWeekBegins(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// method SetWeekdayFormat

func (v *interfaceUser) GoSetWeekdayFormat(flags dbus.Flags, ch chan *dbus.Call, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWeekdayFormat", flags, ch, value)
}

func (v *interfaceUser) SetWeekdayFormat(flags dbus.Flags, value int32) error {
	return (<-v.GoSetWeekdayFormat(flags, make(chan *dbus.Call, 1), value).Done).Err
}

// property ShortDateFormat i

func (v *interfaceUser) ShortDateFormat() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ShortDateFormat",
	}
}

// property DesktopBackgrounds as

func (v *interfaceUser) DesktopBackgrounds() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "DesktopBackgrounds",
	}
}

// property Groups as

func (v *interfaceUser) Groups() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "Groups",
	}
}

// property GreeterBackground s

func (v *interfaceUser) GreeterBackground() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "GreeterBackground",
	}
}

// property HistoryLayout as

func (v *interfaceUser) HistoryLayout() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "HistoryLayout",
	}
}

// property UserName s

func (v *interfaceUser) UserName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "UserName",
	}
}

// property FullName s

func (v *interfaceUser) FullName() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "FullName",
	}
}

// property Use24HourFormat b

func (v *interfaceUser) Use24HourFormat() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Use24HourFormat",
	}
}

// property UUID s

func (v *interfaceUser) UUID() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "UUID",
	}
}

// property IconFile s

func (v *interfaceUser) IconFile() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "IconFile",
	}
}

// property LongDateFormat i

func (v *interfaceUser) LongDateFormat() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "LongDateFormat",
	}
}

// property PasswordLastChange i

func (v *interfaceUser) PasswordLastChange() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "PasswordLastChange",
	}
}

// property LoginTime t

func (v *interfaceUser) LoginTime() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "LoginTime",
	}
}

// property Gid s

func (v *interfaceUser) Gid() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Gid",
	}
}

// property PasswordStatus s

func (v *interfaceUser) PasswordStatus() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "PasswordStatus",
	}
}

// property MaxPasswordAge i

func (v *interfaceUser) MaxPasswordAge() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "MaxPasswordAge",
	}
}

// property AutomaticLogin b

func (v *interfaceUser) AutomaticLogin() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "AutomaticLogin",
	}
}

// property Workspace i

func (v *interfaceUser) Workspace() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "Workspace",
	}
}

// property NoPasswdLogin b

func (v *interfaceUser) NoPasswdLogin() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "NoPasswdLogin",
	}
}

// property Shell s

func (v *interfaceUser) Shell() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Shell",
	}
}

// property ShortTimeFormat i

func (v *interfaceUser) ShortTimeFormat() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "ShortTimeFormat",
	}
}

// property LongTimeFormat i

func (v *interfaceUser) LongTimeFormat() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "LongTimeFormat",
	}
}

// property Layout s

func (v *interfaceUser) Layout() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Layout",
	}
}

// property WeekdayFormat i

func (v *interfaceUser) WeekdayFormat() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "WeekdayFormat",
	}
}

// property CreatedTime t

func (v *interfaceUser) CreatedTime() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "CreatedTime",
	}
}

// property XSession s

func (v *interfaceUser) XSession() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "XSession",
	}
}

// property Locked b

func (v *interfaceUser) Locked() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property HomeDir s

func (v *interfaceUser) HomeDir() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "HomeDir",
	}
}

// property Locale s

func (v *interfaceUser) Locale() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Locale",
	}
}

// property WeekBegins i

func (v *interfaceUser) WeekBegins() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "WeekBegins",
	}
}

// property IconList as

func (v *interfaceUser) IconList() proxy.PropStringArray {
	return &proxy.ImplPropStringArray{
		Impl: v,
		Name: "IconList",
	}
}

// property Uid s

func (v *interfaceUser) Uid() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Uid",
	}
}

// property SystemAccount b

func (v *interfaceUser) SystemAccount() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "SystemAccount",
	}
}

// property AccountType i

func (v *interfaceUser) AccountType() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "AccountType",
	}
}
