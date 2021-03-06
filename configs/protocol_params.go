// Copyright 2015 The go-ethereum Authors

package configs

import "math/big"

const (
	GasLimitBoundDivisor uint64 = 1024      // The bound divisor of the gas limit, used in update calculations.
	MinGasLimit          uint64 = 3000000   // Minimum the gas limit may ever be.
	MaxGasLimit          uint64 = 150000000 // Maximum gas limit of blocks.
	TargetGasLimit       uint64 = 47000000  // The artificial target

	MaximumExtraDataSize  uint64 = 32    // Maximum size extra data may be after Genesis.
	ExpByteGas            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	SloadGas              uint64 = 50    // Multiplied by the number of 32-byte words that are copied (round up) for any *COPY operation and added.
	CallValueTransferGas  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountGas     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxGas                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroGas         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv          uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	SstoreSetGas          uint64 = 20000 // Once per SLOAD operation.
	LogDataGas            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend           uint64 = 2300  // Free gas given at beginning of call.

	Sha3Gas          uint64 = 30    // Once per SHA3 operation.
	Sha3WordGas      uint64 = 6     // Once per word of the SHA3 operation's data.
	SstoreResetGas   uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearGas   uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundGas  uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.
	JumpdestGas      uint64 = 1     // Refunded gas, once per SSTORE operation if the zeroness changes to zero.
	EpochDuration    uint64 = 30000 // Duration between proof-of-work epochs.
	CallGas          uint64 = 40    // Once per CALL operation & message call transaction.
	CreateDataGas    uint64 = 200   //
	CallCreateDepth  uint64 = 1024  // Maximum depth of call/create stack.
	ExpGas           uint64 = 10    // Once per EXP instruction
	LogGas           uint64 = 375   // Per LOG* operation.
	CopyGas          uint64 = 3     //
	StackLimit       uint64 = 1024  // Maximum size of VM stack allowed.
	TierStepGas      uint64 = 0     // Once per operation, for a selection of them.
	LogTopicGas      uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateGas        uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	SuicideRefundGas uint64 = 24000 // Refunded following a suicide operation.
	MemoryGas        uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.
	TxDataNonZeroGas uint64 = 68    // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.

	MaxCodeSize = 24576 // Maximum bytecode to permit for a contract

	// Precompiled contract gas prices

	EcrecoverGas            uint64 = 3000   // Elliptic curve sender recovery gas price
	Sha256BaseGas           uint64 = 60     // Base price for a SHA256 operation
	Sha256PerWordGas        uint64 = 12     // Per-word price for a SHA256 operation
	Ripemd160BaseGas        uint64 = 600    // Base price for a RIPEMD160 operation
	Ripemd160PerWordGas     uint64 = 120    // Per-word price for a RIPEMD160 operation
	IdentityBaseGas         uint64 = 15     // Base price for a data copy operation
	IdentityPerWordGas      uint64 = 3      // Per-work price for a data copy operation
	ModExpQuadCoeffDiv      uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256AddGas             uint64 = 500    // Gas needed for an elliptic curve addition
	Bn256ScalarMulGas       uint64 = 40000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGas     uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGas uint64 = 80000  // Per-point price for an elliptic curve pairing check

	// CPChain primitives
	GetRankGas         uint64 = 1600 // gas price for getting rank operation, need to call contract
	GetMaintenanceGas  uint64 = 600  // gas price  needed for GetMaintenanceGas
	IsProxyGas         uint64 = 2000 // Gas needed for IsProxyGas, need to call contract twice
	GetTxVolumeGas     uint64 = 1200 // Gas needed for GetTxVolumeGas, involving many cryptographic operations
	GetUploadRewardGas uint64 = 1600 // Gas needed for GetUploadRewardGas, need to call contract
	CpuPowValidateGas  uint64 = 200  // Gas needed for CpuPowValidate, involving hash
	MemPowValidateGas  uint64 = 200  // Gas needed for MemPowValidate, involving hash
)

var (
	DurationLimit = big.NewInt(13) // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
)

const (
	Cep1BlocksPerDay = 24 * 60 * 60 / 10
	Cep1BlocksY1     = 366 * Cep1BlocksPerDay // contain a leap day
	Cep1BlocksY2     = 365 * Cep1BlocksPerDay
	Cep1BlocksY3     = 365 * Cep1BlocksPerDay
	Cep1BlocksY4     = 365 * Cep1BlocksPerDay
	Cep1BlocksY5     = 366 * Cep1BlocksPerDay // contain a leap day
)

var (
	Cep1BlockRewardSupplyY1 = new(big.Int).Mul(big.NewInt(40002336), big.NewInt(1e+18))
	Cep1BlockRewardSupplyY2 = new(big.Int).Mul(big.NewInt(29990736), big.NewInt(1e+18))
	Cep1BlockRewardSupplyY3 = new(big.Int).Mul(big.NewInt(22485168), big.NewInt(1e+18))
	Cep1BlockRewardSupplyY4 = new(big.Int).Mul(big.NewInt(16997904), big.NewInt(1e+18))
	Cep1BlockRewardSupplyY5 = new(big.Int).Mul(big.NewInt(127438272), big.NewInt(1e+17))

	// the calculation is based on 10 s a block	is generated.
	Cep1BlockRewardY1 = new(big.Int).Div(Cep1BlockRewardSupplyY1, big.NewInt(Cep1BlocksY1)) // reward 12.65 cpc per block
	Cep1BlockRewardY2 = new(big.Int).Div(Cep1BlockRewardSupplyY2, big.NewInt(Cep1BlocksY2)) // reward 9.51 cpc per block
	Cep1BlockRewardY3 = new(big.Int).Div(Cep1BlockRewardSupplyY3, big.NewInt(Cep1BlocksY3)) // reward 7.13 cpc per block
	Cep1BlockRewardY4 = new(big.Int).Div(Cep1BlockRewardSupplyY4, big.NewInt(Cep1BlocksY4)) // reward 5.39 cpc per block
	Cep1BlockRewardY5 = new(big.Int).Div(Cep1BlockRewardSupplyY5, big.NewInt(Cep1BlocksY5)) // reward 4.03 cpc per block

	Cep1LastBlockY1 = big.NewInt(Cep1BlocksY1)
	Cep1LastBlockY2 = new(big.Int).Add(big.NewInt(Cep1BlocksY2), Cep1LastBlockY1)
	Cep1LastBlockY3 = new(big.Int).Add(big.NewInt(Cep1BlocksY3), Cep1LastBlockY2)
	Cep1LastBlockY4 = new(big.Int).Add(big.NewInt(Cep1BlocksY4), Cep1LastBlockY3)
	Cep1LastBlockY5 = new(big.Int).Add(big.NewInt(Cep1BlocksY5), Cep1LastBlockY4)
)
