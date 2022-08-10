<template>
  <div class="nav">
    <div class="navContent">
      <!-- 主页  -->
      <!-- TODO:click to main page -->
      <div class="homeBtn" @click="handleClickToMainPage">
        <HomeOutlined />
      </div>
      <!-- TODO:后台管理  加入后台系统的链接 -->
      <a :href="blogUrl" class="adminBtn">
        <SettingOutlined />
      </a>
      <!-- 暗黑模式切换 -->
      <!-- 文章单独按钮 -->
      <!-- 其他按钮 -->
      <!-- TODO:获得isActive.设置激活菜单按钮的样式 类似 el-menu -->
      <router-link
        class="navBtn"
        v-for="(item, index) in navArr"
        :key="index"
        :to="item.to"
      >
        {{ item.name }}
      </router-link>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { $ref } from "vue/macros";
import { Drawer } from "ant-design-vue";
import {
  BgColorsOutlined,
  CheckOutlined,
  HomeOutlined,
  MenuOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";
//@ts-ignore
import { blogAdminUrl } from "@/utils/constant";
import { useLinkList } from "./config";
export default defineComponent({
  name: "Nav",
  components: {
    Drawer,
    //icons
    BgColorsOutlined,
    CheckOutlined,
    HomeOutlined,
    MenuOutlined,
    SettingOutlined,
  },
  setup() {
    const handleClickToMainPage = () => {};
    const blogUrl = $ref(blogAdminUrl);
    //获得按钮链接
    const { navArr, secondNavArr, mobileNavArr } = useLinkList();
    return {
      blogUrl,
      handleClickToMainPage,
      navArr,
      secondNavArr,
      mobileNavArr,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/assets/styles/base.scss";
@import "@/assets/styles/style.scss";

.baseBtn {
  height: 44px;
  width: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 22px;
  font-weight: 700;
  margin-right: 20px;
  border-radius: 14px;
  color: $textColor;
  user-select: none;
  @extend .hover;
}

.hiddenNav {
  box-shadow: none !important;
  transform: translate(0, -$navHeight);
}

.nav {
  width: 100%;
  height: $navHeight;
  background-color: $themeColor1;
  position: fixed;
  top: 0;
  z-index: 10;
  box-shadow: 0 0 18px $footerBg;
  @extend .trans;

  .navContent {
    position: relative;
    width: $centerWidth;
    height: 100%;
    margin: 0 auto;
    display: flex;
    justify-content: center;
    align-items: center;

    .homeAndAdmin {
      @extend .baseBtn;
      position: absolute;
      top: 50%;
      right: 0;
      transform: translate(0, -50%);
      font-size: 26px;
      width: 60px;
    }

    .homeBtn {
      @extend .homeAndAdmin;
      cursor: pointer;
      left: 0;
    }

    .adminBtn {
      @extend .homeAndAdmin;
      margin-right: 0;
    }

    .navBtn {
      @extend .baseBtn;
    }
    .navBtn:last-child {
      margin-right: 0;
    }
    .navActive {
      @extend .navBtn;
      background-color: $hoverColor;
    }

    .articlesBtn {
      position: relative;
      @extend .baseBtn;

      .articelsSecond {
        position: absolute;
        top: -160px;
        width: 90px;
        background-color: $themeColor1;
        border-radius: 14px;
        padding: 10px;
        z-index: 0;
        @extend .trans;

        .articelsSecondItem {
          @extend .hover;
          display: flex;
          justify-content: center;
          align-items: center;
          font-size: 18px;
          height: 34px;
          margin-bottom: 10px;
          border-radius: 10px;
          user-select: none;
          background-color: $themeColor2;
          color: $textColor;
        }
        .articelsSecondItem:last-child {
          margin-bottom: 0;
        }
        .sedActive {
          @extend .articelsSecondItem;
          background-color: $hoverColor;
        }
      }
    }
    .articlesBtn:hover .articelsSecond {
      top: 60px;
    }

    .modeBtn {
      @extend .homeAndAdmin;
      right: 80px;
      margin-right: 0;

      .modeOpions {
        position: absolute;
        left: 50%;
        top: -180px;
        transform: translate(-50%, 0);
        width: 80px;
        background-color: $themeColor1;
        border-radius: 14px;
        padding: 10px;
        z-index: 0;
        @extend .trans;

        .modeItem {
          height: 40px;
          background-color: $themeColor;
          margin-bottom: 10px;
          border-radius: 10px;
          display: flex;
          justify-content: center;
          align-items: center;
          font-size: 20px;
          @extend .trans;
          color: #fff;
        }
        .modeItem1,
        .modeItem2 {
          color: #000;
        }

        .modeItem:last-child {
          margin-bottom: 0;
        }
        .modeItem:hover {
          transform: scale(1.07);
          box-shadow: 4px 4px 8px rgba(0, 0, 0, 0.4);
        }
      }
    }
    .modeBtn:hover .modeOpions {
      top: 60px;
    }
  }
}

// 手机端呼出导航的按钮
.mobileNavBtn {
  width: 50rem;
  height: 50rem;
  display: none;
  justify-content: center;
  align-items: center;
  font-size: 22rem;
  color: $textColor;
  position: fixed;
  top: 0;
  right: 0;
  z-index: 99;
}

// 手机端导航
.mobileNavBox {
  .mobileNavItem {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    color: $textColor;
    font-size: 18rem;
    font-family: "dengxian";
    width: 60rem;
    height: 34rem;
    border-radius: 10rem;
    margin-bottom: 16rem;
  }
  .mobileNavItem:last-child {
    margin-bottom: 0;
  }
  .mobileNavActive {
    @extend .mobileNavItem;
    background-color: $hoverColor;
  }
  .modeItem {
    @extend .mobileNavItem;
    border: 2rem solid #ccc;
  }
}

@media screen and (max-width: 1240px) {
  .nav {
    display: none;
  }
  .mobileNavBtn {
    display: flex;
  }
}

.mobile-nav-box {
  display: none;

  .ant-drawer-content-wrapper {
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
</style>
