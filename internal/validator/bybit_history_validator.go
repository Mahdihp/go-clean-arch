package validator

import (
	"github.com/bxcodec/go-clean-arch/internal/bybit_history_service/models"
	"github.com/bxcodec/go-clean-arch/pkg/errmsg"
	"github.com/bxcodec/go-clean-arch/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ByBitHistoryValidator struct {
}

func NewByBitHistoryValidator() ByBitHistoryValidator {
	return ByBitHistoryValidator{}
}

func (v ByBitTradeValidator) ValidateByBitHistoricalSearch(req models.HistoryDto) (map[string]string, error) {
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
