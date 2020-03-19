<template>
  <section class="hero is-center is-fullheight">
    <div class="hero-body">
      <div class="container ">
        <div style="width:350px">
          <Form
            ref="loginForm"
            :model="form"
            :rules="rules"
            @keydown.enter.native="handleSubmit"
          >
            <FormItem prop="username">
              <i-input v-model="form.username" placeholder="请输入用户名">
                <span slot="prepend">
                  <Icon :size="16" type="ios-person"></Icon>
                </span>
              </i-input>
            </FormItem>
            <FormItem prop="password">
              <i-input
                type="password"
                v-model="form.password"
                placeholder="请输入密码"
              >
                <span slot="prepend">
                  <Icon :size="14" type="md-lock"></Icon>
                </span>
              </i-input>
            </FormItem>
            <FormItem>
              <Button @click="handleSubmit" type="primary" long>登录</Button>
            </FormItem>
          </Form>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "LoginForm",
  props: {
    usernameRules: {
      type: Array,
      default: () => {
        return [{ required: true, message: "账号不能为空", trigger: "blur" }];
      }
    },
    passwordRules: {
      type: Array,
      default: () => {
        return [{ required: true, message: "密码不能为空", trigger: "blur" }];
      }
    }
  },
  data() {
    return {
      form: {
        username: "",
        password: ""
      }
    };
  },
  computed: {
    rules() {
      return {
        username: this.usernameRules,
        password: this.passwordRules
      };
    }
  },
  mounted() {
    localStorage.clear();
  },
  methods: {
    handleSubmit() {
      this.$refs.loginForm.validate(async valid => {
        if (valid) {
          const result = await this.$store.dispatch("user/login", this.form);
          if (result.token) {
            await this.$store.dispatch("user/getUserinfo");
            this.$router.push("/home");
          }
        }
      });
    }
  }
};
</script>
