package previewawscloudformationevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress",
		reflect.TypeOf((*CloudFormationHookInvocationProgress)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CloudFormationHookInvocationProgress{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress.CloudFormationHookInvocationProgressProps",
		reflect.TypeOf((*CloudFormationHookInvocationProgress_CloudFormationHookInvocationProgressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress.HookDetail",
		reflect.TypeOf((*CloudFormationHookInvocationProgress_HookDetail)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress.Result",
		reflect.TypeOf((*CloudFormationHookInvocationProgress_Result)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cloudformation.events.CloudFormationHookInvocationProgress.TargetDetail",
		reflect.TypeOf((*CloudFormationHookInvocationProgress_TargetDetail)(nil)).Elem(),
	)
}
