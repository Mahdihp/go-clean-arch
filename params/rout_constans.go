package params

type Route string

const (
	BaseRout Route = "/api"
	VerRoute Route = BaseRout + "/v1"

	OrderBook Route = VerRoute + "/orderbook"
	Order     Route = VerRoute + "/order"
	History   Route = VerRoute + "/history"
)
