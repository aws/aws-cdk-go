package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Service Account.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv34"
//
//   // you can import an existing provider
//   provider := eks.OidcProviderNative_FromOidcProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOidcProviderNative(this, jsii.String("Provider"), &OidcProviderNativeProps{
//   	Url: issuerUrl,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("MyCluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("Cluster"),
//   	OpenIdConnectProvider: provider,
//   	KubectlProviderOptions: &KubectlProviderOptions{
//   		KubectlLayer: kubectlv34.NewKubectlV34Layer(this, jsii.String("kubectl")),
//   	},
//   })
//
//   serviceAccount := cluster.AddServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.GrantReadWrite(serviceAccount)
//
// Experimental.
type ServiceAccount interface {
	constructs.Construct
	awsiam.IPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Return the policy fragment that identifies this principal in a Policy.
	// Experimental.
	PolicyFragment() awsiam.PrincipalPolicyFragment
	// The role which is linked to the service account.
	// Experimental.
	Role() awsiam.IRole
	// The name of the service account.
	// Experimental.
	ServiceAccountName() *string
	// The namespace where the service account is located in.
	// Experimental.
	ServiceAccountNamespace() *string
	// Add to the policy of this principal.
	// Deprecated: use `addToPrincipalPolicy()`.
	AddToPolicy(statement awsiam.PolicyStatement) *bool
	// Add to the policy of this principal.
	// Experimental.
	AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceAccount
type jsiiProxy_ServiceAccount struct {
	internal.Type__constructsConstruct
	internal.Type__awsiamIPrincipal
}

func (j *jsiiProxy_ServiceAccount) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) PolicyFragment() awsiam.PrincipalPolicyFragment {
	var returns awsiam.PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceAccount) ServiceAccountNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNamespace",
		&returns,
	)
	return returns
}


// Experimental.
func NewServiceAccount(scope constructs.Construct, id *string, props *ServiceAccountProps) ServiceAccount {
	_init_.Initialize()

	if err := validateNewServiceAccountParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceAccount{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceAccount_Override(s ServiceAccount, scope constructs.Construct, id *string, props *ServiceAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func ServiceAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.ServiceAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) AddToPolicy(statement awsiam.PolicyStatement) *bool {
	if err := s.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		s,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult {
	if err := s.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

