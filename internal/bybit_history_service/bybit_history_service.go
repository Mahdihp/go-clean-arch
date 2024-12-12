package bybit_history_service

import (
	"github.com/bxcodec/go-clean-arch/config"
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	params2 "github.com/bxcodec/go-clean-arch/internal/bybit_history_service/params"
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
	routeGroup.POST("/search", s.search)

}

func (s ByBitHistoricalServic) search(ctx echo.Context) error {
	var req params2.HistoryDto
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	section := ctx.QueryParam("section")
	if len(section) <= 0 {
		return ctx.JSON(http.StatusCreated, httpmsg.NewMassage(httpmsg.SectionNotFound))
	}

	errorList, err := s.validator.ValidateByBitHistoricalSearch(req)
	var strErrorList = ""
	if err != nil {
		strErrorList += util.MapToString(errorList)
		return ctx.JSON(http.StatusBadRequest, httpmsg.NewStrMassage(strErrorList))
	}
	collection := models.SelectCollection(section)

	user, err := s.userRepo.FindByApiKey(ctx.Request().Context(), req.ApiKey)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, httpmsg.NewMassage(httpmsg.UserNotFound))
	}

	if len(req.Symbol) > 0 && len(req.StartTime) <= 0 || len(req.EndTime) <= 0 {
		history, err := s.historyRepo.FindBySymbol(ctx.Request().Context(), collection, user.ID.String(), req.Symbol, req.PageIndex, req.PageSize)
		if err != nil || len(history) <= 0 {
			return ctx.JSON(http.StatusNotFound, httpmsg.NewMassage(httpmsg.HistoryNotFound))
		}
		return ctx.JSON(http.StatusCreated, history)

	} else if len(req.Symbol) > 0 && len(req.StartTime) > 0 && len(req.EndTime) > 0 {
		history, err := s.historyRepo.FindByBetweenCreatedTimeAndSymbol(ctx.Request().Context(), collection,
			user.ID.String(), req.Symbol, req.StartTime, req.EndTime, req.PageIndex, req.PageSize)
		if err != nil || len(history) <= 0 {
			return ctx.JSON(http.StatusNotFound, httpmsg.NewMassage(httpmsg.HistoryNotFound))
		}
		return ctx.JSON(http.StatusCreated, history)
	}

	return ctx.JSON(http.StatusCreated, httpmsg.NewMassage(httpmsg.BadRequesy))
}
