export ROOT_MOD=github.com/Group-lifelong-youth-training/mygomall

.PHONY: gen-%
gen-%:
	@cd app/$* && cwgo server --type RPC  --service $* --module  ${ROOT_MOD}/app/$*  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl  --idl ../../idl/$*.proto
	@cd rpc_gen && cwgo client --type RPC  --service $* --module  ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/$*.proto
