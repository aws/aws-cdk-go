package awsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A rule added at the top of the rule set to drop spam/virus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var receiptRuleAction IReceiptRuleAction
//   var receiptRuleRef IReceiptRuleRef
//   var receiptRuleSetRef IReceiptRuleSetRef
//
//   dropSpamReceiptRule := awscdk.Aws_ses.NewDropSpamReceiptRule(this, jsii.String("MyDropSpamReceiptRule"), &DropSpamReceiptRuleProps{
//   	RuleSet: receiptRuleSetRef,
//
//   	// the properties below are optional
//   	Actions: []IReceiptRuleAction{
//   		receiptRuleAction,
//   	},
//   	After: receiptRuleRef,
//   	Enabled: jsii.Boolean(false),
//   	ReceiptRuleName: jsii.String("receiptRuleName"),
//   	Recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	ScanEnabled: jsii.Boolean(false),
//   	TlsPolicy: awscdk.*Aws_ses.TlsPolicy_OPTIONAL,
//   })
//
// See: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-action-lambda-example-functions.html
//
type DropSpamReceiptRule interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	Rule() ReceiptRule
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DropSpamReceiptRule
type jsiiProxy_DropSpamReceiptRule struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DropSpamReceiptRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DropSpamReceiptRule) Rule() ReceiptRule {
	var returns ReceiptRule
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}


func NewDropSpamReceiptRule(scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) DropSpamReceiptRule {
	_init_.Initialize()

	if err := validateNewDropSpamReceiptRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DropSpamReceiptRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDropSpamReceiptRule_Override(d DropSpamReceiptRule, scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DropSpamReceiptRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDropSpamReceiptRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DropSpamReceiptRule_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ses.DropSpamReceiptRule",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DropSpamReceiptRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DropSpamReceiptRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

