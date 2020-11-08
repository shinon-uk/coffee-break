<template>
  <div id="app" style="text-align: center">
    <h1>ユーザー一覧</h1>
    <p v-for="user in users">
      {{ user }}
    </p>
    <h1>ユーザーを追加</h1>
    <div v-if="errors.length">
      <p v-for="error in errors">
        {{ error }}
      </p>
    </div>
    <label>
      name：<input v-model="username" name="username" type="text"/><br>
      password：<input v-model="password" name="password" type="text"/><br>
      <input type="submit" value="submit" @click="this.addUser"/>
    </label>
  </div>
</template>

<script>
import axios from "axios"

export default {
  name: "app",
  // data オブジェクトのプロパティの値を変更すると、ビューが反応し、新しい値に一致するように更新
  data() {
    return {
      errors: [],
      users: [],
      username: "",
      password: "",
    }
  },
  // インスタンス作成時の処理
  created: function () {
    this.fetchAllUser()
  },
  // メソッド定義
  methods: {
    fetchAllUser() {
      axios.get('/api/fetchAllUser/')
        .then(response => {
          if (response.status !== 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.users = []
            response = response.data
            response.forEach(user => this.users.push(user));
          }
        })
    },
    addUser() {
      if (!this.checkForm()) {
        return false;
      }
      const formData = new FormData();
      formData.append("username", this.username);
      formData.append("password", this.password);
      axios.post(
        '/api/addUser/',
        formData
      )
        .then(response => {
          if (response.status !== 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.fetchAllUser()
          }
        })
    },
    checkForm: function () {
      if (this.username && this.password) {
        return true;
      }
      this.errors = [];
      if (!this.username) {
        this.errors.push('ユーザー名が入力されてません');
      }
      if (!this.password) {
        this.errors.push('パスワードが入力されてません');
      }
      return false;
    }
  }
};
</script>
