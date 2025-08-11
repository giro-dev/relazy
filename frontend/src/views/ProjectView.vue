<script setup>
import RepoList from '../components/RepoList.vue'

defineProps({
  project: Object
})
</script>

<template>
  <div class="p-4">
    <h1 class="text-xl font-bold mb-4">{{ project.name }}</h1>
    <RepoList :basePath="project.path" />
  </div>
</template>