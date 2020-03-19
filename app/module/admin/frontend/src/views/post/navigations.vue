<template>
  <Row :gutter="16">
    <Edit />
    <List />
  </Row>
</template>

<script>
import Edit from "./navigations/edit";
import List from "./navigations/list";
export default {
  components: { Edit, List }
};
</script>
