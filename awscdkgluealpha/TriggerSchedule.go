package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// Represents a trigger schedule.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   var stack Stack
//   var role IRole
//   var script Code
//
//   job := glue.NewPySparkEtlJob(stack, jsii.String("Job"), &PySparkEtlJobProps{
//   	Role: Role,
//   	Script: Script,
//   })
//   workflow := glue.NewWorkflow(stack, jsii.String("Workflow"))
//
//   workflow.AddScheduledTrigger(jsii.String("WeeklyTrigger"), &ScheduledTriggerOptions{
//   	Actions: []Action{
//   		glue.Action_Job(job),
//   	},
//   	Schedule: glue.TriggerSchedule_Weekly(),
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

// Creates a schedule that fires once a day, at midnight UTC.
//
// Returns: A new TriggerSchedule instance.
// Experimental.
func TriggerSchedule_Daily() TriggerSchedule {
	_init_.Initialize()

	var returns TriggerSchedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.TriggerSchedule",
		"daily",
		nil, // no parameters
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

// Creates a schedule that fires once a week, at midnight UTC on Sunday.
//
// Returns: A new TriggerSchedule instance.
// Experimental.
func TriggerSchedule_Weekly() TriggerSchedule {
	_init_.Initialize()

	var returns TriggerSchedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.TriggerSchedule",
		"weekly",
		nil, // no parameters
		&returns,
	)

	return returns
}

