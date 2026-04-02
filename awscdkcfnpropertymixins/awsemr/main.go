package awsemr

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterMixinProps",
		reflect.TypeOf((*CfnClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin",
		reflect.TypeOf((*CfnClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ApplicationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ApplicationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.AutoTerminationPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_AutoTerminationPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.BootstrapActionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_BootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.CloudWatchAlarmDefinitionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CloudWatchAlarmDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.CloudWatchLogConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_CloudWatchLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ComputeLimitsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ComputeLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.EMRConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EMRConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.HadoopJarStepConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_HadoopJarStepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.InstanceFleetConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.InstanceFleetProvisioningSpecificationsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetProvisioningSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.InstanceFleetResizingSpecificationsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceFleetResizingSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.InstanceGroupConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.InstanceTypeConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_InstanceTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.JobFlowInstancesConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_JobFlowInstancesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.KerberosAttributesProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KerberosAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ManagedScalingPolicyProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ManagedScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.OnDemandCapacityReservationOptionsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandCapacityReservationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.OnDemandProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.OnDemandResizingSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_OnDemandResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.PlacementGroupConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PlacementGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.PlacementTypeProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_PlacementTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ScalingActionProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ScalingConstraintsProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ScalingRuleProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ScalingTriggerProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScalingTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.ScriptBootstrapActionConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_ScriptBootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.SimpleScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SimpleScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.SpotProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SpotProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.SpotResizingSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_SpotResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.StepConfigProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_StepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnClusterPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnClusterPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigMixinProps",
		reflect.TypeOf((*CfnInstanceFleetConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceFleetConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.InstanceFleetProvisioningSpecificationsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceFleetProvisioningSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.InstanceFleetResizingSpecificationsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceFleetResizingSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.InstanceTypeConfigProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_InstanceTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.OnDemandCapacityReservationOptionsProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandCapacityReservationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.OnDemandProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.OnDemandResizingSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_OnDemandResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.SpotProvisioningSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_SpotProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.SpotResizingSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_SpotResizingSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceFleetConfigPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnInstanceFleetConfigPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigMixinProps",
		reflect.TypeOf((*CfnInstanceGroupConfigMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstanceGroupConfigPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.AutoScalingPolicyProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.CloudWatchAlarmDefinitionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_CloudWatchAlarmDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.ConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.EbsConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.MetricDimensionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.ScalingActionProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.ScalingConstraintsProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.ScalingRuleProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.ScalingTriggerProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_ScalingTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.SimpleScalingPolicyConfigurationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_SimpleScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnInstanceGroupConfigPropsMixin.VolumeSpecificationProperty",
		reflect.TypeOf((*CfnInstanceGroupConfigPropsMixin_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStepMixinProps",
		reflect.TypeOf((*CfnStepMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStepPropsMixin",
		reflect.TypeOf((*CfnStepPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStepPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStepPropsMixin.HadoopJarStepConfigProperty",
		reflect.TypeOf((*CfnStepPropsMixin_HadoopJarStepConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStepPropsMixin.KeyValueProperty",
		reflect.TypeOf((*CfnStepPropsMixin_KeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStudioMixinProps",
		reflect.TypeOf((*CfnStudioMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStudioPropsMixin",
		reflect.TypeOf((*CfnStudioPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStudioSessionMappingMixinProps",
		reflect.TypeOf((*CfnStudioSessionMappingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnStudioSessionMappingPropsMixin",
		reflect.TypeOf((*CfnStudioSessionMappingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioSessionMappingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspaceMixinProps",
		reflect.TypeOf((*CfnWALWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspacePropsMixin",
		reflect.TypeOf((*CfnWALWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWALWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
