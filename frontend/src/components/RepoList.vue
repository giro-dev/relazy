<script setup>
import { ref, onMounted, watch } from 'vue'
import { ListLocalRepos } from '../../wailsjs/go/main/App'

const props = defineProps({
  basePath: String
})

const repos = ref([])
const loading = ref(false)

async function loadRepos() {
  loading.value = true
  repos.value = await ListLocalRepos(props.basePath)
  loading.value = false
}

onMounted(loadRepos)
watch(() => props.basePath, loadRepos)
</script>

<template>
  <div>
    <div v-if="loading" class="text-gray-500">Carregant repos...</div>
    <ul v-else class="divide-y divide-gray-200">
      <li v-for="r in repos" :key="r" class="p-2 hover:bg-gray-50 cursor-pointer">
        {{ r }}
      </li>
    </ul>
  </div>
</template>
