package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/ThatGuyKev/ecommerce-services/product/db"
	"github.com/ThatGuyKev/ecommerce-services/product/models"
	pb "github.com/ThatGuyKev/ecommerce-services/product/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) GetProductDetails(ctx context.Context, req *pb.GetProductDetailsReq) (*pb.GetProductDetailsRes, error) {
	fmt.Println("Getting Product Details for: ", req.ProductId)

	var productDetails models.ProductDetails

	s.H.DB.Raw("SELECT id, active_start_date, title, attributes, short_description, long_description, uri, default_price, sale_price FROM product WHERE id = ?", req.ProductId).Scan(&productDetails)

	log.Printf("product details: %s", productDetails.ActiveStartDate)

	return &pb.GetProductDetailsRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetProductsByIds(ctx context.Context, req *pb.GetProductsByIdsReq) (*pb.GetProductsByIdsRes, error) {
	return &pb.GetProductsByIdsRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetCategoryDetails(ctx context.Context, req *pb.GetCategoryDetailsReq) (*pb.GetCategoryDetailsRes, error) {
	return &pb.GetCategoryDetailsRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetCategoriesByIds(ctx context.Context, req *pb.GetCategoriesByIdsReq) (*pb.GetCategoriesByIdsRes, error) {
	return &pb.GetCategoriesByIdsRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetProductAssetsByProductId(ctx context.Context, req *pb.GetProductAssetsByProductIdReq) (*pb.GetProductAssetsByProductIdRes, error) {
	return &pb.GetProductAssetsByProductIdRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetProductAssetById(ctx context.Context, req *pb.GetProductAssetByIdReq) (*pb.GetProductAssetByIdRes, error) {
	return &pb.GetProductAssetByIdRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetProductVariants(ctx context.Context, req *pb.GetProductVariantsReq) (*pb.GetProductVariantsRes, error) {
	return &pb.GetProductVariantsRes{
		Code:    "1",
		Message: "success",
	}, nil
}

func (s *Server) GetProductCategories(ctx context.Context, req *pb.GetProductCategoriesReq) (*pb.GetProductCategoriesRes, error) {
	return &pb.GetProductCategoriesRes{
		Code:    "1",
		Message: "success"}, nil
}
