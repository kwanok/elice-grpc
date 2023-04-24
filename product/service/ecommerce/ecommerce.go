package service

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	productMap map[string]*Product
}

func NewServer() ProductInfoServer {
	return &server{}
}

func (s *server) AddProduct(ctx context.Context, product *Product) (*ProductID, error) {
	id := uuid.New().String()

	product.Id = id
	if s.productMap == nil {
		s.productMap = make(map[string]*Product)
	}

	s.productMap[id] = product

	return &ProductID{Value: id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, id *ProductID) (*Product, error) {
	if s.productMap == nil {
		return nil, status.New(codes.NotFound, "product not found").Err()
	}

	product, ok := s.productMap[id.Value]

	if !ok {
		return nil, status.New(codes.NotFound, "product not found").Err()
	}

	return product, status.New(codes.OK, "").Err()
}
