echo change ulimit to unlimited
ulimit -c unlimited

echo rm scada build file
rm ./scada

echo download dependency library...
go get -d -v ./...

echo scada build...
go build scada

echo scada start
./scada

echo bash...
bash
