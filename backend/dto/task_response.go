package dto

type TaskProgress struct {
	Elapsed  float64   `json:"elapsed"` // 全体経過時間
	Progress []float64 `json:"progress"` // 各リクエストの完了秒
}

type RunTaskResponse struct {
	Count     int          `json:"count"`
	Serial    TaskProgress `json:"serial"`
	Parallel  TaskProgress `json:"parallel"`
}