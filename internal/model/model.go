package model

import "time"

type id struct {
}

type source struct {
	url string
}

type user struct {
	id               int64
	userName         string
	createTime       time.Time
	beginOfTheCourse time.Time
	endOfTheCourse   time.Time
	endChannel       time.Time
}
