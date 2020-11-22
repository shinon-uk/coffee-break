<template>
  <div>
    <h1>ユーザー一覧</h1>
    <p v-for="user in users">
      {{ user }}
    </p>
    <h1>ユーザーを追加</h1>
    <div v-if="errors != null && errors.length > 0" style="color: #ff0000">
      <p v-for="error in errors">
        {{ error }}
      </p>
    </div>
    <label>
      name：<input v-model="username" name="username" type="text"/><br>
      password：<input v-model="password" name="password" type="text"/><br>
      <input type="submit" value="submit" @click="addUser"/>
    </label>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import axios from "axios";

@Component
export default class Page1 extends Vue {
  private users: string[] = [];
  private errors: string[] = [];
  private username = '';
  private password: string = '';

  private created(): void {
    this.fetchAllUser()
  }

  private async fetchAllUser(): Promise<void> {
    const response: any = await axios.get('/api/fetchAllUser/')
      .catch((err: Error) => {
        throw new Error(err.message);
      });
    this.users = []
    response.data.forEach((user: any) => this.users.push(user));
  }

  private async addUser(): Promise<void> {
    if (!this.checkForm()) {
      return;
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
  }

  private checkForm(): boolean {
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
</script>
