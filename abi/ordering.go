package abi

// Code autogenerated. DO NOT EDIT.

import (
	"github.com/tonkeeper/tongo/ton"
)

const (
	IUnknown ContractInterface = iota
	DedustPool
	DedustVault
	Dns
	Editable
	JettonMaster
	JettonWallet
	Locker
	LockerBill
	LockupUniversal
	LockupVesting
	MegatonfiExchange
	MegatonfiRouter
	NftAuctionV1
	NftCollection
	NftItem
	NftSale
	NftSaleV1
	NftSaleV2
	PaymentChannel
	Sbt
	StonfiPool
	StonfiRouter
	StorageContract
	StorageProvider
	SubscriptionV1
	Teleitem
	TonstakePool
	TvPool
	ValidatorController
	Wallet
	WalletHighloadV1R1
	WalletHighloadV1R2
	WalletHighloadV2
	WalletHighloadV2R1
	WalletHighloadV2R2
	WalletV1R1
	WalletV1R2
	WalletV1R3
	WalletV2R1
	WalletV2R2
	WalletV3R1
	WalletV3R2
	WalletV4R1
	WalletV4R2
	WhalesPool
)

func (c ContractInterface) String() string {
	switch c {
	case DedustPool:
		return "dedust_pool"
	case DedustVault:
		return "dedust_vault"
	case Dns:
		return "dns"
	case Editable:
		return "editable"
	case JettonMaster:
		return "jetton_master"
	case JettonWallet:
		return "jetton_wallet"
	case Locker:
		return "locker"
	case LockerBill:
		return "locker_bill"
	case LockupUniversal:
		return "lockup_universal"
	case LockupVesting:
		return "lockup_vesting"
	case MegatonfiExchange:
		return "megatonfi_exchange"
	case MegatonfiRouter:
		return "megatonfi_router"
	case NftAuctionV1:
		return "nft_auction_v1"
	case NftCollection:
		return "nft_collection"
	case NftItem:
		return "nft_item"
	case NftSale:
		return "nft_sale"
	case NftSaleV1:
		return "nft_sale_v1"
	case NftSaleV2:
		return "nft_sale_v2"
	case PaymentChannel:
		return "payment_channel"
	case Sbt:
		return "sbt"
	case StonfiPool:
		return "stonfi_pool"
	case StonfiRouter:
		return "stonfi_router"
	case StorageContract:
		return "storage_contract"
	case StorageProvider:
		return "storage_provider"
	case SubscriptionV1:
		return "subscription_v1"
	case Teleitem:
		return "teleitem"
	case TonstakePool:
		return "tonstake_pool"
	case TvPool:
		return "tv_pool"
	case ValidatorController:
		return "validator_controller"
	case Wallet:
		return "wallet"
	case WalletHighloadV1R1:
		return "wallet_highload_v1r1"
	case WalletHighloadV1R2:
		return "wallet_highload_v1r2"
	case WalletHighloadV2:
		return "wallet_highload_v2"
	case WalletHighloadV2R1:
		return "wallet_highload_v2r1"
	case WalletHighloadV2R2:
		return "wallet_highload_v2r2"
	case WalletV1R1:
		return "wallet_v1r1"
	case WalletV1R2:
		return "wallet_v1r2"
	case WalletV1R3:
		return "wallet_v1r3"
	case WalletV2R1:
		return "wallet_v2r1"
	case WalletV2R2:
		return "wallet_v2r2"
	case WalletV3R1:
		return "wallet_v3r1"
	case WalletV3R2:
		return "wallet_v3r2"
	case WalletV4R1:
		return "wallet_v4r1"
	case WalletV4R2:
		return "wallet_v4r2"
	case WhalesPool:
		return "whales_pool"
	default:
		return "unknown"
	}
}

