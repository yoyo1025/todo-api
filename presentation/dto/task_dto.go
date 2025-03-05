package dto

import "todo-api/domain/model"

type TaskRequest struct {
  Title  string `json:"title"`
  Detail string `json:"detail"`
  Status int64  `json:"status"`
}


type TaskResponse struct {
	ID     uint  `json:"id"`
  UserID int64  `json:"user_id"`
  Title  string `json:"title"`
  Detail string `json:"detail"`
  Status int64  `json:"status"`
}

func ToTaskResponse(t *model.Task) TaskResponse {
	return TaskResponse{
		UserID: t.GetUserId(),
		Title:  t.GetTitle(),
		Detail: t.GetDetail(),
		Status: t.GetStatus(),
	}
}

// スライス変換用
func ToTaskResponses(tasks []*model.Task) []TaskResponse {
	res := make([]TaskResponse, len(tasks))
	for i, t := range tasks {
		res[i] = ToTaskResponse(t)
	}
	return res
}
