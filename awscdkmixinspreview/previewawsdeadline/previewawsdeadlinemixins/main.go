package previewawsdeadlinemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFarmMixinProps",
		reflect.TypeOf((*CfnFarmMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFarmPropsMixin",
		reflect.TypeOf((*CfnFarmPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFarmPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetMixinProps",
		reflect.TypeOf((*CfnFleetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin",
		reflect.TypeOf((*CfnFleetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFleetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.AcceleratorCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.AcceleratorCountRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorCountRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.AcceleratorSelectionProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.AcceleratorTotalMemoryMiBRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_AcceleratorTotalMemoryMiBRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.CustomerManagedFleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CustomerManagedFleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.CustomerManagedWorkerCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_CustomerManagedWorkerCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.Ec2EbsVolumeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_Ec2EbsVolumeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.FleetAmountCapabilityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetAmountCapabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.FleetAttributeCapabilityProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetAttributeCapabilityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.FleetCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.FleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_FleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.HostConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_HostConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.MemoryMiBRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_MemoryMiBRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.ServiceManagedEc2FleetConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2FleetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.ServiceManagedEc2InstanceCapabilitiesProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2InstanceCapabilitiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.ServiceManagedEc2InstanceMarketOptionsProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_ServiceManagedEc2InstanceMarketOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.VCpuCountRangeProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VCpuCountRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnFleetPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnFleetPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnLicenseEndpointMixinProps",
		reflect.TypeOf((*CfnLicenseEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnLicenseEndpointPropsMixin",
		reflect.TypeOf((*CfnLicenseEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnLimitMixinProps",
		reflect.TypeOf((*CfnLimitMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnLimitPropsMixin",
		reflect.TypeOf((*CfnLimitPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLimitPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnMeteredProductMixinProps",
		reflect.TypeOf((*CfnMeteredProductMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnMeteredProductPropsMixin",
		reflect.TypeOf((*CfnMeteredProductPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMeteredProductPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnMonitorMixinProps",
		reflect.TypeOf((*CfnMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnMonitorPropsMixin",
		reflect.TypeOf((*CfnMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueEnvironmentMixinProps",
		reflect.TypeOf((*CfnQueueEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueEnvironmentPropsMixin",
		reflect.TypeOf((*CfnQueueEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueFleetAssociationMixinProps",
		reflect.TypeOf((*CfnQueueFleetAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueFleetAssociationPropsMixin",
		reflect.TypeOf((*CfnQueueFleetAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueFleetAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueLimitAssociationMixinProps",
		reflect.TypeOf((*CfnQueueLimitAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueLimitAssociationPropsMixin",
		reflect.TypeOf((*CfnQueueLimitAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueueLimitAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueueMixinProps",
		reflect.TypeOf((*CfnQueueMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin",
		reflect.TypeOf((*CfnQueuePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnQueuePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin.JobAttachmentSettingsProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_JobAttachmentSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin.JobRunAsUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_JobRunAsUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin.PosixUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_PosixUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnQueuePropsMixin.WindowsUserProperty",
		reflect.TypeOf((*CfnQueuePropsMixin_WindowsUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfileMixinProps",
		reflect.TypeOf((*CfnStorageProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin",
		reflect.TypeOf((*CfnStorageProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStorageProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_deadline.mixins.CfnStorageProfilePropsMixin.FileSystemLocationProperty",
		reflect.TypeOf((*CfnStorageProfilePropsMixin_FileSystemLocationProperty)(nil)).Elem(),
	)
}
