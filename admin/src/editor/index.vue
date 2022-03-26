<template>
  <div class="container">
    <Editor :init="init" v-model="content"></Editor>
  </div>
</template>

<script>
import Editor from '@tinymce/tinymce-vue'
import './tinymce.min.js'
import './icons/default/icons.min.js'
import './themes/silver/theme.min.js'

// import editor plugins

export default {
  components: { Editor },
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      init: {
        height: '500px',
        margin: '0',
        padding: '0',
        branding: false,
        images_upload_handler: async (blobInfo, succFun, failFun) => {
          const formdata = new FormData()
          formdata.append('file', blobInfo.blob(), blobInfo.name())
          const { data: res } = await this.$http.post('upload', formdata)
          succFun(res.url)
        }
      },
      content: this.value
    }
  },
  watch: {
    value (newV) {
      this.content = newV
    },
    content (newV) {
      this.$emit('input', newV)
    }
  }
}
</script>

<style scoped>
@import url('./skins/ui/oxide/skin.min.css');
</style>
