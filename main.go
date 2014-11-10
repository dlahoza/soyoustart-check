package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    "os"
)

func main() {
    SERVER:=os.Args[1]
    resp, err := http.Get("https://ws.ovh.com/dedicated/r2/ws.dispatcher/getAvailability2")
    if err==nil {
	data,_:=ioutil.ReadAll(resp.Body)
	strresp:=string(data)
	strresp=strings.SplitAfter(strresp,SERVER)[1]
	strresp=strings.Split(strresp,"[")[1]
	strresp=strings.Split(strresp,"]")[0]
	strresp=strings.Replace(strresp,"\"__class\":\"dedicatedType:dedicatedAvailability2ZoneStruct\",","",10)
	strresp="["+strresp+"]"
	type zoneType struct {
		Availability string
		Zone string
	}
	var zones []zoneType
	err := json.Unmarshal([]byte(strresp), &zones)
	if err != nil {
	    fmt.Println("error:", err)
	}
	num:=0
	for _,zone:=range zones {
	    if (zone.Availability!="unavailable") {
		fmt.Println("Zone: "+zone.Zone+"\n"+"Availability: "+zone.Availability)
		num++
	    }
	}
	os.Exit(num)
    }
}