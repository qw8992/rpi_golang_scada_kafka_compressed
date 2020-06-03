package main

import (
	"fmt"
	//	"log"
	//	"os"
	"scada/uyeg"
	"time"
)

// UYeGDataCollection 함수는 데이터를 수집하는 함수입니다
func UYeGDataCollection(client *uyeg.ModbusClient, collChan chan<- map[string]interface{}) {
	var errCount, errCountConn = 0, 0
	ticker := time.NewTicker(10 * time.Millisecond)

	for {
		select {
		case <-client.Done1:
			fmt.Println(fmt.Sprintf("=> %s (%s:%d) 데이터 수집 종료", client.Device.MacId, client.Device.Host, client.Device.Port))
			return
		case <-ticker.C:
			readData := client.GetReadHoldingRegisters()

			if readData == nil {
				//test:= readData;
				ticker.Stop()
				errCount = errCount + 1
				fmt.Println(time.Now().In(Loc).Format(TimeFormat), fmt.Sprintf("Failed to read data Try again (%s:%d)..", client.Device.Host, client.Device.Port))
				log1 := fmt.Sprintf("Failed to read data Try again (%s:%d)..", client.Device.Host, client.Device.Port)
				dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, log1))
				if errCount > client.Device.RetryCount {
					client.Handler.Close()
					if client.Connect() {
						fmt.Println(time.Now().In(Loc).Format(TimeFormat), fmt.Sprintf("Succeded to reconnect the connection.. (%s:%d)..", client.Device.Host, client.Device.Port))
						errCount = 0
					} else {
						fmt.Println(time.Now().In(Loc).Format(TimeFormat), fmt.Sprintf("Failed to reconnect the connection.. (%s:%d)..", client.Device.Host, client.Device.Port))
						log2 := fmt.Sprintf("Failed to reconnect the connection.. (%s:%d)..", client.Device.Host, client.Device.Port)
						dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, log2))
						errCountConn = errCountConn + 1

						if errCountConn > client.Device.RetryConnFailedCount {
							derr := make(map[string]interface{})
							derr["Device"] = client.Device
							derr["Error"] = fmt.Sprintf("%s(%s): Connection failed..", client.Device.Name, client.Device.MacId)
							derr["Restart"] = false

							ErrChan <- derr

							dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, derr["Error"].(string)))
						}
					}
				}
				time.Sleep(time.Duration(client.Device.RetryCycle) * time.Millisecond)
				ticker = time.NewTicker(10 * time.Millisecond)
				continue
			} else {
				errCount = 0
			}

			collChan <- readData
		}
	}
}
