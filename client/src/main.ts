// src/main.ts
import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import "./style.css";

const app = createApp(App);
const pinia = createPinia();

// Make sure to use pinia before mounting
app.use(pinia);
app.mount("#app");
