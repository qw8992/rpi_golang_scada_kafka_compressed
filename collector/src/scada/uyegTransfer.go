package main

import (
	"fmt"

	//"log"
	//"os"
	"scada/uyeg"
)

func UYeGTransfer(client *uyeg.ModbusClient, tfChan <-chan []interface{}, chValue chan string) {
	for {
		select {
		case <-client.Done3:
			fmt.Println(fmt.Sprintf("=> %s (%s:%d) 데이터 전송 종료", client.Device.MacId, client.Device.Host, client.Device.Port))
			return
		case data := <-tfChan:
			d := data[0].(map[string]interface{})
			if t, exists := d["time"]; exists {
				bSecT := t.(string)[:len(TimeFormat)-4]
				jsonBytes := client.GetRemapJson(bSecT, data)
				// jsonData := fmt.Sprint(string(jsonBytes))
				// fmt.Println(jsonData)
				jsonData := fmt.Sprint(GetCompressedString(jsonBytes))
				chValue <- jsonData
			}
		}
	}
}
