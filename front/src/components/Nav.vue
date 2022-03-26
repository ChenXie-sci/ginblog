<template>
  <v-card class="mx-auto">
    <v-img src="../assets/Navbg.jpg">
      <v-card-title>
        <v-col align="center">
          <v-avatar size="130" color="grey">
            <img src="../assets/avatar.jpg" alt="" />
          </v-avatar>
          <div class="ma-4 white--text">{{ profileInfo.name }}</div>
        </v-col>
      </v-card-title>
      <v-divider></v-divider>
    </v-img>

    <v-col>
      <div class="ma-3">About Me:</div>
      <div>{{ profileInfo.desc }}</div>
    </v-col>

    <v-divider color="indigo"></v-divider>
    <v-list nav dense>
      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="blue darken-2">{{ 'mdi-qqchat' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{
          profileInfo.qq_chat
        }}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon class="green darken-2">{{ 'mdi-wechat' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{
          profileInfo.wechat
        }}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="red">{{ 'mdi-sina-weibo' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{
          profileInfo.weibo
        }}</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="primary">{{ 'mdi-facebook' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text"
          >18342781298</v-list-item-content
        >
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon color="primary">{{ 'mdi-twitter' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">None</v-list-item-content>
      </v-list-item>

      <v-list-item>
        <v-list-item-icon class="ma-3">
          <v-icon class="primary">{{ 'mdi-email' }}</v-icon>
        </v-list-item-icon>
        <v-list-item-content class="grey--text">{{
          profileInfo.email
        }}</v-list-item-content>
      </v-list-item>
    </v-list>
  </v-card>
</template>

<script>
export default {
  data () {
    return {
      profileInfo: {
        id: 1
      }
    }
  },
  created () {
    this.getProfileInfo()
  },
  methods: {
    // 获取个人设置
    async getProfileInfo () {
      const { data: res } = await this.$http.get(
        `profile/${this.profileInfo.id}`
      )
      this.profileInfo = res.data
      this.$root.$emit('msg', res.data.icp_record)
    }
  }
}
</script>
<style></style>
