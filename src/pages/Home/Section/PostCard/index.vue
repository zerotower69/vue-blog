<template>
  <Card class="card" :isStaic="true" @click="onClick">
    <post-card-loading v-if="loading"></post-card-loading>
    <div v-else>
      <div class="title">{{ title }}</div>
    </div>
  </Card>
</template>

<script lang="ts">
interface Props {
  title?: string;
  content?: string;
  date?: number;
  tags?: string[];
  loading?: boolean;
  onClick?: () => void;
}
import { defineComponent, defineProps } from "vue";
import { $ref, $toRef } from "vue/macros";
import Card from "@/components/Card/index.vue";
import PostCardLoading from "./PostCardLoading/index.vue";
export default defineComponent({
  name: "PostCard",
  components: {
    Card,
    PostCardLoading,
  },
  setup() {
    const props = defineProps<Props>();
    return {
      ...props,
    };
  },
});
</script>
<style lang="scss" scoped>
@import "@/assets/styles/base.scss";
@import "@/assets/styles/style.scss";

.tagBase {
  padding: 0 10px;
  border-radius: 10px;
  color: $textColor;
  height: 30px;
  font-size: 18px;
  background-color: $themeColor1;
  @extend .hover;
}

.card {
  cursor: pointer;

  .title {
    margin: 0;
    height: 80px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 28px;
    font-weight: 700;
    color: $textColor;
  }

  .content {
    margin: 0;
    height: 80px;
    text-indent: 2em;
    color: $textColor;
    font-size: 18px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
  }

  .info {
    position: relative;
    height: 50px;

    .date {
      @extend .tagBase;
      position: absolute;
      bottom: 0;
      display: flex;
      justify-content: center;
      align-items: center;
    }

    .tags {
      position: absolute;
      bottom: 0;
      right: 0;
      max-width: 800px;
      height: 30px;

      .tag {
        @extend .tagBase;
        display: inline-block;
        margin-right: 20px;
        text-align: center;
        line-height: 30px;
      }

      .tag:last-child {
        margin-right: 0;
      }
    }
  }
}

.card:hover {
  transform: scale(1.02);
  cursor: pointer;
}

@media screen and (max-width: 1240px) {
  .card {
    .title {
      height: 50rem;
      line-height: 24rem;
      font-size: 20rem;
    }

    .content {
      height: 72rem;
      text-indent: 2em;
      font-size: 16rem;
    }

    .info {
      display: none;
    }
  }
}
</style>
