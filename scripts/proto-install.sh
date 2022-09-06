#!/bin/bash
PROTO_DIR=$1
PROTO_OUT=$2

NODE_SCRIPT='./node_modules/.bin/proto-loader-gen-types --longs=String  --enums=String --defaults --oneofs --grpcLib=@grpc/grpc-js'
GO_SCRIPT=''
#* generate proto types
# for proto in $(ls -d ${PROTO_DIR}/*/); do
#     ${NODE_SCRIPT} --outDir=${proto%%/} ${proto%%/}/*.proto
# done
${NODE_SCRIPT} --outDir=${PROTO_DIR} ${PROTO_DIR}/*.proto

#* generate doc from proto files
# docker run --rm -v $PWD/_doc:/out -v $PWD/packages/services/_proto/product:/protos pseudomuto/protoc-gen-doc --doc_opt=markdown,docs.md

#* copy the generated ts files to the services
cp -R ${PROTO_DIR} ./packages/api-gateway/src/adapters/

# for genType in $(ls -d ${PROTO_DIR}/*/); do
#     OUT_DIR=${genType//$PROTO_DIR/}
#     cp -a ${genType}/. ./packages/services/svc-${OUT_DIR///}/src/_proto
# done
