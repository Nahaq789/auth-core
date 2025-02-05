# DynamoDB テーブル定義

## users テーブル

| 論理名         | 物理名     | データ型 | キー | GSI     | NULL 許可 | 説明                     |
| -------------- | ---------- | -------- | ---- | ------- | --------- | ------------------------ |
| ユーザー ID    | user_id    | string   | PK   | -       | NO        | Cognito Sub を使用       |
| メールアドレス | email      | string   | -    | GSI1-PK | NO        | ユーザーのメールアドレス |
| 名前           | name       | string   | -    | -       | NO        | ユーザーの表示名         |
| ユーザータイプ | user_type  | string   | -    | -       | NO        | ユーザー種別             |
| 作成日時       | created_at | string   | -    | -       | NO        | レコード作成日時         |
| 更新日時       | updated_at | string   | -    | -       | NO        | レコード更新日時         |

### GSI（グローバルセカンダリインデックス）

1. GSI1（メールアドレスによる検索用）
   - PK: email
   - 属性: すべて

### 補足

- user_id には Cognito の sub を使用することで、Cognito との一貫性を保ちます
- created_at, updated_at を追加し、監査証跡を確保します
- メールアドレスでの検索を効率的に行うため、GSI を設定します
