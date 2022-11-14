package models

import "time"

type Review struct {
	id        int       `json:"id"`
	OrderId   int       `json:"order_id"`
	rate      int       `json:"rate"`
	desc      int       `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
}
