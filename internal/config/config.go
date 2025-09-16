package config

import (
    "api-monitor/internal/checker"
    "gopkg.in/yaml.v3"
    "io/ioutil"
)

func LoadConfig(path string) checker.Config {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    var cfg checker.Config
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        panic(err)
    }

    return cfg
}