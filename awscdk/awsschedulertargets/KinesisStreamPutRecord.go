package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Use an Amazon Kinesis Data Streams as a target for AWS EventBridge Scheduler.
//
// Example:
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   stream := kinesis.NewStream(this, jsii.String("MyStream"))
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewKinesisStreamPutRecord(stream, &KinesisStreamPutRecordProps{
//   		PartitionKey: jsii.String("key"),
//   	}),
//   })
//
type KinesisStreamPutRecord interface {
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

// The jsii proxy struct for KinesisStreamPutRecord
type jsiiProxy_KinesisStreamPutRecord struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_KinesisStreamPutRecord) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewKinesisStreamPutRecord(stream awskinesis.IStream, props *KinesisStreamPutRecordProps) KinesisStreamPutRecord {
	_init_.Initialize()

	if err := validateNewKinesisStreamPutRecordParameters(stream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisStreamPutRecord{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.KinesisStreamPutRecord",
		[]interface{}{stream, props},
		&j,
	)

	return &j
}

func NewKinesisStreamPutRecord_Override(k KinesisStreamPutRecord, stream awskinesis.IStream, props *KinesisStreamPutRecordProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.KinesisStreamPutRecord",
		[]interface{}{stream, props},
		k,
	)
}

func (k *jsiiProxy_KinesisStreamPutRecord) AddTargetActionToRole(role awsiam.IRole) {
	if err := k.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (k *jsiiProxy_KinesisStreamPutRecord) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := k.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		k,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisStreamPutRecord) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := k.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		k,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

