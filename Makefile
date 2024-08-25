test:
	go test -v ./...

cover:
	go test -cover ./... -coverprofile=cover.out.tmp
    # 自動生成コードをカバレッジ対象から外し、カバレッジファイルを作成
	cat cover.out.tmp | grep -v "**_mock.go" | grep -v "wire_gen.go" > cover.out
	rm cover.out.tmp
	go tool cover -html=cover.out -o cover.html
	open cover.html
