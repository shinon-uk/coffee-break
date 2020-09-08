<template>
  <div id="app">
    <h1>dashboard</h1>
    <p>ここに一覧を表示する</p>
    <p>{{ this.contents }}</p>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "app",
  // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
  data() {
    return {
      // コンテンツ情報
      contents: ['test'],
    }
  },
  // インスタンス作成時の処理
  created: function () {
    this.fetchContents()
  },
  // メソッド定義
  methods: {
    // コンテンツ情報を取得する
    fetchContents() {
      axios.get('/api/getContents/')
          .then(response => {
            if (response.status !== 200) {
              throw new Error('レスポンスエラー')
            } else {
              // サーバから取得した情報をdataに設定する
              this.contents = response.data
            }
          })
    },
  }
};
</script>
