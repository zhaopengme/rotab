<template>
  <Card>
    <Row :gutter="16">
      <i-col span="24">
        <h3>阅读设置</h3>
        <br />
        <Form :label-width="120">
          <FormItem label="博客页面至多显示">
            <i-input placeholder="Enter something..."
              ><span slot="append">文章</span></i-input
            >
          </FormItem>
          <FormItem label="Feed中显示最近">
            <i-input placeholder="Enter something..."
              ><span slot="append">个项目</span></i-input
            >
          </FormItem>
          <FormItem>
            <Button type="primary">保存</Button>
          </FormItem>
        </Form>
      </i-col>
    </Row>
  </Card>
</template>

<script>
export default {};
</script>

<style></style>
