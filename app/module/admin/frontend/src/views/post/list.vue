<template>
  <div>
    <Table border :columns="columns" :data="pages.data">
      <template slot-scope="{ row }" slot="name">
        <strong>{{ row.name }}</strong>
      </template>
      <template slot-scope="{ row, index }" slot="action">
        <Button type="primary" size="small" :to="'/post/edit?id=' + row.id">
          编辑
        </Button>
        <Poptip
          confirm
          placement="left"
          title="Are you sure you want to delete this item?"
          @on-ok="deleteHandler(row, index)"
        >
          <Button size="small">Delete</Button>
        </Poptip>
      </template>
    </Table>
    <br />
    <Page :total="pages.total" @on-change="getData" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      columns: [
        {
          title: "id",
          width: 60,
          key: "id"
        },
        {
          title: "标题",
          key: "title"
        },
        {
          title: "发布时间",
          key: "created_at"
        },
        {
          title: "操作",
          slot: "action",
          width: 150,
          align: "center"
        }
      ],
      pages: { data: [], total: 0 }
    };
  },
  mounted() {
    this.getData(1);
  },
  methods: {
    async getData(page) {
      const params = { page: page };
      this.pages = await this.$store.dispatch("posts/list", params);
    },
    deleteHandler(item, idx) {
      this.$store.dispatch("posts/delete", item.id);
      this.pages.data.splice(idx, 1);
    }
  }
};
</script>

<style></style>
