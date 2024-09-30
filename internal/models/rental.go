package models

import "time"

type RentalStatus string

const (
	RentalInProgress RentalStatus = "IN_PROGRESS"
	RentalFinished   RentalStatus = "FINISHED"
	RentalCanceled   RentalStatus = "CANCELED"
)

type Rental struct {
	ID         int64
	RentalUID  string
	Username   string
	PaymentUID string
	CarUID     string
	DateFrom   time.Time
	DateTo     time.Time
	Status     RentalStatus
}
