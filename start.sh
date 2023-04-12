go mod init github.com/a772304419/mysql-operator-demo1

kubebuilder init --domain mysql.zxl.com

go mod tidy

kubebuilder create api --group batch --version v1 --kind Mysql # y y

make install && make run 

kubectl apply -f config/samples/batch_v1_mysql.yaml

kubectl scale mysql mysql-sample --replicas 2

