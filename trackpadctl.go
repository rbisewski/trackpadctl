/*
 * Trackpad Control Tool
 *
 * Description: A simple tool written in golang, for the purposes of
 *              disabling laptop trackpads on bootup.
 *
 *              Specifically, this will work on kernel version 3.0+ or
 *              newer, depending on whether or not your given trackpad
 *              has a driver / module for its device.
 *
 * Author: Robert Bisewski <contact@ibiscybernetics.com>
 */

//
// Package
//
package main

//
// Imports
//
import (
    "fmt"
    "os"
    "os/exec"
    "regexp"
    "strings"
    "strconv"
)

//
// Globals
//
var (

    // shell command location and flags
    shell = "/bin/sh"
    flags = "-c"

    // input command
    cmd_xinput  = "xinput"
    cmd_list    = "list"
    cmd_id_only = "--id-only"
    cmd_disable = "--disable"

    // name of the trackpad device to disable
    trackpad_device_name = "\"FocalTechPS/2 FocalTech Touchpad\""
)

//
// PROGRAM MAIN
//
func main() {

    // String variable to hold eventual output, as well error variable.
    var output []byte;
    var cmd     string   = ""
    var tmp_str string   = ""
    var id_as_str string = ""
    var err error        = nil

    // attempt to exec the xinput command, which output somekind feedback
    output, err = exec.Command(shell, flags, cmd_xinput).Output()

    // safety check, ensure that no error occurred
    if err != nil {
        fmt.Println("Error: xinput not found on system. Exiting...")
        os.Exit(1)
    }

    // attempt to cast the byte output to a string
    tmp_str = string(output)

    // safety check, ensure the command actually gave back string data,
    // since `xinput` by default will print all of the devices
    if len(tmp_str) < 1 {
        fmt.Println("Note: xinput shows no devices on system. Exiting...")
        os.Exit(1)
    }

    // assemble the command to obtain the device ID
    cmd = cmd_xinput + " " + cmd_list + " " + cmd_id_only + " " +
      trackpad_device_name

    // attempt to exec the command, which should return an integer value
    // of which id the device in question is mapped too
    output, err = exec.Command(shell, flags, cmd).Output()

    // safety check, ensure that no error occurred
    if err != nil {
        fmt.Println("The following device may not exist:",
          trackpad_device_name)
        fmt.Println("Error: running xinput --id-only failed. Exiting...")
        os.Exit(1)
    }

    // convert the byte[] into a string
    id_as_str = string(output)

    // remove whitespace from the command output
    id_as_str = strings.TrimSpace(id_as_str)

    // safety check, ensure the response is non-blank
    if len(id_as_str) < 1 {
        fmt.Println("Error: xinput gave non-valid device id. Exiting...")
        os.Exit(1)
    }

    // cast the id to an integer to ensure it is actually an id
    _, err = strconv.Atoi(id_as_str)

    // safety check, ensure that no error occurred
    if err != nil {
        fmt.Println(err)
        fmt.Println("Error: casting ID to uint failed. Exiting...")
        os.Exit(1)
    }

    // since the id is probably valid, attempt to disable the device in
    // question
    cmd = cmd_xinput + " " + cmd_disable + " " + id_as_str
    output, err = exec.Command(shell, flags, cmd).Output()

    // safety check, ensure that no error occurred
    if err != nil {
        fmt.Println("Error: disabling '", trackpad_device_name,
          "' device failed. Exiting...")
        os.Exit(1)
    }

    // dump the bytes to a string
    tmp_str = string(output)

    // if the length of the converted output is 0, then everything probably
    // worked as intended, so exit(0) here
    if len(tmp_str) < 1 {
        os.Exit(0)
    }

    // assemble a regex of the possible error message
    re := regexp.MustCompile("unable to find device")

    // further check, ensure that xinput could actually find the device
    verify := re.FindString(tmp_str)

    // if the error string is present, go ahead and quit
    if len(verify) > 0 {
        fmt.Println("Error: xinput could not find the '",
          trackpad_device_name, "' device. Exiting...")
        os.Exit(1)
    }

    // otherwise everything worked, so return 0 here
    os.Exit(0)
}
