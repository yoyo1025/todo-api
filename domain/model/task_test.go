package model

import (
	"testing"
)

// TestNewTask は NewTask 関数の動作を確認するテスト
func TestNewTask(t *testing.T) {
	task := NewTask(1, "Test Title", "Test Detail", 0)

	// 各フィールドが正しくセットされているか確認
	if task.userId != 1 {
		t.Errorf("Expected userId to be 1, got %d", task.userId)
	}
	if task.title != "Test Title" {
		t.Errorf("Expected title to be 'Test Title', got %s", task.title)
	}
	if task.detail != "Test Detail" {
		t.Errorf("Expected detail to be 'Test Detail', got %s", task.detail)
	}
	if task.status != 0 {
		t.Errorf("Expected status to be 0, got %d", task.status)
	}
}

// TestUpdateTask は Task の Update メソッドが正しく動作するか確認するテスト
func TestUpdateTask(t *testing.T) {
	task := NewTask(1, "Old Title", "Old Detail", 0)
	updatedTask := task.Update("New Title", "New Detail", 1)

	// userId は変更されないことを確認
	if updatedTask.userId != 1 {
		t.Errorf("Expected userId to remain 1, got %d", updatedTask.userId)
	}
	// title, detail, status が更新されているか確認
	if updatedTask.title != "New Title" {
		t.Errorf("Expected title to be 'New Title', got %s", updatedTask.title)
	}
	if updatedTask.detail != "New Detail" {
		t.Errorf("Expected detail to be 'New Detail', got %s", updatedTask.detail)
	}
	if updatedTask.status != 1 {
		t.Errorf("Expected status to be 1, got %d", updatedTask.status)
	}
}

// TestTaskGetters は Task のゲッターメソッドが正しく動作するかを確認するテスト
func TestTaskGetters(t *testing.T) {
	task := NewTask(2, "Getter Title", "Getter Detail", 2)

	// 各ゲッターメソッドの戻り値が正しいか確認
	if task.GetUserId() != 2 {
		t.Errorf("Expected userId to be 2, got %d", task.GetUserId())
	}
	if task.GetTitle() != "Getter Title" {
		t.Errorf("Expected title to be 'Getter Title', got %s", task.GetTitle())
	}
	if task.GetDetail() != "Getter Detail" {
		t.Errorf("Expected detail to be 'Getter Detail', got %s", task.GetDetail())
	}
	if task.GetStatus() != 2 {
		t.Errorf("Expected status to be 2, got %d", task.GetStatus())
	}
}
