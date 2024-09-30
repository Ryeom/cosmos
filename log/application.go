package log

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var ServerLogDesc *os.File
var AccessLogDesc *os.File
var UpTime string

const (
	ProjectName = "cosmos"
	//DefaultLogPath      = "/var/log/"
	ServerLogFileName   = "server.log"
	AccessLogFileName   = "access.log"
	DashboardTimeFormat = "2006-01-02T15:04:05.999999"
	DefaultDateFormat   = "2006-01-02 15:04:05"
)

var Logger *logging.Logger

func MustInitializeApplicationLog() {
	var err error
	defaultLogPath := viper.GetString("cosmos.log-path")
	logPath := defaultLogPath + ProjectName + "/"
	MustCheckDirectoryPath(logPath)
	serverLogPath := logPath + ServerLogFileName
	MustCheckFilePath(serverLogPath)
	accessLogPath := logPath + AccessLogFileName
	MustCheckFilePath(accessLogPath)
	ServerLogDesc, err = os.OpenFile(serverLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	AccessLogDesc, err = os.OpenFile(accessLogPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	Logger = logging.MustGetLogger(ProjectName)
	back1 := logging.NewLogBackend(ServerLogDesc, "", 0)
	format := logging.MustStringFormatter(`%{color}%{time:0102 15:04:05.000} %{shortfunc:15s} ▶ %{level:.5s}%{color:reset} %{shortfile:15s} %{message}`)
	back1Formatter := logging.NewBackendFormatter(back1, format)
	back1Leveled := logging.AddModuleLevel(back1) // 기본로그 외 추가로그
	back1Leveled.SetLevel(logging.ERROR, "")      // 추가로그의 로그 기본 레벨

	logging.SetBackend(back1Formatter)
	logging.SetLevel(logging.DEBUG, ProjectName)

	Logger.Info(banner)
	Logger.Info("Process initialize ... Env :")
}

func MustCheckDirectoryPath(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func MustCheckFilePath(filePath string) {
	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		_, createErr := os.Create(filePath)
		if createErr != nil {
			panic(createErr)
		}
		//file 처리 다시 하기
	}
}

var banner = `
` + strings.Repeat("░", 150) + `

` + strings.Repeat("▅", 150)

func CreateCustomLogConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_custom}", "transaction_id":"${header:transaction-id}", "status":${status}, "method":"${method}"` +
			`, "uri":"${uri}", "remoteIP":"${remote_ip}", "Client-IP":"${header:Client-Ip}", "host":"${host}"` +
			`, "error":"${error}", "I":${bytes_in}, "O":${bytes_out}, "latency":"${latency_human}"}` + "\n",
		//CustomTimeFormat: "15:04:05.000",
		CustomTimeFormat: DashboardTimeFormat,
		Output:           AccessLogDesc,
	}
}
