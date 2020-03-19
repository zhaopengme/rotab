<template>
  <div>
    <Spin size="large" fix v-if="loading"></Spin>
    <Form v-else ref="form" :model="form">
      <FormItem label="标题">
        <i-input v-model="form.title" placeholder="请输入标题"></i-input>
      </FormItem>
      <FormItem label="内容"> </FormItem>
      <Markdown
        v-model="form.content"
        :localCache="false"
        @changed="val => (form.htmlcontent = val)"
      />
      <FormItem label="标签">
        <i-input v-model="form.tags" placeholder="请输入标签"></i-input>
      </FormItem>
      <FormItem label="分类">
        <Categories v-model="form.category_id" />
      </FormItem>
      <FormItem>
        <Button @click="handleSubmit" type="primary">发布</Button>
      </FormItem>
    </Form>
  </div>
</template>

<script>
import Markdown from "../../components/markdown";
import Categories from "../../components/Categories";

export default {
  components: { Markdown, Categories },
  data() {
    return {
      loading: true,
      form: {
        id: "",
        title: "",
        category_id: "",
        tags: "",
        content: "",
        htmlcontent: ""
      }
    };
  },
  mounted() {
    if (!this.$kit.isBlank(this.$route.query.id)) {
      this.getData(this.$route.query.id);
    } else {
      this.loading = false;
    }
  },
  methods: {
    async getData(id) {
      this.form = await this.$store.dispatch("posts/get", id);
      this.loading = false;
    },
    async handleSubmit() {
      await this.$store.dispatch("posts/saveOrUpdate", this.form);
    }
  }
};
</script>
