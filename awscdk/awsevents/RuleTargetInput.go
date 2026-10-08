package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// The input to send to the event target.
//
// This class abstracts over the three different ways CloudFormation lets you shape
// a target's input, so you do not have to choose between them yourself:
//
// * `Input` - a literal value, emitted when the input does not reference the event.
// * `InputPath` - a single JSONPath into the event, emitted by `fromEventPath()`.
// * `InputTransformer` - an `InputPathsMap` plus an `InputTemplate`, emitted
//   automatically when the input references parts of the event via `EventField`.
//
// You pick a factory method based on the *shape* you want the target to receive, and
// whether you reference `EventField` determines which of the above is synthesized.
// For example, `fromObject({ foo: 'bar' })` produces a literal `Input`, while
// `fromObject({ id: EventField.fromPath('$.detail.instance-id') })` produces an
// `InputTransformer` with the path lifted into `InputPathsMap`.
//
// Example:
//   var rule Rule
//   var topic Topic
//
//
//   // Without EventField this is a literal `Input`:
//   //   {"source":"my-app"}
//   rule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
//   	Message: events.RuleTargetInput_FromObject(map[string]*string{
//   		"source": jsii.String("my-app"),
//   	}),
//   }))
//
//   // Referencing the event inserts those fields from each matched event.
//   rule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
//   	Message: events.RuleTargetInput_*FromObject(map[string]*string{
//   		"instance": events.EventField_fromPath(jsii.String("$.detail.instance-id")),
//   		"state": events.EventField_fromPath(jsii.String("$.detail.state")),
//   	}),
//   }))
//
type RuleTargetInput interface {
	// Return the input properties for this input object.
	Bind(rule interfacesawsevents.IRuleRef) *RuleTargetInputProperties
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
//
// This emits CloudFormation's `InputPath` directly, so the target receives only
// the selected fragment of the event. Use this when you want to forward part of
// the event unchanged; use `fromObject()` when you need to reshape it.
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
//
// As with `fromText`, each line is JSON-encoded, so every line is wrapped in
// double quotes in the synthesized template. Whether those quotes are visible
// to the recipient depends on the target service.
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
//
// Returns: RuleTargetInput.
//
// Example:
//   var rule Rule
//   var topic Topic
//
//
//   // Without EventField this is a literal `Input`:
//   //   {"source":"my-app"}
//   rule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
//   	Message: events.RuleTargetInput_FromObject(map[string]*string{
//   		"source": jsii.String("my-app"),
//   	}),
//   }))
//
//   // Referencing the event inserts those fields from each matched event.
//   rule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
//   	Message: events.RuleTargetInput_*FromObject(map[string]*string{
//   		"instance": events.EventField_fromPath(jsii.String("$.detail.instance-id")),
//   		"state": events.EventField_fromPath(jsii.String("$.detail.state")),
//   	}),
//   }))
//
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
//
// The target `Input` field must be valid JSON, so the text is JSON-encoded
// when the template is synthesized. As a result a plain string is wrapped in
// double quotes: `fromText('something')` renders as `Input: '"something"'`.
// The quotes are part of the required JSON encoding, not an extra value added
// by CDK, and cannot be removed.
//
// Whether those quotes are visible to the recipient depends on the target
// service. Some targets deliver the JSON-encoded value as-is, so the recipient
// sees the surrounding quotes (for example, an SNS topic delivering to an
// email subscriber shows `"something"`). To send a structured payload, use
// `RuleTargetInput.fromObject()` instead.
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

func (r *jsiiProxy_RuleTargetInput) Bind(rule interfacesawsevents.IRuleRef) *RuleTargetInputProperties {
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