func ContractInterfaceFromString(s string) ContractInterface {
	switch s {
	case "dedust_pool":
		return DedustPool
	case "dedust_vault":
		return DedustVault
	case "dns":
		return Dns
	case "editable":
		return Editable
	case "jetton_master":
		return JettonMaster
	case "jetton_wallet":
		return JettonWallet
	case "locker":
		return Locker
	case "locker_bill":
		return LockerBill
	case "lockup_universal":
		return LockupUniversal
	case "lockup_vesting":
		return LockupVesting
	case "megatonfi_exchange":
		return MegatonfiExchange
	case "megatonfi_router":
		return MegatonfiRouter
	case "nft_auction_v1":
		return NftAuctionV1
	case "nft_collection":
		return NftCollection
	case "nft_item":
		return NftItem
	case "nft_sale":
		return NftSale
	case "nft_sale_v1":
		return NftSaleV1
	case "nft_sale_v2":
		return NftSaleV2
	case "payment_channel":
		return PaymentChannel
	case "sbt":
		return Sbt
	case "stonfi_pool":
		return StonfiPool
	case "stonfi_router":
		return StonfiRouter
	case "storage_contract":
		return StorageContract
	case "storage_provider":
		return StorageProvider
	case "subscription_v1":
		return SubscriptionV1
	case "teleitem":
		return Teleitem
	case "tonstake_pool":
		return TonstakePool
	case "tv_pool":
		return TvPool
	case "validator_controller":
		return ValidatorController
	case "wallet":
		return Wallet
	case "wallet_highload_v1r1":
		return WalletHighloadV1R1
	case "wallet_highload_v1r2":
		return WalletHighloadV1R2
	case "wallet_highload_v2":
		return WalletHighloadV2
	case "wallet_highload_v2r1":
		return WalletHighloadV2R1
	case "wallet_highload_v2r2":
		return WalletHighloadV2R2
	case "wallet_v1r1":
		return WalletV1R1
	case "wallet_v1r2":
		return WalletV1R2
	case "wallet_v1r3":
		return WalletV1R3
	case "wallet_v2r1":
		return WalletV2R1
	case "wallet_v2r2":
		return WalletV2R2
	case "wallet_v3r1":
		return WalletV3R1
	case "wallet_v3r2":
		return WalletV3R2
	case "wallet_v4r1":
		return WalletV4R1
	case "wallet_v4r2":
		return WalletV4R2
	case "whales_pool":
		return WhalesPool
	default:
		return IUnknown
	}
}

