package cmd

import (
	"errors"
	"fmt"
	//"github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"
	latestapi "github.com/openshift/origin/pkg/api/latest"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	"github.com/spf13/cobra"
	"io"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	
	//log "github.com/golang/glog"
	"strings"
)

func GetBackingServicePlan(bs *backingserviceapi.BackingService, name string) *backingserviceapi.ServicePlan {
	for _, plan := range bs.Spec.Plans {
		if name == plan.Name {
			return &plan
		}
	}
	
	return nil
}

//====================================================
// new
//====================================================

const (
	newBackingServiceInstanceLong = `
Create a new BackingService instance

This command will try to create a backing service instance.
`
	newBackingServiceInstanceExample = `# Create a backingservice instance via backingservice
  $ %[1]s mysql-instance --backingservice=myslql --plan=shared

  # Create a user-provided-service
  $ %[1]s redis-instance -p host=redis.somedomain.com -p port=6379 -p password=H3IIOw0R1D`
)

type NewBackingServiceInstanceOptions struct {
	Name      string
	
	BackingServiceName     string
	BackingServicePlanName string
	Parameters []string
	Mode string
}

func NewCmdNewBackingServiceInstance(fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &NewBackingServiceInstanceOptions{}

	cmd := &cobra.Command{
		Use:     "new-instance NAME --backingservice=BackingServiceName --plan=BackingServicePlanName",
		Short:   "create a new BackingService instance",
		Long:    newBackingServiceInstanceLong,
		Example: fmt.Sprintf(newBackingServiceInstanceExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.complete(cmd, f); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if err := options.Run(cmd, f, out); err != nil {
				kcmdutil.CheckErr(err)
				return
			}
		},
	}

	cmd.Flags().StringVar(&options.BackingServiceName, "backingservice", "", "BackingService Name")
	cmd.Flags().StringVar(&options.BackingServicePlanName, "plan", "", "BackingService Plan Name")
	cmd.Flags().StringSliceVarP(&options.Parameters, "parameters", "p", options.Parameters, "Specify key value pairs of Paremeters variables.")
	// todo: dashboard_url
	
	return cmd
}

func (o *NewBackingServiceInstanceOptions) complete(cmd *cobra.Command, f *clientcmd.Factory) error {
	args := cmd.Flags().Args()
	if len(args) == 0 {
		cmd.Help()
		return errors.New("must have at least 1 argument")
	}

	o.Name = args[0]

	if len(o.BackingServicePlanName) == 0 || len(o.BackingServiceName) == 0{
		if len(o.Parameters)>0{
			o.Mode = backingserviceinstanceapi.UPS

			return nil
		}
		return errors.New("backingservice and plan must be specified.")
	}else{
		o.Mode = "normal"
		if len(o.Parameters) > 0{
			cmd.Help()
			return errors.New("User-provided-service can't be created like this way.")
		}
	}

	return nil
}

func (o *NewBackingServiceInstanceOptions) Run(cmd *cobra.Command, f *clientcmd.Factory, out io.Writer) error {
	client, _, err := f.Clients()
	if err != nil {
		return err
	}

	if o.Mode== backingserviceinstanceapi.UPS {
		bsi := &backingserviceinstanceapi.BackingServiceInstance{}
		bsi.Name = o.Name
		bsi.Spec.BackingServiceName = backingserviceinstanceapi.UPS
		bsi.Spec.BackingServicePlanGuid = backingserviceinstanceapi.UPS
		bsi.Annotations = make(map[string]string)
		bsi.Annotations[backingserviceinstanceapi.UPS] = "true"
		bsi.Spec.Credentials = make(map[string]string)
		bsi.Status.Phase = backingserviceinstanceapi.BackingServiceInstancePhaseUnbound
		for _, para:=range o.Parameters {
			parts := strings.SplitN(para, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid environment variable: %v", para)
			}else{
				bsi.Spec.Credentials[parts[0]] = parts[1]
			}
		}


		namespace, _, err := f.DefaultNamespace()
		if err != nil {
			return err
		}
		_, err = client.BackingServiceInstances(namespace).Create(bsi)
		if err != nil {
			return err
		}

		fmt.Fprintf(out, "User-Provided-Service Instance has been created.\n")

		return nil
	}
	
	//>> todo: maybe better do this is in Create
	bs, err := client.BackingServices("openshift").Get(o.BackingServiceName)
	if err != nil {
		return err
	}
	
	plan := GetBackingServicePlan(bs, o.BackingServicePlanName)
	if plan == nil {
		return fmt.Errorf("plan %s not found", o.BackingServicePlanName)
		//return errors.New("plan not found")
	}
	//<<
	
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	
	backingServiceInstance := &backingserviceinstanceapi.BackingServiceInstance{}
	
	backingServiceInstance.Name = o.Name
	//backingServiceInstance.GenerateName = o.Name
	
	backingServiceInstance.Spec.BackingServiceName = bs.Name // o.BackingServiceName
	//backingServiceInstance.Spec.BackingServiceID = bs.Spec.Id
	backingServiceInstance.Spec.BackingServicePlanGuid = plan.Id // o.BackingServicePlanGuid
	//backingServiceInstance.Spec.BackingServicePlanName = plan.Name
	
	//backingServiceInstance.Status = backingserviceinstanceapi.BackingServiceInstancePhaseCreated
	
	_, err = client.BackingServiceInstances(namespace).Create(backingServiceInstance)
	if err != nil {
		return err
	}
	
	fmt.Fprintf(out, "Backing Service Instance has been created.\n")

	return nil
}

