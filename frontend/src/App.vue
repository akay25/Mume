<script setup>
import { onMounted } from 'vue';
import { useApplicationStore } from '@/stores/application';

const applicationStore = useApplicationStore();

onMounted(async () => {
  await applicationStore.fetchLibaryPath();
  console.log("woprking ", applicationStore.libraryPath);
})
</script>

<template>
  <router-view />
</template>

<style scoped></style>
