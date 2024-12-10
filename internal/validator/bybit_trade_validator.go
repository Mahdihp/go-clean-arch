package validator

import (
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/order"
	"github.com/bxcodec/go-clean-arch/adapter/grpc-proto/position"
	"github.com/bxcodec/go-clean-arch/pkg/errmsg"
	"github.com/bxcodec/go-clean-arch/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ByBitTradeValidator struct {
}

func NewByBitTradeValidator() ByBitTradeValidator {
	return ByBitTradeValidator{}
}

func (v ByBitTradeValidator) ValidateCancel(req *order.CancelOrderRequest) (map[string]string, error) {
	const op = "Validator.ValidateCancel"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateAmend(req *order.AmendOrderRequest) (map[string]string, error) {
	const op = "Validator.ValidateAmend"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateCancelAll(req *order.CancelAllRequest) (map[string]string, error) {
	const op = "Validator.ValidateCreate"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateCreate(req *order.PlaceOrderRequest) (map[string]string, error) {
	const op = "Validator.ValidateCreate"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
		validation.Field(&req.TpslMode, validation.Required),
		validation.Field(&req.PositionIdx, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateTradingStop(req *position.TradingStopRequest) (map[string]string, error) {
	const op = "Validator.ValidateTradingStop"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
		validation.Field(&req.TpslMode, validation.Required),
		validation.Field(&req.PositionIdx, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateSwitchMode(req *position.SwitchModeRequest) (map[string]string, error) {
	const op = "Validator.ValidateSwitchMode"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Mode, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateSwitchIsolated(req *position.SwitchIsolatedRequest) (map[string]string, error) {
	const op = "Validator.ValidateSwitchIsolated"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
		validation.Field(&req.TradeMode, validation.Required),
		validation.Field(&req.BuyLeverage, validation.Required),
		validation.Field(&req.SellLeverage, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateSetLeverage(req *position.SetLeverageRequest) (map[string]string, error) {
	const op = "Validator.ValidateSetLeverage"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
		validation.Field(&req.Symbol, validation.Required),
		validation.Field(&req.BuyLeverage, validation.Required),
		validation.Field(&req.SellLeverage, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
func (v ByBitTradeValidator) ValidateGetPositionInfo(req *position.PositionInfoRequest) (map[string]string, error) {
	const op = "Validator.ValidateCreateRequest"
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required),
	); err != nil {

		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					//fmt.Println(key, value.Error())
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
