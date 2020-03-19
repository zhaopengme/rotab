<template>
  <div>
    <h3>修改密码</h3>
    <br />
    <Form ref="form" :model="form" :rules="rules" :label-width="120">
      <FormItem label="原密码" prop="originPassword">
        <i-input type="password" v-model="form.originPassword"></i-input>
      </FormItem>
      <FormItem label="新密码" prop="password">
        <i-input type="password" v-model="form.password"></i-input>
      </FormItem>
      <FormItem label="重复密码" prop="password2">
        <i-input type="password" v-model="form.password2"></i-input>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="submitHandler">更新</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
export default {
  data() {
    const validatePassCheck = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("Please enter your password again"));
      } else if (this.form.password !== this.form.password2) {
        callback(new Error("The two input passwords do not match!"));
      } else {
        callback();
      }
    };

    return {
      form: {
        originPassword: "",
        password: "",
        password2: ""
      },
      rules: {
        originPassword: [
          {
            required: true,
            message: "The password cannot be empty",
            trigger: "blur"
          }
        ],
        password: [
          {
            required: true,
            message: "The password cannot be empty",
            trigger: "blur"
          }
        ],
        password2: [
          {
            required: true,
            message: "The password cannot be empty",
            trigger: "blur"
          },
          { validator: validatePassCheck, trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    submitHandler() {
      this.$refs.form.validate(async valid => {
        if (valid) {
          const result = await this.$store.dispatch(
            "user/updatePassword",
            this.form
          );
          if (this.$kit.isBlank(result)) {
            this.$Message.success("更新成功!");
            this.$refs.form.resetFields();
          }
        }
      });
    }
  }
};
</script>

<style></style>
