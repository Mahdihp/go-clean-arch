package bybit_history_service

import (
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	"github.com/bxcodec/go-clean-arch/params"
	"github.com/bxcodec/go-clean-arch/pkg/httpmsg"
	"github.com/bxcodec/go-clean-arch/util"
	"github.com/labstack/echo/v4"
	bybit "github.com/wuhewuhe/bybit.go.api"
	"net/http"
)

type ByBitHistoricalServic struct {
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitTradeValidator
	userRepo    UserRepository
	historyRepo HistoryRepository
}

func NewByBitHistoricalServic(cfg config.Config, userRepo UserRepository, historyRepo HistoryRepository) ByBitHistoricalServic {
	return ByBitHistoricalServic{
		Config:      cfg,
		validator:   validator.NewByBitTradeValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
		userRepo:    userRepo,
		historyRepo: historyRepo,
	}
}

func (s ByBitHistoricalServic) SetRoutes(router *echo.Echo) {
	routeGroup := router.Group(string(params.History))
	routeGroup.GET("/search", s.search)

}

func (s ByBitHistoricalServic) search(ctx echo.Context) error {
	var req models.HistoryDto
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	errorList, err := s.validator.ValidateByBitHistoricalSearch(req)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return ctx.JSON(http.StatusBadRequest, httpmsg.NewStrMassage(strErrorList))
	}
	user, err := s.userRepo.FindByApiKey(ctx.Request().Context(), req.ApiKey)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, httpmsg.NewMassage(httpmsg.UserNotFound))
	}
	history, err := s.historyRepo.FindBySymbol(ctx.Request().Context(), user.ID, req.Symbol)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, httpmsg.NewMassage(httpmsg.HistoryNotFound))
	}

	//params := map[string]interface{}{"category": req.Category}
	//res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).GetOrderHistory(ctx.Request().Context())
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, err)
	//}
	return ctx.JSON(http.StatusCreated, history)
}
