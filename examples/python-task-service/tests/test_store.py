import unittest

from task_service import TaskNotFound, TaskStore


class TaskStoreTest(unittest.TestCase):
    def test_create_and_complete_task(self) -> None:
        store = TaskStore()

        created = store.create("write local AGENTS.md")
        completed = store.complete(created.id)

        self.assertTrue(completed.completed)

    def test_complete_missing_task(self) -> None:
        store = TaskStore()

        with self.assertRaises(TaskNotFound):
            store.complete(99)


if __name__ == "__main__":
    unittest.main()
