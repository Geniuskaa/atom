protoc-prepare:
	mkdir -p "internal/service/generated/$(NAME)/v1"

protoc-build:
	protoc \
		-I proto/car_v1 \
		-I proto/validate/ \
		--experimental_allow_proto3_optional \
		--go_out=internal/service/generated/$(NAME)/v1	\
		--go_opt paths=source_relative	\
		--go-grpc_out=internal/service/generated/$(NAME)/v1	\
		--go-grpc_opt paths=source_relative	\
		--validate_opt lang=go,paths=source_relative \
		--validate_out=internal/service/generated/$(NAME)/v1 \
		$(NAME).proto

# При сборке proto файлов использовать эту команду, где после "-" идет
# название вашего proto файла, который вы хотите сбилдить
proto-build:
	make protoc-prepare \
		protoc-build