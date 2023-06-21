package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a new listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationListener applicationListener
//   var applicationTargetGroup applicationTargetGroup
//   var listenerAction listenerAction
//   var listenerCondition listenerCondition
//
//   applicationListenerRule := awscdk.Aws_elasticloadbalancingv2.NewApplicationListenerRule(this, jsii.String("MyApplicationListenerRule"), &ApplicationListenerRuleProps{
//   	Listener: applicationListener,
//   	Priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	Action: listenerAction,
//   	Conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	TargetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   })
//
type ApplicationListenerRule interface {
	constructs.Construct
	// The ARN of this rule.
	ListenerRuleArn() *string
	// The tree node.
	Node() constructs.Node
	// Add a non-standard condition to this rule.
	AddCondition(condition ListenerCondition)
	// Configure the action to perform for this rule.
	ConfigureAction(action ListenerAction)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApplicationListenerRule
type jsiiProxy_ApplicationListenerRule struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ApplicationListenerRule) ListenerRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listenerRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationListenerRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewApplicationListenerRule(scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) ApplicationListenerRule {
	_init_.Initialize()

	if err := validateNewApplicationListenerRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationListenerRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationListenerRule_Override(a ApplicationListenerRule, scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		a,
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
func ApplicationListenerRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationListenerRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationListenerRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListenerRule) AddCondition(condition ListenerCondition) {
	if err := a.validateAddConditionParameters(condition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addCondition",
		[]interface{}{condition},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) ConfigureAction(action ListenerAction) {
	if err := a.validateConfigureActionParameters(action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"configureAction",
		[]interface{}{action},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

