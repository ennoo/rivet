#!/usr/bin/env bash

# 创建CA和申请证书
# ca:
# 生成CA自己的私钥 rootCA.key
openssl genrsa -out rootCA.key 2048
# 根据CA自己的私钥生成自签发的数字证书，该证书里包含CA自己的公钥
openssl req -x509 -new -nodes -key rootCA.key -subj "/CN=localhost" -days 5000 -out rootCA.crt
# server:
mkdir server
# 生成服务端私钥
openssl genrsa -out server/server.key 2048
# 生成 Certificate Sign Request，CSR，证书签名请求
openssl req -new -key server/server.key -subj "/CN=localhost" -out server/server.csr
# 自CA用自己的CA私钥对服务端提交的csr进行签名处理，得到服务端的数字证书 server.crt
openssl x509 -req -in server/server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out server/server.crt -days 5000
# client:
mkdir client
openssl genrsa -out client/client.key 2048
openssl req -new -key client/client.key -subj "/CN=localhost" -out client/client.csr
openssl x509 -req -in client/client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out client/client.crt -days 5000
# golang client:
# golang tls 要校验 ExtKeyUsage，需要在生成client.crt时指定extKeyUsage
# 创建文件client.ext并写入内容extendedKeyUsage=clientAuth
echo "extendedKeyUsage=clientAuth" > client/client.ext
openssl x509 -req -in client/client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -extfile client/client.ext -out client/client.crt -days 5000
# 查看crt内容
openssl x509 -text -in client/client.crt -noout