var methodInvocationOrder = []MethodDescription{
	{
		Name:     "get_asset",
		InvokeFn: GetAsset,
	},
	{
		Name:     "get_assets",
		InvokeFn: GetAssets,
	},
	{
		Name:     "get_auction_info",
		InvokeFn: GetAuctionInfo,
	},
	{
		Name:     "get_authority_address",
		InvokeFn: GetAuthorityAddress,
	},
	{
		Name:     "get_balances",
		InvokeFn: GetBalances,
	},
	{
		Name:     "get_bill_amount",
		InvokeFn: GetBillAmount,
	},
	{
		Name:     "get_channel_data",
		InvokeFn: GetChannelData,
	},
	{
		Name:     "get_collection_data",
		InvokeFn: GetCollectionData,
	},
	{
		Name:     "get_domain",
		InvokeFn: GetDomain,
	},
	{
		Name:     "get_editor",
		InvokeFn: GetEditor,
	},
	{
		Name:     "get_full_domain",
		InvokeFn: GetFullDomain,
	},
	{
		Name:     "get_jetton_data",
		InvokeFn: GetJettonData,
	},
	{
		Name:     "get_last_fill_up_time",
		InvokeFn: GetLastFillUpTime,
	},
	{
		Name:     "get_locker_bill_data",
		InvokeFn: GetLockerBillData,
	},
	{
		Name:     "get_locker_data",
		InvokeFn: GetLockerData,
	},
	{
		Name:     "get_lockup_data",
		InvokeFn: GetLockupData,
	},
	{
		Name:     "get_lp_data",
		InvokeFn: GetLpData,
	},
	{
		Name:     "get_lp_mining_data",
		InvokeFn: GetLpMiningData,
	},
	{
		Name:     "get_lp_swap_data",
		InvokeFn: GetLpSwapData,
	},
	{
		Name:     "get_members_raw",
		InvokeFn: GetMembersRaw,
	},
	{
		Name:     "get_mining_data",
		InvokeFn: GetMiningData,
	},
	{
		Name:     "get_next_proof_info",
		InvokeFn: GetNextProofInfo,
	},
	{
		Name:     "get_nft_data",
		InvokeFn: GetNftData,
	},
	{
		Name:     "get_params",
		InvokeFn: GetParams,
	},
	{
		Name:     "get_plugin_list",
		InvokeFn: GetPluginList,
	},
	{
		Name:     "get_pool_data",
		InvokeFn: GetPoolData,
	},
	{
		Name:     "get_pool_full_data",
		InvokeFn: GetPoolFullData,
	},
	{
		Name:     "get_pool_status",
		InvokeFn: GetPoolStatus,
	},
	{
		Name:     "get_public_key",
		InvokeFn: GetPublicKey,
	},
	{
		Name:     "get_reserves",
		InvokeFn: GetReserves,
	},
	{
		Name:     "get_revoked_time",
		InvokeFn: GetRevokedTime,
	},
	{
		Name:     "get_router_data",
		InvokeFn: GetRouterData,
	},
	{
		Name:     "get_sale_data",
		InvokeFn: GetSaleData,
	},
	{
		Name:     "get_staking_status",
		InvokeFn: GetStakingStatus,
	},
	{
		Name:     "get_storage_contract_data",
		InvokeFn: GetStorageContractData,
	},
	{
		Name:     "get_storage_params",
		InvokeFn: GetStorageParams,
	},
	{
		Name:     "get_subscription_data",
		InvokeFn: GetSubscriptionData,
	},
	{
		Name:     "get_subwallet_id",
		InvokeFn: GetSubwalletId,
	},
	{
		Name:     "get_telemint_auction_config",
		InvokeFn: GetTelemintAuctionConfig,
	},
	{
		Name:     "get_telemint_auction_state",
		InvokeFn: GetTelemintAuctionState,
	},
	{
		Name:     "get_telemint_token_name",
		InvokeFn: GetTelemintTokenName,
	},
	{
		Name:     "get_torrent_hash",
		InvokeFn: GetTorrentHash,
	},
	{
		Name:     "get_validator_controller_data",
		InvokeFn: GetValidatorControllerData,
	},
	{
		Name:     "get_wallet_data",
		InvokeFn: GetWalletData,
	},
	{
		Name:     "get_wallet_params",
		InvokeFn: GetWalletParams,
	},
	{
		Name:     "is_active",
		InvokeFn: IsActive,
	},
	{
		Name:     "list_nominators",
		InvokeFn: ListNominators,
	},
	{
		Name:     "list_votes",
		InvokeFn: ListVotes,
	},
	{
		Name:     "royalty_params",
		InvokeFn: RoyaltyParams,
	},
	{
		Name:     "seqno",
		InvokeFn: Seqno,
	},
}

