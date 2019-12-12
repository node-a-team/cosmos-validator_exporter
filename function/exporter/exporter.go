package exporter

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	t "github.com/node-a-team/cosmos-validator_exporter/types"

	block "github.com/node-a-team/cosmos-validator_exporter/function/block"
	commit "github.com/node-a-team/cosmos-validator_exporter/function/commit"
	minting "github.com/node-a-team/cosmos-validator_exporter/function/minting"
	keyutil "github.com/node-a-team/cosmos-validator_exporter/function/keyutil"
	staking "github.com/node-a-team/cosmos-validator_exporter/function/staking"
	utils "github.com/node-a-team/cosmos-validator_exporter/function/utils"
	validators "github.com/node-a-team/cosmos-validator_exporter/function/validators"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	denomList           = []string{"uatom"}
	gaugesNamespaceList = [...]string{"blockHeight", "currentBlockTime", "precommitRate", "blocksPerYear", "defaultBlockTime", "inflationRateChange", "inflationMax", "inflationMin", "inflationGoalBonded", "defaultBlockTimeInflation", "currentBlockTimeInflation", "proposerWalletAccountNumber", "validatorCount", "notBondedTokens", "bondedTokens", "totalBondedTokens", "bondedRate", "validatorCommitStatus", "proposerPriorityValue", "proposerPriority", "proposingStatus", "votingPower", "delegatorShares", "delegationRatio", "delegatorCount", "selfDelegationAmount", "commissionRate", "commissionMaxRate", "commissionMaxChangeRate", "minSelfDelegation", "jailed"}

	contentsColorInit string = "\033[0m"
)

func newGauge(nameSpace string, name string, help string) prometheus.Gauge {
	result := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "" + nameSpace,
			Name:      "" + name,
			Help:      "" + help,
		})

	return result
}

