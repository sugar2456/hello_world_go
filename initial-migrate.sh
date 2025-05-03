#!/bin/bash

# デフォルト値を設定
DB_HOST=${DB_HOST:-"db"}
DB_PORT=${DB_PORT:-"3306"}
DB_USER=${DB_USER:-"gouser"}
DB_PASSWORD=${DB_PASSWORD:-"gopassword"}
DB_NAME=${DB_NAME:-"godb"}

# データベースURL
DB_URL="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?parseTime=true"
MIGRATION_NAME=${1:-"initial"}

echo "==== Entスキーマからマイグレーションを生成・適用 ===="

# マイグレーションディレクトリを作成
mkdir -p migrations

# Entスキーマの使用とURLエンコーディングのための処理
# "ent://" URLスキームはAtlasがEntスキーマを直接読み込むために使用
atlas migrate diff ${MIGRATION_NAME} \
  --dir file://migrations \
  --to "ent://ent/schema?adapter=mysql" \
  --dev-url "${DB_URL}"

echo "マイグレーション '${MIGRATION_NAME}' を作成しました。内容を確認しています..."
cat migrations/*_${MIGRATION_NAME}.sql

echo "マイグレーションを適用しますか？ (y/n)"
read -r response
if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
  echo "マイグレーションを適用しています..."
  atlas migrate apply \
    --dir file://migrations \
    --url "${DB_URL}"
  
  echo "マイグレーション状態を確認しています..."
  atlas migrate status \
    --dir file://migrations \
    --url "${DB_URL}"
else
  echo "マイグレーションの適用をスキップしました。"
fi