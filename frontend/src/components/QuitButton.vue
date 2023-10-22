<script setup>
import { ref } from 'vue';
import { CloseMe } from '@wailsjs';

const showConfirmationDialog = ref(false);

const openConfirmation = () => {
  console.log("workign");
  showConfirmationDialog.value = true;
}

const cancelQuitOperation = () => {
  showConfirmationDialog.value = false;
}

const confirmQuitOperation = async () => {
  await CloseMe();
}
</script>

<template>
  <div>
    <button @click="openConfirmation" class="p-link layout-topbar-button">
      <vue-feather type="x" size="20" stroke-width="1"></vue-feather>
      <span>Close</span>
    </button>
    <Dialog header="Exit Mume" v-model:visible="showConfirmationDialog" :style="{ width: '350px' }" :modal="true">
      <div class="flex align-items-center justify-content-center">
        <i class="pi pi-exclamation-triangle mr-3" style="font-size: 2rem" />
        <span>Are you sure you want to quit?</span>
      </div>
      <template #footer>
        <Button label="No" icon="pi pi-times" @click="cancelQuitOperation" class="p-button-text" />
        <Button label="Yes" icon="pi pi-check" @click="confirmQuitOperation" class="p-button-text" autofocus />
      </template>
    </Dialog>
  </div>
</template>
