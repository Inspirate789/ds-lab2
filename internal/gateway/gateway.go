package gateway

import (
	"context"
	"github.com/Inspirate789/ds-lab2/internal/models"
	"github.com/Inspirate789/ds-lab2/internal/pkg/app"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/multierr"
	"log/slog"
	"math"
	"strconv"
	"time"
)

type CarsAPI interface {
	app.HealthChecker
	GetCars(ctx context.Context, offset, limit uint64, showAll bool) (res []models.Car, totalCount int64, err error)
	LockCar(ctx context.Context, carUID string) (found bool, err error)
	UnlockCar(ctx context.Context, carUID string) (found bool, err error)
}

type RentalsAPI interface {
	app.HealthChecker
	GetUserRentals(ctx context.Context, username string, offset, limit uint64) (res []models.Rental, totalCount int64, err error)
	GetUserRental(ctx context.Context, rentalUID, username string) (res models.Rental, found, permitted bool, err error)
	CreateRental(carUID, username string, dateFrom, dateTo time.Time) (res models.Rental, err error)
}

type PaymentsAPI interface {
	app.HealthChecker
	CreatePayment(ctx context.Context, price uint64) (res models.Payment, err error)
	CancelPayment(ctx context.Context, paymentUID string) (res models.Payment, found bool, err error)
}

type Gateway struct {
	carsAPI     CarsAPI
	rentalsAPI  RentalsAPI
	paymentsAPI PaymentsAPI
	logger      *slog.Logger
}

func New(carsAPI CarsAPI, rentalsAPI RentalsAPI, paymentsAPI PaymentsAPI, logger *slog.Logger) app.Delivery {
	return &Gateway{
		carsAPI:     carsAPI,
		rentalsAPI:  rentalsAPI,
		paymentsAPI: paymentsAPI,
		logger:      logger,
	}
}

func (gateway *Gateway) HealthCheck(ctx context.Context) error {
	return multierr.Combine(
		gateway.carsAPI.HealthCheck(ctx),
		gateway.rentalsAPI.HealthCheck(ctx),
		gateway.paymentsAPI.HealthCheck(ctx),
	)
}

func (gateway *Gateway) AddHandlers(router fiber.Router) {
	router.Get("/cars", gateway.getCars)
	router.Get("/rental", gateway.getRentals)
	router.Post("/rental", gateway.startCarRental)
	router.Get("/rental/:rentalUID", gateway.getRental)
	router.Post("/rental/:rentalUID/finish", gateway.finishCarRental)
	router.Delete("/rental/:rentalUID", gateway.cancelCarRental)
}

func (gateway *Gateway) getCars(ctx *fiber.Ctx) error {
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		gateway.logger.Debug("car list page not set, use default 0")
		page = 0
	}

	size, err := strconv.ParseUint(ctx.Query("size"), 10, 64)
	if err != nil {
		gateway.logger.Debug("car list size not set, return all cars")
		size = math.MaxInt64
	}

	showAll, err := strconv.ParseBool(ctx.Query("showAll"))
	if err != nil {
		gateway.logger.Debug("'showAll' for car list not set, return only available cars")
		showAll = false
	}

	cars, totalCount, err := gateway.carsAPI.GetCars(ctx.Context(), page*size, size, showAll)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(NewCarsDTO(cars, page*size, size, totalCount).Map())
}

func (gateway *Gateway) getRentals(ctx *fiber.Ctx) error {
	page, err := strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		gateway.logger.Debug("car list page not set, use default 0")
		page = 0
	}

	size, err := strconv.ParseUint(ctx.Query("size"), 10, 64)
	if err != nil {
		gateway.logger.Debug("car list size not set, return all cars")
		size = math.MaxInt64
	}

	username := ctx.Get("X-User-Name")

	rentals, totalCount, err := gateway.rentalsAPI.GetUserRentals(ctx.Context(), username, page*size, size)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(NewRentalsDTO(rentals, page*size, size, totalCount).Map())
}

func (gateway *Gateway) getRental(ctx *fiber.Ctx) error {
	rentalUID := ctx.Params("rentalUid")
	username := ctx.Get("X-User-Name")

	rental, found, permitted, err := gateway.rentalsAPI.GetUserRental(ctx.Context(), rentalUID, username)
	if err != nil {
		return err
	}

	if !found {
		return ctx.Status(fiber.StatusNotFound).JSON(errors.ErrRetalNotFound.Map())
	}

	if !permitted {
		return ctx.Status(fiber.StatusForbidden).JSON(errors.ErrRentalNotPermitted.Map())
	}

	return ctx.Status(fiber.StatusOK).JSON(NewRentalDTO(rental).Map())
}

func (gateway *Gateway) startCarRental(ctx *fiber.Ctx) error {

}

func (gateway *Gateway) cancelCarRental(ctx *fiber.Ctx) error {

}

func (gateway *Gateway) finishCarRental(ctx *fiber.Ctx) error {

}
