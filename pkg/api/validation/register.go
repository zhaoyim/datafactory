package validation

import (
	_ "github.com/openshift/origin/pkg/api/install"

	authorizationvalidation "github.com/openshift/origin/pkg/authorization/api/validation"

	applicationvalidation "github.com/openshift/origin/pkg/application/api/validation"
	backingservicevalidation "github.com/openshift/origin/pkg/backingservice/api/validation"
	backingserviceinstancevalidation "github.com/openshift/origin/pkg/backingserviceinstance/api/validation"
	buildvalidation "github.com/openshift/origin/pkg/build/api/validation"
	deployvalidation "github.com/openshift/origin/pkg/deploy/api/validation"
	imagevalidation "github.com/openshift/origin/pkg/image/api/validation"
	oauthvalidation "github.com/openshift/origin/pkg/oauth/api/validation"
	projectvalidation "github.com/openshift/origin/pkg/project/api/validation"
	routevalidation "github.com/openshift/origin/pkg/route/api/validation"
	sdnvalidation "github.com/openshift/origin/pkg/sdn/api/validation"
	servicebrokervalidation "github.com/openshift/origin/pkg/servicebroker/api/validation"
	templatevalidation "github.com/openshift/origin/pkg/template/api/validation"
	uservalidation "github.com/openshift/origin/pkg/user/api/validation"
	extvalidation "k8s.io/kubernetes/pkg/apis/extensions/validation"
	applicationvalidation "github.com/openshift/origin/pkg/application/api/validation"

	applicationapi "github.com/openshift/origin/pkg/application/api"
	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	backingserviceapi "github.com/openshift/origin/pkg/backingservice/api"
	backingserviceinstanceapi "github.com/openshift/origin/pkg/backingserviceinstance/api"
	buildapi "github.com/openshift/origin/pkg/build/api"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	imageapi "github.com/openshift/origin/pkg/image/api"
	oauthapi "github.com/openshift/origin/pkg/oauth/api"
	projectapi "github.com/openshift/origin/pkg/project/api"
	routeapi "github.com/openshift/origin/pkg/route/api"
	sdnapi "github.com/openshift/origin/pkg/sdn/api"
	servicebrokerapi "github.com/openshift/origin/pkg/servicebroker/api"
	templateapi "github.com/openshift/origin/pkg/template/api"
	userapi "github.com/openshift/origin/pkg/user/api"
	"k8s.io/kubernetes/pkg/apis/extensions"

	// required to be loaded before we register
	_ "github.com/openshift/origin/pkg/api/install"
)

func init() {
	registerAll()
}

