protoc-prepare:
	mkdir -p "internal/service/generated/$(NAME)/v1"

protoc-build:
	protoc -Iproto/v1 \
		--experimental_allow_proto3_optional \
		--go_out=internal/service/generated/$(NAME)/v1	\
		--go_opt paths=source_relative	\
		--go-grpc_out=internal/service/generated/$(NAME)/v1	\
		--go-grpc_opt paths=source_relative	\
		$(NAME).proto

# При сборке proto файлов использовать эту команду, где после "-" идет
# название вашего proto файла, который вы хотите сбилдить
proto-build:
	make protoc-prepare \
		protoc-build