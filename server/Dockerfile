# ビルドステージ
FROM golang:1.23.2 AS builder

WORKDIR /app

# Go Modulesを有効にし、依存関係をキャッシュ
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコンテナにコピー
COPY . .

# Goアプリケーションをビルド
RUN go build -o main .

# 実行ステージ
FROM golang:1.23.2

WORKDIR /app

# ビルドした実行ファイルをコピー
COPY --from=builder /app/main .

# 実行コマンドを設定
CMD ["./main"]
