<template>
  <i-col span="12">
    <Table border :columns="columns" :data="pages.data || []">
      <template slot-scope="{ row }" slot="name">
        <strong>{{ row.name }}</strong>
      </template>
      <template slot-scope="{ row, index }" slot="action">
        <Button
          type="primary"
          size="small"
          style="margin-right: 5px"
          @click="$store.commit('navigations/setData', row)"
          >编辑</Button
        >
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
    <Page :total="pages.total" @on-change="changePage" />
  </i-col>
</template>

<script>
import { mapState } from "vuex";

export default {
  data() {
    return {
      columns: [
        {
          title: "id",
          key: "id"
        },
        {
          title: "标题",
          key: "title"
        },
        {
          title: "status",
          key: "status"
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
      ]
    };
  },
  computed: {
    ...mapState("navigations", {
      pages: state => state.pages
    })
  },
  mounted() {
    this.getData();
  },
  methods: {
    getData() {
      const params = { page: 1 };
      this.$store.dispatch("navigations/list", params);
    },
    changePage(curPage) {
      const params = { page: curPage };
      this.$store.dispatch("navigations/list", params);
    },
    deleteHandler(item, idx) {
      this.$store.commit("navigations/delete", idx);
      this.$store.dispatch("navigations/delete", item.id);
      this.$store.commit("navigations/setData", {});
    }
  }
};
</script>

<style></style>
