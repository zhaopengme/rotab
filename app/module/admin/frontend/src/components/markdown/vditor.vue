<template>
  <div id="vditor" />
</template>

<script>
import Vditor from "vditor";
import "vditor/src/assets/scss/classic.scss";

export default {
  name: "MarkdownEditor",
  props: {
    value: {
      type: String,
      default: ""
    },
    options: {
      type: Object,
      default: () => {
        return {};
      }
    }
  },
  data() {
    return {
      vditor: null
    };
  },
  mounted() {
    this.initVditor();
  },
  methods: {
    initVditor() {
      const options = {
        cache: false,
        width: "100%",
        preview: {
          delay: 1000,
          show: true
        }
      };
      this.vditor = new Vditor("vditor", options);
      // this.$nextTick(() => {
      //   this.isLoading = false;
      // });
    }
  }
};
</script>
