<template>
  <Card>
    <Row :gutter="16">
      <i-col span="24">
        <h3>撰写设置</h3>
        <br />
        <Form :label-width="120">
          <FormItem label="默认分类">
            <Select>
              <Option value="beijing">未分类</Option>
            </Select>
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
