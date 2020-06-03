package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	MYSQL_HOST          string
	MYSQL_PORT          int
	MYSQL_ROOT_PASSWORD string
	MYSQL_DATABASE      string
	MYSQL_USER          string
	MYSQL_PASSWORD      string
	API_URL             string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
		dbConn.NotResultQueryExec(fmt.Sprintf("INSERT INTO E_LOG(MAC_ID, LOG, CREATE_DATE) VALUES ('system', '%s', NOW());", err))
		os.Exit(500)
	}
	return configuration
}

func getFileName() string {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "dev"
	}

	filename := []string{"config.", env, ".json"}
	//fmt.Println(filename)
	_, dirname, _, _ := runtime.Caller(0)
	//filePath := path.Join("./config/", strings.Join(filename, ""))
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	//fmt.Println(filePath)
	return filePath
}
