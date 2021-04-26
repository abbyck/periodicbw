package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/abbyck/periodicbw/speedtest"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Freq string `yaml:"frequency"`
}

func Runner() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		dat, err := ioutil.ReadFile(".pbw-config.yml") // For read access.
		if err != nil {
			logrus.Info("No configuration file found, using the default value as test frequency")
		}
		config := Config{}
		err = yaml.Unmarshal([]byte(dat), &config)
		if err != nil {
			logrus.Warn("error, can not read the config file: %v", err)
		}
		fmt.Printf("Config:\n%v", config)

		c := cron.New()
		c.AddFunc(config.Freq, speedtest.Execute)
		go c.Start()
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
	}
}
