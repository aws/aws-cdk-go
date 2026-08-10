package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an instance of a CloudWatch alarm mute rule.
type IAlarmMuteRule interface {
	interfacesawscloudwatch.IAlarmMuteRuleRef
	awscdk.IResource
	// The ARN of the alarm mute rule.
	AlarmMuteRuleArn() *string
	// The name of the alarm mute rule.
	AlarmMuteRuleName() *string
}

// The jsii proxy for IAlarmMuteRule
type jsiiProxy_IAlarmMuteRule struct {
	internal.Type__interfacesawscloudwatchIAlarmMuteRuleRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAlarmMuteRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAlarmMuteRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) AlarmMuteRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmMuteRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) AlarmMuteRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmMuteRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) AlarmMuteRuleRef() *interfacesawscloudwatch.AlarmMuteRuleReference {
	var returns *interfacesawscloudwatch.AlarmMuteRuleReference
	_jsii_.Get(
		j,
		"alarmMuteRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlarmMuteRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

