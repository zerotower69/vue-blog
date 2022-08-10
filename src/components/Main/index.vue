<template>
  <main>
    <div>
      <error-boundary ref="error"> </error-boundary>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ErrorBoundary from "@/components/ErrorBoundary/index.vue";
export default defineComponent({
  components: {
    ErrorBoundary,
  },
  setup() {},
});
</script>
