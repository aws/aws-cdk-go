package previewawsschedulermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnScheduleGroupMixinProps",
		reflect.TypeOf((*CfnScheduleGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnScheduleGroupPropsMixin",
		reflect.TypeOf((*CfnScheduleGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScheduleGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnScheduleMixinProps",
		reflect.TypeOf((*CfnScheduleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin",
		reflect.TypeOf((*CfnSchedulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.DeadLetterConfigProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_DeadLetterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.EcsParametersProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_EcsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.EventBridgeParametersProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_EventBridgeParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.FlexibleTimeWindowProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_FlexibleTimeWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.KinesisParametersProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_KinesisParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.PlacementConstraintProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_PlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.PlacementStrategyProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_PlacementStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.RetryPolicyProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_RetryPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.SageMakerPipelineParameterProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_SageMakerPipelineParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.SageMakerPipelineParametersProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_SageMakerPipelineParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.SqsParametersProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_SqsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_scheduler.mixins.CfnSchedulePropsMixin.TargetProperty",
		reflect.TypeOf((*CfnSchedulePropsMixin_TargetProperty)(nil)).Elem(),
	)
}
