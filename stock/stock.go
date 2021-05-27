package stock

type Product struct {
	Name string
	IdProduct string
	Price int
	TypeProduct string
}

type Stock struct {
	Name string
	Products []Product
}

func (s Stock) GetProduct(idProduct string) Product {
	for i := 0; i < len(s.Products); i++ {
		if s.Products[i].IdProduct == idProduct {
			return s.Products[i]
		}
	}
	return Product{}
}

func (s Stock) RemoveProduct(idProduct string) {
	for i := 0; i < len(s.Products); i++ {
		if s.Products[i].IdProduct == idProduct {
			s.Products[i] = s.Products[len(s.Products) - 1]
			s.Products[len(s.Products) - 1] = Product{}
			s.Products = s.Products[:len(s.Products) - 1]
		}
	}
}

func (s Stock) AddProduct(product Product) {
	s.Products = append(s.Products, product)
}