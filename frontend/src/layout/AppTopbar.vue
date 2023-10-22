<script setup>
import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
import { useLayout } from '@/layout/composables/layout';
import { useRouter } from 'vue-router';

const { layoutConfig, onMenuButtonClick } = useLayout();

const outsideClickListener = ref(null);
const topbarMenuActive = ref(false);
const router = useRouter();

onMounted(() => {
  bindOutsideClickListener();
});

onBeforeUnmount(() => {
  unbindOutsideClickListener();
});

const logoUrl = computed(() => {
  return `layout/images/${layoutConfig.darkTheme.value ? 'logo-white' : 'logo-dark'}.png`;
});

const onSettingsClick = () => {
  topbarMenuActive.value = false;
  router.push('/settings');
};
const topbarMenuClasses = computed(() => {
  return {
    'layout-topbar-menu-mobile-active': topbarMenuActive.value
  };
});

const bindOutsideClickListener = () => {
  if (!outsideClickListener.value) {
    outsideClickListener.value = (event) => {
      if (isOutsideClicked(event)) {
        topbarMenuActive.value = false;
      }
    };
    document.addEventListener('click', outsideClickListener.value);
  }
};
const unbindOutsideClickListener = () => {
  if (outsideClickListener.value) {
    document.removeEventListener('click', outsideClickListener);
    outsideClickListener.value = null;
  }
};
const isOutsideClicked = (event) => {
  if (!topbarMenuActive.value) return;

  const sidebarEl = document.querySelector('.layout-topbar-menu');
  const topbarEl = document.querySelector('.layout-topbar-menu-button');

  return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
};
</script>

<template>
  <div class="layout-topbar">
    <div class="layout-buttons-container">
      <router-link to="/" class="layout-topbar-logo">
        <img :src="logoUrl" alt="logo" />
      </router-link>

      <div class="layout-topbar-menu-section-middle">
        <router-link to="/" class="p-link layout-menu-button layout-topbar-button">
          <vue-feather type="home" size="20" stroke-width="1"></vue-feather>
        </router-link>
        <router-link to="/songs" class="p-link layout-menu-button layout-topbar-button">
          <vue-feather type="music" size="20" stroke-width="1"></vue-feather>
        </router-link>
        <router-link to="/library" class="p-link layout-menu-button layout-topbar-button">
          <vue-feather type="folder" size="20" stroke-width="1"></vue-feather>
        </router-link>
        <router-link to="/profile" class="p-link layout-menu-button layout-topbar-button">
          <vue-feather type="user" size="20" stroke-width="1"></vue-feather>
        </router-link>
      </div>

      <div class="layout-topbar-menu-section-bottom" :class="topbarMenuClasses">
        <button @click="onSettingsClick()" class="p-link layout-topbar-button">
          <vue-feather type="settings" size="20" stroke-width="1"></vue-feather>
          <span>Settings</span>
        </button>
        <button @click="closeMe()" class="p-link layout-topbar-button">
          <vue-feather type="x" size="20" stroke-width="1"></vue-feather>
          <span>Close</span>
        </button>
      </div>
    </div>
    <div class="layout-curved-edge"></div>
  </div>
</template>

<style lang="scss" scoped></style>
