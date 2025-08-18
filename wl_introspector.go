package main

import (
	"os/exec"
	"strings"

	"github.com/godbus/dbus/v5"
)

func wlGetFocusWindowClass() (string, error) {
	if isGnome {
		return gnomeGetFocusWindowClass()
	}
	if isKDE {
		return kdeGetFocusWindowClass()
	}
	return "", nil
}

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

func kdeGetFocusWindowClass() (string, error) {
	var getActiveWindowCmd = exec.Command("kdotool", "getactivewindow")

	windowActiveId, err := getActiveWindowCmd.Output()
	if err != nil {
		return "", err
	}

	var getWindowClassnameCmd = exec.Command("kdotool", "getwindowclassname", strings.TrimSpace(string(windowActiveId)))
	windowClassname, err := getWindowClassnameCmd.Output()
	if err != nil {
		return "", err
	}

	return string(windowClassname), nil
}
