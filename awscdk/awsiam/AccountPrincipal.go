package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specify AWS account ID as the principal entity in a policy to delegate authority to the account.
//
// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("Cluster"), &DatabaseClusterProps{
//   	Vpc: Vpc,
//   	InstanceType: neptune.InstanceType_R5_LARGE(),
//   	IamAuthentication: jsii.Boolean(true),
//   })
//   role := iam.NewRole(this, jsii.String("DBRole"), &RoleProps{
//   	AssumedBy: iam.NewAccountPrincipal(this.Account),
//   })
//   // Use one of the following statements to grant the role the necessary permissions
//   cluster.GrantConnect(role) // Grant the role neptune-db:* access to the DB
//   cluster.Grant(role, jsii.String("neptune-db:ReadDataViaQuery"), jsii.String("neptune-db:WriteDataViaQuery"))
//
type AccountPrincipal interface {
	ArnPrincipal
	// AWS account ID (i.e. '123456789012').
	AccountId() interface{}
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

// The jsii proxy struct for AccountPrincipal
type jsiiProxy_AccountPrincipal struct {
	jsiiProxy_ArnPrincipal
}

func (j *jsiiProxy_AccountPrincipal) AccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrincipal) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewAccountPrincipal(accountId interface{}) AccountPrincipal {
	_init_.Initialize()

	if err := validateNewAccountPrincipalParameters(accountId); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.AccountPrincipal",
		[]interface{}{accountId},
		&j,
	)

	return &j
}

func NewAccountPrincipal_Override(a AccountPrincipal, accountId interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.AccountPrincipal",
		[]interface{}{accountId},
		a,
	)
}

func (a *jsiiProxy_AccountPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := a.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (a *jsiiProxy_AccountPrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

func (a *jsiiProxy_AccountPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (a *jsiiProxy_AccountPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrincipal) InOrganization(organizationId *string) PrincipalBase {
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

func (a *jsiiProxy_AccountPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		a,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
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

func (a *jsiiProxy_AccountPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

