package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Schedule for canary runs.
//
// Example:
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_4_0(),
//   	EnvironmentVariables: map[string]*string{
//   		"stage": jsii.String("prod"),
//   	},
//   })
//
type Schedule interface {
	// The Schedule expression.
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


// Create a schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
//
// The expression must be in a `rate(number units)` format.
// For example, `Schedule.expression('rate(10 minutes)')`
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// The canary will be executed once.
func Schedule_Once() Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Schedule",
		"once",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construct a schedule from an interval.
//
// Allowed values: 0 (for a single run) or between 1 and 60 minutes.
// To specify a single run, you can use `Schedule.once()`.
func Schedule_Rate(interval awscdk.Duration) Schedule {
	_init_.Initialize()

	if err := validateSchedule_RateParameters(interval); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_synthetics.Schedule",
		"rate",
		[]interface{}{interval},
		&returns,
	)

	return returns
}

