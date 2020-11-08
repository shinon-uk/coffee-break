# coffee-break

## 構築手順
- クローンする
  - `git clone git@github.com:shinon-uk/coffee-break.git`
- ディレクトリに移動する
  - `cd coffee-break`
- 環境構築
  - `make setup`
- DBマイグレート
  - `bee migrate`
- 自動で変更を反映する
  - `npm run dev`
- ブラウザ表示
  - `make open`
  - http://localhost:10080/
