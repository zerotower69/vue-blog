import './index.custom.module.scss';
//import icons
import {
  BgColorsOutlined,
  CheckOutlined,
  HomeOutlined,
  MenuOutlined,
  SettingOutlined,
} from "@ant-design/icons-vue";

//import comp
import { Drawer } from 'ant-design-vue';
import { RouterLink, useRoute } from 'vue-router';
import { blogAdminUrl } from "@/utils/constant";
import { modeMap, modeMapArr } from "@/utils/modeMap";

import { useLinkList } from "./config"

import style from "./index.module.scss";

import type { Component} from 'vue'

interface Props {
  navShow?: boolean;
  setNavShow?: Function;
  mode?: number;
  setMode?: Function;
}

const bodyStyle = window.document.getElementsByTagName('body')[0].style;

const Nav: Component<Props> = ({ navShow,setNavShow,mode,setMode}) => {
    const $route = useRoute();

    
}



