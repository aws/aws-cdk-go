package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
)

// Use an Amazon Kinesis Data Firehose as a target for AWS EventBridge Scheduler.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   var deliveryStream iDeliveryStream
//
//
//   payload := map[string]*string{
//   	"Data": jsii.String("record"),
//   }
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewKinesisDataFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
//   	}),
//   })
//
// Experimental.
type KinesisDataFirehosePutRecord interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for KinesisDataFirehosePutRecord
type jsiiProxy_KinesisDataFirehosePutRecord struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
}

func (j *jsiiProxy_KinesisDataFirehosePutRecord) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataFirehosePutRecord(deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) KinesisDataFirehosePutRecord {
	_init_.Initialize()

	if err := validateNewKinesisDataFirehosePutRecordParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataFirehosePutRecord{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.KinesisDataFirehosePutRecord",
		[]interface{}{deliveryStream, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataFirehosePutRecord_Override(k KinesisDataFirehosePutRecord, deliveryStream awskinesisfirehose.IDeliveryStream, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.KinesisDataFirehosePutRecord",
		[]interface{}{deliveryStream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDataFirehosePutRecord) AddTargetActionToRole(role awsiam.IRole) {
	if err := k.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (k *jsiiProxy_KinesisDataFirehosePutRecord) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := k.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataFirehosePutRecord) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := k.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		k,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

