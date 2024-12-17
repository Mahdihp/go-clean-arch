package params

const (
	BaseRout string = "/api"
	VerRoute string = BaseRout + "/v1"

	OrderBook string = VerRoute + "/orderbook"
	Order     string = VerRoute + "/order"
	History   string = VerRoute + "/history"

	Market_All            string = "all"
	Market_RiskLimit      string = "risklimit"
	Market_InstrumentInfo string = "instrumentinfo"
	Market_Ticker         string = "ticker"

	Market_Spot    string = "spot"
	Market_Linear  string = "linear"
	Market_Inverse string = "inverse"
	Market_Option  string = "option"
)

const (
	And_Opt              string = "$and" // and
	Or_Opt               string = "$or"  // or
	Equal_Opt            string = "$eq"  // ==
	GreaterThan_Opt      string = "$gt"  // >
	GreaterThanEqual_Opt string = "$gte" // >=
	LessThan_Opt         string = "$lt"  // <
	LessThanEqual_Opt    string = "$lte" // <=

	Field_Search_ID        string = "id"
	Field_Search_Symbol    string = "symbol"
	Field_Search_ApiKey    string = "api_key"
	Field_Search_UserId    string = "user_id"
	Field_Search_CreatedAt string = "created_at"
	Field_Search_Category  string = "category"
)
