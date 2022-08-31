package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal that represents an AWS Organization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationPrincipal := awscdk.Aws_iam.NewOrganizationPrincipal(jsii.String("organizationId"))
//
// Experimental.
type OrganizationPrincipal interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() IPrincipal
	// The unique identifier (ID) of an organization (i.e. o-12345abcde).
	// Experimental.
	OrganizationId() *string
	// Return the policy fragment that identifies this principal in a Policy.
	// Experimental.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	// Experimental.
	PrincipalAccount() *string
	// Add the princpial to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	// Experimental.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	// Experimental.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	// Experimental.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	// Experimental.
	DedupeString() *string
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	// Experimental.
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	// Experimental.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	// Experimental.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	// Experimental.
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


// Experimental.
func NewOrganizationPrincipal(organizationId *string) OrganizationPrincipal {
	_init_.Initialize()

	j := jsiiProxy_OrganizationPrincipal{}

	_jsii_.Create(
		"monocdk.aws_iam.OrganizationPrincipal",
		[]interface{}{organizationId},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationPrincipal_Override(o OrganizationPrincipal, organizationId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.OrganizationPrincipal",
		[]interface{}{organizationId},
		o,
	)
}

func (o *jsiiProxy_OrganizationPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	_jsii_.InvokeVoid(
		o,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (o *jsiiProxy_OrganizationPrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

