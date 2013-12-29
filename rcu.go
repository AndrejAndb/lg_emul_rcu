/*
   Alternative script for Control LG smartTV emulator (2012/2013)
   ver 0.1

   Copyright (c) 2013 lg_emul_rcu
   See the file license.txt for copying permission.
*/
package main

import "fmt"
import "net"
import "flag"
import "strings"
import "os"
import "bufio"

const remoteHost = "127.0.0.1"
const remoteEmulControlPort = "54333"
const remoteEmulRemoconPort = "19001"

func int16ToByteArray(c uint16) [2]byte {
	var result [2]byte
	result[0] = byte(c / 256)
	result[1] = byte(c % 256)
	return result
}

func EmulControlCommand(command uint16, arg []byte) (int, error) {
	var result int = 1
	var err error = nil
	conn, err := net.Dial("tcp", remoteHost+":"+remoteEmulControlPort)
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	var commandData = int16ToByteArray(command)
	var argLen = int16ToByteArray(uint16(len(arg)))
	var data []byte
	data = append(data, commandData[:]...)
	data = append(data, argLen[:]...)
	data = append(data, arg[:]...)

	result, err = conn.Write(data)
	if err != nil {
		return result, err
	}

	return result, err
}

func main() {

	fmt.Printf("LG Emul RCU v0.1\n")

	var url *string = flag.String("url", "none", "Url of App for test")
	flag.Parse()

	if !strings.EqualFold(*url, "none") {
		fmt.Printf("> Opening \"%s\"...", *url)
		//EmulControlCommand(2049, []byte("NetCast.RCU.RunBrowser"))
		_, err := EmulControlCommand(1026, []byte(*url))
		//EmulControlCommand(261, []byte("ACK"))
		if err != nil {
			fmt.Printf(" error (%s).\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf(" done.\n")
	}

	fmt.Printf("> Connecting to Remocon...")
	remocon, err := net.Dial("tcp", remoteHost+":"+remoteEmulRemoconPort)
	if err != nil {
		fmt.Printf(" error (%s).\n", err.Error())
		os.Exit(1)
	}
	defer remocon.Close()
	fmt.Printf(" done.\n")

	var exit bool = false
	bio := bufio.NewReader(os.Stdin)
	var n int

	commands := map[string]map[string]string{
		"quit, exit, q": map[string]string{
			"help":    "Quit from program",
			"remocon": "none",
		},
		"?": map[string]string{
			"help":    "Program Help",
			"remocon": "none",
		},
		"help": map[string]string{
			"help":    "Key help",
			"remocon": "201",
		},
		"ratio": map[string]string{
			"help":    "Key ratio",
			"remocon": "202",
		},
		"input": map[string]string{
			"help":    "Key input",
			"remocon": "203",
		},
		"tv": map[string]string{
			"help":    "Key tv",
			"remocon": "204",
		},
		"1": map[string]string{
			"help":    "Key 1",
			"remocon": "301",
		},
		"2": map[string]string{
			"help":    "Key 2",
			"remocon": "302",
		},
		"3": map[string]string{
			"help":    "Key 3",
			"remocon": "303",
		},
		"4": map[string]string{
			"help":    "Key 4",
			"remocon": "304",
		},
		"5": map[string]string{
			"help":    "Key 5",
			"remocon": "305",
		},
		"6": map[string]string{
			"help":    "Key 6",
			"remocon": "306",
		},
		"7": map[string]string{
			"help":    "Key 7",
			"remocon": "307",
		},
		"8": map[string]string{
			"help":    "Key 8",
			"remocon": "308",
		},
		"9": map[string]string{
			"help":    "Key 9",
			"remocon": "309",
		},
		"0": map[string]string{
			"help":    "Key 0",
			"remocon": "310",
		},
		"list": map[string]string{
			"help":    "Key list",
			"remocon": "311",
		},
		"qview": map[string]string{
			"help":    "Key qview",
			"remocon": "312",
		},
		"vol_up": map[string]string{
			"help":    "Key vol_up",
			"remocon": "401",
		},
		"fav": map[string]string{
			"help":    "Key fav",
			"remocon": "402",
		},
		"chup_btn": map[string]string{
			"help":    "Key chup_btn",
			"remocon": "403",
		},
		"3d": map[string]string{
			"help":    "Key 3d",
			"remocon": "404",
		},
		"vol_dw": map[string]string{
			"help":    "Key vol_dw",
			"remocon": "405",
		},
		"mute": map[string]string{
			"help":    "Key mute",
			"remocon": "406",
		},
		"chdw": map[string]string{
			"help":    "Key chdw",
			"remocon": "407",
		},
		"settings": map[string]string{
			"help":    "Key settings",
			"remocon": "501",
		},
		"home": map[string]string{
			"help":    "Key home",
			"remocon": "502",
		},
		"apps": map[string]string{
			"help":    "Key apps",
			"remocon": "503",
		},
		"up": map[string]string{
			"help":    "Key up",
			"remocon": "504",
		},
		"left": map[string]string{
			"help":    "Key left",
			"remocon": "505",
		},
		"center": map[string]string{
			"help":    "Key center",
			"remocon": "506",
		},
		"right": map[string]string{
			"help":    "Key right",
			"remocon": "507",
		},
		"down": map[string]string{
			"help":    "Key down",
			"remocon": "508",
		},
		"back": map[string]string{
			"help":    "Key back",
			"remocon": "509",
		},
		"guide": map[string]string{
			"help":    "Key guide",
			"remocon": "510",
		},
		"exit": map[string]string{
			"help":    "Key exit",
			"remocon": "511",
		},
		"red": map[string]string{
			"help":    "Key red",
			"remocon": "601",
		},
		"green": map[string]string{
			"help":    "Key green",
			"remocon": "602",
		},
		"yellow": map[string]string{
			"help":    "Key yellow",
			"remocon": "603",
		},
		"blue": map[string]string{
			"help":    "Key blue",
			"remocon": "604",
		},
		"text": map[string]string{
			"help":    "Key text",
			"remocon": "701",
		},
		"topt": map[string]string{
			"help":    "Key topt",
			"remocon": "702",
		},
		"qmenu": map[string]string{
			"help":    "Key qmenu",
			"remocon": "703",
		},
		"stop": map[string]string{
			"help":    "Key stop",
			"remocon": "704",
		},
		"play": map[string]string{
			"help":    "Key play",
			"remocon": "705",
		},
		"pause": map[string]string{
			"help":    "Key pause",
			"remocon": "706",
		},
		"rw": map[string]string{
			"help":    "Key rw",
			"remocon": "707",
		},
		"ff": map[string]string{
			"help":    "Key ff",
			"remocon": "708",
		},
		"rec": map[string]string{
			"help":    "Key rec",
			"remocon": "709",
		},
		"energy": map[string]string{
			"help":    "Key energy",
			"remocon": "710",
		},
		"info": map[string]string{
			"help":    "Key info",
			"remocon": "711",
		},
		"simplink": map[string]string{
			"help":    "Key simplink",
			"remocon": "712",
		},
	}

	for !exit {
		fmt.Printf("< ")
		data, _, _ := bio.ReadLine()
		n = len(data)
		s := string(data[:n])
		switch s {
		case "quit", "exit", "q":
			fmt.Printf("> exit\n")
			exit = true
		case "?":
			fmt.Printf("> Commands:\n")
			for k, v := range commands {
				fmt.Printf("  %-15s - %s (%s)\n", k, v["help"], v["remocon"])
			}
		default:
			if command, ok := commands[s]; ok {
				_, err = remocon.Write([]byte(command["remocon"]))
				if err != nil {
					fmt.Printf("Error on send command: '%s'\n", err.Error())
				}
				fmt.Printf("> %s\n", command["help"])
			} else {
				fmt.Printf("> Unknown command\n")
			}
		}
	}
}
