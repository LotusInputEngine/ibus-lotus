package main

import (
	"github.com/godbus/dbus/v5"
)

var wlAppId string

func gnomeGetFocusWindowClass() (string, error) {
	// Install Window Call Extended extension to make this work
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	obj := conn.Object(
		"org.gnome.Shell",
		dbus.ObjectPath("/org/gnome/Shell/Extensions/WindowsExt"),
	)

	var className string
	call := obj.Call("org.gnome.Shell.Extensions.WindowsExt.FocusClass", 0)
	if call.Err != nil {
		return "", call.Err
	}

	if err := call.Store(&className); err != nil {
		return "", err
	}

	return className, nil
}

func wlGetFocusWindowClass() (string, error) {
	if isGnome {
		return gnomeGetFocusWindowClass()
	}
	return "", nil
}
