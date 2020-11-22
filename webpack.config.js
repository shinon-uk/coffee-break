// ファイルおよびディレクトリパスを操作するためのユーティリティを提供する
// @see::https://nodejs.org/api/path.html
const path = require('path');

// プロジェクトルート
const projectRoot = path.resolve(__dirname);

// webpackのローダー。これが無いとwebpackが動かない。
// @see::https://vue-loader.vuejs.org/
const VueLoaderPlugin = require('vue-loader/lib/plugin');

// webpack-dev-serverプログラムにバンドルファイルをファイルシステムに強制的に書き込みます。
// これがないとmain.goで立ち上げたlocalhostに反映されない
// @see::https://www.npmjs.com/package/write-file-webpack-plugin
const WriteFileWebPackPlugin = require('write-file-webpack-plugin');

module.exports = {
  // モード値を production に設定すると最適化された状態で、development に設定するとソースマップ有効でJSファイルが出力される
  mode: 'development',
  // メインとなるJavaScriptファイル（エントリーポイント）
  entry: './static/entry/index.ts',
  // bundleファイルをwebpackがどこにどのような名前で出力すればいいのかを指定する
  output: {
    filename: '[name].js',
    path: path.join(projectRoot, 'static/dist')
  },
  // import文でファイル解決するため
  // これを定義しないと import 文で拡張子を書く必要が生まれる
  resolve: {
    // エイリアスを作成する
    alias: {
      'vue$': 'vue/dist/vue.esm.js',
    },
    // 解決順を指定する
    extensions: ['.js', '.vue', '.ts']
  },
  // webpack-dev-serverのオプションを選択する
  devServer: {
    // 使用するホストを指定する
    host: 'localhost',
    // リクエストをリッスンするポートを指定する
    port: '8080',
    // サーバーに提供するコンテンツを指定する
    contentBase: path.join(__dirname, "static/dist"),
  },
  // @see::https://webpack.js.org/configuration/devtool/
  devtool: "cheap-module-eval-source-map",
  // ローダーの設定
  module: {
    rules: [
      // VueLoaderPluginで必要な設定
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      },
      // TypeScriptに必要な設定
      {
        test: /\.ts$/,
        use: {
          loader: 'ts-loader',
          options: {
            appendTsSuffixTo: [/\.vue$/],
          },
        },
      },
    ]
  },
  // プラグインの設定
  plugins: [
    new VueLoaderPlugin(),
    new WriteFileWebPackPlugin(),
  ]
};