var contractInterfacesOrder = []InterfaceDescription{
	{
		Name: DedustPool,
		Results: []string{
			"GetReserves_DedustResult",
			"GetAssets_DedustResult",
		},
	},
	{
		Name: DedustVault,
		Results: []string{
			"GetAsset_DedustResult",
		},
	},
	{
		Name: JettonMaster,
		Results: []string{
			"GetJettonDataResult",
		},
	},
	{
		Name: JettonWallet,
		Results: []string{
			"GetWalletDataResult",
		},
	},
	{
		Name: TonstakePool,
		Results: []string{
			"GetPoolFullDataResult",
		},
	},
	{
		Name: ValidatorController,
		Results: []string{
			"GetValidatorControllerDataResult",
		},
	},
	{
		Name: Locker,
		Results: []string{
			"GetLockerDataResult",
		},
	},
	{
		Name: LockerBill,
		Results: []string{
			"GetLockerBillDataResult",
		},
	},
	{
		Name: LockupVesting,
		Results: []string{
			"GetLockupDataResult",
		},
	},
	{
		Name: LockupUniversal,
		Results: []string{
			"GetBalancesResult",
		},
	},
	{
		Name: MegatonfiRouter,
		Results: []string{
			"GetMiningData_MegatonResult",
			"GetLpData_MegatonResult",
		},
	},
	{
		Name: MegatonfiExchange,
		Results: []string{
			"GetLpSwapData_MegatonResult",
			"GetLpMiningData_MegatonResult",
		},
	},
	{
		Name: NftSaleV1,
		Results: []string{
			"GetSaleData_BasicResult",
		},
	},
	{
		Name: NftSaleV2,
		Results: []string{
			"GetSaleData_GetgemsResult",
		},
	},
	{
		Name: NftAuctionV1,
		Results: []string{
			"GetSaleData_GetgemsAuctionResult",
		},
	},
	{
		Name: NftCollection,
		Results: []string{
			"GetCollectionDataResult",
		},
	},
	{
		Name: NftItem,
		Results: []string{
			"GetNftDataResult",
		},
	},
	{
		Name: Editable,
		Results: []string{
			"GetEditorResult",
		},
	},
	{
		Name: Sbt,
		Results: []string{
			"GetAuthorityAddressResult",
		},
	},
	{
		Name: PaymentChannel,
		Results: []string{
			"GetChannelDataResult",
		},
	},
	{
		Name: StonfiPool,
		Results: []string{
			"GetPoolData_StonfiResult",
		},
	},
	{
		Name: StonfiRouter,
		Results: []string{
			"GetRouterData_StonfiResult",
		},
	},
	{
		Name: StorageProvider,
		Results: []string{
			"GetWalletParamsResult",
			"GetStorageParamsResult",
			"SeqnoResult",
			"GetPublicKeyResult",
		},
	},
	{
		Name: StorageContract,
		Results: []string{
			"GetStorageContractDataResult",
			"GetTorrentHashResult",
			"IsActiveResult",
			"GetNextProofInfoResult",
		},
	},
	{
		Name: SubscriptionV1,
		Results: []string{
			"GetSubscriptionDataResult",
		},
	},
	{
		Name: Teleitem,
		Results: []string{
			"GetTelemintTokenNameResult",
		},
	},
	{
		Name: TvPool,
		Results: []string{
			"GetPoolData_TfResult",
			"ListNominatorsResult",
			"ListVotesResult",
		},
	},
	{
		Name: WhalesPool,
		Results: []string{
			"GetStakingStatusResult",
			"GetPoolStatusResult",
			"GetParams_WhalesNominatorResult",
		},
	},
}

func (c ContractInterface) recursiveImplements(other ContractInterface) bool {
	switch c {
	case NftAuctionV1:
		return NftSale.Implements(other)
	case NftSaleV1:
		return NftSale.Implements(other)
	case NftSaleV2:
		return NftSale.Implements(other)
	case WalletHighloadV1R1:
		return Wallet.Implements(other)
	case WalletHighloadV1R2:
		return Wallet.Implements(other)
	case WalletHighloadV2:
		return Wallet.Implements(other)
	case WalletHighloadV2R1:
		return Wallet.Implements(other)
	case WalletHighloadV2R2:
		return Wallet.Implements(other)
	case WalletV1R1:
		return Wallet.Implements(other)
	case WalletV1R2:
		return Wallet.Implements(other)
	case WalletV1R3:
		return Wallet.Implements(other)
	case WalletV2R1:
		return Wallet.Implements(other)
	case WalletV2R2:
		return Wallet.Implements(other)
	case WalletV3R1:
		return Wallet.Implements(other)
	case WalletV3R2:
		return Wallet.Implements(other)
	case WalletV4R1:
		return Wallet.Implements(other)
	case WalletV4R2:
		return Wallet.Implements(other)
	}
	return false
}

