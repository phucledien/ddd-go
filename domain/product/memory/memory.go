package memory

import (
	"sync"

	"github.com/google/uuid"

	"github.com/phucledien/tavern/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod
	return nil
}

func (mpr *MemoryProductRepository) Update(upprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(mpr.products, id)
	return nil
}
