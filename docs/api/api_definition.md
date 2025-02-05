# 認証基盤 API エンドポイント定義書

## 基本情報

- **ベース URL**: `https://{domain}/{stage name}/api/v1`
- **認証ヘッダー**: `Authorization: Bearer {idToken}`
- **Content-Type**: `application/json`

## エンドポイント一覧

### 1. サインアップ関連

#### 1.1 メールアドレス重複チェック

- **エンドポイント**: `POST /auth/check_email`
- **認証**: 不要
- **リクエスト**:

```json
{
  "email": "string"
}
```

- **レスポンス**:

```json
{
  "status": "string",
  "message": "string"
}
```

- **説明**:
  - メールアドレスの重複をチェックします
  - 既存ユーザーのメールアドレスと照合します

#### 1.2 ユーザー登録

- **エンドポイント**: `POST /auth/signup`
- **認証**: 不要
- **リクエスト**:

```json
{
  "email": "string",
  "password": "string",
  "userAttributes": {
    "name": "string",
    "custom:attribute": "string"
  }
}
```

- **レスポンス**:

```json
{
  "userId": "string",
  "message": "string"
}
```

- **説明**:
  - ユーザー情報を Cognito に登録します
  - 確認コードがメールで送信されます

#### 1.3 メール確認コード検証

- **エンドポイント**: `POST /auth/verify_email`
- **認証**: 不要
- **リクエスト**:

```json
{
  "userId": "string",
  "email": "string",
  "code": "string"
}
```

- **レスポンス**:

```json
{
  "status": "string",
  "message": "string"
}
```

- **説明**:
  - メールで送信された確認コードを検証します
  - 成功時にユーザーアカウントが有効化されます

### 2. サインイン関連

#### 2.1 認証初期化（SRP）

- **エンドポイント**: `POST /auth/initiate`
- **認証**: 不要
- **リクエスト**:

```json
{
  "username": "string",
  "SRP_A": "string"
}
```

- **レスポンス**:

```json
{
  "challengeName": "string",
  "challengeParameters": {
    "USER_ID_FOR_SRP": "string",
    "SRP_B": "string",
    "SALT": "string",
    "SECRET_BLOCK": "string"
  }
}
```

- **説明**:
  - SRP 認証の初期フェーズを実行します
  - チャレンジパラメータを返却します

#### 2.2 認証応答（SRP）

- **エンドポイント**: `POST /auth/respond_challenge`
- **認証**: 不要
- **リクエスト**:

```json
{
  "username": "string",
  "challengeResponses": {
    "PASSWORD_CLAIM": "string",
    "TIMESTAMP": "string",
    "SECRET_BLOCK": "string"
  }
}
```

- **レスポンス**:

```json
{
  "accessToken": "string",
  "idToken": "string",
  "refreshToken": "string",
  "expiresIn": "number"
}
```

- **説明**:
  - SRP 認証の応答フェーズを実行します
  - 認証成功時にトークン一式を返却します

### 3. トークン管理

#### 3.1 トークンリフレッシュ

- **エンドポイント**: `POST /auth/refresh`
- **認証**: 不要
- **リクエスト**:

```json
{
  "refreshToken": "string"
}
```

- **レスポンス**:

```json
{
  "accessToken": "string",
  "idToken": "string",
  "expiresIn": "number"
}
```

- **説明**:
  - リフレッシュトークンを使用して新しいトークンを取得します
  - アクセストークンと ID トークンが更新されます

#### 3.2 トークン無効化（サインアウト）

- **エンドポイント**: `POST /auth/signout`
- **認証**: 必要
- **リクエスト**:

```json
{
  "refreshToken": "string"
}
```

- **レスポンス**:

```json
{
  "status": "string",
  "message": "string"
}
```

- **説明**:
  - 現在のセッションを終了します
  - リフレッシュトークンを無効化します

## エラーレスポンス

全エンドポイントで共通のエラーレスポンス形式:

```json
{
  "error": {
    "status": "string",
    "code": "string",
    "message": "string",
    "details": "object (optional)"
  }
}
```

### 主要エラーコード

- `AUTH001`: 無効な認証情報
- `AUTH002`: 期限切れトークン
- `AUTH003`: アカウント未確認
- `AUTH004`: レート制限超過
- `AUTH005`: システムエラー
- `AUTH006`: メールアドレス重複

## レート制限

各エンドポイントの制限値:

- サインアップ関連: 10 req/min
- サインイン試行: 5 req/min
- トークンリフレッシュ: 30 req/min
- その他エンドポイント: 60 req/min

## セキュリティ要件

1. **認証トークン**

   - ID トークン有効期限: 1 時間
   - リフレッシュトークン有効期限: 30 日

2. **パスワードポリシー**
   - 最小長: 8 文字
   - 必須要素: 大文字、小文字、数字、特殊文字
