<template>
  <div class="container">
    <a-card>
      <h3>{{ id ? 'Edit article' : 'Add article' }}</h3>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
        <a-form-model-item label="Article title" prop="title">
          <a-input style="width: 300px" v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <a-form-model-item label="Article category" prop="cid">
          <a-select
            style="width: 200px"
            v-model="artInfo.cid"
            placeholder="Please select category"
            @change="cateChange"
          >
            <a-select-option
              v-for="item in CateList"
              :key="item.id"
              :value="item.id"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="Article description" prop="desc">
          <a-input type="textarea" v-model="artInfo.desc"></a-input>
        </a-form-model-item>
        <a-form-model-item label="Article image" prop="img">
          <a-upload
            name="file"
            :action="upUrl"
            :headers="headers"
            @change="upChange"
            listType="picture"
            :defaultFileList="fileList"
          >
            <a-button> <a-icon type="upload" />Click to Upload </a-button>
            <template v-if="id">
              <img
                :src="artInfo.img"
                style="width:120px; height:100px; margin-left: 15px"
              />
            </template>
          </a-upload>
        </a-form-model-item>
        <a-form-model-item label="Article content" prop="content">
          <Editor v-model="artInfo.content"></Editor>
        </a-form-model-item>
        <a-form-model-item>
          <a-button
            type="danger"
            style="margin:15px"
            @click="artOk(artInfo.id)"
            >{{ artInfo.id ? 'Update' : 'Submit' }}</a-button
          >
          <a-button type="primary" @click="addCancel">Cancel</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
import { Url } from '../../plugin/http'
import Editor from '../../editor/index.vue'
export default {
  components: { Editor },
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        desc: '',
        content: '',
        img: ''
      },
      CateList: [],
      upUrl: Url + 'upload',
      headers: {},
      fileList: [],
      artInfoRules: {
        title: [
          { required: true, message: 'Please input title', trigger: 'blur' }
        ],
        cid: [
          {
            required: true,
            message: 'Please input article category',
            trigger: 'blur'
          }
        ],
        desc: [
          {
            required: true,
            message: 'Please input article category',
            trigger: 'change'
          },
          { max: 120, message: 'description max 120 char.', trigger: 'change' }
        ],
        img: [
          {
            required: true,
            message: 'Please input article img',
            trigger: 'blur'
          }
        ],
        content: [
          {
            required: true,
            message: 'Please input article content',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  mounted () {
    this.getCateList()
    this.headers = {
      Authorization: `Bearer ${window.sessionStorage.getItem('token')}`
    }
    if (this.id) {
      this.getArtInfo(this.id)
    }
  },
  methods: {
    async getArtInfo (id) {
      const { data: res } = await this.$http.get(`article/info/${id}`)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    },
    // get Catelist
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.CateList = res.data
    },
    // select category
    cateChange (value) {
      this.artInfo.cid = value
    },
    // upload images
    upChange (info) {
      if (info.file.status !== 'uploading') {
        // console.log(info.file, info.fileList)
      }
      if (info.file.status === 'done') {
        this.$message.success(`${info.file.name} file uploaded successfully`)
        const imgUrl = info.file.response.url
        this.artInfo.img = imgUrl
      } else if (info.file.status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    // upload && update success
    artOk (id) {
      this.$refs.artInfoRef.validate(async valid => {
        if (!valid) return this.$message.error('Parameter doesnot work')
        if (id === 0) {
          const { data: res } = await this.$http.post(
            'article/add',
            this.artInfo
          )
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$router.push('/artlist')
          this.$message.success('Add article success.')
        } else {
          const { data: res } = await this.$http.put(
            `article/${id}`,
            this.artInfo
          )
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$router.push('/artlist')
          this.$message.success('Update article successfully')
        }
      })
    },
    addCancel () {
      this.$refs.artInfoRef.resetFields()
    }
  }
}
</script>
