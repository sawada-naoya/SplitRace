# -----------------------------
# SplitRace 開発用 Makefile
# -----------------------------

# 使用例：make up

# 開発サーバー全体を起動（--build付き）
up:
	docker compose up --build

# 開発サーバーのみ再起動（ビルドなし）
restart:
	docker compose down && docker compose up

# コンテナ停止
down:
	docker compose down

# backendコンテナに入る
backend:
	docker compose exec backend sh

# frontendコンテナに入る
frontend:
	docker compose exec frontend sh

# backendのみ再ビルド
rebuild-backend:
	docker compose build backend

# frontendのみ再ビルド
rebuild-frontend:
	docker compose build frontend

# ログをリアルタイム表示
logs:
	docker compose logs -f
