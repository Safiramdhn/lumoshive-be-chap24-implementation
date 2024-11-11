//go:build wireinject
// +build wireinject

package wire

import (
	"golang-beginner-chap24/database"
	"golang-beginner-chap24/handlers"
	"golang-beginner-chap24/repositories"
	"golang-beginner-chap24/services"

	"github.com/google/wire"
)

var orderHandlerSet = wire.NewSet(
	database.NewPostgresDB,
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
