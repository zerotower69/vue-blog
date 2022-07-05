<template>
    <div>
        <span>{{message}}</span>
    </div>
    <a-button type="primary">aaa</a-button>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
//@ts-ignore
// import {Vue} from 'vue-class-component'
export default defineComponent({
    setup(){
        const message=ref('hello, world!');
        return{
            message
        }
    }
})
</script>