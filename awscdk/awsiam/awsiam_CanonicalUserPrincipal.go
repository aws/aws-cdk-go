package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A policy principal for canonicalUserIds - useful for S3 bucket policies that use Origin Access identities.
//
// See https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html
//
// and
//
// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html
//
// for more details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canonicalUserPrincipal := awscdk.Aws_iam.NewCanonicalUserPrincipal(jsii.String("canonicalUserId"))
//
// Experimental.
type CanonicalUserPrincipal interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// unique identifier assigned by AWS for every account.
	//
	// root user and IAM users for an account all see the same ID.
	// (i.e. 79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be)
	// Experimental.
	CanonicalUserId() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() IPrincipal
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

// The jsii proxy struct for CanonicalUserPrincipal
type jsiiProxy_CanonicalUserPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_CanonicalUserPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanonicalUserPrincipal) CanonicalUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canonicalUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanonicalUserPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanonicalUserPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CanonicalUserPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


// Experimental.
func NewCanonicalUserPrincipal(canonicalUserId *string) CanonicalUserPrincipal {
	_init_.Initialize()

	j := jsiiProxy_CanonicalUserPrincipal{}

	_jsii_.Create(
		"monocdk.aws_iam.CanonicalUserPrincipal",
		[]interface{}{canonicalUserId},
		&j,
	)

	return &j
}

// Experimental.
func NewCanonicalUserPrincipal_Override(c CanonicalUserPrincipal, canonicalUserId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.CanonicalUserPrincipal",
		[]interface{}{canonicalUserId},
		c,
	)
}

func (c *jsiiProxy_CanonicalUserPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	_jsii_.InvokeVoid(
		c,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (c *jsiiProxy_CanonicalUserPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		c,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		c,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CanonicalUserPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		c,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

