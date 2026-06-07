from .store import TaskStore


def main() -> None:
    store = TaskStore()
    task = store.create("add an AGENTS.md file close to the code")
    print(f"{task.id}: {task.title}")


if __name__ == "__main__":
    main()
