package main

import (
	"time"
)

type Fruit struct {
	Id         int64
	Name       string
	Origin     string
	Created_at time.Time
	Color      string
}