//====================================================
// edit
//====================================================
/*
const (
	editBackingServiceInstanceLong = `
Edit a BackingServiceInstance

This command will try to edit a backing service instance.
`
	editBackingServiceInstanceExample = `# Edit a backingserviceinstance with [name BackingServicePlanGuid]
  $ %[1]s mysql_BackingServiceInstance --planid="BackingServicePlanGuid"`
)

type EditBackingServiceInstanceOptions struct {
	Name                   string
	BackingServicePlanGuid string
}

func NewCmdEditBackingServiceInstance(fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &NewBackingServiceInstanceOptions{}

	cmd := &cobra.Command{
		Use:     "edit-backingserviceinstance NAME --plan_guid=BackingServicePlanGuid",
		Short:   "Edit a BackingServiceInstance",
		Long:    editBackingServiceInstanceLong,
		Example: fmt.Sprintf(editBackingServiceInstanceExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.complete(cmd, f); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if err := options.Run(cmd, f, out); err != nil {
				kcmdutil.CheckErr(err)
				return
			}
		},
	}

	cmd.Flags().StringVar(&options.BackingServicePlanGuid, "planid", "", "BackingService Plan GUID")
	
	return cmd
}

func (o *EditBackingServiceInstanceOptions) complete(cmd *cobra.Command, f *clientcmd.Factory) error {
	args := cmd.Flags().Args()
	if len(args) == 0 {
		cmd.Help()
		return errors.New("must have at least 1 argument")
	}

	o.Name = args[0]

	return nil
}

func (o *EditBackingServiceInstanceOptions) Run(cmd *cobra.Command, f *clientcmd.Factory, out io.Writer) error {
	client, _, err := f.Clients()
	if err != nil {
		return err
	}
	
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	
	backingServiceInstance, err := client.BackingServiceInstances(namespace).Get(o.Name)
	if err != nil {
		return err
	}
	
	//>> todo: maybe better do this is in Update
	bs, err := client.BackingServices("openshift").Get(backingServiceInstance.Spec.BackingServiceName)
	if err != nil {
		return err
	}
	
	plan := GetBackingServicePlan(bs, o.BackingServicePlanGuid)
	if plan == nil {
		return errors.New("plan not found")
	}
	//<<
	
	backingServiceInstance.Spec.BackingServicePlanGuid = o.BackingServicePlanGuid
	
	_, err = client.BackingServiceInstances(namespace).Update(backingServiceInstance)
	if err != nil {
		return err
	}
	
	fmt.Fprintf(out, "Backing Service Instance has been updated.\n")

	return nil
}
*/
//====================================================
// bind
//====================================================

const (
	bindBackingServiceInstanceLong = `
Bind a new BackingServiceInstance

This command will try to bind a backing service instance and a deployment config.
`
	bindBackingServiceInstanceExample = `# Bind a new backingserviceinstance with a deploy config [BackingServiceInstanceName DeploymentConfigName]
  $ %[1]s mysql_BackingServiceInstance helloworld_DeploymentConfig`
)

type BindBackingServiceInstanceOptions struct {
	Name                 string
	DeploymentConfigName string
}

func NewCmdBindBackingServiceInstance(fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &BindBackingServiceInstanceOptions{}

	cmd := &cobra.Command{
		Use:     "bind BackingServiceInstanceName DeployConfigName",
		Short:   "bind a BackingServiceInstance and a DeployConfig",
		Long:    bindBackingServiceInstanceLong,
		Example: fmt.Sprintf(bindBackingServiceInstanceExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.complete(cmd, f); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if err := options.Run(cmd, f, out); err != nil {
				kcmdutil.CheckErr(err)
				return
			}
		},
	}
	
	return cmd
}

