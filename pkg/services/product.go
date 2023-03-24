package services

import (
	"context"
	"fmt"
	"net/http"
	"product_svc/pkg/db"
	"product_svc/pkg/models"
	"product_svc/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product models.Product

	if result := s.H.DB.Where("name = ? ", req.Name).First(&product); result.RowsAffected != 0 {
		fmt.Println(req.Name)
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  "Product is exist",
		}, nil
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Total = req.Total

	s.H.DB.Create(&product)

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var product models.Product

	result := s.H.DB.First(&product, req.Id)
	if result.RowsAffected == 0 {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  "Not found product",
		}, nil
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data: &pb.FindOneData{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
			Total: product.Total,
		},
	}, nil
}
