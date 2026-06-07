import { TaskStore } from "./taskStore.js";

const store = new TaskStore();
const task = store.create("add an AGENTS.md file close to the code");

console.log(`${task.id}: ${task.title}`);
