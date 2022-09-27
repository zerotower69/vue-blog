import { defineComponent } from "vue";

import { skills, icp_no, icp_site, source_github } from "@/utils/constant";
import s from "./index.module.scss";
const Footer = defineComponent({
  name: "Footer",
  setup() {
    return () => (
      <footer class={s.footer}>
        <span>
          <span>个人博客系统</span>
          <a
            href={source_github}
            target="_blank"
            ref="noreferrer"
            class={s.text}
          >
            「源代码」
          </a>
        </span>
        <span>
          <a href={icp_site} target="_blank" ref="noreferrer" class={s.text}>
            {icp_no}
          </a>
        </span>
        <span>
          {skills.map((skill, index) => (
            <span
              key={index}
              class={s.siteFrame}
              onClick={() => {
                window.open(skill.url);
              }}
            >
              {skill.name}
            </span>
          ))}
        </span>
      </footer>
    );
  },
});

export default Footer;
