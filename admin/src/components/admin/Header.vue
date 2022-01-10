<template>
    <div class="headerBtn">
        <a-button type="danger" @click="loginOut">Logout</a-button>
    </div>
</template>

<script>
export default {
    methods: {
    loginOut () {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