var knownContracts = map[ton.Bits256]knownContractDescription{
	ton.MustParseHash("0dceed21269d66013e95b19fbb5c55a6f01adad40837baa8e521cde3a02aa46c"): {
		contractInterfaces: []ContractInterface{WalletHighloadV1R2},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("203dd4f358adb49993129aa925cac39916b68a0e4f78d26e8f2c2b69eafa5679"): {
		contractInterfaces: []ContractInterface{WalletHighloadV2R2},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("587cc789eff1c84f46ec3797e45fc809a14ff5ae24f1e0c7a6a99cc9dc9061ff"): {
		contractInterfaces: []ContractInterface{WalletV1R3},
		getMethods: []InvokeFn{
			Seqno,
			GetPublicKey,
		},
	},
	ton.MustParseHash("5c9a5e68c108e18721a07c42f9956bfb39ad77ec6d624b60c576ec88eee65329"): {
		contractInterfaces: []ContractInterface{WalletV2R1},
		getMethods: []InvokeFn{
			Seqno,
		},
	},
	ton.MustParseHash("64dd54805522c5be8a9db59cea0105ccf0d08786ca79beb8cb79e880a8d7322d"): {
		contractInterfaces: []ContractInterface{WalletV4R1},
		getMethods: []InvokeFn{
			GetPluginList,
			GetPublicKey,
			Seqno,
			GetSubwalletId,
		},
	},
	ton.MustParseHash("84dafa449f98a6987789ba232358072bc0f76dc4524002a5d0918b9a75d2d599"): {
		contractInterfaces: []ContractInterface{WalletV3R2},
		getMethods: []InvokeFn{
			GetPublicKey,
			Seqno,
		},
	},
	ton.MustParseHash("8ceb45b3cd4b5cc60eaae1c13b9c092392677fe536b2e9b2d801b62eff931fe1"): {
		contractInterfaces: []ContractInterface{WalletHighloadV2R1},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("9494d1cc8edf12f05671a1a9ba09921096eb50811e1924ec65c3c629fbb80812"): {
		contractInterfaces: []ContractInterface{WalletHighloadV2},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("a0cfc2c48aee16a271f2cfc0b7382d81756cecb1017d077faaab3bb602f6868c"): {
		contractInterfaces: []ContractInterface{WalletV1R1},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("b61041a58a7980b946e8fb9e198e3c904d24799ffa36574ea4251c41a566f581"): {
		contractInterfaces: []ContractInterface{WalletV3R1},
		getMethods: []InvokeFn{
			Seqno,
		},
	},
	ton.MustParseHash("d4902fcc9fad74698fa8e353220a68da0dcf72e32bcb2eb9ee04217c17d3062c"): {
		contractInterfaces: []ContractInterface{WalletV1R2},
		getMethods: []InvokeFn{
			Seqno,
		},
	},
	ton.MustParseHash("d8cdbbb79f2c5caa677ac450770be0351be21e1250486de85cc52aa33dd16484"): {
		contractInterfaces: []ContractInterface{WalletHighloadV1R1},
		getMethods:         []InvokeFn{},
	},
	ton.MustParseHash("fe9530d3243853083ef2ef0b4c2908c0abf6fa1c31ea243aacaa5bf8c7d753f1"): {
		contractInterfaces: []ContractInterface{WalletV2R2},
		getMethods: []InvokeFn{
			Seqno,
			GetPublicKey,
		},
	},
	ton.MustParseHash("feb5ff6820e2ff0d9483e7e0d62c817d846789fb4ae580c878866d959dabd5c0"): {
		contractInterfaces: []ContractInterface{WalletV4R2},
		getMethods: []InvokeFn{
			GetPluginList,
			GetPublicKey,
			Seqno,
			GetSubwalletId,
		},
	},
}

func (c ContractInterface) IntMsgs() []msgDecoderFunc {
	return nil
}
