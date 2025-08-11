<script setup>
import { ref, onMounted } from 'vue'
import { Config } from '../../wailsjs/go/main/App'

const props = defineProps({
  selected: Object
})
const emit = defineEmits(['select'])

const projects = ref([])

onMounted(() => {
  // Replace this with API call if needed
  projects.value = Config.projects || [];
  if (projects.value.length && !props.selected) {
    emit('select', projects.value[0])
  }
})

function selectProject(project) {
  emit('select', project)
}
</script>

<template>
  <aside class="sidebar">
    <ul>
      <li
        v-for="project in projects"
        :key="project.id"
        :class="{ selected: project.id === props.selected?.id }"
        @click="selectProject(project)"
      >
        <span>{{ project.icon }}</span>
        <span>{{ project.name }}</span>
      </li>
    </ul>
  </aside>
</template>

<style>
.sidebar {
  width: 250px;
  background: #f5f5f5;
  padding: 1rem;
}
.selected {
  background: #e0e0e0;
}
</style>