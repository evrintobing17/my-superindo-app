package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/cart"
	"github.com/evrintobing17/my-superindo-app/internal/repository"
)

type cartRepo struct {
	db *repository.Database
}

func NewCartRepository(db *repository.Database) cart.CartRepository {
	return &cartRepo{db: db}
}

// IsCartExists implements cart.CartRepository.
func (c *cartRepo) IsCartExists(ctx context.Context, userID int) bool {
	var count int64
	query := "SELECT COUNT(*) FROM myapp.mapp_cart_user mcu WHERE mcu.user_id = $1"
	err := c.db.Conn.QueryRow(query, userID).Scan(&count)
	fmt.Println(count)
	if err != nil {
		return false
	}
	if count > 0 {
		return false
	}
	return true
}

func (c *cartRepo) IsProductExists(ctx context.Context, cartID int) bool {
	var count int64
	query := "SELECT COUNT(*) FROM myapp.mapp_cart_product mcu WHERE mcu.cart_id = $1"
	err := c.db.Conn.QueryRow(query, cartID).Scan(&count)
	fmt.Println(count)
	if err != nil {
		return false
	}
	if count > 0 {
		return false
	}
	return true
}

func (c *cartRepo) Insert(ctx context.Context, userID int, request models.AddToCardRequest) error {
	var id int
	tx, err := c.db.Conn.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
		return err
	}

	err = tx.QueryRow("INSERT INTO myapp.mapp_cart_user (user_id) VALUES ($1) returning id", userID).Scan(&id)
	if err != nil {
		tx.Rollback()
		log.Fatal("Failed to prepare statement:", err)

		return err
	}

	query := `INSERT INTO myapp.mapp_cart_product (total, cart_id, product_id) VALUES ($1, $2, $3)`

	_, err = tx.Exec(query, request.Total, id, request.ProductID)
	if err != nil {
		log.Fatal("Failed to upsert cart:", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Failed to commit transaction:", err)
		return err
	}

	return nil

}

func (c *cartRepo) Upsert(ctx context.Context, userID int, request models.AddToCardRequest) error {
	var cartid int
	var id int
	err := c.db.Conn.QueryRow("SELECT mcu.id AS cartid FROM myapp.mapp_cart_user mcu WHERE mcu.user_id = $1 AND is_active IS TRUE", userID).Scan(&cartid)
	if err != nil {
		log.Fatal("Failed to Find cart_id :", err)
		return err
	}

	c.db.Conn.QueryRow("SELECT id FROM myapp.mapp_cart_product mcp WHERE mcp.cart_id = $1 AND mcp.product_id = $2 AND is_active IS TRUE", cartid, request.ProductID).Scan(&id)

	query := `UPDATE myapp.mapp_cart_product SET total = $1 WHERE cart_id = $2 AND product_id = $3 AND is_active IS TRUE`

	if id == 0 {
		query = `INSERT INTO myapp.mapp_cart_product (total, cart_id, product_id) VALUES ($1, $2, $3)`
	}
	id = cartid
	_, err = c.db.Conn.Exec(query, request.Total, id, request.ProductID)
	if err != nil {
		log.Fatal("Failed to upsert cart:", err)
		return err
	}

	return nil

}
