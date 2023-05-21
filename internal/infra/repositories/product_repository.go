package repositories

import (
	"github.com/IcaroSilvaFK/brand_monitor_go_lang_test/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	conn *gorm.DB
}

type ProductRepositoryInterface interface {
	Create(p *entities.Product) error
	FindById(id string) (*entities.Product, error)
	Delete(id string) error
	FindAll(query string) ([]*entities.Product, error)
	Update(p *entities.Product, id string) error
}

func NewProductRepository(conn *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{
		conn: conn,
	}
}

func (pdb *ProductRepository) Create(u *entities.Product) error {
	return pdb.conn.Create(u).Error
}
func (pdb *ProductRepository) FindById(id string) (*entities.Product, error) {
	var product entities.Product

	tx := pdb.conn.Find(&product, "id = ?", id)

	return &product, tx.Error
}
func (pdb *ProductRepository) Delete(id string) error {

	return pdb.conn.Where("id = ?", id).Delete(&entities.Product{}).Error
}
func (pdb *ProductRepository) FindAll(query string) ([]*entities.Product, error) {

	var products []*entities.Product

	if err := pdb.conn.Where("name LIKE ?", "%"+query+"%").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (pdb *ProductRepository) Update(p *entities.Product, id string) error {
	return pdb.conn.Where("id = ?", id).Updates(p).Error
}
