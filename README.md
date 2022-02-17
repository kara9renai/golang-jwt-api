# Practice in Go (JWTによる認証)
JWTによる認証を試すときは以下のコマンドで、.envファイルを用意してください。

```
cp .env.example .env
```

次にファイル内に下記のようにして

```
SIGNINGKEY="your_signingkey"
```
SIGNINGKEYを設定してください。(値は任意)

`curl localhost:8080/auth`を実行することで、JWTのトークンが発行されます。その後、JWTのトークンを用いて以下のコマンドを実行すると、認証に成功して期待した内容が得られます。

```
curl localhost:8080/private -H "Authorization:Bearer {jwt_token}"
```