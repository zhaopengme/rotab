<template>
  <Layout class="is-fullheight">
    <Sider class="has-background-white">
      <Sidebar />
    </Sider>
    <Layout>
      <Header class="has-background-white">
        <Headerr />
      </Header>
      <Content>
        <router-view />
      </Content>
      <Footer>Footer</Footer>
    </Layout>
  </Layout>
</template>

<script>
import Sidebar from "./Main/Sidebar.vue";
import Headerr from "./Main/Header.vue";
export default {
  components: { Sidebar, Headerr }
};
</script>
