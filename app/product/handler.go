package main

import (
	"context"
	product "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/product"
	"github.com/Group-lifelong-youth-training/mygomall/app/product/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// CreateProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) CreateProducts(ctx context.Context, req *product.CreateProductsReq) (resp *product.CreateProductsResp, err error) {
	resp, err = service.NewCreateProductsService(ctx).Run(req)

	return resp, err
}

// ReducrProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ReducrProducts(ctx context.Context, req *product.ReduceProductsReq) (resp *product.ReduceProductsResp, err error) {
	resp, err = service.NewReducrProductsService(ctx).Run(req)

	return resp, err
}

// ReduceProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ReduceProducts(ctx context.Context, req *product.ReduceProductsReq) (resp *product.ReduceProductsResp, err error) {
	resp, err = service.NewReduceProductsService(ctx).Run(req)

	return resp, err
}

// UpdateProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProducts(ctx context.Context, req *product.UpdateProductsReq) (resp *product.UpdateProductsResp, err error) {
	resp, err = service.NewUpdateProductsService(ctx).Run(req)

	return resp, err
}
