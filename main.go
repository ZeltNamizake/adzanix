package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Config struct {
	Timezone    string   `json:"timezone"`
	AdzanNormal string   `json:"adzan_normal"`
	AdzanSubuh  string   `json:"adzan_subuh"`
	PrayTimes   map[string]string `json:"pray_times"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	file, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	return config, err
}

func getCurrentTime(tz string) (string, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return "", err
	}

	return time.Now().In(loc).Format("15:04"), nil
}

func playAdzan(file string) {
	exec.Command("mpv", file, "--no-video", "--quiet").Run()
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	lastPlayed := ""

	for {

		now, err := getCurrentTime(config.Timezone)
		if err != nil {
			fmt.Println(err)
			time.Sleep(60 * time.Second)
			continue
		}

		if now != lastPlayed {
			for name, t := range config.PrayTimes {
				if now == t {
					if name == "fajr" {
						playAdzan(config.AdzanSubuh)
					} else {
						playAdzan(config.AdzanNormal)
					}
					lastPlayed = now
					break
				}
			}
		}
		time.Sleep(30 * time.Second)
	}
}
