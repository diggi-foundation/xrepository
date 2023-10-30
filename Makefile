pull:
	go run tools/pull/main.go 
generate:
	swagger generate client --skip-validation -A xrepository -f spec/swagger.json
serve:
	swagger serve spec/swagger.json
