package task

import "testing"

func TestCreateAndCompleteTask(t *testing.T) {
	store := NewStore()

	created := store.Create("write local AGENTS.md")
	if created.ID != 1 {
		t.Fatalf("expected id 1, got %d", created.ID)
	}

	completed, err := store.Complete(created.ID)
	if err != nil {
		t.Fatalf("complete task: %v", err)
	}
	if !completed.Completed {
		t.Fatal("expected task to be completed")
	}
}

func TestCompleteMissingTask(t *testing.T) {
	store := NewStore()

	_, err := store.Complete(99)
	if err != ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
