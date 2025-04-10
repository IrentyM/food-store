package dao

import (
	"order-service/internal/domain"
	"time"
)

type Order struct {
	ID            int       `db:"id"`
	UserID        int       `db:"user_id"`
	Status        string    `db:"status"`
	PaymentStatus string    `db:"payment_status"`
	TotalAmount   float64   `db:"total_amount"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func ToOrder(order domain.Order) Order {
	return Order{
		ID:            order.ID,
		UserID:        order.UserID,
		Status:        string(order.Status),
		PaymentStatus: string(order.PaymentStatus),
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}

func FromOrder(order Order) domain.Order {
	return domain.Order{
		ID:            order.ID,
		UserID:        order.UserID,
		Status:        domain.OrderStatus(order.Status),
		PaymentStatus: domain.PaymentStatus(order.PaymentStatus),
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}
