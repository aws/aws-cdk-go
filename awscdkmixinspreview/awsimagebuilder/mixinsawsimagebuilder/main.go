package mixinsawsimagebuilder

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnComponentMixinProps",
		reflect.TypeOf((*CfnComponentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnComponentPropsMixin",
		reflect.TypeOf((*CfnComponentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnComponentPropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnComponentPropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipeMixinProps",
		reflect.TypeOf((*CfnContainerRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin",
		reflect.TypeOf((*CfnContainerRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnContainerRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.ComponentConfigurationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_ComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.ComponentParameterProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_ComponentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.EbsInstanceBlockDeviceSpecificationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_EbsInstanceBlockDeviceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.InstanceBlockDeviceMappingProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_InstanceBlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnContainerRecipePropsMixin.TargetContainerRepositoryProperty",
		reflect.TypeOf((*CfnContainerRecipePropsMixin_TargetContainerRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationMixinProps",
		reflect.TypeOf((*CfnDistributionConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDistributionConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.DistributionProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_DistributionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.FastLaunchConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.FastLaunchLaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchLaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.FastLaunchSnapshotConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_FastLaunchSnapshotConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.LaunchTemplateConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_LaunchTemplateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnDistributionConfigurationPropsMixin.SsmParameterConfigurationProperty",
		reflect.TypeOf((*CfnDistributionConfigurationPropsMixin_SsmParameterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageMixinProps",
		reflect.TypeOf((*CfnImageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelineMixinProps",
		reflect.TypeOf((*CfnImagePipelineMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin",
		reflect.TypeOf((*CfnImagePipelinePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePipelinePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.AutoDisablePolicyProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_AutoDisablePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.EcrConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_EcrConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.ImageScanningConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ImageScanningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.ImageTestsConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ImageTestsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.PipelineLoggingConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_PipelineLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.WorkflowConfigurationProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_WorkflowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePipelinePropsMixin.WorkflowParameterProperty",
		reflect.TypeOf((*CfnImagePipelinePropsMixin_WorkflowParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin",
		reflect.TypeOf((*CfnImagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.DeletionSettingsProperty",
		reflect.TypeOf((*CfnImagePropsMixin_DeletionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.EcrConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_EcrConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.ImageLoggingConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.ImagePipelineExecutionSettingsProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImagePipelineExecutionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.ImageScanningConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageScanningConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.ImageTestsConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_ImageTestsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnImagePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.WorkflowConfigurationProperty",
		reflect.TypeOf((*CfnImagePropsMixin_WorkflowConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImagePropsMixin.WorkflowParameterProperty",
		reflect.TypeOf((*CfnImagePropsMixin_WorkflowParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipeMixinProps",
		reflect.TypeOf((*CfnImageRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin",
		reflect.TypeOf((*CfnImageRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnImageRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.AdditionalInstanceConfigurationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_AdditionalInstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.ComponentConfigurationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_ComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.ComponentParameterProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_ComponentParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.EbsInstanceBlockDeviceSpecificationProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_EbsInstanceBlockDeviceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.InstanceBlockDeviceMappingProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_InstanceBlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnImageRecipePropsMixin.SystemsManagerAgentProperty",
		reflect.TypeOf((*CfnImageRecipePropsMixin_SystemsManagerAgentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationMixinProps",
		reflect.TypeOf((*CfnInfrastructureConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInfrastructureConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin.InstanceMetadataOptionsProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_InstanceMetadataOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin.LoggingProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_LoggingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin.PlacementProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_PlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnInfrastructureConfigurationPropsMixin.S3LogsProperty",
		reflect.TypeOf((*CfnInfrastructureConfigurationPropsMixin_S3LogsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyMixinProps",
		reflect.TypeOf((*CfnLifecyclePolicyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLifecyclePolicyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.AmiExclusionRulesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_AmiExclusionRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.ExclusionRulesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ExclusionRulesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.IncludeResourcesProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_IncludeResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.LastLaunchedProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_LastLaunchedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.PolicyDetailProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_PolicyDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.RecipeSelectionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_RecipeSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnLifecyclePolicyPropsMixin.ResourceSelectionProperty",
		reflect.TypeOf((*CfnLifecyclePolicyPropsMixin_ResourceSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnWorkflowMixinProps",
		reflect.TypeOf((*CfnWorkflowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnWorkflowPropsMixin",
		reflect.TypeOf((*CfnWorkflowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkflowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_imagebuilder.mixins.CfnWorkflowPropsMixin.LatestVersionProperty",
		reflect.TypeOf((*CfnWorkflowPropsMixin_LatestVersionProperty)(nil)).Elem(),
	)
}
