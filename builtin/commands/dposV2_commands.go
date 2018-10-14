package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/dposv2"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/loomnetwork/go-loom/types"
)

const DPOSV2ContractName = "dposV2"

func ListValidatorsCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "list_validatorsV2",
		Short: "List the current validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.ListValidatorsResponseV2
			err := cli.StaticCallContract(DPOSV2ContractName, "ListValidators", &dposv2.ListValidatorsRequestV2{}, &resp)
			if err != nil {
				return err
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

func ListCandidatesCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "list_candidatesV2",
		Short: "List the registered candidates",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp dposv2.ListCandidateResponseV2
			err := cli.StaticCallContract(DPOSV2ContractName, "ListCandidates", &dposv2.ListCandidateRequestV2{}, &resp)
			if err != nil {
				return err
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

func RegisterCandidateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "register_candidateV2 [public key]",
		Short: "Register a candidate for validator",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, err := cli.ParseBytes(args[0])
			if err != nil {
				return err
			}
			return cli.CallContract(DPOSV2ContractName, "RegisterCandidate", &dposv2.RegisterCandidateRequestV2{
				PubKey: pubKey,
			}, nil)
		},
	}
}

func VoteCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "voteV2 [candidate address] [amount]",
		Short: "Allocate votes to a candidate",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			amount, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV2ContractName, "Vote", &dposv2.VoteRequestV2{
				CandidateAddress: addr.MarshalPB(),
				Amount:           amount,
			}, nil)
		},
	}
}

func DelegateCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "delegateV2 [validator address] [amount]",
		Short: "delegate tokens to a validator",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}
			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV2ContractName, "Delegate", &dposv2.DelegateRequestV2{
				ValidatorAddress: addr.MarshalPB(),
				Amount:           &types.BigUInt{
					Value: *amount,
				},
			}, nil)
		},
	}
}

func CheckDelegationCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "check_delegationV2 [validator address]",
		Short: "check delegation to a particular validator",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			return cli.CallContract(DPOSV2ContractName, "CheckDelegation", &dposv2.CheckDelegationRequestV2{
				ValidatorAddress: addr.MarshalPB(),
			}, nil)
		},
	}
}

func UnbondCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "unbondV2 [validator address] [amount]",
		Short: "De-allocate tokens from a validator",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr, err := cli.ResolveAddress(args[0])
			if err != nil {
				return err
			}

			amount, err := cli.ParseAmount(args[1])
			if err != nil {
				return err
			}
			return cli.CallContract(DPOSV2ContractName, "Unbond", &dposv2.UnbondRequestV2{
				ValidatorAddress: addr.MarshalPB(),
				Amount:           &types.BigUInt{
					Value: *amount,
				},
			}, nil)
		},
	}
}


func ElectCmdV2() *cobra.Command {
	return &cobra.Command{
		Use:   "electV2",
		Short: "Run an election",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.CallContract(DPOSV2ContractName, "Elect", &dposv2.ElectRequestV2{}, nil)
		},
	}
}

func AddDPOSV2(root *cobra.Command) {
	root.AddCommand(
		ListValidatorsCmdV2(),
		RegisterCandidateCmdV2(),
		VoteCmdV2(),
		ElectCmdV2(),
		ListCandidatesCmdV2(),
		DelegateCmdV2(),
		UnbondCmdV2(),
	)
}
