package utils

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidatePlannedDates(plannedDateBegin, plannedDateEnd time.Time) error {
	now := time.Now().UTC()
	if plannedDateBegin.Before(now) {
		return status.Errorf(codes.InvalidArgument, "La fecha y hora de planificacion inicial no puede ser inferior a la fecha y hora actual")
	}

	if now.After(plannedDateEnd) {
		return status.Errorf(codes.InvalidArgument, "La fecha y hora de planificacion final no puede ser inferior a la fecha y hora actual")
	}

	if plannedDateEnd.Before(plannedDateBegin) {
		return status.Errorf(codes.InvalidArgument, "La fecha y hora de planificacion final no puede ser inferior a la planificacion inicial")
	}

	if plannedDateEnd.Sub(plannedDateBegin).Hours() >= 2 {
		return status.Errorf(codes.InvalidArgument, "La planificacion no puede ser mayor a 2 horas")
	}

	return nil
}
