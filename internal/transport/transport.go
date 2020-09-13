package transport

import (
  "fmt"
)

type Config struct {
  name        string
  localPath   string
  remotePath  string
}

func NewConfig(name string) *Config {
  c := Config{
    name: name,
    localPath: "~/.config/nvim",
  }

  return &c
}

func GetConfig(configName string) (config string) {
  fmt.Println(config)

  return config
}

