<script setup>
import RepoList from '../components/RepoList.vue'

defineProps({
  project: Object
})
</script>

<template>
  <div class="p-4">
    <h1 class="text-xl font-bold mb-4">{{ project.Name }}</h1>
    <RepoList :projectId="project.Name" />
  </div>
</template>