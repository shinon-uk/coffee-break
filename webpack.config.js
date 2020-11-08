// ファイルおよびディレクトリパスを操作するためのユーティリティを提供する
// @see::https://nodejs.org/api/path.html
const path = require('path');

// プロジェクトルート
const projectRoot = path.resolve(__dirname);

// webpackバンドルを提供するHTMLファイルの作成を簡素化するwebpackプラグイン
// これが無いとhtmlファイルへの変更が反映されない
// @see::https://webpack.js.org/plugins/html-webpack-plugin/
// @see::https://github.com/jantimon/html-webpack-plugin
const HtmlWebpackPlugin = require('html-webpack-plugin');

// webpackのローダー。これが無いとwebpackが動かない。
// @see::https://vue-loader.vuejs.org/
const VueLoaderPlugin = require('vue-loader/lib/plugin');

// webpack-dev-serverプログラムにバンドルファイルをファイルシステムに強制的に書き込みます。
// これがないとmain.goで立ち上げたlocalhostに反映されない
// @see::https://www.npmjs.com/package/write-file-webpack-plugin
const WriteFileWebPackPlugin = require('write-file-webpack-plugin');

module.exports = {
    // エントリーポイントを指定する
    entry: './static/js/index.js',
    // bundleファイルをwebpackがどこにどのような名前で出力すればいいのかを指定する
    output: {
        filename: '[name].js',
        path: path.join(projectRoot, 'dist')
    },
    // モジュールの解決方法を指定する
    resolve: {
        // エイリアスを作成する
        alias: {
            'vue$': 'vue/dist/vue.esm.js',
        },
        // 解決順を指定する
        extensions: ['.js', '.vue']
    },
    // webpack-dev-serverのオプションを選択する
    devServer: {
        // 使用するホストを指定する
        host: 'localhost',
        // リクエストをリッスンするポートを指定する
        port: '8080',
        // サーバーに提供するコンテンツを指定する
        contentBase: path.join(__dirname, "dist"),
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
        ]
    },
    // プラグインの設定
    plugins: [
        new VueLoaderPlugin(),
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, 'static/index.html')
        }),
        new WriteFileWebPackPlugin(),
    ]
};
