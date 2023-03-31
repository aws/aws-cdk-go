package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal that represents a federated identity provider as from a OpenID Connect provider.
//
// Example:
//   provider := iam.NewOpenIdConnectProvider(this, jsii.String("MyProvider"), &openIdConnectProviderProps{
//   	url: jsii.String("https://openid/connect"),
//   	clientIds: []*string{
//   		jsii.String("myclient1"),
//   		jsii.String("myclient2"),
//   	},
//   })
//   principal := iam.NewOpenIdConnectPrincipal(provider)
//
// Experimental.
type OpenIdConnectPrincipal interface {
	WebIdentityPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The conditions under which the policy is in effect.
	//
	// See [the IAM documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html).
	// Experimental.
	Conditions() *map[string]interface{}
	// federated identity provider (i.e. 'cognito-identity.amazonaws.com' for users authenticated through Cognito).
	// Experimental.
	Federated() *string
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

// The jsii proxy struct for OpenIdConnectPrincipal
type jsiiProxy_OpenIdConnectPrincipal struct {
	jsiiProxy_WebIdentityPrincipal
}

func (j *jsiiProxy_OpenIdConnectPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectPrincipal) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectPrincipal) Federated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenIdConnectPrincipal(openIdConnectProvider IOpenIdConnectProvider, conditions *map[string]interface{}) OpenIdConnectPrincipal {
	_init_.Initialize()

	if err := validateNewOpenIdConnectPrincipalParameters(openIdConnectProvider); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenIdConnectPrincipal{}

	_jsii_.Create(
		"monocdk.aws_iam.OpenIdConnectPrincipal",
		[]interface{}{openIdConnectProvider, conditions},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenIdConnectPrincipal_Override(o OpenIdConnectPrincipal, openIdConnectProvider IOpenIdConnectProvider, conditions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.OpenIdConnectPrincipal",
		[]interface{}{openIdConnectProvider, conditions},
		o,
	)
}

func (o *jsiiProxy_OpenIdConnectPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := o.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (o *jsiiProxy_OpenIdConnectPrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

func (o *jsiiProxy_OpenIdConnectPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (o *jsiiProxy_OpenIdConnectPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		o,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
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

func (o *jsiiProxy_OpenIdConnectPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		o,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

