package main

import (
	"fmt"
	//"log"
	//"os"
	"scada/uyeg"
)

func UYeGTransfer(client *uyeg.ModbusClient, tfChan <-chan []interface{}) {
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
				SendRequest(GetCompressedString(jsonBytes))
				fmt.Println("Mac ", client.Device.MacId, " send Time:", t)
				logs := fmt.Sprintf("Mac %s send Time: %s", client.Device.MacId, t)
				dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO S_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, logs))
			}
		}
	}
}
