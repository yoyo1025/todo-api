package dto

type TaskRequest struct {
  Title  string `json:"title"`
  Detail string `json:"detail"`
  Status int64  `json:"status"`
}