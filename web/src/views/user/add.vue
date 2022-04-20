<template>
  <div class="app-container">
    <el-form ref="userRule" :model="form" :rules="rules" label-width="120px">
      <el-form-item label="姓名" prop="name">
        <el-input v-model="form.name"  placeholder="请填写姓名" />
      </el-form-item>
      <el-form-item label="性别">
        <el-col :span="10">
          <el-select v-model="form.sex" placeholder="请选择性别" @change="getChange">
            <el-option
              v-for="item in sexList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="10" >
          <el-form-item label="年龄" prop="age">
            <el-input v-model="form.age" placeholder="请填写年龄" />
          </el-form-item>
        </el-col>
      </el-form-item>
      <el-form-item label="状态">
        <el-switch
          v-model="form.status"
          :active-value="1"
          :inactive-vlaue="2"
          @change="onChange(form)"
          active-color="#13ce66"
          inactive-color="#ff4949">
        </el-switch>
      </el-form-item>
      <el-form-item label="类型">
        <el-radio-group v-model="form.role">
          <el-radio :label="1">admin</el-radio>
          <el-radio :label="2">user</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.introduction" type="textarea" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">提交</el-button>
        <el-button @click="onClear">清除</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>

  import { getDetail, enable, disable , save} from '@/api/demo'

export default {
  data() {
    return {
      sexList: [
        {value:1,label:'男',text:"男"},
        {value:2,label:'女',text:"女"},
      ],
      form: {
        id: 0,
        name: '',
        age: "",
        role: 1,
        introduction: '',
        status: "2"
      },
      rules: {
        name: [
          { required: true, message: '请输入用户名称', trigger: 'blur' },
          { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
        ],
        age: [
          { required: true, message: '请填写年龄', trigger: 'blur' },
          { type: 'number', message: '请填写数字', trigger: 'blur', transform: (value) => Number(value)},
        ],
      }
    }
  },
  created () {
    if(this.$route.query.id != null) {
      getDetail(this.$route.query.id).then(response => {
        let res = response.data.result
        this.form.id   = res.id
        this.form.name = res.name
        this.form.role = res.role
        this.form.introduction = res.introduction
        this.form.sex = res.sex
        this.form.status = res.status
      })
    }
  },
  methods: {
    getChange(){
      this.form.sex = ""
    },
    onChange(item) {
      if(item.id == 0){
        return
      }
      if(item.status){
        enable(item.id)
      } else {
        disable(item.id)
      }
    },
    onSubmit() {
      this.$refs["userRule"].validate((valid) => {
        if (valid) {
          save(this.form).then(response => {
            this.$message({
              message: '提交成功',
              type: 'success'
            });
            this.$router.push('/user/list')
          })
        } else {
          this.$message({
            message: '参数有误!',
            type: 'warning'
          });
          return false;
        }
      });

    },
    onClear() {
      this.form.name = ""
      this.form.role = 2
      this.form.age = ""
      this.form.introduction = ""
      this.form.sex = ""
      this.form.status = 2
      this.$message({
        message: '清除成功!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

