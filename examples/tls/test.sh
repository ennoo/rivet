#!/usr/bin/env bash

# 创建CA和申请证书
# ca:
# 生成CA自己的私钥 rootCA.key
openssl genrsa -out rootCA.key 2048
# 参数解释：
# -sha256 ：使用sha256算法
# rsa:2048 ：2048位密钥
# -days 365：证书有效期365天
# -subj /CN=*.abc.com ：使用abc.com域名的通配方式作为使用者，这个功能我也是摸索了好久才实现的。一般域证书都只绑定一个域名（二级域名也算独立的一个），使用通配方式，此证书可以加密a.abc.com，b.abc.com等多个子域名
# rootCA.key：密钥文件，自己命名
# rootCA.crt：iphone、mac等使用的证书文件
# 根据CA自己的私钥生成自签发的数字证书，该证书里包含CA自己的公钥
# 证书主题信息一般会显示SSL证书域名CN、单位名称O、部门名称OU、所在国家C、所在省份S和所在市县L。
# 而许多国内CA机构颁发的SSL证书的证书申请单位名称O字段居然是CA的名称。
# 按照X.509证书标准格式规定 SSL证书主题信息中的“O”字段只能写SSL证书申请单位的名称而绝对不能写成CA机构的名词，
# 这不仅不符合证书标准规范而且会给CA机构带来法律风险和给SSL证书申请单位带来品牌伤害和在线信任问题。
# 比如国内某CA机构颁发给某银行的SSL证书的O字段是此CA机构的名称而不是该银行的名称，按照X.509证书标准格式解释该银行属于此CA机构。
# 同样的问题也出现在CNNIC颁发给网易的SSL证书上O字段应该显示网易公司名称但居然显示“CNNIC SSL”这意味着此网站.help.163.com属于CNNIC而不属于网易公司。
# 不仅如此有些还把L字段写成地址S字段写成一串数字等等都是不符合数字证书X.509国际标准的。
# 国际机构颁发SSL证书的风险鉴于国内CA机构颁发的SSL证书不支持各种浏览器浏览器因无法识别而显示警告信息，
# 所以目前国内几乎所有重要的网上银行、网上证券、电子商务 网站和电子政务系统全部部署的是国外CA机构颁发的SSL证书这是有一定风险的。
openssl req -x509 -sha256 -nodes -new -key rootCA.key -subj "/C=CN/ST=Beijing/L=Beijing/O=vh/OU=bc/CN=localhost/emailAddress=aberic@qq.com" -days 5000 -out rootCA.crt
# server:
mkdir server
# 生成服务端私钥
openssl genrsa -out server/server.key 2048
# 生成 Certificate Sign Request，CSR，证书签名请求
openssl req -sha256 -new -key server/server.key -subj "/C=CN/ST=Beijing/L=Beijing/O=vh/OU=bc/CN=localhost/emailAddress=aberic@qq.com" -out server/server.csr
# 检查 csr 的正确性，通过下面的指令可以看到生成的证书的信息，查看是否是sha256RSA等
openssl req -in server/server.csr -text
# 自CA用自己的CA私钥对服务端提交的csr进行签名处理，得到服务端的数字证书 server.crt
openssl x509 -req -in server/server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -sha256 -out server/server.crt -days 5000
# 再次检查 crt 的正确性
openssl x509 -in server/server.crt -text
# client:
mkdir client
openssl genrsa -out client/client.key 2048
openssl req -sha256 -new -key client/client.key -subj "/C=CN/ST=Beijing/L=Beijing/O=vh/OU=bc/CN=localhost/emailAddress=aberic@qq.com" -out client/client.csr
openssl x509 -req -in client/client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -sha256 -out client/client.crt -days 5000
# golang client:
# golang tls 要校验 ExtKeyUsage，需要在生成client.crt时指定extKeyUsage
# 创建文件client.ext并写入内容extendedKeyUsage=clientAuth
echo "extendedKeyUsage=clientAuth" > client/client.ext
openssl x509 -req -in client/client.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -extfile client/client.ext -sha256 -out client/client.crt -days 5000
# 查看crt内容
openssl x509 -text -in client/client.crt -noout