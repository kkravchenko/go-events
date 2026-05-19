package models

import "time"

type Event struct {
    ID        uint      `gorm:"primaryKey"`
    DateTime  time.Time
    Event     string
    LogType   string
}