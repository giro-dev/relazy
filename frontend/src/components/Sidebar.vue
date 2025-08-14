<script setup>
import {onMounted, ref} from 'vue'
import {GetProjects, NewProject} from '../../wailsjs/go/main/App'

const props = defineProps({
  selected: Object
})
const emit = defineEmits(['select'])

const projects = ref([])

async function loadProjects() {
  // This function can be used to load projects if needed
  projects.value = await GetProjects();
}

const showInput = ref(false)
const newProjectName = ref('')

async function createProject() {
  if (!newProjectName.value.trim()) return
  // Replace with your actual API call to create a project
  const newProject = await NewProject(newProjectName.value.trim())
  projects.value.push(newProject)
  emit('select', newProject)
  newProjectName.value = ''
  showInput.value = false
}

onMounted(() => {
  // Replace this with API call if needed
  loadProjects()
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
    <ul class="divide-y divide-gray-200">
      <li
          v-for="project in projects"
          :key="project.id"
          :class="{ selected: project.id === props.selected?.id }"
          class="p-2 hover:bg-gray-50 cursor-pointer"
          @click="selectProject(project)"
      >
        <span>{{ project.Icon }}</span>
        <span>{{ project.Name }}</span>
      </li>
      <li>
        <div v-if="showInput">
          <input v-model="newProjectName" placeholder="Project name"/>
          <div style="display: flex; gap: 0.5rem;">
            <button @click="createProject">Create</button>
            <button @click="showInput = false">X</button>
          </div>
        </div>
        <div v-else>
          <button @click="showInput = true">New Project</button>
        </div>
      </li>
    </ul>
  </aside>
</template>

<style>
.sidebar {
  width: 75px;
  padding: 1rem;
}
</style>