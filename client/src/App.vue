<script setup lang="ts">
import { ref } from "vue";
import TheHeader from "./components/TheHeader.vue";
import OrdersTable from "./components/OrdersTable.vue";
import AllocationsList from "./components/AllocationsList.vue";

const activeTab = ref("orders");
</script>
<template>
  <div class="min-h-screen bg-gray-100">
    <TheHeader />
    <main class="container mx-auto mt-6">
      <OrdersTable v-if="activeTab === 'orders'" />
      <AllocationsList v-else-if="activeTab === 'allocations'" />
    </main>
  </div>
</template>
