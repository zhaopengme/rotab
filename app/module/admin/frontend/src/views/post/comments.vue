<template>
  <div>
    <h2>文章</h2>
    <br />
    <Tabs v-model="value">
      <TabPane label="已通过" name="ok">
        <List status="ok" v-if="value == 'ok'" />
      </TabPane>
      <TabPane label="待审核" name="waiting">
        <List status="waiting" v-if="value == 'waiting'" />
      </TabPane>
      <TabPane label="垃圾" name="spam">
        <List status="spam" v-if="value == 'spam'" />
      </TabPane>
    </Tabs>
  </div>
</template>

<script>
import List from "./comments/list";

export default {
  components: { List },
  data() {
    return {
      value: "ok"
    };
  }
};
</script>
