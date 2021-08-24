package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
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

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to find present working directory:", err)
		os.Exit(1)
	}

	languageNames := make([]string, len(conf.Languages))
	for i, lang := range conf.Languages {
		languageNames[i] = lang.Name
	}

	var chosenLangIdx int
	selectPrompt := &survey.Select{
		Message: "Choose template language",
		Options: languageNames,
		Default: conf.Default,
	}
	if err = survey.AskOne(selectPrompt, &chosenLangIdx); err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		os.Exit(0)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	chosenLang := conf.Languages[chosenLangIdx]

	nameStr := ""
	namePrompt := &survey.Input{
		Message: "Enter filenames to create (space-separated)",
	}
	if err = survey.AskOne(namePrompt, &nameStr); err == terminal.InterruptErr {
		fmt.Println("Interrupted")
		os.Exit(0)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	names := strings.Split(strings.TrimSpace(nameStr), " ")

	ext := chosenLang.Ext()
	for _, name := range names {
		nameWithExt := name + ext
		fullPath := filepath.Join(pwd, nameWithExt)
		chosenLang.CreateFile(fullPath)
	}
}
