<template>
  <div>
    <div :class="s.mobileNavBtn" @click="setVisible(true)">
      <MenuOutlined></MenuOutlined>
    </div>
    <!-- <Drawer
      placement="right"
      title="aaa"
      v-model:visible="visible"
      class="mobile-nav-box"
    >
      <div :class="[s.mobileNavBox]">
        <router-link
          :class="[
            item.to === $route.path ? s.mobileNavActives : s.mobileNavItem,
          ]"
          v-for="(item, index) in navArr"
          :key="index"
          :to="item.to"
        >
          {{ item.name }}
        </router-link>
      </div>
    </Drawer> -->
  </div>
</template>
<script lang="ts">
import style from "./index.module.scss";
import { MenuOutlined } from "@ant-design/icons-vue";
import { Drawer } from "ant-design-vue";
import { defineComponent, ref, defineEmits, onMounted } from "vue";
import { useLinkList } from "./config";
export default defineComponent({
  name: "MobileNav",
  components: {
    MenuOutlined,
    Drawer,
  },
  emits: {},
  setup(props, context) {
    // const visible = ref(false);
    const setVisible = (val: boolean) => {
      //   context.emit("update:visible", val);
      visible.value = val;
    };
    const { mobileNavArr } = useLinkList();
    const navArr = ref(mobileNavArr);
    const visible = ref(false);
    const s = ref(style);
    onMounted(() => {
      console.log();
      visible.value = true;
    });
    return {
      visible,
      navArr,
      setVisible,
      s,
    };
  },
});
</script>
<!-- <style lang="scss" scoped>
@import "@/assets/styles/base.scss";
.mobile-nav-box {
  display: none;
  :deep(.ant-drawer-content-wrapper) {
    width: 90rem !important;

    .ant-drawer-header-close-only {
      display: none;
    }

    .ant-drawer-content {
      background-color: $themeColor;
      color: $textColor;

      .ant-drawer-body {
        padding: 0;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
}

@media screen and (max-width: 1240px) {
  .mobile-nav-box {
    display: block;
  }
}
</style> -->
