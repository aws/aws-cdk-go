package awsautoscalingcommon

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling_common.Alarms",
		reflect.TypeOf((*Alarms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling_common.ArbitraryIntervals",
		reflect.TypeOf((*ArbitraryIntervals)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling_common.CompleteScalingInterval",
		reflect.TypeOf((*CompleteScalingInterval)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_autoscaling_common.IRandomGenerator",
		reflect.TypeOf((*IRandomGenerator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "nextBoolean", GoMethod: "NextBoolean"},
			_jsii_.MemberMethod{JsiiMethod: "nextInt", GoMethod: "NextInt"},
		},
		func() interface{} {
			return &jsiiProxy_IRandomGenerator{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_autoscaling_common.ScalingInterval",
		reflect.TypeOf((*ScalingInterval)(nil)).Elem(),
	)
}
