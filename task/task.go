package task

type Task struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"` // 1 for high or  0 for low
	Retries  int    `json:"retries"`
	Delay    int    `json:"delay"`
}
