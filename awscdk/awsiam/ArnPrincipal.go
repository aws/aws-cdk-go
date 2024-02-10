package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specify a principal by the Amazon Resource Name (ARN).
//
// You can specify AWS accounts, IAM users, Federated SAML users, IAM roles, and specific assumed-role sessions.
// You cannot specify IAM groups or instance profiles as principals.
//
// Example:
//   // Option 2: create your custom mastersRole with scoped assumeBy arn as the Cluster prop. Switch to this role from the AWS console.
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv29"
//   var vpc vpc
//
//
//   mastersRole := iam.NewRole(this, jsii.String("MastersRole"), &RoleProps{
//   	AssumedBy: iam.NewArnPrincipal(jsii.String("arn_for_trusted_principal")),
//   })
//
//   cluster := eks.NewCluster(this, jsii.String("EksCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   	Version: eks.KubernetesVersion_V1_29(),
//   	KubectlLayer: kubectlv29.NewKubectlV29Layer(this, jsii.String("KubectlLayer")),
//   	MastersRole: MastersRole,
//   })
//
//   mastersRole.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("eks:AccessKubernetesApi"),
//   		jsii.String("eks:Describe*"),
//   		jsii.String("eks:List*"),
//   	},
//   	Resources: []*string{
//   		cluster.ClusterArn,
//   	},
//   }))
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
//
type ArnPrincipal interface {
	PrincipalBase
	// Amazon Resource Name (ARN) of the principal entity (i.e. arn:aws:iam::123456789012:user/user-name).
	Arn() *string
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
	// Add the principal to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	DedupeString() *string
	// A convenience method for adding a condition that the principal is part of the specified AWS Organization.
	InOrganization(organizationId *string) PrincipalBase
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	WithSessionTags() PrincipalBase
}

// The jsii proxy struct for ArnPrincipal
type jsiiProxy_ArnPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_ArnPrincipal) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArnPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArnPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArnPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArnPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewArnPrincipal(arn *string) ArnPrincipal {
	_init_.Initialize()

	if err := validateNewArnPrincipalParameters(arn); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArnPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ArnPrincipal",
		[]interface{}{arn},
		&j,
	)

	return &j
}

func NewArnPrincipal_Override(a ArnPrincipal, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ArnPrincipal",
		[]interface{}{arn},
		a,
	)
}

func (a *jsiiProxy_ArnPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := a.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (a *jsiiProxy_ArnPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := a.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := a.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		a,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) InOrganization(organizationId *string) PrincipalBase {
	if err := a.validateInOrganizationParameters(organizationId); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"inOrganization",
		[]interface{}{organizationId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		a,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := a.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArnPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

