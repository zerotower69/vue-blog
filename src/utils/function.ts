import {} from "@vueuse/core";
import { ref, reactive } from "vue";
import type { Ref } from "vue";
/**
 * 生成指定范围内的随机整数，左闭右闭
 * @param {Number} Min
 * @param {Number} Max
 * @return {Number}
 */
export const getRandomNum = (Min: number, Max: number) => {
  const Range = Max - Min + 1;
  const Rand = Math.random();
  return Min + Math.floor(Rand * Range);
};

/**
 * 打乱数组
 * @param {any[]} array
 * @return {any[]}
 */
export const shuffleArray = (array: any[]) => {
  if (!array) return [];
  const res = [...array];
  const len = res.length;
  for (let i = len - 1; i > 0; i--) {
    const randomPos = Math.floor(Math.random() * (i + 1));
    [res[i], res[randomPos]] = [res[randomPos], res[i]];
  }
  return res;
};

export const useRem = (scaleWidth?: number) => {
  if (scaleWidth === undefined) {
    scaleWidth = 750; //设计稿默认为750
  }
  function inMobile() {
    document.getElementsByTagName("html")[0].style.fontSize =
      document.documentElement.clientWidth / (scaleWidth as number) + "px";
  }
  return inMobile;
};

/**
 * 通过一个函数得到响应式数据和设置这个响应对象的方法
 * @param {string} key 值的名字key
 * @param {T} initialValue 初始值
 * @returns {[T,(T)=>void]}  值和设置这个响应数据的函数对象
 */
export function useSafeState<T>(initialValue?: T) {
  let refVal: any = null;
  const valType = typeof initialValue;
  if (valType === "undefined") {
    refVal = ref(null);
  } else if (["string", "number", "boolean"].includes(valType)) {
    refVal = ref(initialValue);
  } else if (["object", "null"].includes(valType)) {
    refVal = reactive(initialValue as any);
  }
  function setValue(val?: T) {
    refVal.value = val;
  }

  return [refVal as Ref<T>, setValue] as const;
}
