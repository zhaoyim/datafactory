package v1

// AUTO-GENERATED FUNCTIONS START HERE
import (
	v1 "github.com/openshift/origin/pkg/application/api/v1"
	apiv1 "github.com/openshift/origin/pkg/authorization/api/v1"
	backingserviceapiv1 "github.com/openshift/origin/pkg/backingservice/api/v1"
	backingserviceinstanceapiv1 "github.com/openshift/origin/pkg/backingserviceinstance/api/v1"
	buildapiv1 "github.com/openshift/origin/pkg/build/api/v1"
	deployapiv1 "github.com/openshift/origin/pkg/deploy/api/v1"
	imageapiv1 "github.com/openshift/origin/pkg/image/api/v1"
	oauthapiv1 "github.com/openshift/origin/pkg/oauth/api/v1"
	projectapiv1 "github.com/openshift/origin/pkg/project/api/v1"
	routeapiv1 "github.com/openshift/origin/pkg/route/api/v1"
	sdnapiv1 "github.com/openshift/origin/pkg/sdn/api/v1"
	servicebrokerapiv1 "github.com/openshift/origin/pkg/servicebroker/api/v1"
	templateapiv1 "github.com/openshift/origin/pkg/template/api/v1"
	userapiv1 "github.com/openshift/origin/pkg/user/api/v1"
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	pkgapiv1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
)

