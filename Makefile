MODULE=github.com/blade-master/push
.PHONY: gen-pushcore
gen-push-core:
	@echo "Generating client files..."
	@cd  rpc_gen &&  cwgo client --type RPC  --service pushcore --module  ${MODULE}/rpc_gen  -I ../idl --idl ../idl/push.proto
	@echo "Client files generated successfully."
	@echo "Generating server files... pass means use define client rpc"
	@cd  app/pushcore &&  cwgo server --type RPC  --service pushcore --module ${MODULE}/app/pushcore  --pass "-use github.com/blade-master/push/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/push.proto
	@echo " Server files generated successfully."