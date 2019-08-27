
#generate root cert
openssl req -x509 -config root.conf -nodes -newkey rsa:2048 -keyout root.key -out root.crt -days 365

#generate intermidiate csr
# openssl req -config intermidiate.conf -nodes -new -newkey rsa:2048 -out intermidiate.csr -keyout intermidiate.key

#Generate intermidiate crt
# openssl x509 -req -days 365 -in intermidiate.csr -extfile intermidiate.conf -CA root.crt -CAkey root.key -CAcreateserial -out intermidiate.crt -extensions v3_req

#Generate end csr
openssl req -config user.conf -nodes -new -newkey rsa:2048 -out user.csr -keyout user.key

#Generate end crt
openssl x509 -req -days 365 -in user.csr -extfile user.conf -CA root.crt -CAkey root.key -CAcreateserial -out user.crt -extensions v3_req

#Generate end csr
openssl req -config pks-superuser.conf -nodes -new -newkey rsa:2048 -out pks-superuser.csr -keyout pks-superuser.key

#Generate end crt
openssl x509 -req -days 365 -in pks-superuser.csr -extfile pks-superuser.conf -CA root.crt -CAkey root.key -CAcreateserial -out pks-superuser.crt -extensions v3_req

#Generate end csr
openssl req -config pao-opsmgr-01.vcf.pvd.pez.pivotal.io.conf -nodes -new -newkey rsa:2048 -out pao-opsmgr-01.vcf.pvd.pez.pivotal.io.csr -keyout pao-opsmgr-01.vcf.pvd.pez.pivotal.io.key

#Generate end crt
openssl x509 -req -days 365 -in pao-opsmgr-01.vcf.pvd.pez.pivotal.io.csr -extfile pao-opsmgr-01.vcf.pvd.pez.pivotal.io.conf -CA root.crt -CAkey root.key -CAcreateserial -out pao-opsmgr-01.vcf.pvd.pez.pivotal.io.crt -extensions v3_req

#Generate end csr
openssl req -config api.pks.vcf.pvd.pez.pivotal.io.conf -nodes -new -newkey rsa:2048 -out api.pks.vcf.pvd.pez.pivotal.io.csr -keyout api.pks.vcf.pvd.pez.pivotal.io.key

#Generate end crt
openssl x509 -req -days 365 -in api.pks.vcf.pvd.pez.pivotal.io.csr -extfile api.pks.vcf.pvd.pez.pivotal.io.conf -CA root.crt -CAkey root.key -CAcreateserial -out api.pks.vcf.pvd.pez.pivotal.io.crt -extensions v3_req

