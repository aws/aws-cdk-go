package awsapplicationautoscaling

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
)

// Schedule for scheduled scaling actions.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &VpcProps{
//   	MaxAzs: jsii.Number(1),
//   })
//   cluster := ecs.NewCluster(this, jsii.String("EcsCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//   scheduledFargateTask := ecsPatterns.NewScheduledFargateTask(this, jsii.String("ScheduledFargateTask"), &ScheduledFargateTaskProps{
//   	Cluster: Cluster,
//   	ScheduledFargateTaskImageOptions: &ScheduledFargateTaskImageOptions{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		MemoryLimitMiB: jsii.Number(512),
//   	},
//   	Schedule: appscaling.Schedule_Expression(jsii.String("rate(1 minute)")),
//   	Tags: []tag{
//   		&tag{
//   			Key: jsii.String("my-tag"),
//   			Value: jsii.String("my-tag-value"),
//   		},
//   	},
//   })
//
type Schedule interface {
	awscdk.Schedule
	// Retrieve the expression for this schedule.
	ExpressionString() *string
	// The timezone of the expression, if applicable.
	TimeZone() awscdk.TimeZone
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	internal.Type__awscdkSchedule
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

func (j *jsiiProxy_Schedule) TimeZone() awscdk.TimeZone {
	var returns awscdk.TimeZone
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a Schedule from a moment in time.
func Schedule_At(moment *time.Time, timeZone awscdk.TimeZone) Schedule {
	_init_.Initialize()

	if err := validateSchedule_AtParameters(moment); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"at",
		[]interface{}{moment, timeZone},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
func Schedule_Cron(options *awscdk.CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string, timeZone awscdk.TimeZone) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"expression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a one-time schedule from a date.
func Schedule_ProtectedAt(date *time.Time, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedAtParameters(date); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"protectedAt",
		[]interface{}{date, timeZone},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
func Schedule_ProtectedCron(options *awscdk.CronOptions, module *string) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedCronParameters(options); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"protectedCron",
		[]interface{}{options, module},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_ProtectedExpression(expression *string, timeZone awscdk.TimeZone) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"protectedExpression",
		[]interface{}{expression, timeZone},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Rates may be defined with any unit of time, but when converted into minutes, the duration must be a positive whole number of minutes.
func Schedule_ProtectedRate(duration awscdk.Duration) awscdk.Schedule {
	_init_.Initialize()

	if err := validateSchedule_ProtectedRateParameters(duration); err != nil {
		panic(err)
	}
	var returns awscdk.Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"protectedRate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
//
// Must be a whole number of seconds.
func Schedule_Rate(duration awscdk.Duration) Schedule {
	_init_.Initialize()

	if err := validateSchedule_RateParameters(duration); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

