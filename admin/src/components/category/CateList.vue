<template>
  <div class="container">
    <h3>Welcome to Category list page.</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addCateVisible = true"
            >Add Cate</a-button
          >
        </a-col>
      </a-row>
    </a-card>

    <a-table
      rowKey="id"
      :columns="columns"
      :pagination="pagination"
      :dataSource="Catelist"
      bordered
      @change="handleTableChange"
    >
      <template slot="action" slot-scope="data">
        <div class="actionSlot">
          <a-button
            type="primary"
            icon="edit"
            style="margin-right: 15px"
            @click="editCate(data.id)"
            >Edit</a-button
          >
          <a-button
            type="danger"
            icon="delete"
            style="margin-right: 15px"
            @click="deleteCate(data.id)"
            >Delete</a-button
          >
        </div>
      </template>
    </a-table>
    <!-- add new Cate table-->
    <a-modal
      closable
      title="Add new Cate"
      :visible="addCateVisible"
      width="60%"
      @ok="addCateOk"
      @cancel="addCateCancel"
      destoryOnclose
    >
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item label="Catename" prop="name">
          <a-input v-model="newCate.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- edit Cate -->
    <a-modal
      title="Edit Cate"
      :visible="editCateVisible"
      width="60%"
      @ok="editCateOk"
      @cancel="editCateCancel"
      destoryOnclose
    >
      <a-form-model :model="CateInfo" :rules="CateRules" ref="addCateRef">
        <a-form-model-item label="Catename" prop="name">
          <a-input v-model="CateInfo.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Catename',
    dataIndex: 'name',
    width: '10%',
    key: 'name',
    align: 'center'
  },
  {
    title: 'Action',
    width: '30%',
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
        showTotal: total => `共${total}条`
      },
      Catelist: [],
      CateInfo: {
        name: '',
        id: 0
      },
      newCate: {
        name: ''
      },
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      editVisible: false,
      CateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.CateInfo.name === '') {
                callback(new Error('Please input Catename'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === '') {
                callback(new Error('Please input Catename'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editCateVisible: false,
      addCateVisible: false
    }
  },
  created () {
    this.getCateList()
  },
  methods: {
    // get Catelist
    async getCateList () {
      const { data: res } = await this.$http.get('category', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
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
      this.getCateList()
    },

    // delete Cate
    deleteCate (id) {
      this.$confirm({
        title: 'Sure to delete this Cate?',
        content: 'Destory all can not recover',
        onOk: async () => {
          const res = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('Success deleted')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('Already cancel delete.')
        }
      })
    },
    // add new Cate
    addCateOk () {
      this.$refs.addCateRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addCateRef.resetFields()
        this.$message.success('Add Cate successfully.')
        this.addCateVisible = false
        await this.getCateList()
      })
    },
    addCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('Edit canceled')
    },
    // edit Cate
    async editCate (id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.CateInfo = res.data
      this.CateInfo.id = id
    },
    editCateOk () {
      this.$refs.addCateRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.put(
          `category/${this.CateInfo.id}`,
          {
            name: this.CateInfo.name
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('Edit Cate successfully.')
        this.EditCateVisible = false
        this.getCateList()
      })
    },
    editCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('Edit canceled')
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
