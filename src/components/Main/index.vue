<template>
  <main>
    <!-- TODO:remove it -->
    <div style="height: 2000px !important">
      <error-boundary ref="error">
        <!-- TODO:利用suspense实现骨架屏 -->
        <!-- <Suspense>
          <template #fallback>
            <div>error</div>
          </template> -->
        <router-view />
        <!-- </Suspense> -->
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