func deepCopy_v1_Application(in v1.Application, out *v1.Application, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_ApplicationSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ApplicationStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ApplicationList(in v1.ApplicationList, out *v1.ApplicationList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]v1.Application, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Application(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ApplicationSpec(in v1.ApplicationSpec, out *v1.ApplicationSpec, c *conversion.Cloner) error {
	out.Name = in.Name
	if in.Items != nil {
		out.Items = make([]v1.Item, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Item(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	out.Destory = in.Destory
	if in.Finalizers != nil {
		out.Finalizers = make([]pkgapiv1.FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = in.Finalizers[i]
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func deepCopy_v1_ApplicationStatus(in v1.ApplicationStatus, out *v1.ApplicationStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	return nil
}

func deepCopy_v1_Item(in v1.Item, out *v1.Item, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.Status = in.Status
	return nil
}

func deepCopy_v1_AuthorizationAttributes(in apiv1.AuthorizationAttributes, out *apiv1.AuthorizationAttributes, c *conversion.Cloner) error {
	out.Namespace = in.Namespace
	out.Verb = in.Verb
	out.Group = in.Group
	out.Version = in.Version
	out.Resource = in.Resource
	out.ResourceName = in.ResourceName
	if newVal, err := c.DeepCopy(in.Content); err != nil {
		return err
	} else {
		out.Content = newVal.(runtime.RawExtension)
	}
	return nil
}

func deepCopy_v1_ClusterPolicy(in apiv1.ClusterPolicy, out *apiv1.ClusterPolicy, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if newVal, err := c.DeepCopy(in.LastModified); err != nil {
		return err
	} else {
		out.LastModified = newVal.(unversioned.Time)
	}
	if in.Roles != nil {
		out.Roles = make([]apiv1.NamedClusterRole, len(in.Roles))
		for i := range in.Roles {
			if err := deepCopy_v1_NamedClusterRole(in.Roles[i], &out.Roles[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Roles = nil
	}
	return nil
}

func deepCopy_v1_ClusterPolicyBinding(in apiv1.ClusterPolicyBinding, out *apiv1.ClusterPolicyBinding, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if newVal, err := c.DeepCopy(in.LastModified); err != nil {
		return err
	} else {
		out.LastModified = newVal.(unversioned.Time)
	}
	if newVal, err := c.DeepCopy(in.PolicyRef); err != nil {
		return err
	} else {
		out.PolicyRef = newVal.(pkgapiv1.ObjectReference)
	}
	if in.RoleBindings != nil {
		out.RoleBindings = make([]apiv1.NamedClusterRoleBinding, len(in.RoleBindings))
		for i := range in.RoleBindings {
			if err := deepCopy_v1_NamedClusterRoleBinding(in.RoleBindings[i], &out.RoleBindings[i], c); err != nil {
				return err
			}
		}
	} else {
		out.RoleBindings = nil
	}
	return nil
}

func deepCopy_v1_ClusterPolicyBindingList(in apiv1.ClusterPolicyBindingList, out *apiv1.ClusterPolicyBindingList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterPolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ClusterPolicyBinding(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ClusterPolicyList(in apiv1.ClusterPolicyList, out *apiv1.ClusterPolicyList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterPolicy, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ClusterPolicy(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ClusterRole(in apiv1.ClusterRole, out *apiv1.ClusterRole, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Rules != nil {
		out.Rules = make([]apiv1.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := deepCopy_v1_PolicyRule(in.Rules[i], &out.Rules[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func deepCopy_v1_ClusterRoleBinding(in apiv1.ClusterRoleBinding, out *apiv1.ClusterRoleBinding, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.UserNames != nil {
		out.UserNames = make([]string, len(in.UserNames))
		for i := range in.UserNames {
			out.UserNames[i] = in.UserNames[i]
		}
	} else {
		out.UserNames = nil
	}
	if in.GroupNames != nil {
		out.GroupNames = make([]string, len(in.GroupNames))
		for i := range in.GroupNames {
			out.GroupNames[i] = in.GroupNames[i]
		}
	} else {
		out.GroupNames = nil
	}
	if in.Subjects != nil {
		out.Subjects = make([]pkgapiv1.ObjectReference, len(in.Subjects))
		for i := range in.Subjects {
			if newVal, err := c.DeepCopy(in.Subjects[i]); err != nil {
				return err
			} else {
				out.Subjects[i] = newVal.(pkgapiv1.ObjectReference)
			}
		}
	} else {
		out.Subjects = nil
	}
	if newVal, err := c.DeepCopy(in.RoleRef); err != nil {
		return err
	} else {
		out.RoleRef = newVal.(pkgapiv1.ObjectReference)
	}
	return nil
}

func deepCopy_v1_ClusterRoleBindingList(in apiv1.ClusterRoleBindingList, out *apiv1.ClusterRoleBindingList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterRoleBinding, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ClusterRoleBinding(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ClusterRoleList(in apiv1.ClusterRoleList, out *apiv1.ClusterRoleList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.ClusterRole, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ClusterRole(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_IsPersonalSubjectAccessReview(in apiv1.IsPersonalSubjectAccessReview, out *apiv1.IsPersonalSubjectAccessReview, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	return nil
}

func deepCopy_v1_LocalResourceAccessReview(in apiv1.LocalResourceAccessReview, out *apiv1.LocalResourceAccessReview, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if err := deepCopy_v1_AuthorizationAttributes(in.AuthorizationAttributes, &out.AuthorizationAttributes, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_LocalSubjectAccessReview(in apiv1.LocalSubjectAccessReview, out *apiv1.LocalSubjectAccessReview, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if err := deepCopy_v1_AuthorizationAttributes(in.AuthorizationAttributes, &out.AuthorizationAttributes, c); err != nil {
		return err
	}
	out.User = in.User
	if in.GroupsSlice != nil {
		out.GroupsSlice = make([]string, len(in.GroupsSlice))
		for i := range in.GroupsSlice {
			out.GroupsSlice[i] = in.GroupsSlice[i]
		}
	} else {
		out.GroupsSlice = nil
	}
	return nil
}

func deepCopy_v1_NamedClusterRole(in apiv1.NamedClusterRole, out *apiv1.NamedClusterRole, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_v1_ClusterRole(in.Role, &out.Role, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_NamedClusterRoleBinding(in apiv1.NamedClusterRoleBinding, out *apiv1.NamedClusterRoleBinding, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_v1_ClusterRoleBinding(in.RoleBinding, &out.RoleBinding, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_NamedRole(in apiv1.NamedRole, out *apiv1.NamedRole, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_v1_Role(in.Role, &out.Role, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_NamedRoleBinding(in apiv1.NamedRoleBinding, out *apiv1.NamedRoleBinding, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := deepCopy_v1_RoleBinding(in.RoleBinding, &out.RoleBinding, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_Policy(in apiv1.Policy, out *apiv1.Policy, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if newVal, err := c.DeepCopy(in.LastModified); err != nil {
		return err
	} else {
		out.LastModified = newVal.(unversioned.Time)
	}
	if in.Roles != nil {
		out.Roles = make([]apiv1.NamedRole, len(in.Roles))
		for i := range in.Roles {
			if err := deepCopy_v1_NamedRole(in.Roles[i], &out.Roles[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Roles = nil
	}
	return nil
}

func deepCopy_v1_PolicyBinding(in apiv1.PolicyBinding, out *apiv1.PolicyBinding, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if newVal, err := c.DeepCopy(in.LastModified); err != nil {
		return err
	} else {
		out.LastModified = newVal.(unversioned.Time)
	}
	if newVal, err := c.DeepCopy(in.PolicyRef); err != nil {
		return err
	} else {
		out.PolicyRef = newVal.(pkgapiv1.ObjectReference)
	}
	if in.RoleBindings != nil {
		out.RoleBindings = make([]apiv1.NamedRoleBinding, len(in.RoleBindings))
		for i := range in.RoleBindings {
			if err := deepCopy_v1_NamedRoleBinding(in.RoleBindings[i], &out.RoleBindings[i], c); err != nil {
				return err
			}
		}
	} else {
		out.RoleBindings = nil
	}
	return nil
}

func deepCopy_v1_PolicyBindingList(in apiv1.PolicyBindingList, out *apiv1.PolicyBindingList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.PolicyBinding, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_PolicyBinding(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_PolicyList(in apiv1.PolicyList, out *apiv1.PolicyList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.Policy, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Policy(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_PolicyRule(in apiv1.PolicyRule, out *apiv1.PolicyRule, c *conversion.Cloner) error {
	if in.Verbs != nil {
		out.Verbs = make([]string, len(in.Verbs))
		for i := range in.Verbs {
			out.Verbs[i] = in.Verbs[i]
		}
	} else {
		out.Verbs = nil
	}
	if newVal, err := c.DeepCopy(in.AttributeRestrictions); err != nil {
		return err
	} else {
		out.AttributeRestrictions = newVal.(runtime.RawExtension)
	}
	if in.APIGroups != nil {
		out.APIGroups = make([]string, len(in.APIGroups))
		for i := range in.APIGroups {
			out.APIGroups[i] = in.APIGroups[i]
		}
	} else {
		out.APIGroups = nil
	}
	if in.Resources != nil {
		out.Resources = make([]string, len(in.Resources))
		for i := range in.Resources {
			out.Resources[i] = in.Resources[i]
		}
	} else {
		out.Resources = nil
	}
	if in.ResourceNames != nil {
		out.ResourceNames = make([]string, len(in.ResourceNames))
		for i := range in.ResourceNames {
			out.ResourceNames[i] = in.ResourceNames[i]
		}
	} else {
		out.ResourceNames = nil
	}
	if in.NonResourceURLsSlice != nil {
		out.NonResourceURLsSlice = make([]string, len(in.NonResourceURLsSlice))
		for i := range in.NonResourceURLsSlice {
			out.NonResourceURLsSlice[i] = in.NonResourceURLsSlice[i]
		}
	} else {
		out.NonResourceURLsSlice = nil
	}
	return nil
}

func deepCopy_v1_ResourceAccessReview(in apiv1.ResourceAccessReview, out *apiv1.ResourceAccessReview, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if err := deepCopy_v1_AuthorizationAttributes(in.AuthorizationAttributes, &out.AuthorizationAttributes, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ResourceAccessReviewResponse(in apiv1.ResourceAccessReviewResponse, out *apiv1.ResourceAccessReviewResponse, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	out.Namespace = in.Namespace
	if in.UsersSlice != nil {
		out.UsersSlice = make([]string, len(in.UsersSlice))
		for i := range in.UsersSlice {
			out.UsersSlice[i] = in.UsersSlice[i]
		}
	} else {
		out.UsersSlice = nil
	}
	if in.GroupsSlice != nil {
		out.GroupsSlice = make([]string, len(in.GroupsSlice))
		for i := range in.GroupsSlice {
			out.GroupsSlice[i] = in.GroupsSlice[i]
		}
	} else {
		out.GroupsSlice = nil
	}
	return nil
}

func deepCopy_v1_Role(in apiv1.Role, out *apiv1.Role, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Rules != nil {
		out.Rules = make([]apiv1.PolicyRule, len(in.Rules))
		for i := range in.Rules {
			if err := deepCopy_v1_PolicyRule(in.Rules[i], &out.Rules[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func deepCopy_v1_RoleBinding(in apiv1.RoleBinding, out *apiv1.RoleBinding, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.UserNames != nil {
		out.UserNames = make([]string, len(in.UserNames))
		for i := range in.UserNames {
			out.UserNames[i] = in.UserNames[i]
		}
	} else {
		out.UserNames = nil
	}
	if in.GroupNames != nil {
		out.GroupNames = make([]string, len(in.GroupNames))
		for i := range in.GroupNames {
			out.GroupNames[i] = in.GroupNames[i]
		}
	} else {
		out.GroupNames = nil
	}
	if in.Subjects != nil {
		out.Subjects = make([]pkgapiv1.ObjectReference, len(in.Subjects))
		for i := range in.Subjects {
			if newVal, err := c.DeepCopy(in.Subjects[i]); err != nil {
				return err
			} else {
				out.Subjects[i] = newVal.(pkgapiv1.ObjectReference)
			}
		}
	} else {
		out.Subjects = nil
	}
	if newVal, err := c.DeepCopy(in.RoleRef); err != nil {
		return err
	} else {
		out.RoleRef = newVal.(pkgapiv1.ObjectReference)
	}
	return nil
}

func deepCopy_v1_RoleBindingList(in apiv1.RoleBindingList, out *apiv1.RoleBindingList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.RoleBinding, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_RoleBinding(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_RoleList(in apiv1.RoleList, out *apiv1.RoleList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]apiv1.Role, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Role(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_SubjectAccessReview(in apiv1.SubjectAccessReview, out *apiv1.SubjectAccessReview, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if err := deepCopy_v1_AuthorizationAttributes(in.AuthorizationAttributes, &out.AuthorizationAttributes, c); err != nil {
		return err
	}
	out.User = in.User
	if in.GroupsSlice != nil {
		out.GroupsSlice = make([]string, len(in.GroupsSlice))
		for i := range in.GroupsSlice {
			out.GroupsSlice[i] = in.GroupsSlice[i]
		}
	} else {
		out.GroupsSlice = nil
	}
	return nil
}

func deepCopy_v1_SubjectAccessReviewResponse(in apiv1.SubjectAccessReviewResponse, out *apiv1.SubjectAccessReviewResponse, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	out.Namespace = in.Namespace
	out.Allowed = in.Allowed
	out.Reason = in.Reason
	return nil
}

func deepCopy_v1_BackingService(in backingserviceapiv1.BackingService, out *backingserviceapiv1.BackingService, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_BackingServiceSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_BackingServiceStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_BackingServiceList(in backingserviceapiv1.BackingServiceList, out *backingserviceapiv1.BackingServiceList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]backingserviceapiv1.BackingService, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_BackingService(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_BackingServiceSpec(in backingserviceapiv1.BackingServiceSpec, out *backingserviceapiv1.BackingServiceSpec, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Id = in.Id
	out.Description = in.Description
	out.Bindable = in.Bindable
	out.PlanUpdateable = in.PlanUpdateable
	if in.Tags != nil {
		out.Tags = make([]string, len(in.Tags))
		for i := range in.Tags {
			out.Tags[i] = in.Tags[i]
		}
	} else {
		out.Tags = nil
	}
	if in.Requires != nil {
		out.Requires = make([]string, len(in.Requires))
		for i := range in.Requires {
			out.Requires[i] = in.Requires[i]
		}
	} else {
		out.Requires = nil
	}
	if in.Metadata != nil {
		out.Metadata = make(map[string]string)
		for key, val := range in.Metadata {
			out.Metadata[key] = val
		}
	} else {
		out.Metadata = nil
	}
	if in.Plans != nil {
		out.Plans = make([]backingserviceapiv1.ServicePlan, len(in.Plans))
		for i := range in.Plans {
			if err := deepCopy_v1_ServicePlan(in.Plans[i], &out.Plans[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Plans = nil
	}
	if in.DashboardClient != nil {
		out.DashboardClient = make(map[string]string)
		for key, val := range in.DashboardClient {
			out.DashboardClient[key] = val
		}
	} else {
		out.DashboardClient = nil
	}
	return nil
}

func deepCopy_v1_BackingServiceStatus(in backingserviceapiv1.BackingServiceStatus, out *backingserviceapiv1.BackingServiceStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	return nil
}

func deepCopy_v1_CustomizeSpec(in backingserviceapiv1.CustomizeSpec, out *backingserviceapiv1.CustomizeSpec, c *conversion.Cloner) error {
	out.Default = in.Default
	out.Max = in.Max
	out.Price = in.Price
	out.Step = in.Step
	out.Unit = in.Unit
	out.Desc = in.Desc
	return nil
}

func deepCopy_v1_ServicePlan(in backingserviceapiv1.ServicePlan, out *backingserviceapiv1.ServicePlan, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Id = in.Id
	out.Description = in.Description
	if err := deepCopy_v1_ServicePlanMetadata(in.Metadata, &out.Metadata, c); err != nil {
		return err
	}
	out.Free = in.Free
	return nil
}

func deepCopy_v1_ServicePlanCost(in backingserviceapiv1.ServicePlanCost, out *backingserviceapiv1.ServicePlanCost, c *conversion.Cloner) error {
	if in.Amount != nil {
		out.Amount = make(map[string]float64)
		for key, val := range in.Amount {
			out.Amount[key] = val
		}
	} else {
		out.Amount = nil
	}
	out.Unit = in.Unit
	return nil
}

func deepCopy_v1_ServicePlanMetadata(in backingserviceapiv1.ServicePlanMetadata, out *backingserviceapiv1.ServicePlanMetadata, c *conversion.Cloner) error {
	if in.Bullets != nil {
		out.Bullets = make([]string, len(in.Bullets))
		for i := range in.Bullets {
			out.Bullets[i] = in.Bullets[i]
		}
	} else {
		out.Bullets = nil
	}
	if in.Costs != nil {
		out.Costs = make([]backingserviceapiv1.ServicePlanCost, len(in.Costs))
		for i := range in.Costs {
			if err := deepCopy_v1_ServicePlanCost(in.Costs[i], &out.Costs[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Costs = nil
	}
	out.DisplayName = in.DisplayName
	if in.Customize != nil {
		out.Customize = make(map[string]backingserviceapiv1.CustomizeSpec)
		for key, val := range in.Customize {
			newVal := new(backingserviceapiv1.CustomizeSpec)
			if err := deepCopy_v1_CustomizeSpec(val, newVal, c); err != nil {
				return err
			}
			out.Customize[key] = *newVal
		}
	} else {
		out.Customize = nil
	}
	return nil
}

func deepCopy_v1_BackingServiceInstance(in backingserviceinstanceapiv1.BackingServiceInstance, out *backingserviceinstanceapiv1.BackingServiceInstance, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_BackingServiceInstanceSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_BackingServiceInstanceStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_BackingServiceInstanceList(in backingserviceinstanceapiv1.BackingServiceInstanceList, out *backingserviceinstanceapiv1.BackingServiceInstanceList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]backingserviceinstanceapiv1.BackingServiceInstance, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_BackingServiceInstance(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_BackingServiceInstanceSpec(in backingserviceinstanceapiv1.BackingServiceInstanceSpec, out *backingserviceinstanceapiv1.BackingServiceInstanceSpec, c *conversion.Cloner) error {
	if err := deepCopy_v1_InstanceProvisioning(in.InstanceProvisioning, &out.InstanceProvisioning, c); err != nil {
		return err
	}
	if err := deepCopy_v1_UserProvidedService(in.UserProvidedService, &out.UserProvidedService, c); err != nil {
		return err
	}
	if in.Binding != nil {
		out.Binding = make([]backingserviceinstanceapiv1.InstanceBinding, len(in.Binding))
		for i := range in.Binding {
			if err := deepCopy_v1_InstanceBinding(in.Binding[i], &out.Binding[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Binding = nil
	}
	out.Bound = in.Bound
	out.InstanceID = in.InstanceID
	if in.Tags != nil {
		out.Tags = make([]string, len(in.Tags))
		for i := range in.Tags {
			out.Tags[i] = in.Tags[i]
		}
	} else {
		out.Tags = nil
	}
	return nil
}

func deepCopy_v1_BackingServiceInstanceStatus(in backingserviceinstanceapiv1.BackingServiceInstanceStatus, out *backingserviceinstanceapiv1.BackingServiceInstanceStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	out.Action = in.Action
	out.Patch = in.Patch
	if in.LastOperation != nil {
		out.LastOperation = new(backingserviceinstanceapiv1.LastOperation)
		if err := deepCopy_v1_LastOperation(*in.LastOperation, out.LastOperation, c); err != nil {
			return err
		}
	} else {
		out.LastOperation = nil
	}
	return nil
}

func deepCopy_v1_BindingRequestOptions(in backingserviceinstanceapiv1.BindingRequestOptions, out *backingserviceinstanceapiv1.BindingRequestOptions, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.BindKind = in.BindKind
	out.BindResourceVersion = in.BindResourceVersion
	out.ResourceName = in.ResourceName
	return nil
}

func deepCopy_v1_InstanceBinding(in backingserviceinstanceapiv1.InstanceBinding, out *backingserviceinstanceapiv1.InstanceBinding, c *conversion.Cloner) error {
	if in.BoundTime != nil {
		if newVal, err := c.DeepCopy(in.BoundTime); err != nil {
			return err
		} else {
			out.BoundTime = newVal.(*unversioned.Time)
		}
	} else {
		out.BoundTime = nil
	}
	out.BindUuid = in.BindUuid
	out.BindDeploymentConfig = in.BindDeploymentConfig
	if in.Credentials != nil {
		out.Credentials = make(map[string]string)
		for key, val := range in.Credentials {
			out.Credentials[key] = val
		}
	} else {
		out.Credentials = nil
	}
	return nil
}

func deepCopy_v1_InstanceProvisioning(in backingserviceinstanceapiv1.InstanceProvisioning, out *backingserviceinstanceapiv1.InstanceProvisioning, c *conversion.Cloner) error {
	out.DashboardUrl = in.DashboardUrl
	out.BackingServiceName = in.BackingServiceName
	out.BackingServiceSpecID = in.BackingServiceSpecID
	out.BackingServicePlanGuid = in.BackingServicePlanGuid
	out.BackingServicePlanName = in.BackingServicePlanName
	if in.Parameters != nil {
		out.Parameters = make(map[string]string)
		for key, val := range in.Parameters {
			out.Parameters[key] = val
		}
	} else {
		out.Parameters = nil
	}
	if in.Creds != nil {
		out.Creds = make(map[string]string)
		for key, val := range in.Creds {
			out.Creds[key] = val
		}
	} else {
		out.Creds = nil
	}
	if in.Accesses != nil {
		out.Accesses = make(map[string][]string)
		for key, val := range in.Accesses {
			if newVal, err := c.DeepCopy(val); err != nil {
				return err
			} else {
				out.Accesses[key] = newVal.([]string)
			}
		}
	} else {
		out.Accesses = nil
	}
	return nil
}

func deepCopy_v1_LastOperation(in backingserviceinstanceapiv1.LastOperation, out *backingserviceinstanceapiv1.LastOperation, c *conversion.Cloner) error {
	out.State = in.State
	out.Description = in.Description
	out.AsyncPollIntervalSeconds = in.AsyncPollIntervalSeconds
	return nil
}

func deepCopy_v1_UserProvidedService(in backingserviceinstanceapiv1.UserProvidedService, out *backingserviceinstanceapiv1.UserProvidedService, c *conversion.Cloner) error {
	if in.Credentials != nil {
		out.Credentials = make(map[string]string)
		for key, val := range in.Credentials {
			out.Credentials[key] = val
		}
	} else {
		out.Credentials = nil
	}
	return nil
}

func deepCopy_v1_BinaryBuildRequestOptions(in buildapiv1.BinaryBuildRequestOptions, out *buildapiv1.BinaryBuildRequestOptions, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.AsFile = in.AsFile
	out.Commit = in.Commit
	out.Message = in.Message
	out.AuthorName = in.AuthorName
	out.AuthorEmail = in.AuthorEmail
	out.CommitterName = in.CommitterName
	out.CommitterEmail = in.CommitterEmail
	return nil
}

func deepCopy_v1_BinaryBuildSource(in buildapiv1.BinaryBuildSource, out *buildapiv1.BinaryBuildSource, c *conversion.Cloner) error {
	out.AsFile = in.AsFile
	return nil
}

func deepCopy_v1_Build(in buildapiv1.Build, out *buildapiv1.Build, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_BuildSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_BuildStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_BuildConfig(in buildapiv1.BuildConfig, out *buildapiv1.BuildConfig, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_BuildConfigSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_BuildConfigStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_BuildConfigList(in buildapiv1.BuildConfigList, out *buildapiv1.BuildConfigList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]buildapiv1.BuildConfig, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_BuildConfig(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_BuildConfigSpec(in buildapiv1.BuildConfigSpec, out *buildapiv1.BuildConfigSpec, c *conversion.Cloner) error {
	if in.Triggers != nil {
		out.Triggers = make([]buildapiv1.BuildTriggerPolicy, len(in.Triggers))
		for i := range in.Triggers {
			if err := deepCopy_v1_BuildTriggerPolicy(in.Triggers[i], &out.Triggers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	if err := deepCopy_v1_BuildSpec(in.BuildSpec, &out.BuildSpec, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_BuildConfigStatus(in buildapiv1.BuildConfigStatus, out *buildapiv1.BuildConfigStatus, c *conversion.Cloner) error {
	out.LastVersion = in.LastVersion
	return nil
}

func deepCopy_v1_BuildList(in buildapiv1.BuildList, out *buildapiv1.BuildList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]buildapiv1.Build, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Build(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_BuildLog(in buildapiv1.BuildLog, out *buildapiv1.BuildLog, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	return nil
}

func deepCopy_v1_BuildLogOptions(in buildapiv1.BuildLogOptions, out *buildapiv1.BuildLogOptions, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	if in.SinceSeconds != nil {
		out.SinceSeconds = new(int64)
		*out.SinceSeconds = *in.SinceSeconds
	} else {
		out.SinceSeconds = nil
	}
	if in.SinceTime != nil {
		if newVal, err := c.DeepCopy(in.SinceTime); err != nil {
			return err
		} else {
			out.SinceTime = newVal.(*unversioned.Time)
		}
	} else {
		out.SinceTime = nil
	}
	out.Timestamps = in.Timestamps
	if in.TailLines != nil {
		out.TailLines = new(int64)
		*out.TailLines = *in.TailLines
	} else {
		out.TailLines = nil
	}
	if in.LimitBytes != nil {
		out.LimitBytes = new(int64)
		*out.LimitBytes = *in.LimitBytes
	} else {
		out.LimitBytes = nil
	}
	out.NoWait = in.NoWait
	if in.Version != nil {
		out.Version = new(int64)
		*out.Version = *in.Version
	} else {
		out.Version = nil
	}
	return nil
}

func deepCopy_v1_BuildOutput(in buildapiv1.BuildOutput, out *buildapiv1.BuildOutput, c *conversion.Cloner) error {
	if in.To != nil {
		if newVal, err := c.DeepCopy(in.To); err != nil {
			return err
		} else {
			out.To = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.To = nil
	}
	if in.PushSecret != nil {
		if newVal, err := c.DeepCopy(in.PushSecret); err != nil {
			return err
		} else {
			out.PushSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.PushSecret = nil
	}
	return nil
}

func deepCopy_v1_BuildPostCommitSpec(in buildapiv1.BuildPostCommitSpec, out *buildapiv1.BuildPostCommitSpec, c *conversion.Cloner) error {
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	if in.Args != nil {
		out.Args = make([]string, len(in.Args))
		for i := range in.Args {
			out.Args[i] = in.Args[i]
		}
	} else {
		out.Args = nil
	}
	out.Script = in.Script
	return nil
}

func deepCopy_v1_BuildRequest(in buildapiv1.BuildRequest, out *buildapiv1.BuildRequest, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Revision != nil {
		out.Revision = new(buildapiv1.SourceRevision)
		if err := deepCopy_v1_SourceRevision(*in.Revision, out.Revision, c); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if in.TriggeredByImage != nil {
		if newVal, err := c.DeepCopy(in.TriggeredByImage); err != nil {
			return err
		} else {
			out.TriggeredByImage = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.TriggeredByImage = nil
	}
	if in.From != nil {
		if newVal, err := c.DeepCopy(in.From); err != nil {
			return err
		} else {
			out.From = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.From = nil
	}
	if in.Binary != nil {
		out.Binary = new(buildapiv1.BinaryBuildSource)
		if err := deepCopy_v1_BinaryBuildSource(*in.Binary, out.Binary, c); err != nil {
			return err
		}
	} else {
		out.Binary = nil
	}
	if in.LastVersion != nil {
		out.LastVersion = new(int)
		*out.LastVersion = *in.LastVersion
	} else {
		out.LastVersion = nil
	}
	if in.Env != nil {
		out.Env = make([]pkgapiv1.EnvVar, len(in.Env))
		for i := range in.Env {
			if newVal, err := c.DeepCopy(in.Env[i]); err != nil {
				return err
			} else {
				out.Env[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Env = nil
	}
	return nil
}

func deepCopy_v1_BuildSource(in buildapiv1.BuildSource, out *buildapiv1.BuildSource, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Binary != nil {
		out.Binary = new(buildapiv1.BinaryBuildSource)
		if err := deepCopy_v1_BinaryBuildSource(*in.Binary, out.Binary, c); err != nil {
			return err
		}
	} else {
		out.Binary = nil
	}
	if in.Dockerfile != nil {
		out.Dockerfile = new(string)
		*out.Dockerfile = *in.Dockerfile
	} else {
		out.Dockerfile = nil
	}
	if in.Git != nil {
		out.Git = new(buildapiv1.GitBuildSource)
		if err := deepCopy_v1_GitBuildSource(*in.Git, out.Git, c); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	if in.Images != nil {
		out.Images = make([]buildapiv1.ImageSource, len(in.Images))
		for i := range in.Images {
			if err := deepCopy_v1_ImageSource(in.Images[i], &out.Images[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	out.ContextDir = in.ContextDir
	if in.SourceSecret != nil {
		if newVal, err := c.DeepCopy(in.SourceSecret); err != nil {
			return err
		} else {
			out.SourceSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.SourceSecret = nil
	}
	if in.Secrets != nil {
		out.Secrets = make([]buildapiv1.SecretBuildSource, len(in.Secrets))
		for i := range in.Secrets {
			if err := deepCopy_v1_SecretBuildSource(in.Secrets[i], &out.Secrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	return nil
}

func deepCopy_v1_BuildSpec(in buildapiv1.BuildSpec, out *buildapiv1.BuildSpec, c *conversion.Cloner) error {
	out.ServiceAccount = in.ServiceAccount
	if err := deepCopy_v1_BuildSource(in.Source, &out.Source, c); err != nil {
		return err
	}
	if in.Revision != nil {
		out.Revision = new(buildapiv1.SourceRevision)
		if err := deepCopy_v1_SourceRevision(*in.Revision, out.Revision, c); err != nil {
			return err
		}
	} else {
		out.Revision = nil
	}
	if err := deepCopy_v1_BuildStrategy(in.Strategy, &out.Strategy, c); err != nil {
		return err
	}
	if err := deepCopy_v1_BuildOutput(in.Output, &out.Output, c); err != nil {
		return err
	}
	if newVal, err := c.DeepCopy(in.Resources); err != nil {
		return err
	} else {
		out.Resources = newVal.(pkgapiv1.ResourceRequirements)
	}
	if err := deepCopy_v1_BuildPostCommitSpec(in.PostCommit, &out.PostCommit, c); err != nil {
		return err
	}
	if in.CompletionDeadlineSeconds != nil {
		out.CompletionDeadlineSeconds = new(int64)
		*out.CompletionDeadlineSeconds = *in.CompletionDeadlineSeconds
	} else {
		out.CompletionDeadlineSeconds = nil
	}
	return nil
}

func deepCopy_v1_BuildStatus(in buildapiv1.BuildStatus, out *buildapiv1.BuildStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	out.Cancelled = in.Cancelled
	out.Reason = in.Reason
	out.Message = in.Message
	if in.StartTimestamp != nil {
		if newVal, err := c.DeepCopy(in.StartTimestamp); err != nil {
			return err
		} else {
			out.StartTimestamp = newVal.(*unversioned.Time)
		}
	} else {
		out.StartTimestamp = nil
	}
	if in.CompletionTimestamp != nil {
		if newVal, err := c.DeepCopy(in.CompletionTimestamp); err != nil {
			return err
		} else {
			out.CompletionTimestamp = newVal.(*unversioned.Time)
		}
	} else {
		out.CompletionTimestamp = nil
	}
	out.Duration = in.Duration
	out.OutputDockerImageReference = in.OutputDockerImageReference
	if in.Config != nil {
		if newVal, err := c.DeepCopy(in.Config); err != nil {
			return err
		} else {
			out.Config = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.Config = nil
	}
	return nil
}

func deepCopy_v1_BuildStrategy(in buildapiv1.BuildStrategy, out *buildapiv1.BuildStrategy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.DockerStrategy != nil {
		out.DockerStrategy = new(buildapiv1.DockerBuildStrategy)
		if err := deepCopy_v1_DockerBuildStrategy(*in.DockerStrategy, out.DockerStrategy, c); err != nil {
			return err
		}
	} else {
		out.DockerStrategy = nil
	}
	if in.SourceStrategy != nil {
		out.SourceStrategy = new(buildapiv1.SourceBuildStrategy)
		if err := deepCopy_v1_SourceBuildStrategy(*in.SourceStrategy, out.SourceStrategy, c); err != nil {
			return err
		}
	} else {
		out.SourceStrategy = nil
	}
	if in.CustomStrategy != nil {
		out.CustomStrategy = new(buildapiv1.CustomBuildStrategy)
		if err := deepCopy_v1_CustomBuildStrategy(*in.CustomStrategy, out.CustomStrategy, c); err != nil {
			return err
		}
	} else {
		out.CustomStrategy = nil
	}
	return nil
}

func deepCopy_v1_BuildTriggerPolicy(in buildapiv1.BuildTriggerPolicy, out *buildapiv1.BuildTriggerPolicy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.GitHubWebHook != nil {
		out.GitHubWebHook = new(buildapiv1.WebHookTrigger)
		if err := deepCopy_v1_WebHookTrigger(*in.GitHubWebHook, out.GitHubWebHook, c); err != nil {
			return err
		}
	} else {
		out.GitHubWebHook = nil
	}
	if in.GenericWebHook != nil {
		out.GenericWebHook = new(buildapiv1.WebHookTrigger)
		if err := deepCopy_v1_WebHookTrigger(*in.GenericWebHook, out.GenericWebHook, c); err != nil {
			return err
		}
	} else {
		out.GenericWebHook = nil
	}
	if in.ImageChange != nil {
		out.ImageChange = new(buildapiv1.ImageChangeTrigger)
		if err := deepCopy_v1_ImageChangeTrigger(*in.ImageChange, out.ImageChange, c); err != nil {
			return err
		}
	} else {
		out.ImageChange = nil
	}
	return nil
}

func deepCopy_v1_CustomBuildStrategy(in buildapiv1.CustomBuildStrategy, out *buildapiv1.CustomBuildStrategy, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	if in.PullSecret != nil {
		if newVal, err := c.DeepCopy(in.PullSecret); err != nil {
			return err
		} else {
			out.PullSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]pkgapiv1.EnvVar, len(in.Env))
		for i := range in.Env {
			if newVal, err := c.DeepCopy(in.Env[i]); err != nil {
				return err
			} else {
				out.Env[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Env = nil
	}
	out.ExposeDockerSocket = in.ExposeDockerSocket
	out.ForcePull = in.ForcePull
	if in.Secrets != nil {
		out.Secrets = make([]buildapiv1.SecretSpec, len(in.Secrets))
		for i := range in.Secrets {
			if err := deepCopy_v1_SecretSpec(in.Secrets[i], &out.Secrets[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Secrets = nil
	}
	out.BuildAPIVersion = in.BuildAPIVersion
	return nil
}

func deepCopy_v1_DockerBuildStrategy(in buildapiv1.DockerBuildStrategy, out *buildapiv1.DockerBuildStrategy, c *conversion.Cloner) error {
	if in.From != nil {
		if newVal, err := c.DeepCopy(in.From); err != nil {
			return err
		} else {
			out.From = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.From = nil
	}
	if in.PullSecret != nil {
		if newVal, err := c.DeepCopy(in.PullSecret); err != nil {
			return err
		} else {
			out.PullSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.PullSecret = nil
	}
	out.NoCache = in.NoCache
	if in.Env != nil {
		out.Env = make([]pkgapiv1.EnvVar, len(in.Env))
		for i := range in.Env {
			if newVal, err := c.DeepCopy(in.Env[i]); err != nil {
				return err
			} else {
				out.Env[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Env = nil
	}
	out.ForcePull = in.ForcePull
	out.DockerfilePath = in.DockerfilePath
	return nil
}

func deepCopy_v1_GitBuildSource(in buildapiv1.GitBuildSource, out *buildapiv1.GitBuildSource, c *conversion.Cloner) error {
	out.URI = in.URI
	out.Ref = in.Ref
	if in.HTTPProxy != nil {
		out.HTTPProxy = new(string)
		*out.HTTPProxy = *in.HTTPProxy
	} else {
		out.HTTPProxy = nil
	}
	if in.HTTPSProxy != nil {
		out.HTTPSProxy = new(string)
		*out.HTTPSProxy = *in.HTTPSProxy
	} else {
		out.HTTPSProxy = nil
	}
	return nil
}

func deepCopy_v1_GitSourceRevision(in buildapiv1.GitSourceRevision, out *buildapiv1.GitSourceRevision, c *conversion.Cloner) error {
	out.Commit = in.Commit
	if err := deepCopy_v1_SourceControlUser(in.Author, &out.Author, c); err != nil {
		return err
	}
	if err := deepCopy_v1_SourceControlUser(in.Committer, &out.Committer, c); err != nil {
		return err
	}
	out.Message = in.Message
	return nil
}

func deepCopy_v1_ImageChangeTrigger(in buildapiv1.ImageChangeTrigger, out *buildapiv1.ImageChangeTrigger, c *conversion.Cloner) error {
	out.LastTriggeredImageID = in.LastTriggeredImageID
	if in.From != nil {
		if newVal, err := c.DeepCopy(in.From); err != nil {
			return err
		} else {
			out.From = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.From = nil
	}
	return nil
}

func deepCopy_v1_ImageSource(in buildapiv1.ImageSource, out *buildapiv1.ImageSource, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	if in.Paths != nil {
		out.Paths = make([]buildapiv1.ImageSourcePath, len(in.Paths))
		for i := range in.Paths {
			if err := deepCopy_v1_ImageSourcePath(in.Paths[i], &out.Paths[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Paths = nil
	}
	if in.PullSecret != nil {
		if newVal, err := c.DeepCopy(in.PullSecret); err != nil {
			return err
		} else {
			out.PullSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.PullSecret = nil
	}
	return nil
}

func deepCopy_v1_ImageSourcePath(in buildapiv1.ImageSourcePath, out *buildapiv1.ImageSourcePath, c *conversion.Cloner) error {
	out.SourcePath = in.SourcePath
	out.DestinationDir = in.DestinationDir
	return nil
}

func deepCopy_v1_SecretBuildSource(in buildapiv1.SecretBuildSource, out *buildapiv1.SecretBuildSource, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Secret); err != nil {
		return err
	} else {
		out.Secret = newVal.(pkgapiv1.LocalObjectReference)
	}
	out.DestinationDir = in.DestinationDir
	return nil
}

func deepCopy_v1_SecretSpec(in buildapiv1.SecretSpec, out *buildapiv1.SecretSpec, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.SecretSource); err != nil {
		return err
	} else {
		out.SecretSource = newVal.(pkgapiv1.LocalObjectReference)
	}
	out.MountPath = in.MountPath
	return nil
}

func deepCopy_v1_SourceBuildStrategy(in buildapiv1.SourceBuildStrategy, out *buildapiv1.SourceBuildStrategy, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	if in.PullSecret != nil {
		if newVal, err := c.DeepCopy(in.PullSecret); err != nil {
			return err
		} else {
			out.PullSecret = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.PullSecret = nil
	}
	if in.Env != nil {
		out.Env = make([]pkgapiv1.EnvVar, len(in.Env))
		for i := range in.Env {
			if newVal, err := c.DeepCopy(in.Env[i]); err != nil {
				return err
			} else {
				out.Env[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Env = nil
	}
	out.Scripts = in.Scripts
	out.Incremental = in.Incremental
	out.ForcePull = in.ForcePull
	return nil
}

func deepCopy_v1_SourceControlUser(in buildapiv1.SourceControlUser, out *buildapiv1.SourceControlUser, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Email = in.Email
	return nil
}

func deepCopy_v1_SourceRevision(in buildapiv1.SourceRevision, out *buildapiv1.SourceRevision, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.Git != nil {
		out.Git = new(buildapiv1.GitSourceRevision)
		if err := deepCopy_v1_GitSourceRevision(*in.Git, out.Git, c); err != nil {
			return err
		}
	} else {
		out.Git = nil
	}
	return nil
}

func deepCopy_v1_WebHookTrigger(in buildapiv1.WebHookTrigger, out *buildapiv1.WebHookTrigger, c *conversion.Cloner) error {
	out.Secret = in.Secret
	return nil
}

func deepCopy_v1_CustomDeploymentStrategyParams(in deployapiv1.CustomDeploymentStrategyParams, out *deployapiv1.CustomDeploymentStrategyParams, c *conversion.Cloner) error {
	out.Image = in.Image
	if in.Environment != nil {
		out.Environment = make([]pkgapiv1.EnvVar, len(in.Environment))
		for i := range in.Environment {
			if newVal, err := c.DeepCopy(in.Environment[i]); err != nil {
				return err
			} else {
				out.Environment[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Environment = nil
	}
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	return nil
}

func deepCopy_v1_DeploymentCause(in deployapiv1.DeploymentCause, out *deployapiv1.DeploymentCause, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.ImageTrigger != nil {
		out.ImageTrigger = new(deployapiv1.DeploymentCauseImageTrigger)
		if err := deepCopy_v1_DeploymentCauseImageTrigger(*in.ImageTrigger, out.ImageTrigger, c); err != nil {
			return err
		}
	} else {
		out.ImageTrigger = nil
	}
	return nil
}

func deepCopy_v1_DeploymentCauseImageTrigger(in deployapiv1.DeploymentCauseImageTrigger, out *deployapiv1.DeploymentCauseImageTrigger, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	return nil
}

func deepCopy_v1_DeploymentConfig(in deployapiv1.DeploymentConfig, out *deployapiv1.DeploymentConfig, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_DeploymentConfigSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_DeploymentConfigStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_DeploymentConfigList(in deployapiv1.DeploymentConfigList, out *deployapiv1.DeploymentConfigList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]deployapiv1.DeploymentConfig, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_DeploymentConfig(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_DeploymentConfigRollback(in deployapiv1.DeploymentConfigRollback, out *deployapiv1.DeploymentConfigRollback, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if err := deepCopy_v1_DeploymentConfigRollbackSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_DeploymentConfigRollbackSpec(in deployapiv1.DeploymentConfigRollbackSpec, out *deployapiv1.DeploymentConfigRollbackSpec, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	out.IncludeTriggers = in.IncludeTriggers
	out.IncludeTemplate = in.IncludeTemplate
	out.IncludeReplicationMeta = in.IncludeReplicationMeta
	out.IncludeStrategy = in.IncludeStrategy
	return nil
}

func deepCopy_v1_DeploymentConfigSpec(in deployapiv1.DeploymentConfigSpec, out *deployapiv1.DeploymentConfigSpec, c *conversion.Cloner) error {
	if err := deepCopy_v1_DeploymentStrategy(in.Strategy, &out.Strategy, c); err != nil {
		return err
	}
	if in.Triggers != nil {
		out.Triggers = make([]deployapiv1.DeploymentTriggerPolicy, len(in.Triggers))
		for i := range in.Triggers {
			if err := deepCopy_v1_DeploymentTriggerPolicy(in.Triggers[i], &out.Triggers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Triggers = nil
	}
	out.Replicas = in.Replicas
	out.Test = in.Test
	if in.Selector != nil {
		out.Selector = make(map[string]string)
		for key, val := range in.Selector {
			out.Selector[key] = val
		}
	} else {
		out.Selector = nil
	}
	if in.Template != nil {
		if newVal, err := c.DeepCopy(in.Template); err != nil {
			return err
		} else {
			out.Template = newVal.(*pkgapiv1.PodTemplateSpec)
		}
	} else {
		out.Template = nil
	}
	return nil
}

func deepCopy_v1_DeploymentConfigStatus(in deployapiv1.DeploymentConfigStatus, out *deployapiv1.DeploymentConfigStatus, c *conversion.Cloner) error {
	out.LatestVersion = in.LatestVersion
	if in.Details != nil {
		out.Details = new(deployapiv1.DeploymentDetails)
		if err := deepCopy_v1_DeploymentDetails(*in.Details, out.Details, c); err != nil {
			return err
		}
	} else {
		out.Details = nil
	}
	return nil
}

func deepCopy_v1_DeploymentDetails(in deployapiv1.DeploymentDetails, out *deployapiv1.DeploymentDetails, c *conversion.Cloner) error {
	out.Message = in.Message
	if in.Causes != nil {
		out.Causes = make([]*deployapiv1.DeploymentCause, len(in.Causes))
		for i := range in.Causes {
			if newVal, err := c.DeepCopy(in.Causes[i]); err != nil {
				return err
			} else if newVal == nil {
				out.Causes[i] = nil
			} else {
				out.Causes[i] = newVal.(*deployapiv1.DeploymentCause)
			}
		}
	} else {
		out.Causes = nil
	}
	return nil
}

func deepCopy_v1_DeploymentLog(in deployapiv1.DeploymentLog, out *deployapiv1.DeploymentLog, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	return nil
}

func deepCopy_v1_DeploymentLogOptions(in deployapiv1.DeploymentLogOptions, out *deployapiv1.DeploymentLogOptions, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	out.Container = in.Container
	out.Follow = in.Follow
	out.Previous = in.Previous
	if in.SinceSeconds != nil {
		out.SinceSeconds = new(int64)
		*out.SinceSeconds = *in.SinceSeconds
	} else {
		out.SinceSeconds = nil
	}
	if in.SinceTime != nil {
		if newVal, err := c.DeepCopy(in.SinceTime); err != nil {
			return err
		} else {
			out.SinceTime = newVal.(*unversioned.Time)
		}
	} else {
		out.SinceTime = nil
	}
	out.Timestamps = in.Timestamps
	if in.TailLines != nil {
		out.TailLines = new(int64)
		*out.TailLines = *in.TailLines
	} else {
		out.TailLines = nil
	}
	if in.LimitBytes != nil {
		out.LimitBytes = new(int64)
		*out.LimitBytes = *in.LimitBytes
	} else {
		out.LimitBytes = nil
	}
	out.NoWait = in.NoWait
	if in.Version != nil {
		out.Version = new(int64)
		*out.Version = *in.Version
	} else {
		out.Version = nil
	}
	return nil
}

func deepCopy_v1_DeploymentStrategy(in deployapiv1.DeploymentStrategy, out *deployapiv1.DeploymentStrategy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.CustomParams != nil {
		out.CustomParams = new(deployapiv1.CustomDeploymentStrategyParams)
		if err := deepCopy_v1_CustomDeploymentStrategyParams(*in.CustomParams, out.CustomParams, c); err != nil {
			return err
		}
	} else {
		out.CustomParams = nil
	}
	if in.RecreateParams != nil {
		out.RecreateParams = new(deployapiv1.RecreateDeploymentStrategyParams)
		if err := deepCopy_v1_RecreateDeploymentStrategyParams(*in.RecreateParams, out.RecreateParams, c); err != nil {
			return err
		}
	} else {
		out.RecreateParams = nil
	}
	if in.RollingParams != nil {
		out.RollingParams = new(deployapiv1.RollingDeploymentStrategyParams)
		if err := deepCopy_v1_RollingDeploymentStrategyParams(*in.RollingParams, out.RollingParams, c); err != nil {
			return err
		}
	} else {
		out.RollingParams = nil
	}
	if newVal, err := c.DeepCopy(in.Resources); err != nil {
		return err
	} else {
		out.Resources = newVal.(pkgapiv1.ResourceRequirements)
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	return nil
}

func deepCopy_v1_DeploymentTriggerImageChangeParams(in deployapiv1.DeploymentTriggerImageChangeParams, out *deployapiv1.DeploymentTriggerImageChangeParams, c *conversion.Cloner) error {
	out.Automatic = in.Automatic
	if in.ContainerNames != nil {
		out.ContainerNames = make([]string, len(in.ContainerNames))
		for i := range in.ContainerNames {
			out.ContainerNames[i] = in.ContainerNames[i]
		}
	} else {
		out.ContainerNames = nil
	}
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	out.LastTriggeredImage = in.LastTriggeredImage
	return nil
}

func deepCopy_v1_DeploymentTriggerPolicy(in deployapiv1.DeploymentTriggerPolicy, out *deployapiv1.DeploymentTriggerPolicy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.ImageChangeParams != nil {
		out.ImageChangeParams = new(deployapiv1.DeploymentTriggerImageChangeParams)
		if err := deepCopy_v1_DeploymentTriggerImageChangeParams(*in.ImageChangeParams, out.ImageChangeParams, c); err != nil {
			return err
		}
	} else {
		out.ImageChangeParams = nil
	}
	return nil
}

func deepCopy_v1_ExecNewPodHook(in deployapiv1.ExecNewPodHook, out *deployapiv1.ExecNewPodHook, c *conversion.Cloner) error {
	if in.Command != nil {
		out.Command = make([]string, len(in.Command))
		for i := range in.Command {
			out.Command[i] = in.Command[i]
		}
	} else {
		out.Command = nil
	}
	if in.Env != nil {
		out.Env = make([]pkgapiv1.EnvVar, len(in.Env))
		for i := range in.Env {
			if newVal, err := c.DeepCopy(in.Env[i]); err != nil {
				return err
			} else {
				out.Env[i] = newVal.(pkgapiv1.EnvVar)
			}
		}
	} else {
		out.Env = nil
	}
	out.ContainerName = in.ContainerName
	if in.Volumes != nil {
		out.Volumes = make([]string, len(in.Volumes))
		for i := range in.Volumes {
			out.Volumes[i] = in.Volumes[i]
		}
	} else {
		out.Volumes = nil
	}
	return nil
}

func deepCopy_v1_LifecycleHook(in deployapiv1.LifecycleHook, out *deployapiv1.LifecycleHook, c *conversion.Cloner) error {
	out.FailurePolicy = in.FailurePolicy
	if in.ExecNewPod != nil {
		out.ExecNewPod = new(deployapiv1.ExecNewPodHook)
		if err := deepCopy_v1_ExecNewPodHook(*in.ExecNewPod, out.ExecNewPod, c); err != nil {
			return err
		}
	} else {
		out.ExecNewPod = nil
	}
	if in.TagImages != nil {
		out.TagImages = make([]deployapiv1.TagImageHook, len(in.TagImages))
		for i := range in.TagImages {
			if err := deepCopy_v1_TagImageHook(in.TagImages[i], &out.TagImages[i], c); err != nil {
				return err
			}
		}
	} else {
		out.TagImages = nil
	}
	return nil
}

func deepCopy_v1_RecreateDeploymentStrategyParams(in deployapiv1.RecreateDeploymentStrategyParams, out *deployapiv1.RecreateDeploymentStrategyParams, c *conversion.Cloner) error {
	if in.TimeoutSeconds != nil {
		out.TimeoutSeconds = new(int64)
		*out.TimeoutSeconds = *in.TimeoutSeconds
	} else {
		out.TimeoutSeconds = nil
	}
	if in.Pre != nil {
		out.Pre = new(deployapiv1.LifecycleHook)
		if err := deepCopy_v1_LifecycleHook(*in.Pre, out.Pre, c); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Mid != nil {
		out.Mid = new(deployapiv1.LifecycleHook)
		if err := deepCopy_v1_LifecycleHook(*in.Mid, out.Mid, c); err != nil {
			return err
		}
	} else {
		out.Mid = nil
	}
	if in.Post != nil {
		out.Post = new(deployapiv1.LifecycleHook)
		if err := deepCopy_v1_LifecycleHook(*in.Post, out.Post, c); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func deepCopy_v1_RollingDeploymentStrategyParams(in deployapiv1.RollingDeploymentStrategyParams, out *deployapiv1.RollingDeploymentStrategyParams, c *conversion.Cloner) error {
	if in.UpdatePeriodSeconds != nil {
		out.UpdatePeriodSeconds = new(int64)
		*out.UpdatePeriodSeconds = *in.UpdatePeriodSeconds
	} else {
		out.UpdatePeriodSeconds = nil
	}
	if in.IntervalSeconds != nil {
		out.IntervalSeconds = new(int64)
		*out.IntervalSeconds = *in.IntervalSeconds
	} else {
		out.IntervalSeconds = nil
	}
	if in.TimeoutSeconds != nil {
		out.TimeoutSeconds = new(int64)
		*out.TimeoutSeconds = *in.TimeoutSeconds
	} else {
		out.TimeoutSeconds = nil
	}
	if in.MaxUnavailable != nil {
		if newVal, err := c.DeepCopy(in.MaxUnavailable); err != nil {
			return err
		} else {
			out.MaxUnavailable = newVal.(*intstr.IntOrString)
		}
	} else {
		out.MaxUnavailable = nil
	}
	if in.MaxSurge != nil {
		if newVal, err := c.DeepCopy(in.MaxSurge); err != nil {
			return err
		} else {
			out.MaxSurge = newVal.(*intstr.IntOrString)
		}
	} else {
		out.MaxSurge = nil
	}
	if in.UpdatePercent != nil {
		out.UpdatePercent = new(int)
		*out.UpdatePercent = *in.UpdatePercent
	} else {
		out.UpdatePercent = nil
	}
	if in.Pre != nil {
		out.Pre = new(deployapiv1.LifecycleHook)
		if err := deepCopy_v1_LifecycleHook(*in.Pre, out.Pre, c); err != nil {
			return err
		}
	} else {
		out.Pre = nil
	}
	if in.Post != nil {
		out.Post = new(deployapiv1.LifecycleHook)
		if err := deepCopy_v1_LifecycleHook(*in.Post, out.Post, c); err != nil {
			return err
		}
	} else {
		out.Post = nil
	}
	return nil
}

func deepCopy_v1_TagImageHook(in deployapiv1.TagImageHook, out *deployapiv1.TagImageHook, c *conversion.Cloner) error {
	out.ContainerName = in.ContainerName
	if newVal, err := c.DeepCopy(in.To); err != nil {
		return err
	} else {
		out.To = newVal.(pkgapiv1.ObjectReference)
	}
	return nil
}

func deepCopy_v1_Image(in imageapiv1.Image, out *imageapiv1.Image, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.DockerImageReference = in.DockerImageReference
	if newVal, err := c.DeepCopy(in.DockerImageMetadata); err != nil {
		return err
	} else {
		out.DockerImageMetadata = newVal.(runtime.RawExtension)
	}
	out.DockerImageMetadataVersion = in.DockerImageMetadataVersion
	out.DockerImageManifest = in.DockerImageManifest
	if in.DockerImageLayers != nil {
		out.DockerImageLayers = make([]imageapiv1.ImageLayer, len(in.DockerImageLayers))
		for i := range in.DockerImageLayers {
			if err := deepCopy_v1_ImageLayer(in.DockerImageLayers[i], &out.DockerImageLayers[i], c); err != nil {
				return err
			}
		}
	} else {
		out.DockerImageLayers = nil
	}
	return nil
}

func deepCopy_v1_ImageImportSpec(in imageapiv1.ImageImportSpec, out *imageapiv1.ImageImportSpec, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	if in.To != nil {
		if newVal, err := c.DeepCopy(in.To); err != nil {
			return err
		} else {
			out.To = newVal.(*pkgapiv1.LocalObjectReference)
		}
	} else {
		out.To = nil
	}
	if err := deepCopy_v1_TagImportPolicy(in.ImportPolicy, &out.ImportPolicy, c); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}

func deepCopy_v1_ImageImportStatus(in imageapiv1.ImageImportStatus, out *imageapiv1.ImageImportStatus, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Status); err != nil {
		return err
	} else {
		out.Status = newVal.(unversioned.Status)
	}
	if in.Image != nil {
		out.Image = new(imageapiv1.Image)
		if err := deepCopy_v1_Image(*in.Image, out.Image, c); err != nil {
			return err
		}
	} else {
		out.Image = nil
	}
	out.Tag = in.Tag
	return nil
}

func deepCopy_v1_ImageLayer(in imageapiv1.ImageLayer, out *imageapiv1.ImageLayer, c *conversion.Cloner) error {
	out.Name = in.Name
	out.Size = in.Size
	return nil
}

func deepCopy_v1_ImageList(in imageapiv1.ImageList, out *imageapiv1.ImageList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]imageapiv1.Image, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Image(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ImageStream(in imageapiv1.ImageStream, out *imageapiv1.ImageStream, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_ImageStreamSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ImageStreamStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ImageStreamImage(in imageapiv1.ImageStreamImage, out *imageapiv1.ImageStreamImage, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_Image(in.Image, &out.Image, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ImageStreamImport(in imageapiv1.ImageStreamImport, out *imageapiv1.ImageStreamImport, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_ImageStreamImportSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ImageStreamImportStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ImageStreamImportSpec(in imageapiv1.ImageStreamImportSpec, out *imageapiv1.ImageStreamImportSpec, c *conversion.Cloner) error {
	out.Import = in.Import
	if in.Repository != nil {
		out.Repository = new(imageapiv1.RepositoryImportSpec)
		if err := deepCopy_v1_RepositoryImportSpec(*in.Repository, out.Repository, c); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		out.Images = make([]imageapiv1.ImageImportSpec, len(in.Images))
		for i := range in.Images {
			if err := deepCopy_v1_ImageImportSpec(in.Images[i], &out.Images[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}

func deepCopy_v1_ImageStreamImportStatus(in imageapiv1.ImageStreamImportStatus, out *imageapiv1.ImageStreamImportStatus, c *conversion.Cloner) error {
	if in.Import != nil {
		out.Import = new(imageapiv1.ImageStream)
		if err := deepCopy_v1_ImageStream(*in.Import, out.Import, c); err != nil {
			return err
		}
	} else {
		out.Import = nil
	}
	if in.Repository != nil {
		out.Repository = new(imageapiv1.RepositoryImportStatus)
		if err := deepCopy_v1_RepositoryImportStatus(*in.Repository, out.Repository, c); err != nil {
			return err
		}
	} else {
		out.Repository = nil
	}
	if in.Images != nil {
		out.Images = make([]imageapiv1.ImageImportStatus, len(in.Images))
		for i := range in.Images {
			if err := deepCopy_v1_ImageImportStatus(in.Images[i], &out.Images[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	return nil
}

func deepCopy_v1_ImageStreamList(in imageapiv1.ImageStreamList, out *imageapiv1.ImageStreamList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]imageapiv1.ImageStream, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ImageStream(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ImageStreamMapping(in imageapiv1.ImageStreamMapping, out *imageapiv1.ImageStreamMapping, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_Image(in.Image, &out.Image, c); err != nil {
		return err
	}
	out.Tag = in.Tag
	return nil
}

func deepCopy_v1_ImageStreamSpec(in imageapiv1.ImageStreamSpec, out *imageapiv1.ImageStreamSpec, c *conversion.Cloner) error {
	out.DockerImageRepository = in.DockerImageRepository
	if in.Tags != nil {
		out.Tags = make([]imageapiv1.TagReference, len(in.Tags))
		for i := range in.Tags {
			if err := deepCopy_v1_TagReference(in.Tags[i], &out.Tags[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Tags = nil
	}
	return nil
}

func deepCopy_v1_ImageStreamStatus(in imageapiv1.ImageStreamStatus, out *imageapiv1.ImageStreamStatus, c *conversion.Cloner) error {
	out.DockerImageRepository = in.DockerImageRepository
	if in.Tags != nil {
		out.Tags = make([]imageapiv1.NamedTagEventList, len(in.Tags))
		for i := range in.Tags {
			if err := deepCopy_v1_NamedTagEventList(in.Tags[i], &out.Tags[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Tags = nil
	}
	return nil
}

func deepCopy_v1_ImageStreamTag(in imageapiv1.ImageStreamTag, out *imageapiv1.ImageStreamTag, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Tag != nil {
		out.Tag = new(imageapiv1.TagReference)
		if err := deepCopy_v1_TagReference(*in.Tag, out.Tag, c); err != nil {
			return err
		}
	} else {
		out.Tag = nil
	}
	out.Generation = in.Generation
	if in.Conditions != nil {
		out.Conditions = make([]imageapiv1.TagEventCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_v1_TagEventCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if err := deepCopy_v1_Image(in.Image, &out.Image, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ImageStreamTagList(in imageapiv1.ImageStreamTagList, out *imageapiv1.ImageStreamTagList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]imageapiv1.ImageStreamTag, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ImageStreamTag(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_NamedTagEventList(in imageapiv1.NamedTagEventList, out *imageapiv1.NamedTagEventList, c *conversion.Cloner) error {
	out.Tag = in.Tag
	if in.Items != nil {
		out.Items = make([]imageapiv1.TagEvent, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_TagEvent(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	if in.Conditions != nil {
		out.Conditions = make([]imageapiv1.TagEventCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_v1_TagEventCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func deepCopy_v1_RepositoryImportSpec(in imageapiv1.RepositoryImportSpec, out *imageapiv1.RepositoryImportSpec, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.From); err != nil {
		return err
	} else {
		out.From = newVal.(pkgapiv1.ObjectReference)
	}
	if err := deepCopy_v1_TagImportPolicy(in.ImportPolicy, &out.ImportPolicy, c); err != nil {
		return err
	}
	out.IncludeManifest = in.IncludeManifest
	return nil
}

func deepCopy_v1_RepositoryImportStatus(in imageapiv1.RepositoryImportStatus, out *imageapiv1.RepositoryImportStatus, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Status); err != nil {
		return err
	} else {
		out.Status = newVal.(unversioned.Status)
	}
	if in.Images != nil {
		out.Images = make([]imageapiv1.ImageImportStatus, len(in.Images))
		for i := range in.Images {
			if err := deepCopy_v1_ImageImportStatus(in.Images[i], &out.Images[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Images = nil
	}
	if in.AdditionalTags != nil {
		out.AdditionalTags = make([]string, len(in.AdditionalTags))
		for i := range in.AdditionalTags {
			out.AdditionalTags[i] = in.AdditionalTags[i]
		}
	} else {
		out.AdditionalTags = nil
	}
	return nil
}

func deepCopy_v1_TagEvent(in imageapiv1.TagEvent, out *imageapiv1.TagEvent, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.Created); err != nil {
		return err
	} else {
		out.Created = newVal.(unversioned.Time)
	}
	out.DockerImageReference = in.DockerImageReference
	out.Image = in.Image
	out.Generation = in.Generation
	return nil
}

func deepCopy_v1_TagEventCondition(in imageapiv1.TagEventCondition, out *imageapiv1.TagEventCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	if newVal, err := c.DeepCopy(in.LastTransitionTime); err != nil {
		return err
	} else {
		out.LastTransitionTime = newVal.(unversioned.Time)
	}
	out.Reason = in.Reason
	out.Message = in.Message
	out.Generation = in.Generation
	return nil
}

func deepCopy_v1_TagImportPolicy(in imageapiv1.TagImportPolicy, out *imageapiv1.TagImportPolicy, c *conversion.Cloner) error {
	out.Insecure = in.Insecure
	out.Scheduled = in.Scheduled
	return nil
}

func deepCopy_v1_TagReference(in imageapiv1.TagReference, out *imageapiv1.TagReference, c *conversion.Cloner) error {
	out.Name = in.Name
	if in.Annotations != nil {
		out.Annotations = make(map[string]string)
		for key, val := range in.Annotations {
			out.Annotations[key] = val
		}
	} else {
		out.Annotations = nil
	}
	if in.From != nil {
		if newVal, err := c.DeepCopy(in.From); err != nil {
			return err
		} else {
			out.From = newVal.(*pkgapiv1.ObjectReference)
		}
	} else {
		out.From = nil
	}
	out.Reference = in.Reference
	if in.Generation != nil {
		out.Generation = new(int64)
		*out.Generation = *in.Generation
	} else {
		out.Generation = nil
	}
	if err := deepCopy_v1_TagImportPolicy(in.ImportPolicy, &out.ImportPolicy, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_OAuthAccessToken(in oauthapiv1.OAuthAccessToken, out *oauthapiv1.OAuthAccessToken, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	out.AuthorizeToken = in.AuthorizeToken
	out.RefreshToken = in.RefreshToken
	return nil
}

func deepCopy_v1_OAuthAccessTokenList(in oauthapiv1.OAuthAccessTokenList, out *oauthapiv1.OAuthAccessTokenList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthAccessToken, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_OAuthAccessToken(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_OAuthAuthorizeToken(in oauthapiv1.OAuthAuthorizeToken, out *oauthapiv1.OAuthAuthorizeToken, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.ClientName = in.ClientName
	out.ExpiresIn = in.ExpiresIn
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	out.RedirectURI = in.RedirectURI
	out.State = in.State
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	return nil
}

func deepCopy_v1_OAuthAuthorizeTokenList(in oauthapiv1.OAuthAuthorizeTokenList, out *oauthapiv1.OAuthAuthorizeTokenList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthAuthorizeToken, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_OAuthAuthorizeToken(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_OAuthClient(in oauthapiv1.OAuthClient, out *oauthapiv1.OAuthClient, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.Secret = in.Secret
	out.RespondWithChallenges = in.RespondWithChallenges
	if in.RedirectURIs != nil {
		out.RedirectURIs = make([]string, len(in.RedirectURIs))
		for i := range in.RedirectURIs {
			out.RedirectURIs[i] = in.RedirectURIs[i]
		}
	} else {
		out.RedirectURIs = nil
	}
	return nil
}

func deepCopy_v1_OAuthClientAuthorization(in oauthapiv1.OAuthClientAuthorization, out *oauthapiv1.OAuthClientAuthorization, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.ClientName = in.ClientName
	out.UserName = in.UserName
	out.UserUID = in.UserUID
	if in.Scopes != nil {
		out.Scopes = make([]string, len(in.Scopes))
		for i := range in.Scopes {
			out.Scopes[i] = in.Scopes[i]
		}
	} else {
		out.Scopes = nil
	}
	return nil
}

func deepCopy_v1_OAuthClientAuthorizationList(in oauthapiv1.OAuthClientAuthorizationList, out *oauthapiv1.OAuthClientAuthorizationList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthClientAuthorization, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_OAuthClientAuthorization(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_OAuthClientList(in oauthapiv1.OAuthClientList, out *oauthapiv1.OAuthClientList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]oauthapiv1.OAuthClient, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_OAuthClient(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_Project(in projectapiv1.Project, out *projectapiv1.Project, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_ProjectSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ProjectStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ProjectList(in projectapiv1.ProjectList, out *projectapiv1.ProjectList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]projectapiv1.Project, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Project(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ProjectRequest(in projectapiv1.ProjectRequest, out *projectapiv1.ProjectRequest, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	return nil
}

func deepCopy_v1_ProjectSpec(in projectapiv1.ProjectSpec, out *projectapiv1.ProjectSpec, c *conversion.Cloner) error {
	if in.Finalizers != nil {
		out.Finalizers = make([]pkgapiv1.FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = in.Finalizers[i]
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func deepCopy_v1_ProjectStatus(in projectapiv1.ProjectStatus, out *projectapiv1.ProjectStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	return nil
}

func deepCopy_v1_Route(in routeapiv1.Route, out *routeapiv1.Route, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_RouteSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_RouteStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_RouteIngress(in routeapiv1.RouteIngress, out *routeapiv1.RouteIngress, c *conversion.Cloner) error {
	out.Host = in.Host
	out.RouterName = in.RouterName
	if in.Conditions != nil {
		out.Conditions = make([]routeapiv1.RouteIngressCondition, len(in.Conditions))
		for i := range in.Conditions {
			if err := deepCopy_v1_RouteIngressCondition(in.Conditions[i], &out.Conditions[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

func deepCopy_v1_RouteIngressCondition(in routeapiv1.RouteIngressCondition, out *routeapiv1.RouteIngressCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	out.Reason = in.Reason
	out.Message = in.Message
	if in.LastTransitionTime != nil {
		if newVal, err := c.DeepCopy(in.LastTransitionTime); err != nil {
			return err
		} else {
			out.LastTransitionTime = newVal.(*unversioned.Time)
		}
	} else {
		out.LastTransitionTime = nil
	}
	return nil
}

func deepCopy_v1_RouteList(in routeapiv1.RouteList, out *routeapiv1.RouteList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]routeapiv1.Route, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Route(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_RoutePort(in routeapiv1.RoutePort, out *routeapiv1.RoutePort, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TargetPort); err != nil {
		return err
	} else {
		out.TargetPort = newVal.(intstr.IntOrString)
	}
	return nil
}

func deepCopy_v1_RouteSpec(in routeapiv1.RouteSpec, out *routeapiv1.RouteSpec, c *conversion.Cloner) error {
	out.Host = in.Host
	out.Path = in.Path
	if newVal, err := c.DeepCopy(in.To); err != nil {
		return err
	} else {
		out.To = newVal.(pkgapiv1.ObjectReference)
	}
	if in.Port != nil {
		out.Port = new(routeapiv1.RoutePort)
		if err := deepCopy_v1_RoutePort(*in.Port, out.Port, c); err != nil {
			return err
		}
	} else {
		out.Port = nil
	}
	if in.TLS != nil {
		out.TLS = new(routeapiv1.TLSConfig)
		if err := deepCopy_v1_TLSConfig(*in.TLS, out.TLS, c); err != nil {
			return err
		}
	} else {
		out.TLS = nil
	}
	return nil
}

func deepCopy_v1_RouteStatus(in routeapiv1.RouteStatus, out *routeapiv1.RouteStatus, c *conversion.Cloner) error {
	if in.Ingress != nil {
		out.Ingress = make([]routeapiv1.RouteIngress, len(in.Ingress))
		for i := range in.Ingress {
			if err := deepCopy_v1_RouteIngress(in.Ingress[i], &out.Ingress[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func deepCopy_v1_TLSConfig(in routeapiv1.TLSConfig, out *routeapiv1.TLSConfig, c *conversion.Cloner) error {
	out.Termination = in.Termination
	out.Certificate = in.Certificate
	out.Key = in.Key
	out.CACertificate = in.CACertificate
	out.DestinationCACertificate = in.DestinationCACertificate
	out.InsecureEdgeTerminationPolicy = in.InsecureEdgeTerminationPolicy
	return nil
}

func deepCopy_v1_ClusterNetwork(in sdnapiv1.ClusterNetwork, out *sdnapiv1.ClusterNetwork, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.Network = in.Network
	out.HostSubnetLength = in.HostSubnetLength
	out.ServiceNetwork = in.ServiceNetwork
	return nil
}

func deepCopy_v1_ClusterNetworkList(in sdnapiv1.ClusterNetworkList, out *sdnapiv1.ClusterNetworkList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]sdnapiv1.ClusterNetwork, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ClusterNetwork(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_HostSubnet(in sdnapiv1.HostSubnet, out *sdnapiv1.HostSubnet, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.Host = in.Host
	out.HostIP = in.HostIP
	out.Subnet = in.Subnet
	return nil
}

func deepCopy_v1_HostSubnetList(in sdnapiv1.HostSubnetList, out *sdnapiv1.HostSubnetList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]sdnapiv1.HostSubnet, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_HostSubnet(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_NetNamespace(in sdnapiv1.NetNamespace, out *sdnapiv1.NetNamespace, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.NetName = in.NetName
	out.NetID = in.NetID
	return nil
}

func deepCopy_v1_NetNamespaceList(in sdnapiv1.NetNamespaceList, out *sdnapiv1.NetNamespaceList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]sdnapiv1.NetNamespace, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_NetNamespace(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ServiceBroker(in servicebrokerapiv1.ServiceBroker, out *servicebrokerapiv1.ServiceBroker, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if err := deepCopy_v1_ServiceBrokerSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := deepCopy_v1_ServiceBrokerStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func deepCopy_v1_ServiceBrokerList(in servicebrokerapiv1.ServiceBrokerList, out *servicebrokerapiv1.ServiceBrokerList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]servicebrokerapiv1.ServiceBroker, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_ServiceBroker(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_ServiceBrokerSpec(in servicebrokerapiv1.ServiceBrokerSpec, out *servicebrokerapiv1.ServiceBrokerSpec, c *conversion.Cloner) error {
	out.Url = in.Url
	out.Name = in.Name
	out.UserName = in.UserName
	out.Password = in.Password
	if in.Finalizers != nil {
		out.Finalizers = make([]pkgapiv1.FinalizerName, len(in.Finalizers))
		for i := range in.Finalizers {
			out.Finalizers[i] = in.Finalizers[i]
		}
	} else {
		out.Finalizers = nil
	}
	return nil
}

func deepCopy_v1_ServiceBrokerStatus(in servicebrokerapiv1.ServiceBrokerStatus, out *servicebrokerapiv1.ServiceBrokerStatus, c *conversion.Cloner) error {
	out.Phase = in.Phase
	return nil
}

func deepCopy_v1_Parameter(in templateapiv1.Parameter, out *templateapiv1.Parameter, c *conversion.Cloner) error {
	out.Name = in.Name
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Value = in.Value
	out.Generate = in.Generate
	out.From = in.From
	out.Required = in.Required
	return nil
}

func deepCopy_v1_Template(in templateapiv1.Template, out *templateapiv1.Template, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Objects != nil {
		out.Objects = make([]runtime.RawExtension, len(in.Objects))
		for i := range in.Objects {
			if newVal, err := c.DeepCopy(in.Objects[i]); err != nil {
				return err
			} else {
				out.Objects[i] = newVal.(runtime.RawExtension)
			}
		}
	} else {
		out.Objects = nil
	}
	if in.Parameters != nil {
		out.Parameters = make([]templateapiv1.Parameter, len(in.Parameters))
		for i := range in.Parameters {
			if err := deepCopy_v1_Parameter(in.Parameters[i], &out.Parameters[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Parameters = nil
	}
	if in.Labels != nil {
		out.Labels = make(map[string]string)
		for key, val := range in.Labels {
			out.Labels[key] = val
		}
	} else {
		out.Labels = nil
	}
	return nil
}

func deepCopy_v1_TemplateList(in templateapiv1.TemplateList, out *templateapiv1.TemplateList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]templateapiv1.Template, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Template(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_Group(in userapiv1.Group, out *userapiv1.Group, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if in.Users != nil {
		out.Users = make([]string, len(in.Users))
		for i := range in.Users {
			out.Users[i] = in.Users[i]
		}
	} else {
		out.Users = nil
	}
	return nil
}

func deepCopy_v1_GroupList(in userapiv1.GroupList, out *userapiv1.GroupList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]userapiv1.Group, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Group(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_Identity(in userapiv1.Identity, out *userapiv1.Identity, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.ProviderName = in.ProviderName
	out.ProviderUserName = in.ProviderUserName
	if newVal, err := c.DeepCopy(in.User); err != nil {
		return err
	} else {
		out.User = newVal.(pkgapiv1.ObjectReference)
	}
	if in.Extra != nil {
		out.Extra = make(map[string]string)
		for key, val := range in.Extra {
			out.Extra[key] = val
		}
	} else {
		out.Extra = nil
	}
	return nil
}

func deepCopy_v1_IdentityList(in userapiv1.IdentityList, out *userapiv1.IdentityList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]userapiv1.Identity, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_Identity(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func deepCopy_v1_User(in userapiv1.User, out *userapiv1.User, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	out.FullName = in.FullName
	if in.Identities != nil {
		out.Identities = make([]string, len(in.Identities))
		for i := range in.Identities {
			out.Identities[i] = in.Identities[i]
		}
	} else {
		out.Identities = nil
	}
	if in.Groups != nil {
		out.Groups = make([]string, len(in.Groups))
		for i := range in.Groups {
			out.Groups[i] = in.Groups[i]
		}
	} else {
		out.Groups = nil
	}
	return nil
}

func deepCopy_v1_UserIdentityMapping(in userapiv1.UserIdentityMapping, out *userapiv1.UserIdentityMapping, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ObjectMeta); err != nil {
		return err
	} else {
		out.ObjectMeta = newVal.(pkgapiv1.ObjectMeta)
	}
	if newVal, err := c.DeepCopy(in.Identity); err != nil {
		return err
	} else {
		out.Identity = newVal.(pkgapiv1.ObjectReference)
	}
	if newVal, err := c.DeepCopy(in.User); err != nil {
		return err
	} else {
		out.User = newVal.(pkgapiv1.ObjectReference)
	}
	return nil
}

func deepCopy_v1_UserList(in userapiv1.UserList, out *userapiv1.UserList, c *conversion.Cloner) error {
	if newVal, err := c.DeepCopy(in.TypeMeta); err != nil {
		return err
	} else {
		out.TypeMeta = newVal.(unversioned.TypeMeta)
	}
	if newVal, err := c.DeepCopy(in.ListMeta); err != nil {
		return err
	} else {
		out.ListMeta = newVal.(unversioned.ListMeta)
	}
	if in.Items != nil {
		out.Items = make([]userapiv1.User, len(in.Items))
		for i := range in.Items {
			if err := deepCopy_v1_User(in.Items[i], &out.Items[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func init() {
	err := api.Scheme.AddGeneratedDeepCopyFuncs(
		deepCopy_v1_Application,
		deepCopy_v1_ApplicationList,
		deepCopy_v1_ApplicationSpec,
		deepCopy_v1_ApplicationStatus,
		deepCopy_v1_Item,
		deepCopy_v1_AuthorizationAttributes,
		deepCopy_v1_ClusterPolicy,
		deepCopy_v1_ClusterPolicyBinding,
		deepCopy_v1_ClusterPolicyBindingList,
		deepCopy_v1_ClusterPolicyList,
		deepCopy_v1_ClusterRole,
		deepCopy_v1_ClusterRoleBinding,
		deepCopy_v1_ClusterRoleBindingList,
		deepCopy_v1_ClusterRoleList,
		deepCopy_v1_IsPersonalSubjectAccessReview,
		deepCopy_v1_LocalResourceAccessReview,
		deepCopy_v1_LocalSubjectAccessReview,
		deepCopy_v1_NamedClusterRole,
		deepCopy_v1_NamedClusterRoleBinding,
		deepCopy_v1_NamedRole,
		deepCopy_v1_NamedRoleBinding,
		deepCopy_v1_Policy,
		deepCopy_v1_PolicyBinding,
		deepCopy_v1_PolicyBindingList,
		deepCopy_v1_PolicyList,
		deepCopy_v1_PolicyRule,
		deepCopy_v1_ResourceAccessReview,
		deepCopy_v1_ResourceAccessReviewResponse,
		deepCopy_v1_Role,
		deepCopy_v1_RoleBinding,
		deepCopy_v1_RoleBindingList,
		deepCopy_v1_RoleList,
		deepCopy_v1_SubjectAccessReview,
		deepCopy_v1_SubjectAccessReviewResponse,
		deepCopy_v1_BackingService,
		deepCopy_v1_BackingServiceList,
		deepCopy_v1_BackingServiceSpec,
		deepCopy_v1_BackingServiceStatus,
		deepCopy_v1_CustomizeSpec,
		deepCopy_v1_ServicePlan,
		deepCopy_v1_ServicePlanCost,
		deepCopy_v1_ServicePlanMetadata,
		deepCopy_v1_BackingServiceInstance,
		deepCopy_v1_BackingServiceInstanceList,
		deepCopy_v1_BackingServiceInstanceSpec,
		deepCopy_v1_BackingServiceInstanceStatus,
		deepCopy_v1_BindingRequestOptions,
		deepCopy_v1_InstanceBinding,
		deepCopy_v1_InstanceProvisioning,
		deepCopy_v1_LastOperation,
		deepCopy_v1_UserProvidedService,
		deepCopy_v1_BinaryBuildRequestOptions,
		deepCopy_v1_BinaryBuildSource,
		deepCopy_v1_Build,
		deepCopy_v1_BuildConfig,
		deepCopy_v1_BuildConfigList,
		deepCopy_v1_BuildConfigSpec,
		deepCopy_v1_BuildConfigStatus,
		deepCopy_v1_BuildList,
		deepCopy_v1_BuildLog,
		deepCopy_v1_BuildLogOptions,
		deepCopy_v1_BuildOutput,
		deepCopy_v1_BuildPostCommitSpec,
		deepCopy_v1_BuildRequest,
		deepCopy_v1_BuildSource,
		deepCopy_v1_BuildSpec,
		deepCopy_v1_BuildStatus,
		deepCopy_v1_BuildStrategy,
		deepCopy_v1_BuildTriggerPolicy,
		deepCopy_v1_CustomBuildStrategy,
		deepCopy_v1_DockerBuildStrategy,
		deepCopy_v1_GitBuildSource,
		deepCopy_v1_GitSourceRevision,
		deepCopy_v1_ImageChangeTrigger,
		deepCopy_v1_ImageSource,
		deepCopy_v1_ImageSourcePath,
		deepCopy_v1_SecretBuildSource,
		deepCopy_v1_SecretSpec,
		deepCopy_v1_SourceBuildStrategy,
		deepCopy_v1_SourceControlUser,
		deepCopy_v1_SourceRevision,
		deepCopy_v1_WebHookTrigger,
		deepCopy_v1_CustomDeploymentStrategyParams,
		deepCopy_v1_DeploymentCause,
		deepCopy_v1_DeploymentCauseImageTrigger,
		deepCopy_v1_DeploymentConfig,
		deepCopy_v1_DeploymentConfigList,
		deepCopy_v1_DeploymentConfigRollback,
		deepCopy_v1_DeploymentConfigRollbackSpec,
		deepCopy_v1_DeploymentConfigSpec,
		deepCopy_v1_DeploymentConfigStatus,
		deepCopy_v1_DeploymentDetails,
		deepCopy_v1_DeploymentLog,
		deepCopy_v1_DeploymentLogOptions,
		deepCopy_v1_DeploymentStrategy,
		deepCopy_v1_DeploymentTriggerImageChangeParams,
		deepCopy_v1_DeploymentTriggerPolicy,
		deepCopy_v1_ExecNewPodHook,
		deepCopy_v1_LifecycleHook,
		deepCopy_v1_RecreateDeploymentStrategyParams,
		deepCopy_v1_RollingDeploymentStrategyParams,
		deepCopy_v1_TagImageHook,
		deepCopy_v1_Image,
		deepCopy_v1_ImageImportSpec,
		deepCopy_v1_ImageImportStatus,
		deepCopy_v1_ImageLayer,
		deepCopy_v1_ImageList,
		deepCopy_v1_ImageStream,
		deepCopy_v1_ImageStreamImage,
		deepCopy_v1_ImageStreamImport,
		deepCopy_v1_ImageStreamImportSpec,
		deepCopy_v1_ImageStreamImportStatus,
		deepCopy_v1_ImageStreamList,
		deepCopy_v1_ImageStreamMapping,
		deepCopy_v1_ImageStreamSpec,
		deepCopy_v1_ImageStreamStatus,
		deepCopy_v1_ImageStreamTag,
		deepCopy_v1_ImageStreamTagList,
		deepCopy_v1_NamedTagEventList,
		deepCopy_v1_RepositoryImportSpec,
		deepCopy_v1_RepositoryImportStatus,
		deepCopy_v1_TagEvent,
		deepCopy_v1_TagEventCondition,
		deepCopy_v1_TagImportPolicy,
		deepCopy_v1_TagReference,
		deepCopy_v1_OAuthAccessToken,
		deepCopy_v1_OAuthAccessTokenList,
		deepCopy_v1_OAuthAuthorizeToken,
		deepCopy_v1_OAuthAuthorizeTokenList,
		deepCopy_v1_OAuthClient,
		deepCopy_v1_OAuthClientAuthorization,
		deepCopy_v1_OAuthClientAuthorizationList,
		deepCopy_v1_OAuthClientList,
		deepCopy_v1_Project,
		deepCopy_v1_ProjectList,
		deepCopy_v1_ProjectRequest,
		deepCopy_v1_ProjectSpec,
		deepCopy_v1_ProjectStatus,
		deepCopy_v1_Route,
		deepCopy_v1_RouteIngress,
		deepCopy_v1_RouteIngressCondition,
		deepCopy_v1_RouteList,
		deepCopy_v1_RoutePort,
		deepCopy_v1_RouteSpec,
		deepCopy_v1_RouteStatus,
		deepCopy_v1_TLSConfig,
		deepCopy_v1_ClusterNetwork,
		deepCopy_v1_ClusterNetworkList,
		deepCopy_v1_HostSubnet,
		deepCopy_v1_HostSubnetList,
		deepCopy_v1_NetNamespace,
		deepCopy_v1_NetNamespaceList,
		deepCopy_v1_ServiceBroker,
		deepCopy_v1_ServiceBrokerList,
		deepCopy_v1_ServiceBrokerSpec,
		deepCopy_v1_ServiceBrokerStatus,
		deepCopy_v1_Parameter,
		deepCopy_v1_Template,
		deepCopy_v1_TemplateList,
		deepCopy_v1_Group,
		deepCopy_v1_GroupList,
		deepCopy_v1_Identity,
		deepCopy_v1_IdentityList,
		deepCopy_v1_User,
		deepCopy_v1_UserIdentityMapping,
		deepCopy_v1_UserList,
	)
	if err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

// AUTO-GENERATED FUNCTIONS END HERE
