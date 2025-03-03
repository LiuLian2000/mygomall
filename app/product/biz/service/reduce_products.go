package service

import (
	"context"
	product "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/product"
)

type ReduceProductsService struct {
	ctx context.Context
} // NewReduceProductsService new ReduceProductsService
func NewReduceProductsService(ctx context.Context) *ReduceProductsService {
	return &ReduceProductsService{ctx: ctx}
}

// Run create note info
func (s *ReduceProductsService) Run(req *product.ReduceProductsReq) (resp *product.ReduceProductsResp, err error) {
	// Finish your business logic.

	return
}
