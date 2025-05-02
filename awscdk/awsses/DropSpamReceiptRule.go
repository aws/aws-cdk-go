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
//
//   var receiptRule receiptRule
//   var receiptRuleAction iReceiptRuleAction
//   var receiptRuleSet receiptRuleSet
//
//   dropSpamReceiptRule := awscdk.Aws_ses.NewDropSpamReceiptRule(this, jsii.String("MyDropSpamReceiptRule"), &DropSpamReceiptRuleProps{
//   	RuleSet: receiptRuleSet,
//
//   	// the properties below are optional
//   	Actions: []*iReceiptRuleAction{
//   		receiptRuleAction,
//   	},
//   	After: receiptRule,
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

