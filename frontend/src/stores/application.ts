import { defineStore } from 'pinia';
import { GetLibraryPath, SetLibraryPath } from '@wailsjs';

export const useApplicationStore = defineStore('application', {
  state: () => ({
    platform: "linux",
    libraryPath: "",
  }),
  actions: {
    async fetchLibaryPath() {
      this.libraryPath = await GetLibraryPath();
    }
  }
});