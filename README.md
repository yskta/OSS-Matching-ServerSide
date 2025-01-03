### アプリの起動
- `go run cmd/app/main.go`: アプリの起動
### DBのスキーマからデータベースアクセス用のモデルの自動生成
- `xo schema "postgres://postgres:postgres@localhost:5432/oss_matching?sslmode=disable" -o internal/model`：