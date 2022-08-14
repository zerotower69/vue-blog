import { defineComponent } from "vue";

import Aside from "./Aside";
import Section from "./Section";

import PageTitle from "@/components/PageTitle";

import { siteTitle } from "@/utils/constant";

import s from "./index.module.scss";

import { onMounted } from "vue";
import { useSafeState } from "@/utils/function";

const getPoem = require("jinrishici");
const Home = defineComponent({
  name: "Home",
  components: {
    PageTitle,
  },
  setup(props, { slots }) {
    const [poem, setPoem] = useSafeState("");
    onMounted(() => {
      getPoem.load(
        (res: {
          data: {
            content: string;
          };
        }) => setPoem(res.data.content)
      );
    });
    return () => (
      <>
        <PageTitle
          title={siteTitle}
          desc={poem.value || ""}
          className={s.homeTitle}
        />
        <div class={s.body}>
          <Aside />
          <Section />
        </div>
      </>
    );
  },
});

export default Home;