func Exporter() {


	var gauges []prometheus.Gauge = make([]prometheus.Gauge, len(gaugesNamespaceList))
	var gaugesDenom []prometheus.Gauge = make([]prometheus.Gauge, len(denomList)*3)

	for i := 0; i < len(gaugesNamespaceList); i++ {
		gauges[i] = newGauge("Cosmos", gaugesNamespaceList[i], "")
		prometheus.MustRegister(gauges[i])
	}

	count := 0
	for i := 0; i < len(denomList)*3; i += 3 {
		gaugesDenom[i] = newGauge("Cosmos_rewards", denomList[count], "")
		gaugesDenom[i+1] = newGauge("Cosmos_commission", denomList[count], "")
		gaugesDenom[i+2] = newGauge("Cosmos_balances", denomList[count], "")
		prometheus.MustRegister(gaugesDenom[i])
		prometheus.MustRegister(gaugesDenom[i+1])
		prometheus.MustRegister(gaugesDenom[i+2])

		count++
	}

	gaugesForLabel := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "Cosmos",
			Name:      "labels",
			Help:      "",
		},
		[]string{"chainId", "moniker", "validatorPubKey", "operatorAddress", "accountAddress", "consHexAddress"},
	)
	prometheus.MustRegister(gaugesForLabel)

	// for csv file export
	// Run once at first, then run once every 10000block
	fileExportChecker := 0
	baseBlockForFileExport := 10000

	for {

		blockStatus := block.BlockStatus()

		currentBlockHeight, _ := strconv.Atoi(blockStatus.Block.Header.Height)
//		consensusStatus := consensus.ConsensusStatus()
		commitStatus := commit.CommitStatus(currentBlockHeight)
		mintingParamsStatus := minting.MintingParamsStatus()



		// validators
		validatorCountOrigin, validatorsetsStatus := validators.ValidatorsetsStatus()
		validatorsStatus := validators.ValidatorsStatus()

		// block
		chainId := blockStatus.Block.Header.Chain_id
		blockTime := blockStatus.Block.Header.Time.Format("060102 15:04:05")
		blockHeight := utils.StringToFloat64(blockStatus.Block.Header.Height)
		currentBlockTime := block.CalcBlockTime(blockStatus)

		// minting
		blocksPerYear := utils.StringToFloat64(mintingParamsStatus.Blocks_per_year)
		defaultBlockTime := (60*60*24*365.25)/blocksPerYear
		inflationRateChange := utils.StringToFloat64(mintingParamsStatus.Inflation_rate_change)
		inflationMax := utils.StringToFloat64(mintingParamsStatus.Inflation_max)
		inflationMin := utils.StringToFloat64(mintingParamsStatus.Inflation_min)
		inflationGoalBonded := utils.StringToFloat64(mintingParamsStatus.Goal_bonded)
		defaultBlockTimeInflation := minting.GetInflation()
		currentBlockTimeInflation := minting.GetInflation()*(defaultBlockTime/currentBlockTime)

		// commit
		precommitRate := commit.PrecommitRate(commitStatus) * 100
		validatorCount := float64(validatorCountOrigin)
		proposerConsHexAddress := blockStatus.Block.Header.Proposer_address
		proposerMoniker := block.ProposerMoniker(blockStatus.Block.Header.Proposer_address, validatorsetsStatus, validatorsStatus)
		proposerWalletAccountNumber := float64(0.0)


		// staking
		notBondedTokensOrigin, bondedTokensOrigin := staking.GetStakingPool()
		totalBondedTokensOrigin := notBondedTokensOrigin + bondedTokensOrigin

		notBondedTokens := notBondedTokensOrigin / math.Pow10(6)
		bondedTokens := bondedTokensOrigin / math.Pow10(6)
		totalBondedTokens := totalBondedTokensOrigin / math.Pow10(6)
		bondedRate := bondedTokensOrigin / totalBondedTokensOrigin


		// csv file export(validatorsAccountNumber)
		if fileExportChecker == 0 || int(blockHeight)%baseBlockForFileExport == 0 {
			validators.ValidatorsAccountNumber(blockHeight, validatorsetsStatus, validatorsStatus)
			fileExportChecker++
		}

		// sorting
		keys := []string{}

		for key := range validatorsetsStatus {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			a, _ := strconv.Atoi(validatorsetsStatus[keys[i]][3])
			b, _ := strconv.Atoi(validatorsetsStatus[keys[j]][3])
			return a < b
		})





		// validator_pubkey: gaiad tendermint show-validator -> priv_validator_key.json
		for _, validatorPubKey := range keys {

			// validatorsStatus#1
			operatorAddress := validatorsStatus[validatorPubKey][1]

			// our validator
			if operatorAddress == t.OperatorAddr {

				// get proposerWalletAccountNumber
				for _, proposerValidatorPubKey := range keys {

					operatorAddress := validatorsStatus[proposerValidatorPubKey][1]
					consBech32Address := validatorsetsStatus[proposerValidatorPubKey][0]
					consHexAddress := keyutil.RunFromBech32(consBech32Address)

					if proposerConsHexAddress == consHexAddress {
						accountAddress := keyutil.OperAddrToOtherAddr(operatorAddress)[0]
						_, proposerWalletAccountNumber = validators.ValidatorAccount(accountAddress)
					}
				}

				// validatorsetsStatus
				consBech32Address := validatorsetsStatus[validatorPubKey][0]
				votingPower := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][1])
				proposerPriorityValue := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][2])
				proposerPriority := utils.StringToFloat64(validatorsetsStatus[validatorPubKey][3])

				// validatorsStatus#2
				moniker := validatorsStatus[validatorPubKey][0]

				jailed := utils.BoolStringToFloat64(validatorsStatus[validatorPubKey][2])
				//				tokens := utils.StringToFloat64(validatorsStatus[validatorPubKey][3])/math.Pow10(6)
				delegatorShares := utils.StringToFloat64(validatorsStatus[validatorPubKey][4]) / math.Pow10(6)
				commissionRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][5])
				commissionMaxRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][6])
				commissionMaxChangeRate := utils.StringToFloat64(validatorsStatus[validatorPubKey][7])
				//				commission_updateTime := validatorsStatus[validatorPubKey][8]
				minSelfDelegation := utils.StringToFloat64(validatorsStatus[validatorPubKey][9])
				//				unbonding_height := validatorsStatus[validatorPubKey][10]
				//				unbonding_time := validatorsStatus[validatorPubKey][11]
				//				identity := validatorsStatus[validatorPubKey][12]
				//				websote := validatorsStatus[validatorPubKey][13]
				//				details := validatorsStatus[validatorPubKey][14]

				// keyutil
				accountAddress := keyutil.OperAddrToOtherAddr(operatorAddress)[0]
				consHexAddress := keyutil.RunFromBech32(consBech32Address)

				// etc
				proposingStatus := float64(utils.GetPoposingCheck(blockStatus.Block.Header.Proposer_address, consHexAddress))
				delegatorCount, selfDelegationAmountOrigin := validators.ValidatorDelegatorNumber(operatorAddress, accountAddress)
				delegationRatio := delegatorShares / bondedTokens
				selfDelegationAmount := selfDelegationAmountOrigin / math.Pow10(6)
				validatorCommitStatus := float64(commit.ValidatorPrecommitStatus(commitStatus, consHexAddress))
				rewards, commission := validators.ValidatorRewards(operatorAddress)
				balances, walletAccountNumber := validators.ValidatorAccount(accountAddress)

				// print
				if t.OutputPrint {
					fmt.Printf("\n\n\033[1m\033[7m\033[32m[ ############ Chain_id: %s ############ ]\n\n"+contentsColorInit, chainId)
					fmt.Printf("\033[1m> Height: \033[32m%0.0f\n"+contentsColorInit, blockHeight)

					fmt.Printf("  - Time: %s UTC\n", blockTime)
					fmt.Printf("  - BlockTime: %0.2fs\n", currentBlockTime)

					fmt.Printf("  - Proposer: %s(%s)\n", proposerMoniker, proposerConsHexAddress)
					fmt.Printf("  - PrecommitRate: %f\n", precommitRate)

					fmt.Println("\n  - notBondedTokens: ", notBondedTokens)
					fmt.Println("  - bondedTokens: ", bondedTokens)
					fmt.Println("  - totalTokens: ", totalBondedTokens)
					fmt.Println("  - bondedRate: ", bondedTokens/totalBondedTokens)
					fmt.Println("  - blocksPerYear: ", blocksPerYear)
					fmt.Println("  - defaultBlockTime: ", defaultBlockTime)
					fmt.Println("  - inflationRateChange: ", inflationRateChange)
					fmt.Println("  - inflationMax: ", inflationMax)
					fmt.Println("  - inflationMin: ", inflationMin)
					fmt.Println("  - inflationGoalBonded : ", inflationGoalBonded )
					fmt.Println("  - defaultBlockTimeInflation: ", defaultBlockTimeInflation)
					fmt.Println("  - currentBlockTimeInflation: ", currentBlockTimeInflation)

					fmt.Printf("\n\n\033[1m> Moniker: \033[33m%s\n"+contentsColorInit, moniker)
					fmt.Println("  - validatorCount: ", validatorCount)
					fmt.Println("  - validatorPubKey: ", validatorPubKey)
					fmt.Println("  - operatorAddress: ", operatorAddress)
					fmt.Println("  - accountAddress: ", accountAddress)
					fmt.Println("  - consBech32Address: ", consBech32Address)
					fmt.Println("  - consHexAddress: ", consHexAddress)

					fmt.Println("\n  - validatorCommitStatus: ", validatorCommitStatus)
					fmt.Println("  - proposerPriorityValue: ", proposerPriorityValue)
					fmt.Println("  - proposerPriority: ", proposerPriority)
					fmt.Println("  - proposingStatus: ", proposingStatus)
					fmt.Printf("  - votingPower: %f\n", votingPower)
					fmt.Println("  - jailed: ", jailed)
					//				fmt.Println("  - tokens: ", tokens)
					fmt.Println("  - delegatorShares: ", delegatorShares)
					fmt.Printf("  - delegationRatio: %0.4f\n", delegationRatio)
					fmt.Println("  - delegatorCount: ", delegatorCount)
					fmt.Println("  - selfDelegationAmount: ", selfDelegationAmount)
					fmt.Printf("  - commissionRate: %0.4f\n", commissionRate)
					fmt.Printf("  - commissionMaxRate: %0.4f\n", commissionMaxRate)
					fmt.Printf("  - commissionMaxChangeRate: %0.4f\n", commissionMaxChangeRate)
					fmt.Println("  - minSelfDelegation: ", minSelfDelegation)
					fmt.Println("  - walletAccountNumber: ", walletAccountNumber)
				}
				count := 0
				for i := 0; i < len(denomList)*3; i += 3 {
					gaugesDenom[i].Set(utils.GetAmount(rewards, denomList[count]))
					gaugesDenom[i+1].Set(utils.GetAmount(commission, denomList[count]))
					gaugesDenom[i+2].Set(utils.GetAmount(balances, denomList[count]))

					if t.OutputPrint {
						fmt.Println("\n  - rewards_"+denomList[count]+": ", utils.GetAmount(rewards, denomList[count]))
						fmt.Println("  - commission_"+denomList[count]+": ", utils.GetAmount(commission, denomList[count]))
						fmt.Println("  - balances_"+denomList[count]+": ", utils.GetAmount(balances, denomList[count]))
					}

					count++
				}

				// prometheus giages value
				gaugesValue := [...]float64{blockHeight, currentBlockTime, precommitRate, blocksPerYear, defaultBlockTime, inflationRateChange, inflationMax, inflationMin, inflationGoalBonded, defaultBlockTimeInflation, currentBlockTimeInflation, proposerWalletAccountNumber, validatorCount, notBondedTokens, bondedTokens, totalBondedTokens, bondedRate, validatorCommitStatus, proposerPriorityValue, proposerPriority, proposingStatus, votingPower, delegatorShares, delegationRatio, delegatorCount, selfDelegationAmount, commissionRate, commissionMaxRate, commissionMaxChangeRate, minSelfDelegation, jailed}


				for i := 0; i < len(gaugesNamespaceList); i++ {
					gauges[i].Set(gaugesValue[i])
				}

				gaugesForLabel.WithLabelValues(chainId, moniker, validatorPubKey, operatorAddress, accountAddress, consHexAddress).Add(0)
			}

		}
		time.Sleep(2 * time.Second)
	}
}
