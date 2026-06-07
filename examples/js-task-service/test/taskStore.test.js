import assert from "node:assert/strict";
import test from "node:test";

import { TaskNotFoundError, TaskStore } from "../src/taskStore.js";

test("creates and completes a task", () => {
  const store = new TaskStore();

  const created = store.create("write local AGENTS.md");
  const completed = store.complete(created.id);

  assert.equal(completed.completed, true);
});

test("throws for missing tasks", () => {
  const store = new TaskStore();

  assert.throws(() => store.complete(99), TaskNotFoundError);
});
