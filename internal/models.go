package internal

import "time"

type Room struct {
	ID        int64     `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Question struct {
	ID        int64     `db:"id" json:"id"`
	RoomID    string    `db:"room_id" json:"room_id"`
	Content   string    `db:"content" json:"content"`
	Votes     int64     `db:"votes" json:"votes"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type ConfusionEvent struct {
	ID        int64     `db:"id" json:"id"`
	RoomID    int64     `db:"room_id" json:"room_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
