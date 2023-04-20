package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Base class for policy principals.
//
// Example:
//   principal := iam.NewAccountPrincipal(jsii.String("123456789000")).WithConditions(map[string]interface{}{
//   	"StringEquals": map[string]*string{
//   		"foo": jsii.String("baz"),
//   	},
//   })
//
// Experimental.
type PrincipalBase interface {
	IAssumeRolePrincipal
	IComparablePrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
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

// The jsii proxy struct for PrincipalBase
type jsiiProxy_PrincipalBase struct {
	jsiiProxy_IAssumeRolePrincipal
	jsiiProxy_IComparablePrincipal
}

func (j *jsiiProxy_PrincipalBase) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalBase) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalBase) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrincipalBase) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrincipalBase_Override(p PrincipalBase) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.PrincipalBase",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_PrincipalBase) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := p.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (p *jsiiProxy_PrincipalBase) AddToPolicy(statement PolicyStatement) *bool {
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

func (p *jsiiProxy_PrincipalBase) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := p.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		p,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalBase) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalBase) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrincipalBase) WithConditions(conditions *map[string]interface{}) PrincipalBase {
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

func (p *jsiiProxy_PrincipalBase) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		p,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

