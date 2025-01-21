### アプリの起動
- `go run cmd/app/main.go`: アプリの起動
### DBのスキーマからデータベースアクセス用のモデルの自動生成
- `xo schema "postgres://postgres:postgres@localhost:5432/oss_matching?sslmode=disable" -o internal/model`：
### ヘルスチェック
- `curl http://localhost:8080/health`
### swaggerのドキュメント生成コマンド
- `swag init -g cmd/app/main.go --parseDependency`
### swaggerのドキュメントアクセスURL
- `http://localhost:8080/swagger/index.html`