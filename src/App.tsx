import s from "./App.module.scss";
import classNames from "classnames";
import Nav from "@/components/Nav";
import Main from "@/components/Main/index.vue";
import Footer from "@/components/Footer";

import { BackTop } from "ant-design-vue";

import { useStore } from "@/store";
import { onMounted, onUnmounted, defineComponent } from "vue";
import type { UseScrollReturn } from "@vueuse/core";
import { vScroll } from "@vueuse/components";
import { useRem } from "@/utils/function";
const App = defineComponent({
  components: {
    Nav,
    Main,
    Footer,
    BackTop,
  },
  directives: { scroll: vScroll },
  setup(props) {
    //set bg color;
    const $store = useStore();
    const BgClasses = [s.bg0, s.bg1, s.bg2];
    const callback = useRem(450);
    onMounted(() => {
      if ($store.mode === undefined) {
        $store.setMode(1);
      }
      $store.setNavShow(true);
      // $store.setMode();
      //set rem
      callback();
      window.addEventListener("resize", callback);
    });
    onUnmounted(() => {
      window.removeEventListener("resize", callback);
    });

    // scroll callback
    function scrollCallback(state: UseScrollReturn) {
      const { x, y, isScrolling, arrivedState, directions } = state;
      const { top, bottom } = directions;
      console.log(arrivedState.bottom, arrivedState.left);
      if (top) {
        $store.setNavShow(true);
      } else if (!top && y.value == 0) {
        $store.setNavShow(true);
      } else if (bottom) {
        $store.setNavShow(false);
      }
    }
    return () => (
      <div class={classNames(s.AppBox, BgClasses[$store.mode!])}>
        <div v-scroll={scrollCallback}>
          <Nav />
          <Main></Main>
          <Footer></Footer>
          <BackTop></BackTop>
        </div>
      </div>
    );
  },
});

export default App;
