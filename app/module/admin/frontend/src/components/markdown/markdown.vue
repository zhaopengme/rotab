<template>
  <div id="editor" class="vditor" />
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
      const that = this;
      const options = {
        cache: false,
        width: "100%",
        height: "0",
        tab: "\t",
        counter: "999999",
        typewriterMode: true,
        preview: {
          delay: 100,
          show: true
        },
        upload: {
          max: 5 * 1024 * 1024,
          // linkToImgUrl: 'https://sm.ms/api/upload',
          handler(file) {
            let formData = new FormData();
            for (let i in file) {
              formData.append("file", file[i]);
            }
            let request = new XMLHttpRequest();
            request.open("POST", "http://localhost:8199/file/upload");
            request.onload = that.onloadCallback;
            request.send(formData);
          }
        },
        input: this.changeHandler
      };
      var v = new Vditor("editor", options);
      window.v = v;
      this.vditor = v;
      this.$nextTick(() => {
        this.vditor.focus();
        this.vditor.setValue(this.value);
      });
    },

    onloadCallback(oEvent) {
      const currentTarget = oEvent.currentTarget;
      if (currentTarget.status !== 200) {
        return this.$message({
          type: "error",
          message: currentTarget.status + " " + currentTarget.statusText
        });
      }
      let resp = JSON.parse(currentTarget.response);
      if (resp.code === 0) {
        this.vditor.insertValue(`![${resp.data.name}](${resp.data.url})`);
      } else {
        console.log(resp);
      }
    },
    setDefaultText() {
      const savedMdContent = localStorage.getItem("vditorvditor") || "";
      if (!savedMdContent.trim()) {
        // localStorage.setItem("vditorvditor", defaultText);
      }
    },
    changeHandler(mdval) {
      this.$emit("input", mdval);
      this.$emit("on-change", mdval);
      this.$emit("changed", this.vditor.getHTML());
    }
  }
};
</script>

<style lang="scss">
.vditor {
  .vditor-wysiwyg {
    padding: 3px 5px 3px 5px !important;
  }
  .vditor-reset p {
    margin-top: auto;
    margin-bottom: auto;
  }
}
</style>
