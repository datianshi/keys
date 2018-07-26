
#generate root cert
openssl req -x509 -config root.conf -nodes -newkey rsa:2048 -keyout root.key -out root.crt -days 365

#generate intermidiate csr
 openssl req -config intermidiate.conf -nodes -new -newkey rsa:2048 -out intermidiate.csr -keyout intermidiate.key

#Generate intermidiate crt
 openssl x509 -req -days 365 -in intermidiate.csr -extfile intermidiate.conf -CA root.crt -CAkey root.key -CAcreateserial -out intermidiate.crt -extensions v3_req

#Generate end csr
openssl req -config user.conf -nodes -new -newkey rsa:2048 -out user.csr -keyout user.key

#Generate end crt
openssl x509 -req -days 365 -in user.csr -extfile user.conf -CA intermidiate.crt -CAkey intermidiate.key -CAcreateserial -out user.crt -extensions v3_req
