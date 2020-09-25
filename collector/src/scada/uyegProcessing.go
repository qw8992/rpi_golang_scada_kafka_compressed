package main

import (
	"fmt"
	//"log"
	"math"
	//"os"
	"scada/uyeg"
	"strings"
	"time"
)

// UYeGProcessing 함수는 수집된 데이터를 정제, 처리하는 함수
func UYeGProcessing(client *uyeg.ModbusClient, collChan <-chan map[string]interface{}, tfChan chan<- []interface{}) {

	var queue ItemQueue
	if queue.items == nil {
		queue = ItemQueue{}
		queue.New()
	}

	//var queue *Queue = New()
	go QueueProcess(client, &queue, tfChan)

	for {
		select {
		case <-client.Done2:
			fmt.Println(fmt.Sprintf("=> %s (%s:%d) 데이터 처리 종료", client.Device.MacId, client.Device.Host, client.Device.Port))
			return
		case data := <-collChan:
			queue.Enqueue(data)

		}
		time.Sleep(1 * time.Millisecond)
	}
}

func QueueProcess(client *uyeg.ModbusClient, queue *ItemQueue, tfChan chan<- []interface{}) {
	syncMap := SyncMap{v: make(map[string]interface{})}
	ds := make([]interface{}, 0, 100)   // 미리 공간 할당해둠1
	var lastData map[string]interface{} // 미리 공간 할당해둠2

	//var preTime  *string = new(string)

	for {
		for len(queue.items) > 0 {
			data := (*queue).Dequeue()

			t := (*data).(map[string]interface{})["time"].(time.Time).Truncate(time.Duration(client.Device.ProcessInterval) * time.Millisecond).Format(TimeFormat)

			if v := syncMap.Get(t); v != nil { // 데이터가 있는경우

				tv := make(map[string]interface{})
				for k, v := range v.(map[string]interface{}) {
					var tmp float64
					if strings.Contains(k, "time") {
						tv[k] = v
						continue
					} else if strings.Contains(k, "Volt") {
						tmp = math.Min(v.(float64), (*data).(map[string]interface{})[k].(float64))
					} else {
						tmp = math.Max(v.(float64), (*data).(map[string]interface{})[k].(float64))
					}
					tv[k] = tmp
				}
				syncMap.Set(t, tv)

				//}
			} else {
				(*data).(map[string]interface{})["time"] = t
				syncMap.Set(t, (*data).(map[string]interface{}))

				tmillisecond := t[len(t)-4:]
				t2, _ := time.Parse(TimeFormat[:len(TimeFormat)-4], t)
				//fmt.Println("nowtime : ", t, "tmillisecond: ", tmillisecond, ", t2:", t2, ", len(syncMap) = ", syncMap.Size())
				if tmillisecond == ".000" && syncMap.Size() >= 10 {
					sMap := syncMap.GetMap()
					bSecT, _ := time.Parse(TimeFormat[:len(TimeFormat)-4], t2.Add(-1 * time.Second).Format(TimeFormat)[:len(TimeFormat)-4])
					//	fmt.Println("nowtime : ", t, "t2: ", t2, "bSecT: ", bSecT)

					// .000 부터 데이터 비교.
					for i := 0; i < 1000/client.Device.ProcessInterval; i++ {
						vT := bSecT.Add(time.Duration(i*client.Device.ProcessInterval) * time.Millisecond).Format(TimeFormat)
						if val, exists := sMap[vT]; exists == true { // 데이터가 있는 경우.
							value := val.(map[string]interface{})
							value["status"] = true
							//							fmt.Println("1 :", ds)
							ds = append(ds, value)
							//							fmt.Println("2 :", ds)
							lastData = val.(map[string]interface{}) // 마지막 데이터를 초기화 시킨다.
							syncMap.Delete(vT)                      // 추가한 데이터는 삭제한다.
						} else { // 데이터가 없는 경우.
							if lastData != nil {
								ld := CopyMap(lastData)
								ld["time"] = bSecT.Format(TimeFormat)
								ld["status"] = false
								//								fmt.Println("1 :", ds)
								ds = append(ds, ld)
								//								fmt.Println("2 :", ds)
							}
							fmt.Println(" No Data ", vT)

							derr := make(map[string]interface{})
							derr["Device"] = client.Device
							derr["Error"] = fmt.Sprintf(" No Data ", vT)
							derr["Restart"] = false

							ErrChan <- derr
							dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('%s', '%s', NOW());", client.Device.MacId, derr["Error"].(string)))
						}
					}
					tfChan <- ds
					ds = ds[:0] // 데이터 삭제
				}
			}
			time.Sleep(1 * time.Millisecond)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