func (o *BindBackingServiceInstanceOptions) complete(cmd *cobra.Command, f *clientcmd.Factory) error {
	args := cmd.Flags().Args()
	if len(args) < 2 {
		cmd.Help()
		return errors.New("must have at least 2 arguments")
	}

	o.Name                 = args[0]
	o.DeploymentConfigName = args[1]

	return nil
}

func (o *BindBackingServiceInstanceOptions) Run(cmd *cobra.Command, f *clientcmd.Factory, out io.Writer) error {
	client, _, err := f.Clients()
	if err != nil {
		return err
	}
	
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	
	//>> todo: maybe better do this is in CreateBinding
	/*
	_, err = client.BackingServiceInstances(namespace).Get(o.Name)
	if err != nil {
		return err
	}
	
	_, err = client.DeploymentConfigs(namespace).Get(o.DeploymentConfigName)
	if err != nil {
		return err
	}
	*/
	//<<
	
	bro := backingserviceinstanceapi.NewBindingRequestOptions(
		backingserviceinstanceapi.BindKind_DeploymentConfig, 
		latestapi.Version.Version,
		o.DeploymentConfigName)
	bro.Name = o.Name 
	bro.Namespace = namespace
	
	err = client.BackingServiceInstances(namespace).CreateBinding(o.Name, bro)
	if err != nil {
		return err
	}
	
	fmt.Fprintf(out, "Backing Service Instance has been bound.\n")

	return nil
}

//====================================================
// unbind
//====================================================

const (
	unbindBackingServiceInstanceLong = `
Unbind a new BackingServiceInstance

This command will try to unbind a backing service instance and a deployment config.
`
	unbindBackingServiceInstanceExample = `# Unbind a new backingserviceinstance with and deploy config [BackingServiceInstanceName DeploymentConfigName]
  $ %[1]s mysql_BackingServiceInstance helloworld_DeploymentConfig`
)

type UnbindBackingServiceInstanceOptions struct {
	Name                 string
	DeploymentConfigName string
}

func NewCmdUnbindBackingServiceInstance(fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &UnbindBackingServiceInstanceOptions{}

	cmd := &cobra.Command{
		Use:     "unbind BackingServiceInstanceName DeployConfigName",
		Short:   "unbind a BackingServiceInstance and a DeployConfig",
		Long:    unbindBackingServiceInstanceLong,
		Example: fmt.Sprintf(unbindBackingServiceInstanceExample, fullName),
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.complete(cmd, f); err != nil {
				kcmdutil.CheckErr(err)
				return
			}

			if err := options.Run(cmd, f, out); err != nil {
				kcmdutil.CheckErr(err)
				return
			}
		},
	}
	
	return cmd
}

func (o *UnbindBackingServiceInstanceOptions) complete(cmd *cobra.Command, f *clientcmd.Factory) error {
	args := cmd.Flags().Args()
	if len(args) < 2 {
		cmd.Help()
		return errors.New("must have at least 2 arguments")
	}

	o.Name                 = args[0]
	o.DeploymentConfigName = args[1]

	return nil
}

func (o *UnbindBackingServiceInstanceOptions) Run(cmd *cobra.Command, f *clientcmd.Factory, out io.Writer) error {
	client, _, err := f.Clients()
	if err != nil {
		return err
	}
	
	namespace, _, err := f.DefaultNamespace()
	if err != nil {
		return err
	}
	
	//>> todo: maybe better do this is in DeleteBinding
	/*
	_, err = client.BackingServiceInstances(namespace).Get(o.Name)
	if err != nil {
		return err
	}
	
	_, err = client.DeploymentConfigs(namespace).Get(o.DeploymentConfigName)
	if err != nil {
		return err
	}
	*/
	//<<

	bro := backingserviceinstanceapi.NewBindingRequestOptions(
		backingserviceinstanceapi.BindKind_DeploymentConfig,
		latestapi.Version.Version,
		o.DeploymentConfigName)
	bro.Name = o.Name
	bro.Namespace = namespace
	
	//err = client.BackingServiceInstances(namespace).DeleteBinding(o.Name)
	err = client.BackingServiceInstances(namespace).UpdateBinding(o.Name,bro)
	if err != nil {
		return err
	}
	
	fmt.Fprintf(out, "Backing Service Instance has been unbound.\n")

	return nil
}


