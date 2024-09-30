package util

import (
	"errors"
	"github.com/Ryeom/cosmos/log"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strconv"
	"strings"
)

/* 1. mode 설정 : local dev stage prod */

func MustSetArguments() error {
	if len(os.Args) < 2 {
		err := errors.New("Process mode is not exists. arguments length : " + strconv.Itoa(len(os.Args)))
		return err
	}
	allowEnv := append([]string{"local", "dev", "stage", "prod"})
	if !Contains(allowEnv, os.Args[1]) {
		err := errors.New("Process mode is not allowed." + strconv.Itoa(len(os.Args)))
		return err
	}
	return nil
}

func MustInitializeSetting() {
	var configPath string
	var err error
	var fileName = os.Args[1] + "-settings" // 기본 설정 파일명
	var extension = "toml"
	if len(os.Args) == 2 {
		log.Logger.Error("it works in build file path")
		configPath, err = os.Getwd()
		if err != nil {
			log.Logger.Error("work directory ", err)
			err = errors.New("Failed reading work directory. " + err.Error())
		}
	} else {
		pathDepth := strings.Split(os.Args[2], "/")
		file := strings.Split(pathDepth[len(pathDepth)-1:][0], ".")
		fileName = file[0]
		extension = file[len(file)-1]
		configPath = os.Args[2][0 : len(os.Args[2])-len(fileName)-len(extension)-2]
	}
	viper.SetConfigName(fileName)
	viper.SetConfigType(extension)
	viper.AddConfigPath(configPath)
	log.Logger.Info("Reading configuration from", configPath)

	err = viper.ReadInConfig()
	if err != nil {
		log.Logger.Error("Failed reading configuration .. ", err)
		err = errors.New("Failed reading configuration. " + err.Error())

	}

	goos := runtime.GOOS
	mark := false
	if goos == "windows" || goos == "darwin" {
		mark = true
	}
	cosmosKey := viper.GetString("cosmos.key")
	for k, v := range viper.AllKeys() {
		if strings.HasPrefix(v, "cosmos.") {
			continue
		}
		value := viper.GetString(v)
		if value == "" {
			continue
		}

		originalValue, decErr := DecryptAES(value, []byte(cosmosKey))
		if decErr != nil {
			log.Logger.Error("Failed reading configuration value Decrypt .. ", decErr, k, v)
			err = errors.New("Failed reading configuration value Decrypt. " + decErr.Error())

		}
		viper.Set(v, originalValue)
		if mark {
			log.Logger.Info("[set config]", v, " : ", originalValue)
		}
	}
	// 기타 강제 설정
	ip := GetLocalIP()
	viper.SetDefault("cosmos.local-ip", ip)

}
