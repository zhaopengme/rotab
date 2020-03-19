<template>
  <div>
    <Table border :columns="columns" :data="pages.data || []">
      <template slot-scope="{ row }" slot="name">
        <strong>{{ row.name }}</strong>
      </template>
      <template slot-scope="{ row }" slot="action">
        <ButtonGroup size="small">
          <Button size="small" type="text" @click="changeStatus(row.id, 'ok')">
            通过
          </Button>
          <Button
            size="small"
            type="text"
            @click="changeStatus(row.id, 'waiting')"
          >
            待审核
          </Button>
          <Button
            size="small"
            type="text"
            @click="changeStatus(row.id, 'spam')"
          >
            垃圾
          </Button>
          <Button size="small" type="text">
            编辑
          </Button>
          <Button size="small" type="text">
            回复
          </Button>
          <Button size="small" type="text" @click="deleteHandler(row.id)">
            删除
          </Button>
        </ButtonGroup>
      </template>
    </Table>
    <br />
    <Page :total="pages.total" @on-change="getData" />
  </div>
</template>

<script>
export default {
  props: ["status"],
  data() {
    return {
      columns: [
        {
          title: "id",
          key: "id"
        },
        {
          title: "文章标题",
          key: "title"
        },
        {
          title: "评论内容",
          key: "content"
        },
        {
          title: "发布时间",
          key: "created_at"
        },
        {
          title: "操作",
          slot: "action",
          align: "center",
          width: "330"
        }
      ],
      pages: { limit: 10, page: 1, total: 0, data: [] }
    };
  },
  mounted() {
    this.getData(1);
  },
  methods: {
    async getData(curPage) {
      const params = { page: curPage, status: this.status };
      let result = await this.$store.dispatch("comments/list", params);
      result.page = curPage;
      this.pages = result;
    },
    async changeStatus(id, status) {
      const params = { id: id, status: status };
      await this.$store.dispatch("comments/saveOrUpdate", params);
      this.getData(this.pages.page);
    },
    async deleteHandler(id) {
      await this.$store.dispatch("comments/delete", id);
      this.getData(this.pages.page);
    }
  }
};
</script>

<style></style>
