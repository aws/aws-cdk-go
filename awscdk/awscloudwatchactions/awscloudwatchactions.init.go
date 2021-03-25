package awscloudwatchactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_cloudwatch_actions.ApplicationScalingAction",
		reflect.TypeOf((*ApplicationScalingAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationScalingAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudwatchIAlarmAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cloudwatch_actions.AutoScalingAction",
		reflect.TypeOf((*AutoScalingAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AutoScalingAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudwatchIAlarmAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cloudwatch_actions.Ec2Action",
		reflect.TypeOf((*Ec2Action)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2Action{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudwatchIAlarmAction)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_cloudwatch_actions.Ec2InstanceAction",
		reflect.TypeOf((*Ec2InstanceAction)(nil)).Elem(),
		map[string]interface{}{
			"STOP": Ec2InstanceAction_STOP,
			"TERMINATE": Ec2InstanceAction_TERMINATE,
			"RECOVER": Ec2InstanceAction_RECOVER,
			"REBOOT": Ec2InstanceAction_REBOOT,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_cloudwatch_actions.SnsAction",
		reflect.TypeOf((*SnsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SnsAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudwatchIAlarmAction)
			return &j
		},
	)
}
