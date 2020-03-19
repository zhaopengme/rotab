<template>
  <Row :gutter="16">
    <i-col span="12">
      <h3>添加标签</h3>
      <br />
      <Form :label-width="120">
        <FormItem label="名称">
          <i-input placeholder="Enter something..."></i-input>
        </FormItem>
        <FormItem label="别名">
          <i-input placeholder="Enter something..."></i-input>
        </FormItem>
        <FormItem label="描述">
          <i-input
            type="textarea"
            :autosize="{ minRows: 2, maxRows: 5 }"
            placeholder="Enter something..."
          ></i-input>
        </FormItem>
        <FormItem>
          <Button type="primary">添加标签</Button>
        </FormItem>
      </Form>
    </i-col>
    <i-col span="12">
      <Table border :columns="columns12" :data="data6">
        <template slot-scope="{ row }" slot="name">
          <strong>{{ row.name }}</strong>
        </template>
        <template slot-scope="{ row, index }" slot="action">
          <Button
            type="primary"
            size="small"
            style="margin-right: 5px"
            @click="show(index)"
            >编辑</Button
          >
          <Button type="error" size="small" @click="remove(index)"
            >Delete</Button
          >
        </template>
      </Table>
      <br />
      <Page :total="100" />
    </i-col>
  </Row>
</template>

<script>
export default {
  data() {
    return {
      columns12: [
        {
          title: "评论",
          slot: "name"
        },
        {
          title: "回复至",
          key: "age"
        },
        {
          title: "作者",
          key: "age"
        },
        {
          title: "发布时间",
          key: "address"
        },
        {
          title: "操作",
          slot: "action",
          width: 150,
          align: "center"
        }
      ],
      data6: [
        {
          name: "John Brown",
          age: 18,
          address: "New York No. 1 Lake Park"
        },
        {
          name: "Jim Green",
          age: 24,
          address: "London No. 1 Lake Park"
        },
        {
          name: "Joe Black",
          age: 30,
          address: "Sydney No. 1 Lake Park"
        },
        {
          name: "Jon Snow",
          age: 26,
          address: "Ottawa No. 2 Lake Park"
        }
      ]
    };
  },
  methods: {
    show(index) {
      this.$Modal.info({
        title: "User Info",
        content: `Name：${this.data6[index].name}<br>Age：${this.data6[index].age}<br>Address：${this.data6[index].address}`
      });
    },
    remove(index) {
      this.data6.splice(index, 1);
    }
  }
};
</script>
