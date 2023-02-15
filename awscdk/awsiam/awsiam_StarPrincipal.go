package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal that uses a literal '*' in the IAM JSON language.
//
// Some services behave differently when you specify `Principal: "*"`
// or `Principal: { AWS: "*" }` in their resource policy.
//
// `StarPrincipal` renders to `Principal: *`. Most of the time, you
// should use `AnyPrincipal` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   starPrincipal := awscdk.Aws_iam.NewStarPrincipal()
//
type StarPrincipal interface {
	PrincipalBase
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

// The jsii proxy struct for StarPrincipal
type jsiiProxy_StarPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_StarPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StarPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StarPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StarPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewStarPrincipal() StarPrincipal {
	_init_.Initialize()

	j := jsiiProxy_StarPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.StarPrincipal",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStarPrincipal_Override(s StarPrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.StarPrincipal",
		nil, // no parameters
		s,
	)
}

func (s *jsiiProxy_StarPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := s.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (s *jsiiProxy_StarPrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

func (s *jsiiProxy_StarPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := s.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StarPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StarPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StarPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StarPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := s.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StarPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

