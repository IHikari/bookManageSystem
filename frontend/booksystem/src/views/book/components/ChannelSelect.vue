<script setup>
import { bookGetChannelsService } from '../../../api/book.js'
import { ref } from 'vue'

defineProps({
  modelValue: {
    type: [Number, String]
  },
  width: {
    type: String
  }
})

const emit = defineEmits(['update:modelValue'])

const channelList = ref([])
const getChannelList = async () => {
  const res = await bookGetChannelsService()
  channelList.value = res.data.data
}
getChannelList()
</script>

<template>
  <!--
     el-select下拉选框
     (1)label:展示给用户看的
     (2)value:收集起来提交给后台的
     -->
  <el-select
    :modelValue="modelValue"
    @update:modelValue="emit('update:modelValue', $event)"
    :style="{ width }"
  >
    <el-option
      v-for="channel in channelList"
      :key="channel.ID"
      :label="channel.name"
      :value="channel.ID"
    ></el-option>
  </el-select>
</template>

<style scoped></style>
