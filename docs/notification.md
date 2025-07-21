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

## 🎯 改善ポイント（Next step）

- Worker の並列処理化
- リトライロジックの実装
- CloudWatch Logs 連携
- モバイルアプリとの統合

## API リクエスト例

```bash
curl -X POST http://localhost:8080/device \
  -H "Content-Type: application/json" \
  -d '{"user_id":"u1","title":"Hello","message":"Test"}'
```
