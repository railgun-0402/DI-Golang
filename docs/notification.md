## 概要

- モバイル通知システムを構成している
- 技術スタックは以下
  - 非同期メッセージング：AWS SQS
  - 通知配信：AWS SNS
  - 通知履歴保存：DynamoDB

## システム構成

```yaml
Client (curl, モバイルなど)
|
| HTTP POST /device
v
APIサーバー (cmd/server)
|
| SQSにメッセージ送信
v
AWS SQS
|
| メッセージ受信
v
Worker (cmd/worker)
|
| SNSにPublish
| DynamoDBに保存
```

## 🎯 各コンポーネントの役割

### ✅ API サーバー

- /device に POST されたリクエストを SQS に積むだけ
- リクエストの検証や永続化はしない

### ✅ SQS

- 非同期処理用のキュー
- メッセージを一時的に保持し、Worker が取り出す

### ✅ Worker

- SQS からメッセージを受信
- SNS に Publish
- DynamoDB に通知履歴を保存
- メッセージを SQS から削除

### ✅ SNS

- 1 対多の通知配信を担当
- メール・SMS・モバイルデバイスにメッセージを配信

### ✅ DynamoDB

- 通知履歴を永続化

## 🚀 実行方法

### API サーバーを起動

```bash
go run main.go --mode=server
```

### Worker を起動

```bash
go run main.go --mode=worker
```

## 🔍 動作確認フロー

### 1️⃣ SNS サブスクライバーの設定

- サブスクライバーとして登録(メルアドの場合)

```bash
aws sns subscribe \
  --topic-arn $SNS_TOPIC_ARN \
  --protocol email \
  --notification-endpoint youremail@example.com
```

- メールが届くので、確認リンクをクリックして「Confirmed」に！

```bash
aws sns list-subscriptions-by-topic --topic-arn $SNS_TOPIC_ARN
```

→SubscriptionArn が PendingConfirmation ではなく実際の ARN になっていれば OK！

### 2️⃣ API リクエスト例

```bash
curl -X POST http://localhost:8080/notifications \
  -H "Content-Type: application/json" \
  -d '{"user_id":"u1","title":"Hello","message":"Test"}'
```

- ✅ SQS にメッセージが積まれる
- ✅ Worker がメッセージを取り出す

### 3️⃣ Worker の動作確認

- Worker のターミナルに以下が出る
- SQS からメッセージを受信
- SNS に Publish したログ
- DynamoDB に保存したログ
- SQS からメッセージを削除したログ

### 4️⃣ メール受信確認

- メールサブスクライバーにメッセージが届く

## 🎯 改善ポイント（Next step）

- Worker の並列処理化
- DynamoDB の履歴一覧 API を作る(CRUD まで仕上げてしまってもよさそう)
- リトライロジックの実装
- CloudWatch Logs 連携
- モバイルアプリのデバイストークンを SNS サブスクライバーとして登録して Push 通知
