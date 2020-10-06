package main

import (
	"fmt"
	"scada/uyeg"
)

func UYeGStartFunc(client *uyeg.ModbusClient) {
	fmt.Println(client.Device.Host, "uyeg start func")
	defer func() {
		v := recover()

		if v != nil {
			derr := make(map[string]interface{})
			derr["Device"] = client.Device
			derr["Error"] = v
			derr["Restart"] = true

			ErrChan <- derr
		}
	}()

	if !client.Connect() {
		derr := make(map[string]interface{})
		derr["Device"] = client.Device
		derr["Error"] = fmt.Sprintf("%s(%s): Connection failed", client.Device.Name, client.Device.MacId)
		derr["Restart"] = false

		ErrChan <- derr
		dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, derr["Error"].(string)))
	}

	collChan := make(chan map[string]interface{}, 20)
	tfChan := make(chan []interface{}, 20)
	chValue := make(chan string)

	go kafka(chValue, client)
	go UYeGTransfer(client, tfChan, chValue)
	go UYeGProcessing(client, collChan, tfChan)
	go UYeGDataCollection(client, collChan)
}
