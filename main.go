package main

import (
	"github.com/abbyck/periodicbw/commands"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	log.Info("Started periodicbw")
	// runner.Runner()
	// resp := commands.Execute(os.Args[1:])

	// if resp.Err != nil {
	// 	log.Error("Error occured")
	// 	os.Exit(-1)
	// }
	commands.Execute()
}
