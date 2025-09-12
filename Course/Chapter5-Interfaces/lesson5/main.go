package main

import (
	"fmt"
)

func (e email) cost() int {
	length := len(e.body)
	if !e.isSubscribed {
		return length * 5
	} else {
		return length * 2
	}
}

func (e email) format() string {
	subscribed := ""
	if e.isSubscribed {
		subscribed = "Subscribed"
	} else {
		subscribed = "Not Subscribed"
	}
	return fmt.Sprintf(
		"'%s' | %s",
		e.body,
		subscribed,
	)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
