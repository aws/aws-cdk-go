package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// Service Account.
//
// Example:
//   // or create a new one using an existing issuer url
//   var issuerUrl string
//   // you can import an existing provider
//   provider := eks.openIdConnectProvider.fromOpenIdConnectProviderArn(this, jsii.String("Provider"), jsii.String("arn:aws:iam::123456:oidc-provider/oidc.eks.eu-west-1.amazonaws.com/id/AB123456ABC"))
//   provider2 := eks.NewOpenIdConnectProvider(this, jsii.String("Provider"), &openIdConnectProviderProps{
//   	url: issuerUrl,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("MyCluster"), &clusterAttributes{
//   	clusterName: jsii.String("Cluster"),
//   	openIdConnectProvider: provider,
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456:role/service-role/k8sservicerole"),
//   })
//
//   serviceAccount := cluster.addServiceAccount(jsii.String("MyServiceAccount"))
//
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//   bucket.grantReadWrite(serviceAccount)
//
// Experimental.
type ServiceAccount interface {
	awscdk.Construct
	awsiam.IPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ServiceAccount
type jsiiProxy_ServiceAccount struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_ServiceAccount) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_eks.ServiceAccount",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewServiceAccount_Override(s ServiceAccount, scope constructs.Construct, id *string, props *ServiceAccountProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks.ServiceAccount",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ServiceAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks.ServiceAccount",
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

func (s *jsiiProxy_ServiceAccount) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceAccount) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ServiceAccount) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceAccount) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceAccount) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
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

func (s *jsiiProxy_ServiceAccount) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