func registerAll() {
	Validator.MustRegister(&authorizationapi.SubjectAccessReview{}, authorizationvalidation.ValidateSubjectAccessReview, nil)
	Validator.MustRegister(&authorizationapi.ResourceAccessReview{}, authorizationvalidation.ValidateResourceAccessReview, nil)
	Validator.MustRegister(&authorizationapi.LocalSubjectAccessReview{}, authorizationvalidation.ValidateLocalSubjectAccessReview, nil)
	Validator.MustRegister(&authorizationapi.LocalResourceAccessReview{}, authorizationvalidation.ValidateLocalResourceAccessReview, nil)

	Validator.MustRegister(&authorizationapi.Policy{}, authorizationvalidation.ValidateLocalPolicy, authorizationvalidation.ValidateLocalPolicyUpdate)
	Validator.MustRegister(&authorizationapi.PolicyBinding{}, authorizationvalidation.ValidateLocalPolicyBinding, authorizationvalidation.ValidateLocalPolicyBindingUpdate)
	Validator.MustRegister(&authorizationapi.Role{}, authorizationvalidation.ValidateLocalRole, authorizationvalidation.ValidateLocalRoleUpdate)
	Validator.MustRegister(&authorizationapi.RoleBinding{}, authorizationvalidation.ValidateLocalRoleBinding, authorizationvalidation.ValidateLocalRoleBindingUpdate)

	Validator.MustRegister(&authorizationapi.ClusterPolicy{}, authorizationvalidation.ValidateClusterPolicy, authorizationvalidation.ValidateClusterPolicyUpdate)
	Validator.MustRegister(&authorizationapi.ClusterPolicyBinding{}, authorizationvalidation.ValidateClusterPolicyBinding, authorizationvalidation.ValidateClusterPolicyBindingUpdate)
	Validator.MustRegister(&authorizationapi.ClusterRole{}, authorizationvalidation.ValidateClusterRole, authorizationvalidation.ValidateClusterRoleUpdate)
	Validator.MustRegister(&authorizationapi.ClusterRoleBinding{}, authorizationvalidation.ValidateClusterRoleBinding, authorizationvalidation.ValidateClusterRoleBindingUpdate)

	Validator.MustRegister(&buildapi.Build{}, buildvalidation.ValidateBuild, buildvalidation.ValidateBuildUpdate)
	Validator.MustRegister(&buildapi.BuildConfig{}, buildvalidation.ValidateBuildConfig, buildvalidation.ValidateBuildConfigUpdate)
	Validator.MustRegister(&buildapi.BuildRequest{}, buildvalidation.ValidateBuildRequest, nil)
	Validator.MustRegister(&buildapi.BuildLogOptions{}, buildvalidation.ValidateBuildLogOptions, nil)

	Validator.MustRegister(&deployapi.DeploymentConfig{}, deployvalidation.ValidateDeploymentConfig, deployvalidation.ValidateDeploymentConfigUpdate)
	Validator.MustRegister(&deployapi.DeploymentConfigRollback{}, deployvalidation.ValidateDeploymentConfigRollback, nil)
	Validator.MustRegister(&deployapi.DeploymentLogOptions{}, deployvalidation.ValidateDeploymentLogOptions, nil)
	Validator.MustRegister(&extensions.Scale{}, extvalidation.ValidateScale, nil)

	Validator.MustRegister(&imageapi.Image{}, imagevalidation.ValidateImage, imagevalidation.ValidateImageUpdate)
	Validator.MustRegister(&imageapi.ImageStream{}, imagevalidation.ValidateImageStream, imagevalidation.ValidateImageStreamUpdate)
	Validator.MustRegister(&imageapi.ImageStreamImport{}, imagevalidation.ValidateImageStreamImport, nil)
	Validator.MustRegister(&imageapi.ImageStreamMapping{}, imagevalidation.ValidateImageStreamMapping, nil)
	Validator.MustRegister(&imageapi.ImageStreamTag{}, imagevalidation.ValidateImageStreamTag, imagevalidation.ValidateImageStreamTagUpdate)

	Validator.MustRegister(&oauthapi.OAuthAccessToken{}, oauthvalidation.ValidateAccessToken, nil)
	Validator.MustRegister(&oauthapi.OAuthAuthorizeToken{}, oauthvalidation.ValidateAuthorizeToken, nil)
	Validator.MustRegister(&oauthapi.OAuthClient{}, oauthvalidation.ValidateClient, oauthvalidation.ValidateClientUpdate)
	Validator.MustRegister(&oauthapi.OAuthClientAuthorization{}, oauthvalidation.ValidateClientAuthorization, oauthvalidation.ValidateClientAuthorizationUpdate)

	Validator.MustRegister(&projectapi.Project{}, projectvalidation.ValidateProject, projectvalidation.ValidateProjectUpdate)
	Validator.MustRegister(&projectapi.ProjectRequest{}, projectvalidation.ValidateProjectRequest, nil)

	Validator.MustRegister(&routeapi.Route{}, routevalidation.ValidateRoute, routevalidation.ValidateRouteUpdate)

	Validator.MustRegister(&sdnapi.ClusterNetwork{}, sdnvalidation.ValidateClusterNetwork, sdnvalidation.ValidateClusterNetworkUpdate)
	Validator.MustRegister(&sdnapi.HostSubnet{}, sdnvalidation.ValidateHostSubnet, sdnvalidation.ValidateHostSubnetUpdate)
	Validator.MustRegister(&sdnapi.NetNamespace{}, sdnvalidation.ValidateNetNamespace, sdnvalidation.ValidateNetNamespaceUpdate)

	Validator.MustRegister(&templateapi.Template{}, templatevalidation.ValidateTemplate, templatevalidation.ValidateTemplateUpdate)

	Validator.MustRegister(&userapi.User{}, uservalidation.ValidateUser, uservalidation.ValidateUserUpdate)
	Validator.MustRegister(&userapi.Identity{}, uservalidation.ValidateIdentity, uservalidation.ValidateIdentityUpdate)
	Validator.MustRegister(&userapi.UserIdentityMapping{}, uservalidation.ValidateUserIdentityMapping, uservalidation.ValidateUserIdentityMappingUpdate)
	Validator.MustRegister(&userapi.Group{}, uservalidation.ValidateGroup, uservalidation.ValidateGroupUpdate)
	Validator.MustRegister(&applicationapi.Application{}, applicationvalidation.ValidateApplication, applicationvalidation.ValidateApplicationUpdate)
	Validator.MustRegister(&servicebrokerapi.ServiceBroker{}, servicebrokervalidation.ValidateServiceBroker, servicebrokervalidation.ValidateServiceBrokerUpdate)
	Validator.MustRegister(&backingserviceapi.BackingService{}, backingservicevalidation.ValidateBackingService, backingservicevalidation.ValidateBackingServiceUpdate)
	Validator.MustRegister(&backingserviceinstanceapi.BackingServiceInstance{}, backingserviceinstancevalidation.ValidateBackingServiceInstance, backingserviceinstancevalidation.ValidateBackingServiceInstanceUpdate)
	//Validator.MustRegister(&backingserviceinstanceapi.BindingRequest{}, backingserviceinstancevalidation.ValidateBackingServiceInstanceBindingRequest, nil)
	Validator.MustRegister(&backingserviceinstanceapi.BindingRequestOptions{}, backingserviceinstancevalidation.ValidateBackingServiceInstanceBindingRequestOptions, nil)
	Validator.MustRegister(&applicationapi.Application{}, applicationapi.ValidateApplicationProxy, applicationapi.ValidateApplicationUpdate)
}

