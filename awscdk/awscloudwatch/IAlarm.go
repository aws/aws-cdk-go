package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a CloudWatch Alarm.
type IAlarm interface {
	IAlarmRule
	awscdk.IResource
	// Alarm ARN (i.e. arn:aws:cloudwatch:<region>:<account-id>:alarm:Foo).
	AlarmArn() *string
	// Name of the alarm.
	AlarmName() *string
}

// The jsii proxy for IAlarm
type jsiiProxy_IAlarm struct {
	jsiiProxy_IAlarmRule
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAlarm) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAlarm) RenderAlarmRule() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"renderAlarmRule",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAlarm) AlarmArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarm) AlarmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarm) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarm) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

