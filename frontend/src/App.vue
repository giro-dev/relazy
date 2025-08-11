<script setup>
import { ref } from 'vue'
import Sidebar from './components/Sidebar.vue'
import ProjectView from './views/ProjectView.vue'

const selectedProject = ref(null)

function selectProject(project) {
  selectedProject.value = project
}
</script>

<template>
  <div class="flex h-screen">
    <Sidebar
        :selected="selectedProject"
        @select="selectProject"
    />
    <div class="flex  bg-gray-100">
      <ProjectView :project="selectedProject" />
    </div>
  </div>
</template>

<style>
body {
  margin: 0;
  font-family: sans-serif;
}
</style>
