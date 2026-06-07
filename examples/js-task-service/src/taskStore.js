export class TaskNotFoundError extends Error {
  constructor(id) {
    super(`task ${id} was not found`);
    this.name = "TaskNotFoundError";
  }
}

export class TaskStore {
  #nextId = 1;
  #tasks = new Map();

  create(title) {
    const task = { id: this.#nextId, title, completed: false };
    this.#tasks.set(task.id, task);
    this.#nextId += 1;
    return task;
  }

  complete(id) {
    const task = this.#tasks.get(id);
    if (!task) {
      throw new TaskNotFoundError(id);
    }
    const completed = { ...task, completed: true };
    this.#tasks.set(id, completed);
    return completed;
  }

  list() {
    return [...this.#tasks.values()];
  }
}
