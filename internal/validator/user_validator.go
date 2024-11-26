package user_validator

import (
	user_repository "github.com/bxcodec/go-clean-arch/internal/repository/account"
	"github.com/bxcodec/go-clean-arch/params"
)

type UserValidator struct {
	userRepo user_repository.AccountRepository
}

func New(userRepo user_repository.AccountRepository) UserValidator {
	return UserValidator{userRepo: userRepo}
}

func (v UserValidator) ValidateRegisterRequest(req params.UserLoginRequest) (map[string]string, error) {
	const op = "Validator.ValidateRegisterRequest"

	//if err := validation.ValidateStruct(&req,
	//	// TODO - add 3 to config
	//	validation.Field(&req.FirstName,
	//		validation.Required,
	//		validation.Length(3, 50)),
	//
	//	validation.Field(&req.LastName,
	//		validation.Required,
	//		validation.Length(3, 50)),
	//
	//	validation.Field(&req.NationalCode,
	//		validation.Required,
	//		validation.Length(1, 11)),
	//
	//	validation.Field(&req.Address,
	//		validation.Required,
	//		validation.Length(1, 500)),
	//
	//	validation.Field(&req.RoleId,
	//		validation.Required),
	//
	//	validation.Field(&req.Password,
	//		validation.Required,
	//		validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),
	//
	//	validation.Field(&req.Mobile,
	//		validation.Required,
	//		validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
	//		validation.By(v.checkPhoneNumberUniqueness)),
	//); err != nil {
	//	fieldErrors := make(map[string]string)
	//
	//	errV, ok := err.(validation.Errors)
	//	if ok {
	//		for key, value := range errV {
	//			if value != nil {
	//				fieldErrors[key] = value.Error()
	//			}
	//		}
	//	}
	//
	//	return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
	//		WithKind(richerror.KindInvalid).
	//		WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	//}

	return nil, nil
}
