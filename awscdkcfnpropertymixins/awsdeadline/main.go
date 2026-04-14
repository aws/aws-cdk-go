package awsdeadline

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFarmMixinProps",
		reflect.TypeOf((*CfnFarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFarmPropsMixin",
		reflect.TypeOf((*CfnFarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.AcceleratorCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.AcceleratorCountRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorCountRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.AcceleratorSelectionProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.AcceleratorTotalMemoryMiBRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorTotalMemoryMiBRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.CustomerManagedAutoScalingConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CustomerManagedAutoScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.CustomerManagedFleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CustomerManagedFleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.CustomerManagedWorkerCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CustomerManagedWorkerCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.Ec2EbsVolumeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_Ec2EbsVolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.FleetAmountCapabilityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetAmountCapabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.FleetAttributeCapabilityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetAttributeCapabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.FleetCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.FleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.HostConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_HostConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.MemoryMiBRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_MemoryMiBRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.ServiceManagedEc2AutoScalingConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2AutoScalingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.ServiceManagedEc2FleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2FleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.ServiceManagedEc2InstanceCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2InstanceCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.ServiceManagedEc2InstanceMarketOptionsProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2InstanceMarketOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.VCpuCountRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VCpuCountRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnFleetPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnLicenseEndpointMixinProps",
		reflect.TypeOf((*CfnLicenseEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnLicenseEndpointPropsMixin",
		reflect.TypeOf((*CfnLicenseEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnLimitMixinProps",
		reflect.TypeOf((*CfnLimitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnLimitPropsMixin",
		reflect.TypeOf((*CfnLimitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLimitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnMeteredProductMixinProps",
		reflect.TypeOf((*CfnMeteredProductMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnMeteredProductPropsMixin",
		reflect.TypeOf((*CfnMeteredProductPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMeteredProductPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnMonitorMixinProps",
		reflect.TypeOf((*CfnMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnMonitorPropsMixin",
		reflect.TypeOf((*CfnMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueEnvironmentMixinProps",
		reflect.TypeOf((*CfnQueueEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueEnvironmentPropsMixin",
		reflect.TypeOf((*CfnQueueEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueFleetAssociationMixinProps",
		reflect.TypeOf((*CfnQueueFleetAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueFleetAssociationPropsMixin",
		reflect.TypeOf((*CfnQueueFleetAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueFleetAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueLimitAssociationMixinProps",
		reflect.TypeOf((*CfnQueueLimitAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueLimitAssociationPropsMixin",
		reflect.TypeOf((*CfnQueueLimitAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueLimitAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueueMixinProps",
		reflect.TypeOf((*CfnQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueuePropsMixin",
		reflect.TypeOf((*CfnQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueuePropsMixin.JobAttachmentSettingsProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_JobAttachmentSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueuePropsMixin.JobRunAsUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_JobRunAsUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueuePropsMixin.PosixUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_PosixUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnQueuePropsMixin.WindowsUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_WindowsUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnStorageProfileMixinProps",
		reflect.TypeOf((*CfnStorageProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnStorageProfilePropsMixin",
		reflect.TypeOf((*CfnStorageProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStorageProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_deadline.CfnStorageProfilePropsMixin.FileSystemLocationProperty",
		reflect.TypeOf((*CfnStorageProfilePropsMixin_FileSystemLocationProperty)(nil)).Elem(),
	)
}
