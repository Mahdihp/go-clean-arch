// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain/ent/account"
	"github.com/bxcodec/go-clean-arch/domain/ent/schema"
	"github.com/bxcodec/go-clean-arch/domain/ent/wallet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescUsername is the schema descriptor for username field.
	accountDescUsername := accountFields[1].Descriptor()
	// bybit_ws.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	account.UsernameValidator = accountDescUsername.Validators[0].(func(string) error)
	// accountDescPassword is the schema descriptor for password field.
	accountDescPassword := accountFields[2].Descriptor()
	// bybit_ws.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	account.PasswordValidator = accountDescPassword.Validators[0].(func(string) error)
	// accountDescSalt is the schema descriptor for salt field.
	accountDescSalt := accountFields[3].Descriptor()
	// bybit_ws.SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	account.SaltValidator = accountDescSalt.Validators[0].(func(string) error)
	// accountDescDisplayName is the schema descriptor for display_name field.
	accountDescDisplayName := accountFields[4].Descriptor()
	// bybit_ws.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	account.DisplayNameValidator = accountDescDisplayName.Validators[0].(func(string) error)
	// accountDescIsActive is the schema descriptor for is_active field.
	accountDescIsActive := accountFields[5].Descriptor()
	// bybit_ws.DefaultIsActive holds the default value on creation for the is_active field.
	account.DefaultIsActive = accountDescIsActive.Default.(bool)
	// accountDescDeleted is the schema descriptor for deleted field.
	accountDescDeleted := accountFields[6].Descriptor()
	// bybit_ws.DefaultDeleted holds the default value on creation for the deleted field.
	account.DefaultDeleted = accountDescDeleted.Default.(bool)
	// accountDescCreatedAt is the schema descriptor for created_at field.
	accountDescCreatedAt := accountFields[7].Descriptor()
	// bybit_ws.DefaultCreatedAt holds the default value on creation for the created_at field.
	account.DefaultCreatedAt = accountDescCreatedAt.Default.(func() time.Time)
	// accountDescUpdatedAt is the schema descriptor for updated_at field.
	accountDescUpdatedAt := accountFields[8].Descriptor()
	// bybit_ws.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	account.DefaultUpdatedAt = accountDescUpdatedAt.Default.(time.Time)
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescName is the schema descriptor for name field.
	walletDescName := walletFields[1].Descriptor()
	// wallet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	wallet.NameValidator = walletDescName.Validators[0].(func(string) error)
	// walletDescIsActive is the schema descriptor for is_active field.
	walletDescIsActive := walletFields[2].Descriptor()
	// wallet.DefaultIsActive holds the default value on creation for the is_active field.
	wallet.DefaultIsActive = walletDescIsActive.Default.(bool)
	// walletDescDeleted is the schema descriptor for deleted field.
	walletDescDeleted := walletFields[3].Descriptor()
	// wallet.DefaultDeleted holds the default value on creation for the deleted field.
	wallet.DefaultDeleted = walletDescDeleted.Default.(bool)
	// walletDescCreatedAt is the schema descriptor for created_at field.
	walletDescCreatedAt := walletFields[4].Descriptor()
	// wallet.DefaultCreatedAt holds the default value on creation for the created_at field.
	wallet.DefaultCreatedAt = walletDescCreatedAt.Default.(func() time.Time)
	// walletDescUpdatedAt is the schema descriptor for updated_at field.
	walletDescUpdatedAt := walletFields[5].Descriptor()
	// wallet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	wallet.DefaultUpdatedAt = walletDescUpdatedAt.Default.(time.Time)
}
