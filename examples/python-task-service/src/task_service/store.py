from dataclasses import dataclass


class TaskNotFound(Exception):
    """Raised when a task ID does not exist in the store."""


@dataclass(frozen=True)
class Task:
    id: int
    title: str
    completed: bool = False


class TaskStore:
    def __init__(self) -> None:
        self._next_id = 1
        self._tasks: dict[int, Task] = {}

    def create(self, title: str) -> Task:
        task = Task(id=self._next_id, title=title)
        self._tasks[task.id] = task
        self._next_id += 1
        return task

    def complete(self, task_id: int) -> Task:
        task = self._tasks.get(task_id)
        if task is None:
            raise TaskNotFound(f"task {task_id} was not found")
        completed = Task(id=task.id, title=task.title, completed=True)
        self._tasks[task_id] = completed
        return completed

    def list(self) -> list[Task]:
        return list(self._tasks.values())
