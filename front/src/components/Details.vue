<template>
  <div class="container">
    <div class="d-flex justify-center pa-3 text-h3 font-weight-bold">
      {{ artInfo.title }}
      <div class="d-flex mx-10 justify-center">
        <v-icon class="mr-1" color="indigo" small>
          {{ 'mdi-calendar-month' }}
        </v-icon>
        <span>{{ artInfo.CreatedAt | dateformat('YYYY-MM-DD') }}</span>
      </div>
      <v-divider class="pa-3 ma-3"></v-divider>
      <v-alert
        class="ma-4"
        elevation="1"
        color="indigo"
        dark
        border="left"
        outlined
        >{{ artInfo.desc }}</v-alert
      >
      <div
        class="content ma-5 pa-3 text-justify"
        v-html="artInfo.content"
      ></div>
    </div>
  </div>
</template>
<script>
export default {
  props: ['id'],
  data () {
    return {
      artInfo: {},
      total: 0,
      queryParam: {
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getArtInfo()
  },
  methods: {
    async getArtInfo () {
      const { data: res } = await this.$http.get(`article/info/${this.id}`)
      this.artInfo = res.data
    }
  }
}
</script>
<style scoped>
.content >>> img,
span,
p {
  width: 100%;
}
</style>
