<template>
  <div
    :class="{ card: true, center: loading, active: !isStatic }"
    @click="onClick"
  >
    <a-skeleton v-if="loading" :paragraph="{ rows: 1 }"></a-skeleton>
    <slot v-else></slot>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@vue/runtime-core";
interface Props {
  loading?: boolean;
  isStatic?: boolean;
  onClick?: (a: unknown) => void;
}
export default defineComponent({
  name: "Card",
  setup() {
    const props = withDefaults(defineProps<Props>(), {
      loading: false,
      isStatic: false,
      onClick: () => {},
    });
    return {
      ...props,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "@/assets/styles/base.scss";
@import "@/assets/styles/style.scss";

.card {
  margin-bottom: $space;
  border-radius: $bigRadius;
  overflow: hidden;
  user-select: none;
  background-color: $themeColor;
  color: $textColor;
  padding: $space;
  @extend .trans;
}

.active:hover {
  transform: scale(1.04);
}

.center {
  display: flex !important;
  justify-content: center !important;
  align-items: center !important;
}

@media screen and (max-width: 1240px) {
  .card {
    margin-bottom: 10rem;
    border-radius: 18rem;
    padding: 10rem;
  }
}
</style>
