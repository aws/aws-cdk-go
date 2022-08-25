package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal representing all identities in all accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anyone := awscdk.Aws_iam.NewAnyone()
//
// Deprecated: use `AnyPrincipal`.
type Anyone interface {
	AnyPrincipal
	// Amazon Resource Name (ARN) of the principal entity (i.e. arn:aws:iam::123456789012:user/user-name).
	// Deprecated: use `AnyPrincipal`.
	Arn() *string
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Deprecated: use `AnyPrincipal`.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Deprecated: use `AnyPrincipal`.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	// Deprecated: use `AnyPrincipal`.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	// Deprecated: use `AnyPrincipal`.
	PrincipalAccount() *string
	// Add the princpial to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	// Deprecated: use `AnyPrincipal`.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	// Deprecated: use `AnyPrincipal`.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	// Deprecated: use `AnyPrincipal`.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	// Deprecated: use `AnyPrincipal`.
	DedupeString() *string
	// A convenience method for adding a condition that the principal is part of the specified AWS Organization.
	// Deprecated: use `AnyPrincipal`.
	InOrganization(organizationId *string) PrincipalBase
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	// Deprecated: use `AnyPrincipal`.
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	// Deprecated: use `AnyPrincipal`.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	// Deprecated: use `AnyPrincipal`.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	// Deprecated: use `AnyPrincipal`.
	WithSessionTags() PrincipalBase
}

// The jsii proxy struct for Anyone
type jsiiProxy_Anyone struct {
	jsiiProxy_AnyPrincipal
}

func (j *jsiiProxy_Anyone) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Anyone) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Anyone) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Anyone) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Anyone) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


// Deprecated: use `AnyPrincipal`.
func NewAnyone() Anyone {
	_init_.Initialize()

	j := jsiiProxy_Anyone{}

	_jsii_.Create(
		"monocdk.aws_iam.Anyone",
		nil, // no parameters
		&j,
	)

	return &j
}

// Deprecated: use `AnyPrincipal`.
func NewAnyone_Override(a Anyone) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.Anyone",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_Anyone) AddToAssumeRolePolicy(document PolicyDocument) {
	_jsii_.InvokeVoid(
		a,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (a *jsiiProxy_Anyone) AddToPolicy(statement PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		a,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) InOrganization(organizationId *string) PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"inOrganization",
		[]interface{}{organizationId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		a,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Anyone) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

