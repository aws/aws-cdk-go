package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal that represents a federated identity provider as Web Identity such as Cognito, Amazon, Facebook, Google, etc.
//
// Example:
//   principal := iam.NewWebIdentityPrincipal(jsii.String("cognito-identity.amazonaws.com"), map[string]interface{}{
//   	"StringEquals": map[string]*string{
//   		"cognito-identity.amazonaws.com:aud": jsii.String("us-east-2:12345678-abcd-abcd-abcd-123456"),
//   	},
//   	"ForAnyValue:StringLike": map[string]*string{
//   		"cognito-identity.amazonaws.com:amr": jsii.String("unauthenticated"),
//   	},
//   })
//
type WebIdentityPrincipal interface {
	FederatedPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The conditions under which the policy is in effect.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html
	//
	Conditions() *map[string]interface{}
	// federated identity provider (i.e. 'cognito-identity.amazonaws.com' for users authenticated through Cognito).
	Federated() *string
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

// The jsii proxy struct for WebIdentityPrincipal
type jsiiProxy_WebIdentityPrincipal struct {
	jsiiProxy_FederatedPrincipal
}

func (j *jsiiProxy_WebIdentityPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebIdentityPrincipal) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebIdentityPrincipal) Federated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebIdentityPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebIdentityPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebIdentityPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewWebIdentityPrincipal(identityProvider *string, conditions *map[string]interface{}) WebIdentityPrincipal {
	_init_.Initialize()

	if err := validateNewWebIdentityPrincipalParameters(identityProvider); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebIdentityPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.WebIdentityPrincipal",
		[]interface{}{identityProvider, conditions},
		&j,
	)

	return &j
}

func NewWebIdentityPrincipal_Override(w WebIdentityPrincipal, identityProvider *string, conditions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.WebIdentityPrincipal",
		[]interface{}{identityProvider, conditions},
		w,
	)
}

func (w *jsiiProxy_WebIdentityPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := w.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (w *jsiiProxy_WebIdentityPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := w.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		w,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := w.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		w,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		w,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := w.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		w,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebIdentityPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		w,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

