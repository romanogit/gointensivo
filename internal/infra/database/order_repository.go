package database

import (
	"github.com/romanogit/gointensivo/internal/entity"

	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository {
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("Insert into Orders (id, price, tax, final_price) values (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	return err
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int

	err := r.Db.QueryRow("select count(1) from orders").Scan(&total)

	return total, err
}

