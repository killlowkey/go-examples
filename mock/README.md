1. 安装 mock：`git get github.com/golang/mock`
2. 安装 mock-gen：`go install github.com/golang/mock/mockgen@latest`
3. 生成 mock 文件：`mockgen -source=interface.go -destination=mock_interface.go -package=mockExample`