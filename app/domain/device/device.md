# device & notification

## 概要

デバイストークン管理＋ Notification 送信 API（Go/Echo, DDD+CleanArch, AWS SQS+SNS, 非同期）
のハンズオン設計＋実装

## アーキテクチャの概要

```java
モバイルアプリ（別途作成）
      ↑↓
Echo API Server (Go)
      ↑↓
DynamoDBにデバイストークン保存
      ↑↓
SQSキュー ← Notificationをenqueue
      ↑↓
Worker (Go, 別プロセス)がSQSをポーリングし、SNSで配信
      ↑↓
AWS SNSがAPNs/FCMに通知
```
