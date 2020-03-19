<template>
  <Row :gutter="16">
    <i-col span="24">
      <h3>常规选项</h3>
      <br />
      <Form :label-width="120">
        <FormItem label="站点标题">
          <i-input v-model="option.title"></i-input>
        </FormItem>
        <FormItem label="副标题">
          <i-input v-model="option.subTitle"></i-input>
        </FormItem>
        <FormItem label="描述">
          <i-input
            type="textarea"
            v-model="option.description"
            :autosize="{ minRows: 2, maxRows: 5 }"
          ></i-input>
        </FormItem>
        <FormItem>
          <Button type="primary" @click="submitHandler">保存</Button>
        </FormItem>
      </Form>
    </i-col>
  </Row>
</template>

<script>
export default {
  data() {
    return {
      option: {}
    };
  },
  computed: {
    formParams: function() {
      return {
        category: "general",
        title: this.option.title,
        subTitle: this.option.subTitle,
        description: this.option.description
      };
    }
  },
  mounted() {
    this.getData();
  },
  methods: {
    async getData() {
      const params = {
        category: "general"
      };
      this.option = await this.$store.dispatch("option/getByCategory", params);
    },
    async submitHandler() {
      const result = await this.$store.dispatch(
        "option/updateGeneralOption",
        this.formParams
      );
      console.log(result);
    }
  }
};
</script>
