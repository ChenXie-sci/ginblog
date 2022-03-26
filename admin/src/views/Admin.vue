<template>
  <div class="container">
    <a-layout>
      <a-layout-sider>
        <Nav />
      </a-layout-sider>
      <a-layout>
        <a-layout-header>
          <Header />
        </a-layout-header>
        <a-layout-content>
          <router-view :key="$route.path"></router-view>
        </a-layout-content>
        <a-layout-footer>
          <Footer />
        </a-layout-footer>
      </a-layout>
    </a-layout>
  </div>
</template>

<script>
import Nav from '../components/admin/Nav.vue'
import Footer from '../components/admin/Footer.vue'
import Header from '../components/admin/Header.vue'

export default {
  components: { Nav, Footer, Header }
}
</script>

<style scoped>
.container {
  height: 700px;
}
</style>
