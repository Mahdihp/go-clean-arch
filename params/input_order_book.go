package params

type InputOrderBook struct {
	ReqID string   `json:"req_id"`
	Op    string   `json:"op"`
	Args  []string `json:"args"`
}
