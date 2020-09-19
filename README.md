# coffee-break

## 構築手順
- クローンする
  - `git clone git@github.com:shinon-uk/coffee-break.git`
- ディレクトリに移動する
  - `cd coffee-break`
- コンテナを立ち上げる
  - `docker-compose up -d --build`
- パッケージをインストールし、webpackを実行してブラウザを確認する<br>
  - `npm install`
  - `npm run webpack`
  - http://localhost:10080/
- `webpack-dev-server`を実行するとwebpackを実行しなくても自動で変更が反映される<br>
  - `npm run dev`
