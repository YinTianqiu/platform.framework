protos=$(ls *.proto)
for proto in ${protos}
do  
    echo ${proto}
    protoc --proto_path=$GOPATH/src:. --micro_out=$GOPATH/src/  --go_out=$GOPATH/src/ ${proto}
    result=$?
    if [ ${result} -ne 0 ]; then
         echo "generate fail ${proto} > ${result}"
         exit ${result}
    fi
done