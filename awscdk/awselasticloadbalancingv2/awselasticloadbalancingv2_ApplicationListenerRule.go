package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
//   applicationListenerRule := awscdk.Aws_elasticloadbalancingv2.NewApplicationListenerRule(this, jsii.String("MyApplicationListenerRule"), &applicationListenerRuleProps{
//   	listener: applicationListener,
//   	priority: jsii.Number(123),
//
//   	// the properties below are optional
//   	action: listenerAction,
//   	conditions: []*listenerCondition{
//   		listenerCondition,
//   	},
//   	fixedResponse: &fixedResponse{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		contentType: awscdk.*Aws_elasticloadbalancingv2.contentType_TEXT_PLAIN,
//   		messageBody: jsii.String("messageBody"),
//   	},
//   	hostHeader: jsii.String("hostHeader"),
//   	pathPattern: jsii.String("pathPattern"),
//   	pathPatterns: []*string{
//   		jsii.String("pathPatterns"),
//   	},
//   	redirectResponse: &redirectResponse{
//   		statusCode: jsii.String("statusCode"),
//
//   		// the properties below are optional
//   		host: jsii.String("host"),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: jsii.String("protocol"),
//   		query: jsii.String("query"),
//   	},
//   	targetGroups: []iApplicationTargetGroup{
//   		applicationTargetGroup,
//   	},
//   })
//
// Experimental.
type ApplicationListenerRule interface {
	awscdk.Construct
	// The ARN of this rule.
	// Experimental.
	ListenerRuleArn() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Add a non-standard condition to this rule.
	// Experimental.
	AddCondition(condition ListenerCondition)
	// Add a fixed response.
	// Deprecated: Use configureAction instead.
	AddFixedResponse(fixedResponse *FixedResponse)
	// Add a redirect response.
	// Deprecated: Use configureAction instead.
	AddRedirectResponse(redirectResponse *RedirectResponse)
	// Add a TargetGroup to load balance to.
	// Deprecated: Use configureAction instead.
	AddTargetGroup(targetGroup IApplicationTargetGroup)
	// Configure the action to perform for this rule.
	// Experimental.
	ConfigureAction(action ListenerAction)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Add a non-standard condition to this rule.
	//
	// If the condition conflicts with an already set condition, it will be overwritten by the one you specified.
	// Deprecated: use `addCondition` instead.
	SetCondition(field *string, values *[]*string)
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the rule.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ApplicationListenerRule
type jsiiProxy_ApplicationListenerRule struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_ApplicationListenerRule) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplicationListenerRule(scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) ApplicationListenerRule {
	_init_.Initialize()

	if err := validateNewApplicationListenerRuleParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationListenerRule{}

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplicationListenerRule_Override(a ApplicationListenerRule, scope constructs.Construct, id *string, props *ApplicationListenerRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_elasticloadbalancingv2.ApplicationListenerRule",
		[]interface{}{scope, id, props},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func ApplicationListenerRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationListenerRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_elasticloadbalancingv2.ApplicationListenerRule",
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

func (a *jsiiProxy_ApplicationListenerRule) AddFixedResponse(fixedResponse *FixedResponse) {
	if err := a.validateAddFixedResponseParameters(fixedResponse); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addFixedResponse",
		[]interface{}{fixedResponse},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) AddRedirectResponse(redirectResponse *RedirectResponse) {
	if err := a.validateAddRedirectResponseParameters(redirectResponse); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addRedirectResponse",
		[]interface{}{redirectResponse},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) AddTargetGroup(targetGroup IApplicationTargetGroup) {
	if err := a.validateAddTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addTargetGroup",
		[]interface{}{targetGroup},
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

func (a *jsiiProxy_ApplicationListenerRule) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationListenerRule) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationListenerRule) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationListenerRule) SetCondition(field *string, values *[]*string) {
	if err := a.validateSetConditionParameters(field); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"setCondition",
		[]interface{}{field, values},
	)
}

func (a *jsiiProxy_ApplicationListenerRule) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
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

func (a *jsiiProxy_ApplicationListenerRule) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

