<template>
  <Row :gutter="16">
    <Edit />
    <List />
  </Row>
</template>

<script>
import Edit from "./categories/edit";
import List from "./categories/list";
export default {
  components: { Edit, List }
};
</script>
