<template>
  <i-col span="12">
    <h3>导航管理</h3>
    <br />
    <Form :label-width="120">
      <FormItem label="标题">
        <i-input
          v-model="form.title"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem label="链接">
        <i-input v-model="form.url" placeholder="Enter something..."></i-input>
      </FormItem>
      <FormItem label="打开方式">
        <Select v-model="form.open_method">
          <Option value="self">当前页</Option>
          <Option value="_balnk">新页面</Option>
        </Select>
      </FormItem>
      <FormItem label="是否使用">
        <RadioGroup v-model="form.status">
          <Radio label="no">禁用</Radio>
          <Radio label="yes">启用</Radio>
        </RadioGroup>
      </FormItem>
      <FormItem label="排序">
        <i-input
          v-model="form.number"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="submitHandler">确定</Button>
      </FormItem>
    </Form>
  </i-col>
</template>

<script>
import { mapState } from "vuex";

export default {
  data() {
    return {
      form: {
        id: "",
        title: "",
        url: "",
        status: "",
        open_method: "",
        number: ""
      }
    };
  },
  computed: {
    ...mapState("navigations", {
      item: state => state.item
    })
  },
  watch: {
    item: function(newVal) {
      this.form.id = newVal.id;
      this.form.title = newVal.title;
      this.form.url = newVal.url;
      this.form.open_method = newVal.open_method;
      this.form.number = newVal.number;
      this.form.status = newVal.status;
    }
  },
  methods: {
    async submitHandler() {
      const result = await this.$store.dispatch(
        "navigations/saveOrUpdate",
        this.form
      );
      console.log(result);
      this.$store.dispatch("navigations/list");
      this.$store.commit("navigations/setData", {});
    }
  }
};
</script>
