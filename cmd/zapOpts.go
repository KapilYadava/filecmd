package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type Options struct {
	Name    string
	Phone   string
	Address string
}

func Command() *cobra.Command {
	var (
		zapOpts = zap.Options{Development: true}
		opts    Options
	)

	zapOptions := &cobra.Command{
		Use:   "user",
		Short: "get user info",
		Long:  "lookup for user informations.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := zap.New(zap.UseFlagOptions(&zapOpts))
			ctrl.SetLogger(logger)
			cmd.SetContext(ctrl.LoggerInto(cmd.Context(), ctrl.Log))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd.Context(), opts)
		},
	}

	flags := flag.NewFlagSet("", 0)
	zapOpts.BindFlags(flags)
	zapOptions.PersistentFlags().AddGoFlagSet(flags)

	opts.AddFlags(zapOptions.Flags())
	opts.MarkFlagsRequired(zapOptions)
	return zapOptions
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Name, "Name", "Kapil", "name of user")
	fs.StringVar(&o.Address, "Address", "Sahaberi", "address of user")
	fs.StringVar(&o.Phone, "Phone", o.Phone, "user phone number")
}

func (o *Options) MarkFlagsRequired(cmd *cobra.Command) {
	_ = cmd.MarkFlagRequired("supported-machine-classes")
}

func init() {
	rootCmd.AddCommand(Command())
}

func Run(ctx context.Context, opts Options) error {
	if opts.Name == "Kapil" {
		return fmt.Errorf("unkown user")
	}

	fmt.Println(opts.Name, opts.Address, opts.Phone)
	return nil
}
