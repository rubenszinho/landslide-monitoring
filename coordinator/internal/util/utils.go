package util

import "strings"

func SplitTopic(topic string) []string {
	return strings.Split(topic, "/")
}
