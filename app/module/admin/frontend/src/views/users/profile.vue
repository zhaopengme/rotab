<template>
  <div>
    <h3>个人资料</h3>
    <br />
    <Form :label-width="120">
      <FormItem label="用户名">
        <i-input v-model="username" disabled></i-input>
      </FormItem>
      <FormItem label="显示名称">
        <i-input
          v-model="form.nickname"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem label="邮箱">
        <i-input
          v-model="form.email"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem label="描述">
        <i-input
          v-model="form.description"
          type="textarea"
          :autosize="{ minRows: 2, maxRows: 5 }"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="submitHandler">更新</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  data() {
    return {
      form: {
        nickname: "",
        email: "",
        description: ""
      }
    };
  },
  computed: {
    ...mapState("user", {
      username: state => state.username
    })
  },
  mounted() {
    this.getData();
  },
  methods: {
    async getData() {
      this.form = await this.$store.dispatch("user/getUserinfo");
    },
    async submitHandler() {
      await this.$store.dispatch("user/updateUserinfo", this.form);
      this.getData();
    }
  }
};
</script>

<style></style>
