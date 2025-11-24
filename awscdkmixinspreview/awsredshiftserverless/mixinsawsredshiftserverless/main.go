package mixinsawsredshiftserverless

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespaceMixinProps",
		reflect.TypeOf((*CfnNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin",
		reflect.TypeOf((*CfnNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin.NamespaceProperty",
		reflect.TypeOf((*CfnNamespacePropsMixin_NamespaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnNamespacePropsMixin.SnapshotCopyConfigurationProperty",
		reflect.TypeOf((*CfnNamespacePropsMixin_SnapshotCopyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnSnapshotMixinProps",
		reflect.TypeOf((*CfnSnapshotMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnSnapshotPropsMixin",
		reflect.TypeOf((*CfnSnapshotPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSnapshotPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnSnapshotPropsMixin.SnapshotProperty",
		reflect.TypeOf((*CfnSnapshotPropsMixin_SnapshotProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupMixinProps",
		reflect.TypeOf((*CfnWorkgroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin",
		reflect.TypeOf((*CfnWorkgroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkgroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.ConfigParameterProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_ConfigParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.EndpointProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_EndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.NetworkInterfaceProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_NetworkInterfaceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.PerformanceTargetProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_PerformanceTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.VpcEndpointProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_VpcEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_redshiftserverless.mixins.CfnWorkgroupPropsMixin.WorkgroupProperty",
		reflect.TypeOf((*CfnWorkgroupPropsMixin_WorkgroupProperty)(nil)).Elem(),
	)
}
