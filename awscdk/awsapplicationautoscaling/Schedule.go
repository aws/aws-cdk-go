package awsapplicationautoscaling

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
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
	// Retrieve the expression for this schedule.
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


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a Schedule from a moment in time.
func Schedule_At(moment *time.Time) Schedule {
	_init_.Initialize()

	if err := validateSchedule_AtParameters(moment); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"at",
		[]interface{}{moment},
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
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	if err := validateSchedule_ExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationautoscaling.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
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

