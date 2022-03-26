<template>
  <a-layout-sider class="footer" breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'My Blog' }}</span>
    </div>
    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index"
        ><a-icon type="dashboard"></a-icon><span>Dashboard</span></a-menu-item
      >
      <a-sub-menu>
        <span slot="title"><a-icon type="file" /><span>Article</span></span>
        <a-menu-item key="addart"
          ><a-icon type="form"></a-icon><span>Write article</span></a-menu-item
        >
        <a-menu-item key="artlist"
          ><a-icon type="snippets"></a-icon
          ><span>Article table</span></a-menu-item
        >
      </a-sub-menu>
      <a-menu-item key="catelist"
        ><a-icon type="book"></a-icon><span>Category table</span></a-menu-item
      >
      <a-menu-item key="userlist"
        ><a-icon type="user"></a-icon><span>User table</span></a-menu-item
      >
      <a-menu-item key="profile"
        ><a-icon type="setting"></a-icon
        ><span>Personal setting</span></a-menu-item
      >
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage (item) {
      this.$router.push('/' + item.key).catch(error => error)
    }
  }
}
</script>

<style scoped>
.footer {
  background-color: #eee;
  text-align: center;
}
.footer span {
  font-size: 18px;
}

.log {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
