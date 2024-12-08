package ports

import "github.com/hongchengp/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	 Charge (domain.order)
}