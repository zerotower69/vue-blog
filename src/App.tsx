import s from "./App.module.scss";

import Nav from "@/components/Nav/index.vue";
import Main from "@/components/Main/index.vue";
import Footer from "@/components/Footer/index.vue";

const App = () => {
  return (
    <div class={s.AppBox}>
      <Nav></Nav>
    </div>
  );
};

export default App;
