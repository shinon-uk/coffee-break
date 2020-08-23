# coffee-break

## 構築手順
- クローンする
  - `git clone git@github.com:shinon-uk/coffee-break.git`
- ディレクトリに移動する
  - `cd coffee-break`
- コンテナを立ち上げる
  - `docker-compose up -d --build`
- パッケージをインストールし、webpackを実行する<br>
  - `npm install`
  - `npm run webpack`

- ブラウザで確認する<br>
  - http://localhost:10080/

- `webpack-dev-server`で確認する<br>
  - `npm run dev`
  - http://localhost:8000/
