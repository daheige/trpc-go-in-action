# 生成trpc-go代码
gen-pb:
	trpc create \
		--protocol=grpc \
		-p proto/helloworld.proto \
		-o pb -f \
		--rpconly \
		--nogomod=true \
		--mock=false

.PHONY: gen-pb
