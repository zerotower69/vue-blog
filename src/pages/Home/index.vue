<template>
  <div class="body"></div>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core";

export default defineComponent({
  name: "Home",
  setup() {},
});
</script>
<style lang="scss" scoped>
.homeTitle {
  height: 100vh;
}

.body {
  display: flex;
}

@media screen and (max-width: 1240px) {
  .homeTitle {
    height: 50vh;
  }
}
</style>
