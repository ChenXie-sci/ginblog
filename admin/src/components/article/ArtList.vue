<template>
  <div class="container">
    <h3>Welcome to Article list page.</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.title"
            placeholder="Input Article name"
            @search="getArtList"
            allowClear
            enter-button
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="$router.push(`/addart`)"
            >Add Art</a-button
          >
        </a-col>

        <a-col :span="6" :offset="4">
          <a-select
            defaultValue="Please select category"
            style="width: 200px"
            @change="handleTableChange"
          >
            <a-select-option
              v-for="item in Catelist"
              :key="item.id"
              :value="item.id"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-col>
      </a-row>
    </a-card>
    <a-table
      rowKey="ID"
      :columns="columns"
      :pagination="pagination"
      :dataSource="Artlist"
      bordered
      @change="CateChange"
    >
      <span slot="img" class="img" slot-scope="img">
        <img :src="img" />
      </span>
      <template slot="action" slot-scope="data">
        <div class="actionSlot">
          <a-button
            size="small"
            type="primary"
            icon="edit"
            style="margin-right: 15px"
            @click="$router.push(`/addart/${data.ID}`)"
            >Edit</a-button
          >
          <a-button
            size="small"
            type="danger"
            icon="delete"
            style="margin-right: 15px"
            @click="deleteArt(data.ID)"
            >Delete</a-button
          >
        </div>
      </template>
    </a-table>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Category name',
    dataIndex: 'Category.name',
    width: '10%',
    key: 'name',
    align: 'center'
  },
  {
    title: 'Article title',
    dataIndex: 'title',
    width: '10%',
    key: 'title',
    align: 'center'
  },
  {
    title: 'Article Description',
    dataIndex: 'img',
    width: '20%',
    key: 'img',
    align: 'center'
  },
  {
    title: 'Article image',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center',
    scopeSlots: { customRender: 'img' }
  },
  {
    title: 'Action',
    width: '15%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  data () {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: total => `Total ${total} Record`
      },
      Artlist: [],
      Catelist: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getArtList()
    this.getCateList()
  },
  methods: {
    // get Artlist
    async getArtList () {
      const { data: res } = await this.$http.get('article', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      // console.log(res)
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Artlist = res.data
      this.pagination.total = res.total
    },
    // get Catelist
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Catelist = res.data
      this.pagination.total = res.total
    },

    // change pagetable
    handleTableChange (pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pagesize = pagination.pagesize
      this.queryParam.pagesize = pagination.pagesize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArtList()
    },

    // delete Art
    deleteArt (id) {
      this.$confirm({
        title: 'Sure to delete this Art?',
        content: 'Destroy all can not recover',
        onOk: async () => {
          const res = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('Success deleted')
          this.getArtList()
        },
        onCancel: () => {
          this.$message.info('Already concel delete.')
        }
      })
    },
    // add new Art
    addArtOk () {
      this.$refs.addArtRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.post('article/add', {
          Artname: this.newArt.Artname,
          password: this.newArt.password,
          role: this.newArt.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('Add Art successfully.')
        this.addArtVisible = false
        this.getArtList()
      })
    },
    addArtCancel () {
      this.$refs.addArtRef.resetFields()
      this.addArtVisible = false
    },
    adminChange (value) {
      this.ArtInfo.role = Number(value)
      console.log(this.ArtInfo.role)
    },
    // edit Art
    async editArt (id) {
      this.editArtVisible = true
      const { data: res } = await this.$http.put(`article/${id}`)
      this.ArtInfo = res.data
      this.ArtInfo.id = id
    },
    editArtOk () {
      this.$refs.addArtRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.put(
          `article/${this.ArtInfo.id}`,
          {
            Artname: this.ArtInfo.Artname,
            role: this.ArtInfo.role
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('Edit Art successfully.')
        this.EditArtVisible = false
        this.getArtList()
      })
    },
    editArtCancel () {
      this.$refs.addArtRef.resetFields()
      this.editArtVisible = false
      this.$message.info('Edit canceled')
    },
    CateChange (value) {
      this.getCateArt(value)
    },
    async getCateArt (id) {
      const { data: res } = await this.$http.get(`article/list/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.Artlist = res.data
      this.pagination.total = res.total
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

.ArtImg {
  height: 100%;
  width: 100%;
}
.ArtImg img {
  width: 100px;
  height: 80px;
}
</style>
