package model

import (
	"fmt"
	"push-notice/helper"
)

type YunyingEvent struct {
	Id         int    `db:"id"`
	Eventkey   string `db:"event_key"`
	Starttime  int    `db:"start_time"`
	Expiretime int    `db:"expire_time"`
	Extendinfo string `db:"extend_info"`
}

func GetEventList() ([]YunyingEvent, error) {
	database := helper.Db
	var list []YunyingEvent
	err := database.Select(&list, "select id,event_key,start_time,expire_time,extend_info from yunying_event")

	if err != nil {
		fmt.Println("fail", err)
		return nil, err
	}

	return list, nil

}
