package cmd

import (
	"errors"
	"fmt"
	"github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"
	"github.com/spf13/cobra"
	"io"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

const (
	deleteApplicationLong = `
application is a group of resource
delete application is used to delete resources in the same application
`
	deleteApplicationExample = `# delete a new application with [name deletelabel]
  $ %[1]s  mobile_app  --onlylabel=true
 `
)

type DeleteApplicationOptions struct {
	Name      string
	OnlyLabel bool

	Client  client.Interface
	KClient kclient.Interface

	Out io.Writer
}

func NewCmdDeleteApplication(fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &DeleteApplicationOptions{}
	options.Out = out

	cmd := &cobra.Command{
		Use:     `delete-application NAME [--onlylabel=true]`,
		Short:   "delete a existed application",
		Long:    deleteApplicationLong,
		Example: fmt.Sprintf(deleteApplicationExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			if err = options.complete(cmd, f); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if options.Client, options.KClient, err = f.Clients(); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if err := options.Run(f); err != nil {
				fmt.Println("run err %s", err.Error())
			}

		},
	}

	cmd.Flags().BoolVarP(&options.OnlyLabel, "onlylabel", "l", false, "only delete label")

	return cmd
}

func (o *DeleteApplicationOptions) complete(cmd *cobra.Command, f *clientcmd.Factory) error {
	args := cmd.Flags().Args()
	if len(args) != 1 {
		cmd.Help()
		return errors.New("must have exactly one argument")
	}

	o.Name = args[0]

	return nil
}

func (o *DeleteApplicationOptions) Run(f *clientcmd.Factory) error {
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}

	app, err := o.Client.Applications(namespace).Get(o.Name)
	if err != nil {
		return err
	}

	if o.OnlyLabel {
		if err = o.Client.Applications(namespace).Delete(app.Name); err != nil {
			return err
		}

		fmt.Printf("application %s deleted", app.Name)
		return nil
	}

	app.Spec.Destory = true
	if _, err := o.Client.Applications(namespace).Update(app); err != nil {
		return err
	}

	fmt.Printf("application %s deleted", app.Name)
	return nil
}
