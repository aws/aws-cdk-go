package awsworkspacesinstances

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnVolumeAssociationMixinProps",
		reflect.TypeOf((*CfnVolumeAssociationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnVolumeAssociationPropsMixin",
		reflect.TypeOf((*CfnVolumeAssociationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVolumeAssociationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnVolumeMixinProps",
		reflect.TypeOf((*CfnVolumeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnVolumePropsMixin",
		reflect.TypeOf((*CfnVolumePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVolumePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnVolumePropsMixin.TagSpecificationProperty",
		reflect.TypeOf((*CfnVolumePropsMixin_TagSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstanceMixinProps",
		reflect.TypeOf((*CfnWorkspaceInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspaceInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.BlockDeviceMappingProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_BlockDeviceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.CapacityReservationSpecificationProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_CapacityReservationSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.CapacityReservationTargetProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_CapacityReservationTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.CpuOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_CpuOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.CreditSpecificationRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_CreditSpecificationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.EC2ManagedInstanceProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_EC2ManagedInstanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.EbsBlockDeviceProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_EbsBlockDeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.EnclaveOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_EnclaveOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.HibernationOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_HibernationOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.IamInstanceProfileSpecificationProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_IamInstanceProfileSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.InstanceMaintenanceOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_InstanceMaintenanceOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.InstanceMarketOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_InstanceMarketOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.InstanceMetadataOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_InstanceMetadataOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.InstanceNetworkInterfaceSpecificationProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_InstanceNetworkInterfaceSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.InstanceNetworkPerformanceOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_InstanceNetworkPerformanceOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.LicenseConfigurationRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_LicenseConfigurationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.ManagedInstanceProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_ManagedInstanceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.PlacementProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_PlacementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.PrivateDnsNameOptionsRequestProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_PrivateDnsNameOptionsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.RunInstancesMonitoringEnabledProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_RunInstancesMonitoringEnabledProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.SpotMarketOptionsProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_SpotMarketOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_workspacesinstances.CfnWorkspaceInstancePropsMixin.TagSpecificationProperty",
		reflect.TypeOf((*CfnWorkspaceInstancePropsMixin_TagSpecificationProperty)(nil)).Elem(),
	)
}
