package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Cedar authorization policy statement.
//
// A statement names the principal, action and resource it applies to, and optionally
// conditions that narrow it further. All three parts are required, so a statement is
// complete as soon as it is constructed.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var engine PolicyEngine
//
//
//   // Example 4: raw Cedar, for features the API does not model
//   // Example 4: raw Cedar, for features the API does not model
//   awscdk.NewPolicy(this, jsii.String("CustomPolicy"), &PolicyProps{
//   	PolicyEngine: engine,
//   	Statement: awscdk.PolicyStatement_FromCedar(jsii.String("permit(principal, action, resource) when { context.custom > 10 };")),
//   })
//
type PolicyStatement interface {
	// Generate the Cedar policy statement string.
	//
	// This is called internally by the Policy construct.
	//
	// Returns: Valid Cedar policy statement.
	ToCedar() *string
}

// The jsii proxy struct for PolicyStatement
type jsiiProxy_PolicyStatement struct {
	_ byte // padding
}

func NewPolicyStatement(props *PolicyStatementProps) PolicyStatement {
	_init_.Initialize()

	if err := validateNewPolicyStatementParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyStatement{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyStatement",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPolicyStatement_Override(p PolicyStatement, props *PolicyStatementProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyStatement",
		[]interface{}{props},
		p,
	)
}

// Create a statement from raw Cedar source.
//
// Use this for Cedar features this API does not model, or to migrate an existing
// policy.
//
// The source is used exactly as given. This method does not escape, quote, or
// validate it, so it is treated as trusted input and you own its correctness and
// its safety. Do not build the string by joining values that come from outside
// your application: a value containing a double quote can close a string literal
// early and add policy statements you did not write. Pass such values through
// `PolicyCondition` and the principal, action and resource factories instead,
// which reject that case at synthesis time. Service-side validation does not help,
// because an injected policy is still valid Cedar.
func PolicyStatement_FromCedar(cedarStatement *string) PolicyStatement {
	_init_.Initialize()

	if err := validatePolicyStatement_FromCedarParameters(cedarStatement); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyStatement",
		"fromCedar",
		[]interface{}{cedarStatement},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ToCedar() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toCedar",
		nil, // no parameters
		&returns,
	)

	return returns
}

