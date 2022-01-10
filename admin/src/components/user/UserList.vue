<template>
  <div class="container">
    <h3>Welcome to user list page.</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="input username"
            @search="getUserList"
            allowClear
            enter-button
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUserVisible = true"
            >Add user</a-button
          >
        </a-col>
      </a-row>
    </a-card>
    <a-table
      rowKey="username"
      :columns="columns"
      :pagination="pagination"
      :dataSource="userlist"
      bordered
      @change="handleTableChange"
    >
      <span slot="role" slot-scope="role">{{
        role === 1 ? 'admin' : 'user'
      }}</span>
      <template slot="action" slot-scope="data">
        <div class="actionSlot">
          <a-button
            type="primary"
            icon="edit"
            style="margin-right: 15px"
            @click="editUser(data.ID)"
            >Edit</a-button
          >
          <a-button
            type="danger"
            icon="delete"
            style="margin-right: 15px"
            @click="deleteUser(data.ID)"
            >Delete</a-button
          >
          <a-button type="info" icon="info">Change password</a-button>
        </div>
      </template>
    </a-table>
    <!-- add new user table-->
    <a-modal
      title="Add new User"
      :visible="addUserVisible"
      width="60%"
      @ok="addUserOk"
      @cancel="addUserCancel"
      destoryOnclose
    >
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="Username" prop="username">
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="Password" prop="password">
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item
          has-feedback
          label="Confirm Password"
          prop="checkPass"
        >
          <a-input-password v-model="newUser.checkPass"></a-input-password>
        </a-form-model-item>
        <a-form-model-item label="Is admin?" prop="role">
          <a-select defaultValue="2" style="120px" @change="adminChange">
            <a-select-option key="1" value="1">Yes</a-select-option>
            <a-select-option key="2" value="2">No</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- edit user -->
    <a-modal
      title="Edit User"
      :visible="editUserVisible"
      width="60%"
      @ok="editUserOk"
      @cancel="editUserCancel"
      destoryOnclose
    >
      <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
        <a-form-model-item label="Username" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="Is admin?" prop="role">
          <a-select defaultValue="2" style="120px" @change="adminChange">
            <a-select-option key="1" value="1">Yes</a-select-option>
            <a-select-option key="2" value="2">No</a-select-option>
          </a-select>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
import day from 'dayjs'

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Username',
    dataIndex: 'username',
    width: '10%',
    key: 'username',
    align: 'center'
  },
  {
    title: 'Rigister Time',
    dataIndex: 'CreatedAt',
    width: '20%',
    align: 'center',
    customRender: val => {
      return val ? day(val).format('YYYY-MM-DD HH:mm') : 'nodata'
    }
  },
  {
    title: 'Role',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
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
        showTotal: total => `Total ${total} Record`
      },
      userlist: [],
      userInfo: {
        username: '',
        password: '',
        role: 2,
        checkPass: ''
      },
      newUser: {
        username: '',
        password: '',
        role: 2,
        checkPass: ''
      },
      changePassword: {
        id: 0,
        password: '',
        checkPass: ''
      },
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      editVisible: false,
      visible: false,
      addUserVisible: false,
      editUserVisible: false,
      userRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username === '') {
                callback(new Error('Please input username'))
              }
              if (
                [...this.userInfo.username].length < 4 ||
                [...this.userInfo.username].length > 12
              ) {
                callback(new Error('Username should be 4 to 12 length.'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('Please input password'))
              }
              if (
                [...this.userInfo.password].length < 6 ||
                [...this.userInfo.password].length > 20
              ) {
                callback(new Error('Password should be 6 to 20 length.'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkPass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('Please input password'))
              }
              if (this.userInfo.password !== this.userInfo.checkPass) {
                callback(new Error('Password are not the same'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addUserRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username === '') {
                callback(new Error('Please input username'))
              }
              if (
                [...this.newUser.username].length < 4 ||
                [...this.newUser.username].length > 12
              ) {
                callback(new Error('Username should be 4 to 12 length.'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password === '') {
                callback(new Error('Please input password'))
              }
              if (
                [...this.newUser.password].length < 6 ||
                [...this.newUser.password].length > 20
              ) {
                callback(new Error('Password should be 6 to 20 length.'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkPass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password === '') {
                callback(new Error('Please input password'))
              }
              if (this.newUser.password !== this.newUser.checkPass) {
                callback(new Error('Password are not the same'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created () {
    this.getUserList()
  },
  methods: {
    // get userlist
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.userlist = res.data
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
      this.getUserList()
    },

    // delete user
    deleteUser (id) {
      this.$confirm({
        title: 'Sure to delete this user?',
        content: 'Destroy all can not recover',
        onOk: async () => {
          const res = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) {
            return this.$message.error(res.message)
          }
          this.$message.success('Success deleted')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('Already concel delete.')
        }
      })
    },
    // add new user
    addUserOk () {
      this.$refs.addUserRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('Add user successfully.')
        this.addUserVisible = false
        this.getUserList()
      })
    },
    addUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
    },
    adminChange (value) {
      this.userInfo.role = Number(value)
      console.log(this.userInfo.role)
    },
    // edit user
    async editUser (id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserOk () {
      this.$refs.addUserRef.validate(async valid => {
        if (!valid) {
          return this.$message.error('Input wrong, please reinput.')
        }
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$message.success('Edit user successfully.')
        this.EditUserVisible = false
        this.getUserList()
      })
    },
    editUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.editUserVisible = false
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
