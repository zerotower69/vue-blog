<template>
  <h1 v-if="hasError"></h1>
  <div v-else>
    <slot>Something went wrong.</slot>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "ErrorBoundary",
  setup() {
    const hasError = ref(false);
    const getDerivedStateFromError = () => {
      hasError.value = true;
    };
    return {
      hasError,
      getDerivedStateFromError,
    };
  },
});
</script>
