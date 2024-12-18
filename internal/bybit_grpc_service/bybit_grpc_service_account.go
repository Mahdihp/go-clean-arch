package bybit_grpc_service

import (
	"context"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/account"
	"github.com/bxcodec/go-clean-arch/config"
	params_http "github.com/bxcodec/go-clean-arch/internal/bybit_grpc_service/params"
	"github.com/bxcodec/go-clean-arch/internal/validator"
	bybit "github.com/wuhewuhe/bybit.go.api"
)

type ByBitHttpServerAccount struct {
	account.UnimplementedAccountServicServer
	Config      config.Config
	byBitClient *bybit.Client
	validator   validator.ByBitMarketValidator
}

func NewByBitHttpServerAccount(cfg config.Config) ByBitHttpServerAccount {
	return ByBitHttpServerAccount{
		Config:      cfg,
		validator:   validator.NewByBitMarketValidator(),
		byBitClient: bybit.NewBybitHttpClient(cfg.ByBitWs.ApiKey, cfg.ByBitWs.ApiSecret, bybit.WithBaseURL(bybit.MAINNET)),
	}
}

// Get Wallet Balance
// Obtain wallet balance, query asset information of each currency.
// By default, currency information with assets or liabilities of 0 is not returned.
// https://bybit-exchange.github.io/docs/v5/account/wallet-balance
func (s *ByBitHttpServerAccount) GetWalletBalance(ctx context.Context, in *account.GetWalletBalanceRequest) (*account.GetWalletBalanceResponse, error) {

	params := map[string]interface{}{"accountType": in.AccountType, "coin": in.Coin}
	res, err := s.byBitClient.NewClassicalBybitServiceWithParams(params).GetAccountWallet(ctx)
	if err != nil {
		return nil, err
	}
	response := params_http.ToGetWalletBalanceResponse(res)
	return &response, nil
}
