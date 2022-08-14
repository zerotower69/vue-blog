import { defineComponent, defineProps, withDefaults } from "vue";
import classNames from "classnames";
import s from "./index.module.scss";

interface Props {
  title?: string;
  desc?: string;
  className?: string;
}

const PageTitle = defineComponent({
  name: "PageTitle",
  props: {
    title: String,
    desc: String,
    className: String,
  },
  setup(props, context) {
    return () => (
      <div class={classNames(s.box, props.className)}>
        <div class={s.title}>{props.title}</div>
        {props.desc && <div class={s.desc}>{props.desc}</div>}
        {context.slots.default && context.slots.default()}
      </div>
    );
  },
});

export default PageTitle;
