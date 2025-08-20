package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Enables session tags on role assumptions from a principal.
//
// For more information on session tags, see:
// https://docs.aws.amazon.com/IAM/latest/UserGuide/id_session-tags.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var principal iPrincipal
//
//   sessionTagsPrincipal := awscdk.Aws_iam.NewSessionTagsPrincipal(principal)
//
type SessionTagsPrincipal interface {
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
	AddToAssumeRolePolicy(doc PolicyDocument)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult
	// Append the given string to the wrapped principal's dedupe string (if available).
	AppendDedupe(append *string) *string
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

// The jsii proxy struct for SessionTagsPrincipal
type jsiiProxy_SessionTagsPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_SessionTagsPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SessionTagsPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SessionTagsPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SessionTagsPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewSessionTagsPrincipal(principal IPrincipal) SessionTagsPrincipal {
	_init_.Initialize()

	if err := validateNewSessionTagsPrincipalParameters(principal); err != nil {
		panic(err)
	}
	j := jsiiProxy_SessionTagsPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.SessionTagsPrincipal",
		[]interface{}{principal},
		&j,
	)

	return &j
}

func NewSessionTagsPrincipal_Override(s SessionTagsPrincipal, principal IPrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.SessionTagsPrincipal",
		[]interface{}{principal},
		s,
	)
}

func (s *jsiiProxy_SessionTagsPrincipal) AddToAssumeRolePolicy(doc PolicyDocument) {
	if err := s.validateAddToAssumeRolePolicyParameters(doc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToAssumeRolePolicy",
		[]interface{}{doc},
	)
}

func (s *jsiiProxy_SessionTagsPrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

func (s *jsiiProxy_SessionTagsPrincipal) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := s.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SessionTagsPrincipal) AppendDedupe(append *string) *string {
	if err := s.validateAppendDedupeParameters(append); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"appendDedupe",
		[]interface{}{append},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SessionTagsPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SessionTagsPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SessionTagsPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SessionTagsPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
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

func (s *jsiiProxy_SessionTagsPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

