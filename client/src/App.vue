<template>
  <div class="min-h-screen bg-gray-100">
    <TheHeader v-model:activeTab="activeTab" />
    <main class="container mx-auto mt-6 px-4">
      <OrdersTable v-if="activeTab === 'orders'" />
      <AllocationsList v-else-if="activeTab === 'allocations'" />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import TheHeader from "./components/TheHeader.vue";
import OrdersTable from "./components/OrdersTable.vue";
import AllocationsList from "./components/AllocationsList.vue";

const activeTab = ref<"orders" | "allocations">("orders");
</script>
