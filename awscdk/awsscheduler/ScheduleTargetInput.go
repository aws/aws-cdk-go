package awsscheduler

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The text or well-formed JSON input passed to the target of the schedule.
//
// Tokens and ContextAttribute may be used in the input.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   payload := map[string]*string{
//   	"message": jsii.String("Hello scheduler!"),
//   }
//
//   target := targets.NewSnsPublish(topic, &ScheduleTargetBaseProps{
//   	Input: awscdk.ScheduleTargetInput_FromObject(payload),
//   })
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: Target,
//   })
//
type ScheduleTargetInput interface {
	// Return the input properties for this input object.
	Bind(schedule ISchedule) *string
}

// The jsii proxy struct for ScheduleTargetInput
type jsiiProxy_ScheduleTargetInput struct {
	_ byte // padding
}

func NewScheduleTargetInput_Override(s ScheduleTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler.ScheduleTargetInput",
		nil, // no parameters
		s,
	)
}

// Pass a JSON object to the target.
//
// The object will be transformed into a well-formed JSON string in the final template.
func ScheduleTargetInput_FromObject(obj interface{}) ScheduleTargetInput {
	_init_.Initialize()

	if err := validateScheduleTargetInput_FromObjectParameters(obj); err != nil {
		panic(err)
	}
	var returns ScheduleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleTargetInput",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Pass simple text to the target.
//
// For passing complex values like JSON object to a target use method
// `ScheduleTargetInput.fromObject()` instead.
func ScheduleTargetInput_FromText(text *string) ScheduleTargetInput {
	_init_.Initialize()

	if err := validateScheduleTargetInput_FromTextParameters(text); err != nil {
		panic(err)
	}
	var returns ScheduleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_scheduler.ScheduleTargetInput",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScheduleTargetInput) Bind(schedule ISchedule) *string {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

