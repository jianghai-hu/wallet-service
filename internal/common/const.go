package common

//nolint:revive,stylecheck // fix in future
const (
	ID_GENERATOR_MACHINE_ID = 911
	DEFAULT_WALLET_EXT_INFO = "This is wallet detail info"

	MONEY_ACTION_TYPE_MONEY_IN  = 0
	MONEY_ACTION_TYPE_MONEY_OUT = 1

	ORDER_TYPE_DEPOSIT  = 0
	ORDER_TYPE_WITHDRAW = 1
	ORDER_TYPE_TRANSFER = 2

	TRANSACTION_STATUS_PENDING  = 0
	TRANSACTION_STATUS_COMPLETE = 1
	TRANSACTION_STATUS_FAILED   = 2
	TRANSACTION_STATUS_CANCELED = 3

	COMMAND_DEPOSIT  = "/api/v1/wallet/deposit"
	COMMAND_WITHDRAW = "/api/v1/wallet/withdraw"
	COMMAND_TRANSFER = "/api/v1/wallet/transfer"

	LOG_PATH = "./log"
)
