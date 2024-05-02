test:
	go test -tags=test -v -count=1 ./app/...

run_query:
	go build -tags=build -o query ./cmd/query.go
	chmod +x ./query
	./query

run_parse:
	go build -tags=build -o parse ./cmd/parse.go
	chmod +x ./parse
	./run_parse.sh
