import "./index.custom.module.scss";
import s from "./index.module.scss";
import { useLinkList } from "./config";
//import icons
import {
  BgColorsOutlined,
  CheckOutlined,
  HomeOutlined,
  MenuOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";

import classNames from "classnames";
//import comp
import { Drawer } from "ant-design-vue";
import MobileNav from "./mobileNav.vue";
import { RouterLink, useRouter, useRoute } from "vue-router";
import { blogAdminUrl } from "@/utils/constant";
import { modeMap, modeMapArr } from "@/utils/modeMap";

import { defineComponent, useCssModule, onMounted, ref, reactive } from "vue";
// import { $ref } from "vue/macros";
import { useStore } from "@/store";
import { createGlobalState } from "@vueuse/core";

const modeOptions = ["rgb(19,38,36)", "rgb(110,180,214)", "rgb(171,194,208)"];
export interface Nav {
  navShow?: boolean;
  setNavShow?: () => void;
  mode?: number;
  setMode?: () => void;
}

const bodyStyle = window.document.getElementsByTagName("body")[0].style;

const NavComp = defineComponent<Nav>({
  name: "Nav",
  components: {
    Drawer,
    MobileNav,
  },
  setup(props) {
    const $route = useRoute();
    const $router = useRouter();
    const $store = useStore();
    const { navArr, secondNavArr, mobileNavArr } = useLinkList();
    onMounted(() => {
      // console.log(style);
    });

    //   const [visible, { toggle,setLeft,setRight}] = useToggle(false);
    // control visible;
    const useState = createGlobalState(() => {
      return {
        visible: false,
      };
    });
    const state = useState();
    const visible = ref(false);
    const setVisible = (val: boolean) => {
      visible.value = val;
    };
    return () => (
      <>
        {/* pc导航 */}
        <nav class={classNames(s.nav, { [s.hiddenNav]: !$store.navShow })}>
          <div class={s.navContent}>
            {/* 主页 */}
            <div class={s.homeBtn} onClick={() => $router.push("/")}>
              <HomeOutlined />
            </div>

            {/* 后台管理 */}
            <a
              class={s.adminBtn}
              href={blogAdminUrl}
              target="_blank"
              rel="noreferrer"
            >
              <SettingOutlined />
            </a>

            {/* 黑暗模式切换 */}
            <div class={classNames(s.modeBtn)}>
              <BgColorsOutlined />
              <div class={classNames(s.modeOptions)}>
                {modeOptions.map((backgroundColor, index) => (
                  <div
                    key={index}
                    style={{ backgroundColor }}
                    class={classNames(s.modeItem, s[`modeItem${index}`])}
                    onClick={() => $store.setMode(index)}
                  >
                    <CheckOutlined
                      style={{
                        display: $store.mode === index ? "block" : "none",
                      }}
                    />
                  </div>
                ))}
              </div>
            </div>

            {/* 文章单独按钮 */}
            <div class={s.articlesBtn}>
              <div class={s.articlesSecond}>
                {secondNavArr.map((item, index) => (
                  <RouterLink
                    class={
                      item.to == $route.path
                        ? s.sedActive
                        : s.articlesSecondItem
                    }
                    to={item.to}
                    key={index}
                  >
                    {item.name}
                  </RouterLink>
                ))}
              </div>
              <span>文章</span>
            </div>

            {/* 其他按钮 */}
            {navArr.map((item, index) => (
              <RouterLink
                class={s.navBtn}
                //   class={({ isActive }) => (isActive ? s.navActive : s.navBtn)}
                to={item.to}
                key={index}
              >
                {item.name}
              </RouterLink>
            ))}
          </div>
        </nav>
        {/* 移动端导航 */}
        {/* 点击打开 */}
        <div class={s.mobileNavBtn} onclick={setVisible(true)}>
          <MenuOutlined></MenuOutlined>
        </div>
        {/* <MobileNav></MobileNav> */}
        <Drawer
          placement="right"
          onClose={() => {
            visible.value = false;
          }}
          visible={visible.value}
          class={classNames(s.drawerClass)}
        >
          <div class={s.mobileNavBox}>
            {mobileNavArr.map((item, index) => (
              <RouterLink
                key={index}
                to={item.to}
                class={
                  item.to === $route.path ? s.mobileNavActive : s.mobileNavItem
                }
              >
                {item.name}
              </RouterLink>
            ))}
            {modeOptions.map((backgroundColor, index) => (
              <div
                key={index}
                style={{ backgroundColor }}
                class={classNames(s.modeItem, s[`modeItem${index}`])}
                onClick={() => $store.setMode(index)}
              >
                {$store.mode === index && <CheckOutlined />}
              </div>
            ))}
          </div>
        </Drawer>
      </>
    );
  },
});

export default NavComp;
