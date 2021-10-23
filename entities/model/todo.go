package model

// Todo Specificationの定義（例…LimitDateが登録日（登録した日）より前の場合エラーを返す）
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	LimitDate string `json:"limitDate"`
	Status    bool   `json:"status"`
}
