package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-clean-arch/internal/bybit_history_service/params"
	"go-clean-arch/pkg/errmsg"
	"go-clean-arch/pkg/richerror"
)

type ByBitHistoryValidator struct {
}

func NewByBitHistoryValidator() ByBitHistoryValidator {
	return ByBitHistoryValidator{}
}

func (v ByBitTradeValidator) ValidateByBitHistoricalSearch(req params.HistoryDto) (map[string]string, error) {
	const op = "Validator.ValidateByBitHistoricalSearch"

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
