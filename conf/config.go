package conf

import (
	"io"
	"os"
	"path"
)

type Config struct {
	Languages []Language
	Default   string
}

type Language struct {
	Name string
	Path string
}

func (lang Language) Ext() string {
	return path.Ext(lang.Path)
}

func (lang Language) CreateFile(newPath string) error {
	oldFile, err := os.Open(lang.Path)
	if err != nil {
		return err
	}
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	io.Copy(newFile, oldFile)
	return nil
}
