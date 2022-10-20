package config

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func init() {
	systemConfigDir, _ := os.UserConfigDir()
	configDir := fmt.Sprintf("%s/%s", systemConfigDir, "droidjs")
	filepath := fmt.Sprintf("%s/%s", configDir, "config.yaml")

	_, err := os.Stat(configDir)
	if os.IsNotExist(err) {
		if err := os.Mkdir(configDir, os.ModePerm); err != nil { // perm 0666
			panic(err)
		}
	}
	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(configDir) // call multiple times to add many search paths
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = viper.WriteConfigAs(filepath)
			if err != nil {
				panic(err)
			}
		} else {
			// Config file was found but another error was produced
		}
	}

}
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
func SetupConfig(args []string, options map[string]string) int {
	edit, _ := strconv.ParseBool(options["edit"])
	fmt.Printf("edit==>%v\n", edit)
	address := viper.GetString("server.address")
	token := viper.GetString("server.token")
	if !edit {
		address = ""
		token = ""
		//return 0
	}
	method := promptui.Select{
		Label: "选择配置方式",
		Items: []string{"手动输入", "App扫码"},
	}

	_, result, err := method.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 1
	}

	fmt.Printf("You choose %q\n", result)
	inputAddress := promptui.Prompt{
		Label:   "服务器地址",
		Default: address,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("错误的服务器地址")
			}
			return nil
		},
	}
	resultAddress, err := inputAddress.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 1
	}
	fmt.Printf("api server: %s\n", resultAddress)

	inputToken := promptui.Prompt{
		Label:   "访问token",
		Default: token,
		Validate: func(input string) error {
			if input == "" {
				return errors.New("token错误")
			}
			return nil
		},
	}
	resultToken, err := inputToken.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 1
	}
	fmt.Printf("api token: %s\n", resultToken)
	viper.Set("server.address", resultAddress)
	viper.Set("server.token", resultToken)
	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}
	println("已保存新的配置")
	return 0
}
