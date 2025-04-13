// Copyright (C) 2025 Alisa Frelia
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ip-change-notifier/ipaddress"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type DiscordMessage struct {
	Content string `json:"content"`
}

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})

	if err := godotenv.Load(); err != nil {
		log.Warn("No .env file found. Proceeding without it.")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if level, err := logrus.ParseLevel(logLevel); err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL != "" {
		log.Hooks.Add(&DiscordHook{WebhookURL: webhookURL})
	}
}

type DiscordHook struct {
	WebhookURL string
}

func (hook *DiscordHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *DiscordHook) Fire(entry *logrus.Entry) error {
	message, err := entry.String()
	if err != nil {
		return err
	}

	return postToDiscord(message)
}

func postToDiscord(message string) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		return fmt.Errorf("DISCORD_WEBHOOK_URL is not set in the environment")
	}

	discordMessage := DiscordMessage{
		Content: message,
	}
	msgBytes, _ := json.Marshal(discordMessage)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return fmt.Errorf("unable to send message to Discord: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	c := cron.New()

	schedule := os.Getenv("SCHEDULE")
	if schedule == "" {
		schedule = "@every 1h"
	}
	log.Info("IP address check schedule is " + schedule)

	_, err := c.AddFunc(schedule, func() {
		log.Info("Running scheduled IP change detection")
		result, err := ipaddress.DetectChange()
		if err != nil {
			log.Error("Error detecting IP change: ", err)
			return
		}

		log.Info("Public IP fetched successfully: ", result.CurrentIP)

		if !result.Changed {
			log.Info("IP has not changed. No action required.")
			return
		}

		err = postToDiscord(fmt.Sprintf("Public IP is: %s", result.CurrentIP))
		if err != nil {
			log.Error("Failed to post a notification to Discord: ", err)
			return
		}

		log.Info("Message posted to Discord successfully.")
	})
	if err != nil {
		log.Fatal("Failed to schedule IP change detection: ", err)
	}

	c.Start()

	// Keep the application running
	select {}
}
