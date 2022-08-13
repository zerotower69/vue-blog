import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    component: () => import("@/pages/Home/index.vue"),
  },
  {
    path: "/about",
    name: "about",
    component: () =>
      import(/* webpackChunkName: "about" */ "@/pages/About/index.vue"),
  },
  {
    path: "/articles",
    name: "ArticleList",
    component: () => import("@/pages/Articles/index.vue"),
  },
  {
    path: "/artDetail",
    name: "artDetail",
    component: () => import("@/pages/ArtDetail/index.vue"),
  },
  {
    path: "/classes",
    name: "Classes",
    component: () => import("@/pages/Classes/index.vue"),
  },
  {
    path: "/tags",
    name: "Tags",
    component: () => import("@/pages/Tags/tags.vue"),
  },
  {
    path: "/gallery",
    name: "Gallery",
    component: () => import("@/pages/Gallery/index.vue"),
  },
  {
    path: "/img",
    name: "Img",
    component: () => import("@/pages/Img/index.vue"),
  },
  {
    path: "/say",
    name: "Say",
    component: () => import("@/pages/Say/say.vue"),
  },
  {
    path: "/msg",
    name: "Msg",
    component: () => import("@/pages/Msg/msg.vue"),
  },
  {
    path: "/link",
    name: "Link",
    component: () => import("@/pages/Link/index.vue"),
  },
  {
    path: "/show",
    name: "Show",
    component: () => import("@/pages/Show/show.vue"),
  },
  {
    path: "/log",
    name: "Log",
    component: () => import("@/pages/Log/index.vue"),
  },
  {
    path: "/post",
    name: "Post",
    component: () => import("@/pages/Post/post.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
