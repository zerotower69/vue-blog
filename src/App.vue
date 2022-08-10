<template>
  <div class="AppBox">
    <Nav></Nav>
    <Main></Main>
    <Footer></Footer>
    <!-- back to top -->
    <a-back-top>
      <div class="ant-back-top-inner">UP</div>
    </a-back-top>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import Nav from "@/components/Nav/index.vue";
import Main from "@/components/Main/index.vue";
import Footer from "@/components/Footer/index.vue";
export default defineComponent({
  name: "App",
  components: {
    Nav,
    Main,
    Footer,
  },
});
</script>
<style lang="scss">
@import "./global.custom.scss";

@font-face {
  font-family: "FiraCode";
  font-display: auto;
  src: url(@/assets/fonts/FiraCode-Regular.ttf);
  font-weight: normal;
  font-style: normal;
}

.AppBox {
  font-family: "dengxian";
  background-position: center 0;
  background-size: cover;
  background-repeat: no-repeat;
  background-attachment: fixed;
}

.bg0 {
  background-image: url(@/assets/imgs/bg0.webp);
}
.bg1 {
  background-image: url(@/assets/imgs/bg1.webp);
}
.bg2 {
  background-image: url(@/assets/imgs/bg2.webp);
}
</style>
