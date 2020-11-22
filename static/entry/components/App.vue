<template>
  <div id="app" style="text-align: center">
    <router-link to="/page1">ページ１へ</router-link>
    <router-link to="/page2">ページ２へ</router-link>
    <router-view/>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';

@Component
export default class App extends Vue {
}
</script>
