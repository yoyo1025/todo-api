package model

type Task struct {
	userId int64
	title string
	detail string
	status int64
}

func NewTask(userId int64, title, detail string, status int64) *Task {
	return &Task{
		userId: userId,
		title: title,
		detail: detail,
		status: status,
	}
}

// モデルの内容を詰め替える
func (t *Task) Update(title, detail string, status int64) *Task {
	return &Task{
		userId: t.userId,
		title: title,
		detail: detail,
		status: status,
	}
}

func (t *Task) GetUserId() int64 {
	return t.userId
}

func (t *Task) GetTitle() string {
	return t.title
}

func (t *Task) GetDetail() string {
	return t.detail
}

func (t *Task) GetStatus() int64 {
	return t.status
}