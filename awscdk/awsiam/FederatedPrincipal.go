package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Principal entity that represents a federated identity provider such as Amazon Cognito, that can be used to provide temporary security credentials to users who have been authenticated.
//
// Additional condition keys are available when the temporary security credentials are used to make a request.
// You can use these keys to write policies that limit the access of federated users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//
//   federatedPrincipal := awscdk.Aws_iam.NewFederatedPrincipal(jsii.String("federated"), map[string]interface{}{
//   	"conditionsKey": conditions,
//   }, jsii.String("assumeRoleAction"))
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_iam-condition-keys.html#condition-keys-wif
//
type FederatedPrincipal interface {
	PrincipalBase
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

// The jsii proxy struct for FederatedPrincipal
type jsiiProxy_FederatedPrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_FederatedPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedPrincipal) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedPrincipal) Federated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FederatedPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewFederatedPrincipal(federated *string, conditions *map[string]interface{}, assumeRoleAction *string) FederatedPrincipal {
	_init_.Initialize()

	if err := validateNewFederatedPrincipalParameters(federated); err != nil {
		panic(err)
	}
	j := jsiiProxy_FederatedPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.FederatedPrincipal",
		[]interface{}{federated, conditions, assumeRoleAction},
		&j,
	)

	return &j
}

func NewFederatedPrincipal_Override(f FederatedPrincipal, federated *string, conditions *map[string]interface{}, assumeRoleAction *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.FederatedPrincipal",
		[]interface{}{federated, conditions, assumeRoleAction},
		f,
	)
}

func (f *jsiiProxy_FederatedPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := f.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (f *jsiiProxy_FederatedPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := f.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		f,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := f.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		f,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		f,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := f.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		f,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FederatedPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		f,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

