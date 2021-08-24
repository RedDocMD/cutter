package main

import (
	"fmt"
	"os"

	"github.com/RedDocMD/cutter/conf"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.cutter")
	viper.AddConfigPath("$HOME/.config/cutter")
	viper.AddConfigPath(".")

	var conf conf.Config

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Fprintln(os.Stderr, "Failed to find config file")
			os.Exit(127)
		} else {
			fmt.Fprintln(os.Stderr, "Other fatal error:", err)
			os.Exit(1)
		}
	}

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse config file:", err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", conf)
}
