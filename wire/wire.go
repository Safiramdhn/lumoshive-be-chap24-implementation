//go:build wireinject
// +build wireinject

package wire

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/handlers"
	"golang-beginner-chap24/middleware"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"

	"github.com/google/wire"
)

var orderHandlerSet = wire.NewSet(
	database.NewPostgresDB,
	utils.IntiLogger,
	repositories.NewOrderRepository,
	services.NewOrderService,
	handlers.NewOrderHandler,
)

var paymentMethods = wire.NewSet(
	database.NewPostgresDB,
	repositories.NewPaymentMethodRepository,
	services.NewPaymentMethodService,
	handlers.NewPaymentMethodHandler,
)

func InitializeOrderHandler() *handlers.OrderHandler {
	wire.Build(orderHandlerSet)
	return nil
}

func InitializePaymentMethodHandler() *handlers.PaymentMethodHandler {
	wire.Build(paymentMethods)
	return nil
}

var middlewareSet = wire.NewSet(
	utils.IntiLogger,
	middleware.NewMiddleware,
)

func IniitalMiddleware() *middleware.Middleware {
	wire.Build(middlewareSet)
	return nil
}

var orderItemSet = wire.NewSet(
	database.NewPostgresDB,
	utils.IntiLogger,
	repositories.NewOrderItemRepo,
	services.NewOrderItemService,
	handlers.NewOrderItemHandler,
)

func InitializeOrderItemHandler() *handlers.OrderItemHandler {
	wire.Build(orderItemSet)
	return nil
}
