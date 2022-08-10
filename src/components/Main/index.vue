<template>
  <main>
    <div>
      <error-boundary ref="error">
        <!-- TODO:利用suspense实现骨架屏 -->
        <suspense>
          <router-view></router-view>
        </suspense>
      </error-boundary>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ErrorBoundary from "@/components/ErrorBoundary/index.vue";
export default defineComponent({
  name: "Main",
  components: {
    ErrorBoundary,
  },
  setup() {},
});
</script>
