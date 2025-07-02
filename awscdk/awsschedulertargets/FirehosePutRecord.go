package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use an Amazon Data Firehose as a target for AWS EventBridge Scheduler.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var deliveryStream iDeliveryStream
//
//
//   payload := map[string]*string{
//   	"Data": jsii.String("record"),
//   }
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
//   		Input: awscdk.ScheduleTargetInput_FromObject(payload),
//   	}),
//   })
//
type FirehosePutRecord interface {
	ScheduleTargetBase
	awsscheduler.IScheduleTarget
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for FirehosePutRecord
type jsiiProxy_FirehosePutRecord struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_FirehosePutRecord) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewFirehosePutRecord(deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) FirehosePutRecord {
	_init_.Initialize()

	if err := validateNewFirehosePutRecordParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehosePutRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.FirehosePutRecord",
		[]interface{}{deliveryStream, props},
		&j,
	)

	return &j
}

func NewFirehosePutRecord_Override(f FirehosePutRecord, deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.FirehosePutRecord",
		[]interface{}{deliveryStream, props},
		f,
	)
}

func (f *jsiiProxy_FirehosePutRecord) AddTargetActionToRole(role awsiam.IRole) {
	if err := f.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (f *jsiiProxy_FirehosePutRecord) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := f.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirehosePutRecord) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := f.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		f,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

