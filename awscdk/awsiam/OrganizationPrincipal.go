package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal that represents an AWS Organization.
//
// Example:
//   // Grant permissions to an entire AWS organization
//   var fn function
//
//   org := iam.NewOrganizationPrincipal(jsii.String("o-xxxxxxxxxx"))
//
//   fn.GrantInvoke(org)
//
type OrganizationPrincipal interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// The unique identifier (ID) of an organization (i.e. o-12345abcde).
	OrganizationId() *string
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

// The jsii proxy struct for OrganizationPrincipal
type jsiiProxy_OrganizationPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_OrganizationPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPrincipal) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewOrganizationPrincipal(organizationId *string) OrganizationPrincipal {
	_init_.Initialize()

	if err := validateNewOrganizationPrincipalParameters(organizationId); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.OrganizationPrincipal",
		[]interface{}{organizationId},
		&j,
	)

	return &j
}

func NewOrganizationPrincipal_Override(o OrganizationPrincipal, organizationId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.OrganizationPrincipal",
		[]interface{}{organizationId},
		o,
	)
}

func (o *jsiiProxy_OrganizationPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := o.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (o *jsiiProxy_OrganizationPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := o.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		o,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := o.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		o,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		o,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := o.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		o,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		o,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

