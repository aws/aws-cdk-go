package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Represents a trigger schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   triggerSchedule := glue_alpha.TriggerSchedule_Cron(&CronOptions{
//   	Day: jsii.String("day"),
//   	Hour: jsii.String("hour"),
//   	Minute: jsii.String("minute"),
//   	Month: jsii.String("month"),
//   	WeekDay: jsii.String("weekDay"),
//   	Year: jsii.String("year"),
//   })
//
// Experimental.
type TriggerSchedule interface {
	// The expression string for the schedule.
	// Experimental.
	ExpressionString() *string
}

// The jsii proxy struct for TriggerSchedule
type jsiiProxy_TriggerSchedule struct {
	_ byte // padding
}

func (j *jsiiProxy_TriggerSchedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


// Creates a new TriggerSchedule instance with a cron expression.
//
// Returns: A new TriggerSchedule instance.
// Experimental.
func TriggerSchedule_Cron(options *awsevents.CronOptions) TriggerSchedule {
	_init_.Initialize()

	if err := validateTriggerSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns TriggerSchedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.TriggerSchedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a new TriggerSchedule instance with a custom expression.
//
// Returns: A new TriggerSchedule instance.
// Experimental.
func TriggerSchedule_Expression(expression *string) TriggerSchedule {
	_init_.Initialize()

	if err := validateTriggerSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns TriggerSchedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.TriggerSchedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

