package main

import (
	"context"
	pb "service/ecommerce"

	uuid "github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type pbService struct {
	pb.UnimplementedProductInfoServiceServer
	produceMap map[string]*pb.Product
}

func (s *pbService) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID: %v", err)
	}

	in.Id = out.String()
	if s.produceMap == nil {
		s.produceMap = make(map[string]*pb.Product)
	}
	s.produceMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *pbService) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.produceMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product with ID %s does not exist.", in.Value)
}
