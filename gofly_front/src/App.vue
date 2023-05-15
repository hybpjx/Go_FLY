<script setup lang="ts">
</script>

<template>
  <div class="app-w">
    {{ app.getConfig("baseUrl") }}
  </div>
</template>

<style scoped>
.app-w {
  display: flex;
  color: red;
}
</style>
