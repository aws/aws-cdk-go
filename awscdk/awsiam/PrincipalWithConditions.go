package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An IAM principal with additional conditions specifying when the policy is in effect.
//
// For more information about conditions, see:
// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var conditions interface{}
//   var principal iPrincipal
//
//   principalWithConditions := awscdk.Aws_iam.NewPrincipalWithConditions(principal, map[string]interface{}{
//   	"conditionsKey": conditions,
//   })
//
type PrincipalWithConditions interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The conditions under which the policy is in effect.
	//
	// See [the IAM documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html).
	Conditions() *map[string]interface{}
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
	// Add a condition to the principal.
	AddCondition(key *string, value interface{})
	// Adds multiple conditions to the principal.
	//
	// Values from the conditions parameter will overwrite existing values with the same operator
	// and key.
	AddConditions(conditions *map[string]interface{})
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

// The jsii proxy struct for PrincipalWithConditions
type jsiiProxy_PrincipalWithConditions struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_PrincipalWithConditions) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalWithConditions) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalWithConditions) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalWithConditions) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalWithConditions) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewPrincipalWithConditions(principal IPrincipal, conditions *map[string]interface{}) PrincipalWithConditions {
	_init_.Initialize()

	if err := validateNewPrincipalWithConditionsParameters(principal, conditions); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrincipalWithConditions{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PrincipalWithConditions",
		[]interface{}{principal, conditions},
		&j,
	)

	return &j
}

func NewPrincipalWithConditions_Override(p PrincipalWithConditions, principal IPrincipal, conditions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PrincipalWithConditions",
		[]interface{}{principal, conditions},
		p,
	)
}

func (p *jsiiProxy_PrincipalWithConditions) AddCondition(key *string, value interface{}) {
	if err := p.validateAddConditionParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCondition",
		[]interface{}{key, value},
	)
}

func (p *jsiiProxy_PrincipalWithConditions) AddConditions(conditions *map[string]interface{}) {
	if err := p.validateAddConditionsParameters(conditions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addConditions",
		[]interface{}{conditions},
	)
}

func (p *jsiiProxy_PrincipalWithConditions) AddToAssumeRolePolicy(doc PolicyDocument) {
	if err := p.validateAddToAssumeRolePolicyParameters(doc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addToAssumeRolePolicy",
		[]interface{}{doc},
	)
}

func (p *jsiiProxy_PrincipalWithConditions) AddToPolicy(statement PolicyStatement) *bool {
	if err := p.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		p,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := p.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		p,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) AppendDedupe(append *string) *string {
	if err := p.validateAppendDedupeParameters(append); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"appendDedupe",
		[]interface{}{append},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := p.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		p,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalWithConditions) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		p,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

