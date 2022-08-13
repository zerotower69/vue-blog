import { defineStore } from "pinia";
import { storeState } from "./interface";
import { useLocalStorage } from "@vueuse/core";
import type { RemovableRef } from "@vueuse/core";
export const useStore = defineStore("main", {
  state: (): storeState => {
    return {
      navShow: false,
      artSum: 0,
      name: useLocalStorage("name", ""),
      link: useLocalStorage("link", ""),
      email: useLocalStorage("email", ""),
      avatar: useLocalStorage("avatar", ""),
      mode: useLocalStorage("localMode", 1),
    };
  },
  actions: {
    //navbar showing or not
    setNavShow(data: boolean) {
      this.navShow = data;
    },
    //set the account of article
    setArtSum(data: number) {
      this.artSum = data;
    },
    //set user name
    setName(data: string) {
      this.name = data;
    },
    //set email
    setEmail(data: string) {
      this.email = data;
    },
    //set link
    setLink(data: string) {
      this.link = data;
    },
    //set avatar
    setAvatar(data: string) {
      this.avatar = data;
    },
    //set mode
    setMode(data: number) {
      this.mode = data;
    },
  },
});
