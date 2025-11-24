package mixinsawsodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterMixinProps",
		reflect.TypeOf((*CfnCloudAutonomousVmClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin",
		reflect.TypeOf((*CfnCloudAutonomousVmClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudAutonomousVmClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudAutonomousVmClusterPropsMixin.MaintenanceWindowProperty",
		reflect.TypeOf((*CfnCloudAutonomousVmClusterPropsMixin_MaintenanceWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructureMixinProps",
		reflect.TypeOf((*CfnCloudExadataInfrastructureMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin",
		reflect.TypeOf((*CfnCloudExadataInfrastructurePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudExadataInfrastructurePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin.CustomerContactProperty",
		reflect.TypeOf((*CfnCloudExadataInfrastructurePropsMixin_CustomerContactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudExadataInfrastructurePropsMixin.MaintenanceWindowProperty",
		reflect.TypeOf((*CfnCloudExadataInfrastructurePropsMixin_MaintenanceWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterMixinProps",
		reflect.TypeOf((*CfnCloudVmClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin",
		reflect.TypeOf((*CfnCloudVmClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudVmClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin.DataCollectionOptionsProperty",
		reflect.TypeOf((*CfnCloudVmClusterPropsMixin_DataCollectionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnCloudVmClusterPropsMixin.DbNodeProperty",
		reflect.TypeOf((*CfnCloudVmClusterPropsMixin_DbNodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkMixinProps",
		reflect.TypeOf((*CfnOdbNetworkMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOdbNetworkPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin.ManagedS3BackupAccessProperty",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin_ManagedS3BackupAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin.ManagedServicesProperty",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin_ManagedServicesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin.S3AccessProperty",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin_S3AccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin.ServiceNetworkEndpointProperty",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin_ServiceNetworkEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbNetworkPropsMixin.ZeroEtlAccessProperty",
		reflect.TypeOf((*CfnOdbNetworkPropsMixin_ZeroEtlAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbPeeringConnectionMixinProps",
		reflect.TypeOf((*CfnOdbPeeringConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_odb.mixins.CfnOdbPeeringConnectionPropsMixin",
		reflect.TypeOf((*CfnOdbPeeringConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOdbPeeringConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
