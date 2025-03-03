package service

import (
	"context"
	"testing"
	product "github.com/Group-lifelong-youth-training/mygomall/rpc_gen/kitex_gen/product"
)

func TestReduceProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewReduceProductsService(ctx)
	// init req and assert value

	req := &product.ReduceProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
