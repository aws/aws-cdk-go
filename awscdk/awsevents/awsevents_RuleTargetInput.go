package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The input to send to the event target.
//
// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import sfn "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   role := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
//   })
//   stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &stateMachineProps{
//   	definition: sfn.NewWait(this, jsii.String("Hello"), &waitProps{
//   		time: sfn.waitTime.duration(cdk.*duration.seconds(jsii.Number(10))),
//   	}),
//   })
//
//   rule.addTarget(targets.NewSfnStateMachine(stateMachine, &sfnStateMachineProps{
//   	input: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	deadLetterQueue: dlq,
//   	role: role,
//   }))
//
type RuleTargetInput interface {
	// Return the input properties for this input object.
	Bind(rule IRule) *RuleTargetInputProperties
}

// The jsii proxy struct for RuleTargetInput
type jsiiProxy_RuleTargetInput struct {
	_ byte // padding
}

func NewRuleTargetInput_Override(r RuleTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		nil, // no parameters
		r,
	)
}

// Take the event target input from a path in the event JSON.
func RuleTargetInput_FromEventPath(path *string) RuleTargetInput {
	_init_.Initialize()

	if err := validateRuleTargetInput_FromEventPathParameters(path); err != nil {
		panic(err)
	}
	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromEventPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Pass text to the event target, splitting on newlines.
//
// This is only useful when passing to a target that does not
// take a single argument.
//
// May contain strings returned by `EventField.from()` to substitute in parts
// of the matched event.
func RuleTargetInput_FromMultilineText(text *string) RuleTargetInput {
	_init_.Initialize()

	if err := validateRuleTargetInput_FromMultilineTextParameters(text); err != nil {
		panic(err)
	}
	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromMultilineText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

// Pass a JSON object to the event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
func RuleTargetInput_FromObject(obj interface{}) RuleTargetInput {
	_init_.Initialize()

	if err := validateRuleTargetInput_FromObjectParameters(obj); err != nil {
		panic(err)
	}
	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Pass text to the event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
//
// The Rule Target input value will be a single string: the string you pass
// here.  Do not use this method to pass a complex value like a JSON object to
// a Rule Target.  Use `RuleTargetInput.fromObject()` instead.
func RuleTargetInput_FromText(text *string) RuleTargetInput {
	_init_.Initialize()

	if err := validateRuleTargetInput_FromTextParameters(text); err != nil {
		panic(err)
	}
	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RuleTargetInput) Bind(rule IRule) *RuleTargetInputProperties {
	if err := r.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *RuleTargetInputProperties

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

