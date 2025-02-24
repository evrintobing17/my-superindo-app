package repository

import (
	"context"
	"log"

	"github.com/evrintobing17/my-superindo-app/internal/models"
	"github.com/evrintobing17/my-superindo-app/internal/module/product"
	"github.com/evrintobing17/my-superindo-app/internal/repository"
)

type productRepo struct {
	db *repository.Database
}

func NewProductRepository(db *repository.Database) product.ProductRespository {
	return &productRepo{db: db}
}

// GetList implements product.ProductRespository.
func (p *productRepo) GetList(ctx context.Context, categoryID *int) (models.GetListProductResp, error) {
	var product []models.Product
	var resp models.GetListProductResp
	var args []interface{}
	query := "SELECT id, title, category_id, description FROM myapp.products WHERE 1=1"
	if categoryID != nil {
		query += " AND category_id = $1"
		args = append(args, categoryID)
	}
	rows, err := p.db.Conn.Query(query, args...)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}

	for rows.Next() {
		var (
			id          int
			title       string
			categoryID  int
			description string
		)

		err := rows.Scan(&id, &title, &categoryID, &description)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}

		product = append(product, models.Product{
			ID:          id,
			Title:       title,
			CategoryID:  categoryID,
			Description: description,
		})
	}
	resp.Product = product

	return resp, nil
}

func (p *productRepo) GetProductByProductID(ctx context.Context, id int) (models.Product, error) {
	var resp models.Product
	query := "SELECT id, title, category_id, description FROM myapp.products WHERE id = $1"
	err := p.db.Conn.QueryRow(query, id).Scan(&resp.ID, &resp.Title, &resp.CategoryID, &resp.Description)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	return resp, nil
}
