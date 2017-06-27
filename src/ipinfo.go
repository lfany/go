package main

import (
	"github.com/revh/ipinfo"
	"encoding/json"
	"fmt"
)

func main() {
	ip, err := ipinfo.MyIP()
	foreignIP, err := ipinfo.ForeignIP("8.8.8.8")
	ip = foreignIP
	if err != nil {
		println("err: ", err.Error())
	}

	marshal, err := json.Marshal(ip)
	if err != nil {
		println("err: ", err.Error())
	}
	var ipinfo ipinfo.IPInfo
	err = json.Unmarshal(marshal, &ipinfo)
	if err != nil {
		println("err: ", err.Error())
	}
	fmt.Println("IP: ", ipinfo.IP)
	fmt.Println("Hostname: ", ipinfo.Hostname)
	fmt.Println("City: ", ipinfo.City)
	fmt.Println("Region: ", ipinfo.Region)
	fmt.Println("Country: ", ipinfo.Country)
	fmt.Println("Loc: ", ipinfo.Loc)
	fmt.Println("Org: ", ipinfo.Org)
	fmt.Println("Postal: ", ipinfo.Postal)
}
