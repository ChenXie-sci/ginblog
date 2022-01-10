import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ArticleList from '../components/ArticleList.vue'
import Details from '../components/Details.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: Home,
    component: Home,
    children: [{ path: '/', component: ArticleList, meta: { title: 'Welcome to ginblog.' } },
    { path: 'detail/:id', component: Details, meta: { title: 'Article content' }, props: true }]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})
export default router
