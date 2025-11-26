package previewawsemrmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ApplicationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.AutoTerminationPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AutoTerminationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.BootstrapActionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.CloudWatchAlarmDefinitionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CloudWatchAlarmDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ComputeLimitsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ComputeLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.HadoopJarStepConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_HadoopJarStepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.InstanceFleetConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.InstanceFleetProvisioningSpecificationsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetProvisioningSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.InstanceFleetResizingSpecificationsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetResizingSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.InstanceGroupConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.InstanceTypeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.JobFlowInstancesConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JobFlowInstancesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.KerberosAttributesProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KerberosAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ManagedScalingPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ManagedScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.OnDemandCapacityReservationOptionsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandCapacityReservationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.OnDemandProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.OnDemandResizingSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.PlacementGroupConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PlacementGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.PlacementTypeProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PlacementTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ScalingActionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ScalingConstraintsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ScalingRuleProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ScalingTriggerProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.ScriptBootstrapActionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScriptBootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.SimpleScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SimpleScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.SpotProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SpotProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.SpotResizingSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SpotResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.StepConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnClusterPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigMixinProps",
		reflect.TypeOf((*CfnInstanceFleetConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceFleetConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.InstanceFleetProvisioningSpecificationsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceFleetProvisioningSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.InstanceFleetResizingSpecificationsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceFleetResizingSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.InstanceTypeConfigProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.OnDemandCapacityReservationOptionsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandCapacityReservationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.OnDemandProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.OnDemandResizingSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.SpotProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_SpotProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.SpotResizingSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_SpotResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceFleetConfigPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigMixinProps",
		reflect.TypeOf((*CfnInstanceGroupConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceGroupConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.CloudWatchAlarmDefinitionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_CloudWatchAlarmDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.ScalingActionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.ScalingConstraintsProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.ScalingRuleProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.ScalingTriggerProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.SimpleScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_SimpleScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnInstanceGroupConfigPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStepMixinProps",
		reflect.TypeOf((*CfnStepMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStepPropsMixin",
		reflect.TypeOf((*CfnStepPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStepPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStepPropsMixin.HadoopJarStepConfigProperty",
		reflect.TypeOf((*CfnStepPropsMixin_HadoopJarStepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStepPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnStepPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioMixinProps",
		reflect.TypeOf((*CfnStudioMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioPropsMixin",
		reflect.TypeOf((*CfnStudioPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingMixinProps",
		reflect.TypeOf((*CfnStudioSessionMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnStudioSessionMappingPropsMixin",
		reflect.TypeOf((*CfnStudioSessionMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioSessionMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnWALWorkspaceMixinProps",
		reflect.TypeOf((*CfnWALWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emr.mixins.CfnWALWorkspacePropsMixin",
		reflect.TypeOf((*CfnWALWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWALWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
