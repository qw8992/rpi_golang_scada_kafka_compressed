package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"scada/db"
	"scada/uyeg"
	"time"

	"github.com/valyala/fasthttp"
)

// Loc 변수는 서울 타임존을 의미
var Loc, _ = time.LoadLocation("Asia/Seoul")

// TimeFormat 을 위한 변수
var TimeFormat = "2006-01-02 15:04:05.000"

// GetEnabledDevices 함수는 Enabled 상태의 장치들을 디비로부터 가져옴.
func GetEnabledDevices(dbConn *db.DataBase) map[int]uyeg.Device {
	rows, err := dbConn.Conn.Query(`
	SELECT id,GATEWAY_ID, MAC_ID, NAME, HOST, PORT, UNIT_ID, REMAP_VERSION, PROCESS_INTERVAL, RETRY_CYCLE, RETRY_COUNT, RETRY_CONN_FAILED_COUNT
	FROM DEVICE WHERE ENABLED=True;
	`)

	if err != nil {
		fmt.Println(err.Error())
		return map[int]uyeg.Device{}
	}

	ds := map[int]uyeg.Device{}
	for rows.Next() {
		var device uyeg.Device
		err := rows.Scan(
			&device.Id,
			&device.GatewayId,
			&device.MacId,
			&device.Name,
			&device.Host,
			&device.Port,
			&device.UnitId,
			&device.Version,
			&device.ProcessInterval,
			&device.RetryCycle,
			&device.RetryCount,
			&device.RetryConnFailedCount,
		)

		if err != nil {
			fmt.Println(err.Error())
			dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('system', '%s', NOW());", err.Error()))
			continue
		}

		ds[device.Id] = device
	}

	return ds
}

// GetCompressedString 함수는 압축된 스트링을 반환.
func GetCompressedString(data []byte) string {
	defer func() {
		v := recover()
		if v != nil {
			log.Println("GetCompressedString:", v)
			dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('system', '%s', NOW());", v))
		}
	}()

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		log.Panic(err)
	}
	if err := gz.Close(); err != nil {
		log.Panic(err)
	}
	// fmt.Println(b.Bytes())
	return string(b.Bytes())
}

// SendRequest 함수는 Request를 날린다.
func SendRequest(packet string) {
	defer func() {
		v := recover()
		if v != nil {
			fmt.Println("SendRequest:", v)
			dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('system', '%s', NOW());", v))
		}
	}()

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(conf.API_URL)
	req.Header.Add("Content-Type", "text/plain")
	req.Header.SetMethod("POST")
	req.SetBody([]byte(packet))

	res := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, res)

	fmt.Println(req)
}
