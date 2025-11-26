package previewawsautoscalingevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents",
		reflect.TypeOf((*AutoScalingGroupEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceLaunchLifecycleActionPattern", GoMethod: "EC2InstanceLaunchLifecycleActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceLaunchSuccessfulPattern", GoMethod: "EC2InstanceLaunchSuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceLaunchUnsuccessfulPattern", GoMethod: "EC2InstanceLaunchUnsuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceTerminateLifecycleActionPattern", GoMethod: "EC2InstanceTerminateLifecycleActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceTerminateSuccessfulPattern", GoMethod: "EC2InstanceTerminateSuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eC2InstanceTerminateUnsuccessfulPattern", GoMethod: "EC2InstanceTerminateUnsuccessfulPattern"},
		},
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchLifecycleAction",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchLifecycleAction.EC2InstanceLaunchLifecycleActionProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchSuccessful",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchSuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceLaunchSuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchSuccessful.Details",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchSuccessful.EC2InstanceLaunchSuccessfulProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchUnsuccessful",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchUnsuccessful.Details",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceLaunchUnsuccessful.EC2InstanceLaunchUnsuccessfulProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateLifecycleAction",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateLifecycleAction.EC2InstanceTerminateLifecycleActionProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateSuccessful",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateSuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateSuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateSuccessful.Details",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateSuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateSuccessful.EC2InstanceTerminateSuccessfulProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateUnsuccessful",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateUnsuccessful.Details",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents.EC2InstanceTerminateUnsuccessful.EC2InstanceTerminateUnsuccessfulProps",
		reflect.TypeOf((*AutoScalingGroupEvents_EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps)(nil)).Elem(),
	)
}
