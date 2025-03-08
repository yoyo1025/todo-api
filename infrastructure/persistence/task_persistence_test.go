package persistence

import (
	"testing"

	"todo-api/database"
	"todo-api/domain/model"
	"todo-api/infrastructure/record"
)

func TestCreateTask(t *testing.T) {
	db := database.SetupTestDB()
	repo := NewTaskCommandPersistence(db)
	task := model.NewTask(1, "Test Title", "Test Detail", 0)

	if err := repo.Create(task); err != nil {
		t.Fatalf("Create task failed: %v", err)
	}

	var rec record.TaskRecord
	db.First(&rec)

	if rec.UserID != task.GetUserId() {
		t.Errorf("Expected UserID %d, got %d", task.GetUserId(), rec.UserID)
	}
	if rec.Title != task.GetTitle() {
		t.Errorf("Expected Title %s, got %s", task.GetTitle(), rec.Title)
	}
}

func TestUpdateTask(t *testing.T) {
	db := database.SetupTestDB()
	repo := NewTaskCommandPersistence(db)
	task := model.NewTask(1, "Old Title", "Old Detail", 0)
	db.Create(&record.TaskRecord{UserID: task.GetUserId(), Title: task.GetTitle(), Detail: task.GetDetail(), Status: task.GetStatus()})

	updatedTask := model.NewTask(1, "New Title", "New Detail", 1)
	if err := repo.Update(1, updatedTask); err != nil {
		t.Fatalf("Update task failed: %v", err)
	}

	var rec record.TaskRecord
	db.First(&rec)

	if rec.Title != updatedTask.GetTitle() {
		t.Errorf("Expected Title %s, got %s", updatedTask.GetTitle(), rec.Title)
	}
	if rec.Detail != updatedTask.GetDetail() {
		t.Errorf("Expected Detail %s, got %s", updatedTask.GetDetail(), rec.Detail)
	}
	if rec.Status != updatedTask.GetStatus() {
		t.Errorf("Expected Status %d, got %d", updatedTask.GetStatus(), rec.Status)
	}
}

func TestFindAllTasks(t *testing.T) {
	db := database.SetupTestDB()
	queryRepo := NewTaskQueryPersistence(db)
	db.Create(&record.TaskRecord{UserID: 1, Title: "Task 1", Detail: "Detail 1", Status: 0})
	db.Create(&record.TaskRecord{UserID: 1, Title: "Task 2", Detail: "Detail 2", Status: 1})

tasks, err := queryRepo.FindAllTask(1)
	if err != nil {
		t.Fatalf("FindAll failed: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestFindTaskById(t *testing.T) {
	db := database.SetupTestDB()
	queryRepo := NewTaskQueryPersistence(db)
	db.Create(&record.TaskRecord{UserID: 1, Title: "Task 1", Detail: "Detail 1", Status: 0})

	task, err := queryRepo.FindTaskById(1)
	if err != nil {
		t.Fatalf("FindById failed: %v", err)
	}
	if task.Title != "Task 1" {
		t.Errorf("Expected Title 'Task 1', got %s", task.Title)
	}
}
