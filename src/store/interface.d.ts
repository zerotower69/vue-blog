import { RemovableRef } from "@vueuse/core";

export interface storeState {
  navShow: boolean | RemovableRef<boolean>;
  artSum: number | RemovableRef<number>;
  name: string | RemovableRef<string>;
  link: string | RemovableRef<string>;
  email: string | RemovableRef<string>;
  avatar: string | RemovableRef<string>;
  mode: number | RemovableRef<number>;
}
