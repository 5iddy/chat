package cli

import (
	"chat/x/profile/types"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-profile [name] [bio] [website]",
		Short: "Create a new profile",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var indexName string
			var argBio string
			var argWebsite string

			argsLength := len(args)

			if argsLength >= 1 {
				indexName = args[0]
			}

			if argsLength >= 2 {
				argBio = args[1]
			}

			if argsLength >= 3 {
				argWebsite = args[2]
			}
			// Get indexes

			// Get value arguments
			// argCastPosts = strings.Split(args[3], listSeparator)
			// argPosts := make([]uint64, len(argCastPosts))
			// for i, arg := range argCastPosts {
			// 	value, err := cast.ToUint64E(arg)
			// 	if err != nil {
			// 		return err
			// 	}
			// 	argPosts[i] = value
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProfile(
				clientCtx.GetFromAddress().String(),
				indexName,
				argBio,
				argWebsite,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-profile [name] [bio] [website]",
		Short: "Update a profile",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexName := args[0]

			// Get value arguments
			argBio := args[1]
			argWebsite := args[2]
			argCastPosts := strings.Split(args[3], listSeparator)
			argPosts := make([]uint64, len(argCastPosts))
			for i, arg := range argCastPosts {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				argPosts[i] = value
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateProfile(
				clientCtx.GetFromAddress().String(),
				indexName,
				argBio,
				argWebsite,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-profile [name]",
		Short: "Delete a profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexName := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteProfile(
				clientCtx.GetFromAddress().String(),
				indexName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
