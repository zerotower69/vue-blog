import "./index.custom.scss";

import s from "./index.scss";
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
import { RouterLink, useRouter } from "vue-router";
import { blogAdminUrl } from "@/utils/constant";
import { modeMap, modeMapArr } from "@/utils/modeMap";

import { defineComponent, useCssModule, onMounted, ref, reactive } from "vue";
// import { $ref } from "vue/macros";
import { useStore } from "@/store";
import { useEventListener, useLocalStorage, useScroll } from "@vueuse/core";

const modeOptions = ["rgb(19,38,36)", "rgb(110,180,214)", "rgb(171,194,208)"];
export interface Nav {
  navShow?: boolean;
  setNavShow?: () => void;
  mode?: number;
  setMode?: () => void;
}

// const style = useCssModule("./index.scss");

const bodyStyle = window.document.getElementsByTagName("body")[0].style;

const NavComp = defineComponent<Nav>({
  name: "Nav",
  setup(props) {
    const $router = useRouter();
    const $store = useStore();
    const { navArr, secondNavArr, mobileNavArr } = useLinkList();
    onMounted(() => {
      // console.log(style);
    });

    //   const [visible, { toggle,setLeft,setRight}] = useToggle(false);

    return () => (
      <>
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
      </>
    );
  },
});

export default NavComp;
