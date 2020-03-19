<template>
  <Select :value="value" @on-change="changeHandler">
    <Option v-for="item in list" :key="item.id" :value="item.id">
      {{ item.title }}
    </Option>
  </Select>
</template>

<script>
export default {
  props: ["value"],
  data() {
    return {
      list: []
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    async getData() {
      this.list = await this.$store.dispatch("categories/all");
    },
    changeHandler(categoryId) {
      this.$emit("input", categoryId);
    }
  }
};
</script>

<style></style>
