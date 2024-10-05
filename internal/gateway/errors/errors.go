package errors

type GatewayError string

func (e GatewayError) Error() string {
	return string(e)
}

func (e GatewayError) Map() map[string]any {
	return map[string]any{"message": e.Error()}
}

const (
	ErrPaymentNotFound    GatewayError = "payment not found"
	ErrRentalNotFound     GatewayError = "rental not found"
	ErrRentalNotPermitted GatewayError = "rental not permitted"
)

func ErrInvalidRentalRequest(msg string) GatewayError {
	return GatewayError("invalid rental request: " + msg)
}

func ErrInvalidDateFrom(msg string) GatewayError {
	return GatewayError("invalid period start date: " + msg)
}

func ErrInvalidDateTo(msg string) GatewayError {
	return GatewayError("invalid period end date: " + msg)
}

func ErrInvalidRentalPeriod(dateFrom, dateTo string) GatewayError {
	return GatewayError("invalid period start date: [" + dateFrom + ", " + dateTo + "]")
}

func ErrRollbackWrap(err error) GatewayError {
	return GatewayError("rollback: " + err.Error())
}
