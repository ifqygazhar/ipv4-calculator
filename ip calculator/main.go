package main

import (
	"fmt"
	"log"
)

func main() {
	ascii := fmt.Sprintln(`

	█ █▀█ █░█ █░█   █▀▀ ▄▀█ █░░ █▀▀ █░█ █░░ ▄▀█ ▀█▀ █▀█ █▀█
	█ █▀▀ ▀▄▀ ▀▀█   █▄▄ █▀█ █▄▄ █▄▄ █▄█ █▄▄ █▀█ ░█░ █▄█ █▀▄
	Coded by : Ifqy gifha azhar

	contoh ip : 192.168.1.1
	contoh prefix: 24
	`)
	fmt.Println(ascii)

	fmt.Print("masukan ip address :")
	var ip string
	fmt.Scanln(&ip)
	fmt.Print("masukan prefix subnetmask :")
	var subnet int
	fmt.Scanln(&subnet)

	if len(ip) < 11 && len(ip) < 10 {
		log.Fatal("error bukan ipv4")
	}
	switch subnet {
	case 9:

		fmt.Println("address", ip, "subnetmask 255.128.0.0 host 8388606")
	case 10:

		fmt.Println("address", ip, "subnetmask 255.192.0.0 host 4194302")
	case 11:

		fmt.Println("address", ip, "subnetmask 255.224.0.0 host 2097150")
	case 12:

		fmt.Println("address", ip, "subnetmask 255.240.0.0 host 1048574")
	case 13:

		fmt.Println("address", ip, "subnetmask 255.248.0.0 host 524286")
	case 14:

		fmt.Println("address", ip, "subnetmask 255.252.0.0 host 262142")
	case 15:

		fmt.Println("address", ip, "subnetmask 255.254.0.0 host 131070")
	case 16:

		fmt.Println("address", ip, "subnetmask 255.255.0.0 host 65534")
	case 17:

		fmt.Println("address", ip, "subnetmask 255.255.128.0 host 32766")
	case 18:

		fmt.Println("address", ip, "subnetmask 255.255.192.0 host 16382")
	case 19:

		fmt.Println("address", ip, "subnetmask 255.255.224.0 host 8190")
	case 20:

		fmt.Println("address", ip, "subnetmask 255.255.240.0 host 4094")
	case 21:

		fmt.Println("address", ip, "subnetmask 255.255.248.0 host 2046")
	case 22:

		fmt.Println("address", ip, "subnetmask 255.255.252.0 host 1022")
	case 23:

		fmt.Println("address", ip, "subnetmask 255.255.254.0 host 510")
	case 24:

		fmt.Println("address", ip, "subnetmask 255.255.255.0 host 254")
	case 25:

		fmt.Println("address", ip, "subnetmask 255.255.255.128 host 126")
	case 26:

		fmt.Println("address", ip, "subnetmask 255.255.255.192 host 62")
	case 27:

		fmt.Println("address", ip, "subnetmask 255.255.255.224 host 30")
	case 28:

		fmt.Println("address", ip, "subnetmask 255.255.255.240 host 14")
	case 29:

		fmt.Println("address", ip, "subnetmask 255.255.255.248 host 6")
	case 30:

		fmt.Println("address", ip, "subnetmask 255.255.255.252 host 2")
	default:
		fmt.Println("yang anda masukan bukan angka atau angka prefix subnetmask kurang dari 9 atau lebih dari 30")
	}

}
