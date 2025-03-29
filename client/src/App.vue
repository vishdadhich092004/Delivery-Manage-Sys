<template>
  <div class="min-h-screen bg-gray-100">
    <TheHeader v-model:activeTab="activeTab" />
    <main class="container mx-auto mt-6 px-4">
      <button
        @click="testBackendConnection"
        class="mb-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
      >
        Test Backend Connection
      </button>
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
import { testBakcend } from "./services/api";

const activeTab = ref<"orders" | "allocations">("orders");

const testBackendConnection = async () => {
  try {
    const result = await testBakcend();
    alert("Backend connection successful: " + JSON.stringify(result));
  } catch (error) {
    alert("Backend connection failed: " + error);
  }
};
</script>
