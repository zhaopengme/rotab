<template>
  <i-col span="12">
    <h3>新分类目录</h3>
    <br />
    <Form :label-width="120">
      <FormItem label="名称">
        <i-input
          v-model="form.title"
          placeholder="Enter something..."
        ></i-input>
      </FormItem>
      <FormItem label="别名">
        <i-input v-model="form.slug" placeholder="Enter something..."></i-input>
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
        slug: "",
        description: ""
      }
    };
  },
  computed: {
    ...mapState("categories", {
      item: state => state.item
    })
  },
  watch: {
    item: function(newVal) {
      this.form.id = newVal.id;
      this.form.title = newVal.title;
      this.form.slug = newVal.slug;
      this.form.description = newVal.description;
    }
  },
  methods: {
    async submitHandler() {
      const result = await this.$store.dispatch(
        "categories/saveOrUpdate",
        this.form
      );
      console.log(result);
      this.$store.dispatch("categories/list");
      this.$store.commit("categories/setData", {});
    }
  }
};
</script>
