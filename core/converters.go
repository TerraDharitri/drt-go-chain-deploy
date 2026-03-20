package core

import (
	"encoding/hex"
	"math/big"

	drtCore "github.com/TerraDharitri/drt-go-chain-core/core"
	vmcommonBuiltInFunctions "github.com/TerraDharitri/drt-go-chain-vm-common/builtInFunctions"
	"github.com/TerraDharitri/drt-go-chain/config"
	"github.com/TerraDharitri/drt-go-chain/process"
	"github.com/TerraDharitri/drt-go-chain/process/mock"
	"github.com/TerraDharitri/drt-go-chain/process/smartContract/hooks"
	"github.com/TerraDharitri/drt-go-chain/testscommon"
	dataRetrieverMock "github.com/TerraDharitri/drt-go-chain/testscommon/dataRetriever"
	"github.com/TerraDharitri/drt-go-chain/testscommon/enableEpochsHandlerMock"
	"github.com/TerraDharitri/drt-go-chain/testscommon/epochNotifier"
	"github.com/TerraDharitri/drt-go-chain/testscommon/genericMocks"
	mockState "github.com/TerraDharitri/drt-go-chain/testscommon/state"
)

// ConvertToPositiveBigInt will try to convert the provided string to its big int corresponding value. Only
// positive numbers are allowed
func ConvertToPositiveBigInt(value string) (*big.Int, error) {
	valueNumber, isNumber := big.NewInt(0).SetString(value, 10)
	if !isNumber {
		return nil, ErrStringIsNotANumber
	}

	if valueNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, ErrNegativeValue
	}

	return valueNumber, nil
}

// GenerateSCAddress will generate the resulting SC address from the provided public key string and nonce
func GenerateSCAddress(
	pkString string,
	nonce uint64,
	vmType string,
	converter drtCore.PubkeyConverter,
) (string, error) {
	blockchainHook, err := generateBlockchainHook(converter)
	if err != nil {
		return "", err
	}

	pk, err := converter.Decode(pkString)
	if err != nil {
		return "", err
	}

	vmTypeBytes, err := hex.DecodeString(vmType)
	if err != nil {
		return "", err
	}

	scResultingAddressBytes, err := blockchainHook.NewAddress(pk, nonce, vmTypeBytes)
	if err != nil {
		return "", err
	}

	encodedAddress, err := converter.Encode(scResultingAddressBytes)
	if err != nil {
		return "", nil
	}

	return encodedAddress, nil
}

func generateBlockchainHook(converter drtCore.PubkeyConverter) (process.BlockChainHookHandler, error) {
	builtInFuncs := vmcommonBuiltInFunctions.NewBuiltInFunctionContainer()
	datapool := dataRetrieverMock.NewPoolsHolderMock()
	arg := hooks.ArgBlockChainHook{
		Accounts:              &mockState.AccountsStub{},
		PubkeyConv:            converter,
		StorageService:        genericMocks.NewChainStorerMock(0),
		DataPool:              datapool,
		BlockChain:            &testscommon.ChainHandlerMock{},
		ShardCoordinator:      mock.NewOneShardCoordinatorMock(),
		Marshalizer:           &mock.MarshalizerMock{},
		Uint64Converter:       &mock.Uint64ByteSliceConverterMock{},
		BuiltInFunctions:      builtInFuncs,
		NFTStorageHandler:     &testscommon.SimpleNFTStorageHandlerStub{},
		GlobalSettingsHandler: &testscommon.DCDTGlobalSettingsHandlerStub{},
		CompiledSCPool:        datapool.SmartContracts(),
		ConfigSCStorage:       config.StorageConfig{},
		EnableEpochs:          config.EnableEpochs{},
		EpochNotifier:         &epochNotifier.EpochNotifierStub{},
		EnableEpochsHandler:   &enableEpochsHandlerMock.EnableEpochsHandlerStub{},
		WorkingDir:            "",
		NilCompiledSCStore:    true,
		GasSchedule: &testscommon.GasScheduleNotifierMock{
			GasSchedule: make(map[string]map[string]uint64),
			LatestGasScheduleCalled: func() map[string]map[string]uint64 {
				return make(map[string]map[string]uint64)
			},
			LatestGasScheduleCopyCalled: func() map[string]map[string]uint64 {
				return make(map[string]map[string]uint64)
			},
		},
		Counter:                  &testscommon.BlockChainHookCounterStub{},
		MissingTrieNodesNotifier: &testscommon.MissingTrieNodesNotifierStub{},
		EpochStartTrigger:        &testscommon.EpochStartTriggerStub{},
		RoundHandler:             &testscommon.RoundHandlerMock{},
	}

	return hooks.NewBlockChainHookImpl(arg)
}
