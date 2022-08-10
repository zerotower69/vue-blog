<template>
  <footer class="footer">
    <span>
      个人博客系统
      <a :href="source_github" target="_blank" ref="noreferrer" class="text">
        「源代码」
      </a>
    </span>
    <span>
      <a :href="icp_site" target="_blank" ref="noreferrer" class="text">
        {{ icp_no }}</a
      >
    </span>
    <span>
      <span v-for="(item, index) in frameArr" class="siteFrame" :key="index">
        {{ item }}
      </span>
    </span>
  </footer>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
import { $ref } from "vue/macros";
import * as CONSTANT from "@/utils/constant";
export default defineComponent({
  name: "Footer",
  setup() {
    const frameArr = $ref([
      "React",
      "React Router",
      "Redux",
      "Webpack",
      "AntD",
      "ahooks",
      "CloudBase",
    ]);
    const icp_no = $ref(CONSTANT.icp_no);
    const icp_site = $ref(CONSTANT.icp_site);
    const source_github = $ref(CONSTANT.source_github);
    return {
      frameArr,
      icp_no,
      icp_site,
      source_github,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/assets/styles/base.scss";
@import "@/assets/styles/style.scss";

.footer {
  height: $footerHeight;
  background-color: $footerBg;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
  color: $textColor;
  font-size: 16px;
  margin-top: 2 * $space;
  font-weight: 700;

  .text {
    color: $textColor;
    @extend .trans;
  }
  .text:hover {
    color: $hoverColor;
  }

  .siteFrame {
    background-color: $themeColor;
    padding: 2px 8px;
    margin-right: 10px;
    border-radius: 6px;
    user-select: none;
    @extend .hover;
  }
  .siteFrame:last-child {
    margin-right: 0;
  }
}

@media screen and (max-width: 1240px) {
  .footer {
    font-weight: 400;
    height: 70rem;
    font-size: 14rem;
    margin-top: 20rem;

    .siteFrame {
      display: none;
    }
  }
}
</style>
