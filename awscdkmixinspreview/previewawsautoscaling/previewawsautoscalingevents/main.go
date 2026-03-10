package previewawsautoscalingevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.CustomizedMetricSpecification",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_CustomizedMetricSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.CustomizedMetricSpecificationItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_CustomizedMetricSpecificationItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.InstancesDistribution",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_InstancesDistribution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.LaunchTemplate",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_LaunchTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.LaunchTemplate1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_LaunchTemplate1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.LaunchTemplateSpecification",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_LaunchTemplateSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.MixedInstancesPolicy",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_MixedInstancesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.PredefinedMetricSpecification",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_PredefinedMetricSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.RequestParametersItem1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.TargetTrackingConfiguration",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_TargetTrackingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.AutoScalingGroupEvents",
		reflect.TypeOf((*AutoScalingGroupEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceLaunchLifecycleActionPattern", GoMethod: "Ec2InstanceLaunchLifecycleActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceLaunchSuccessfulPattern", GoMethod: "Ec2InstanceLaunchSuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceLaunchUnsuccessfulPattern", GoMethod: "Ec2InstanceLaunchUnsuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceTerminateLifecycleActionPattern", GoMethod: "Ec2InstanceTerminateLifecycleActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceTerminateSuccessfulPattern", GoMethod: "Ec2InstanceTerminateSuccessfulPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ec2InstanceTerminateUnsuccessfulPattern", GoMethod: "Ec2InstanceTerminateUnsuccessfulPattern"},
		},
		func() interface{} {
			return &jsiiProxy_AutoScalingGroupEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchLifecycleAction",
		reflect.TypeOf((*EC2InstanceLaunchLifecycleAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceLaunchLifecycleAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchLifecycleAction.EC2InstanceLaunchLifecycleActionProps",
		reflect.TypeOf((*EC2InstanceLaunchLifecycleAction_EC2InstanceLaunchLifecycleActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful",
		reflect.TypeOf((*EC2InstanceLaunchSuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceLaunchSuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful.Details",
		reflect.TypeOf((*EC2InstanceLaunchSuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchSuccessful.EC2InstanceLaunchSuccessfulProps",
		reflect.TypeOf((*EC2InstanceLaunchSuccessful_EC2InstanceLaunchSuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful",
		reflect.TypeOf((*EC2InstanceLaunchUnsuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceLaunchUnsuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful.Details",
		reflect.TypeOf((*EC2InstanceLaunchUnsuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceLaunchUnsuccessful.EC2InstanceLaunchUnsuccessfulProps",
		reflect.TypeOf((*EC2InstanceLaunchUnsuccessful_EC2InstanceLaunchUnsuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateLifecycleAction",
		reflect.TypeOf((*EC2InstanceTerminateLifecycleAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceTerminateLifecycleAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateLifecycleAction.EC2InstanceTerminateLifecycleActionProps",
		reflect.TypeOf((*EC2InstanceTerminateLifecycleAction_EC2InstanceTerminateLifecycleActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful",
		reflect.TypeOf((*EC2InstanceTerminateSuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceTerminateSuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful.Details",
		reflect.TypeOf((*EC2InstanceTerminateSuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateSuccessful.EC2InstanceTerminateSuccessfulProps",
		reflect.TypeOf((*EC2InstanceTerminateSuccessful_EC2InstanceTerminateSuccessfulProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful",
		reflect.TypeOf((*EC2InstanceTerminateUnsuccessful)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EC2InstanceTerminateUnsuccessful{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful.Details",
		reflect.TypeOf((*EC2InstanceTerminateUnsuccessful_Details)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_autoscaling.events.EC2InstanceTerminateUnsuccessful.EC2InstanceTerminateUnsuccessfulProps",
		reflect.TypeOf((*EC2InstanceTerminateUnsuccessful_EC2InstanceTerminateUnsuccessfulProps)(nil)).Elem(),
	)
}
