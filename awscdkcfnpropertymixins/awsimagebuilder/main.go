package awsimagebuilder

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnComponentMixinProps",
		reflect.TypeOf((*CfnComponentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnComponentPropsMixin",
		reflect.TypeOf((*CfnComponentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnComponentPropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnComponentPropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipeMixinProps",
		reflect.TypeOf((*CfnContainerRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin",
		reflect.TypeOf((*CfnContainerRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.ComponentConfigurationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_ComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.ComponentParameterProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_ComponentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.EbsInstanceBlockDeviceSpecificationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_EbsInstanceBlockDeviceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.InstanceBlockDeviceMappingProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_InstanceBlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnContainerRecipePropsMixin.TargetContainerRepositoryProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_TargetContainerRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationMixinProps",
		reflect.TypeOf((*CfnDistributionConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.DistributionProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_DistributionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.FastLaunchConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.FastLaunchLaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchLaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.FastLaunchSnapshotConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchSnapshotConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.LaunchTemplateConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_LaunchTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnDistributionConfigurationPropsMixin.SsmParameterConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_SsmParameterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageMixinProps",
		reflect.TypeOf((*CfnImageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelineMixinProps",
		reflect.TypeOf((*CfnImagePipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin",
		reflect.TypeOf((*CfnImagePipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.AutoDisablePolicyProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_AutoDisablePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.EcrConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_EcrConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.ImageScanningConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ImageScanningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.ImageTestsConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ImageTestsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.PipelineLoggingConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_PipelineLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.WorkflowConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_WorkflowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePipelinePropsMixin.WorkflowParameterProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_WorkflowParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin",
		reflect.TypeOf((*CfnImagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.DeletionSettingsProperty",
		reflect.TypeOf((*CfnImagePropsMixin_DeletionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.EcrConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_EcrConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.ImageLoggingConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.ImagePipelineExecutionSettingsProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImagePipelineExecutionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.ImageScanningConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageScanningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.ImageTestsConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageTestsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnImagePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.WorkflowConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_WorkflowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImagePropsMixin.WorkflowParameterProperty",
		reflect.TypeOf((*CfnImagePropsMixin_WorkflowParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipeMixinProps",
		reflect.TypeOf((*CfnImageRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin",
		reflect.TypeOf((*CfnImageRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImageRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.AdditionalInstanceConfigurationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_AdditionalInstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.ComponentConfigurationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_ComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.ComponentParameterProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_ComponentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.EbsInstanceBlockDeviceSpecificationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_EbsInstanceBlockDeviceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.InstanceBlockDeviceMappingProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_InstanceBlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnImageRecipePropsMixin.SystemsManagerAgentProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_SystemsManagerAgentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationMixinProps",
		reflect.TypeOf((*CfnInfrastructureConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationPropsMixin",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInfrastructureConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationPropsMixin.InstanceMetadataOptionsProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_InstanceMetadataOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationPropsMixin.PlacementProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_PlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnInfrastructureConfigurationPropsMixin.S3LogsProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_S3LogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyMixinProps",
		reflect.TypeOf((*CfnLifecyclePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecyclePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.AmiExclusionRulesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_AmiExclusionRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.ExclusionRulesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ExclusionRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.IncludeResourcesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_IncludeResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.LastLaunchedProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_LastLaunchedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.PolicyDetailProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_PolicyDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.RecipeSelectionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_RecipeSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnLifecyclePolicyPropsMixin.ResourceSelectionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ResourceSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowMixinProps",
		reflect.TypeOf((*CfnWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowPropsMixin",
		reflect.TypeOf((*CfnWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowPropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
}
