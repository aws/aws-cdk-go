package awsautoscalinghooktargets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling_hooktargets.FunctionHook",
		reflect.TypeOf((*FunctionHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FunctionHook{}
			_jsii_.InitJsiiProxy(&j.Type__awsautoscalingILifecycleHookTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling_hooktargets.QueueHook",
		reflect.TypeOf((*QueueHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_QueueHook{}
			_jsii_.InitJsiiProxy(&j.Type__awsautoscalingILifecycleHookTarget)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_autoscaling_hooktargets.TopicHook",
		reflect.TypeOf((*TopicHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_TopicHook{}
			_jsii_.InitJsiiProxy(&j.Type__awsautoscalingILifecycleHookTarget)
			return &j
		},
	)
}
