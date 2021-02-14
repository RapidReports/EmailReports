package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	gfunc "github.com/arctheowl/EmailReports/GmailFunctions"
	"go.uber.org/zap"
)

//Config is the configuration set out in the config.json file
//It will include
type Config struct {
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	for {
		logger := createLogger()
		logger.Info("Starting Server",
			"Category 1", "yes",
		)
		//config := LoadConfiguration("config.json")
		mail := gfunc.SelectMail()
		CsvData := gfunc.GetAttachmentData(mail)
		fmt.Println(CsvData)

		time.Sleep(30 * time.Second)
	}
}

//LoadConfiguration - will load config.json file from this folder to define the Config struct
func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

//CreateLogFile creates the log file that logs will be added to.
func createLogger() *zap.SugaredLogger {

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"./logs/production.log"}
	config.ErrorOutputPaths = []string{"./logs/production_err.log"}

	logger, _ := config.Build()
	Sugar := logger.Sugar()
	return Sugar
}
