<template>
  <Row :gutter="16">
    <i-col span="24">
      <h3>讨论设置</h3>
      <br />
      <Form :label-width="120">
        <FormItem label="是否支持评论">
          <RadioGroup>
            <Radio label="male">打开</Radio>
            <Radio label="female">关闭</Radio>
          </RadioGroup>
        </FormItem>
        <FormItem>
          <Button type="primary">保存</Button>
        </FormItem>
      </Form>
    </i-col>
  </Row>
</template>

<script>
export default {};
</script>

<style></style>
