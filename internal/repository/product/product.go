package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/piresc/product-app/internal/entity"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type ProductRepo struct {
	cfg *entity.Config
	db  *sqlx.DB
}

func NewProductRepository(
	cfg *entity.Config,
) *ProductRepo {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.SSLMode)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	log.Println("Connected to database")
	return &ProductRepo{
		cfg: cfg,
		db:  db,
	}
}

// Create inserts a new product into the database
func (r *ProductRepo) Create(product *entity.Product) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(`INSERT INTO products (name, description, price, variety, rating, stock, created_at, updated_at) 
        VALUES (:name, :description, :price, :variety, :rating, :stock, :created_at, :updated_at)`, product)
	return err
}

// GetAll retrieves all products from the database
func (r *ProductRepo) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Select(&products, "SELECT id, name, description, price, variety, rating, stock, created_at, updated_at FROM products")
	return products, err
}

// GetByID retrieves a product by its ID
func (r *ProductRepo) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Get(&product, "SELECT id, name, description, price, variety, rating, stock, created_at, updated_at FROM products WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}
	return &product, nil
}

// GetByVariant retrieves products by their variant
func (r *ProductRepo) GetByVariant(variant string) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Select(&products, "SELECT id, name, description, price, variety, rating, stock, created_at, updated_at FROM products WHERE variety = $1", variant)
	return products, err
}

// GetByTitle retrieves products by their title
func (r *ProductRepo) GetByTitle(title string) ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Select(&products, "SELECT id, name, description, price, variety, rating, stock, created_at, updated_at FROM products WHERE name ILIKE $1", "%"+title+"%")
	return products, err
}

// Update modifies an existing product in the database
func (r *ProductRepo) Update(product *entity.Product) error {
	product.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(`UPDATE products SET name = :name, description = :description, 
		price = :price, variety = :variety, rating = :rating, stock = :stock, 
		updated_at = :updated_at 
		WHERE id = :id`, product)
	return err
}

// Delete removes a product from the database by its ID
func (r *ProductRepo) Delete(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Get(&product, "SELECT id, name, description, price, variety, rating, stock, created_at, updated_at FROM products WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}
	_, err = r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error deleting product with id %d", id)
	}
	return &product, nil
}
