package previewawsroute53mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnCidrCollectionMixinProps",
		reflect.TypeOf((*CfnCidrCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnCidrCollectionPropsMixin",
		reflect.TypeOf((*CfnCidrCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCidrCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnCidrCollectionPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnCidrCollectionPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnDNSSECMixinProps",
		reflect.TypeOf((*CfnDNSSECMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnDNSSECPropsMixin",
		reflect.TypeOf((*CfnDNSSECPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDNSSECPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckMixinProps",
		reflect.TypeOf((*CfnHealthCheckMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin",
		reflect.TypeOf((*CfnHealthCheckPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHealthCheckPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin.AlarmIdentifierProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_AlarmIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin.HealthCheckConfigProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_HealthCheckConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHealthCheckPropsMixin.HealthCheckTagProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_HealthCheckTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZoneMixinProps",
		reflect.TypeOf((*CfnHostedZoneMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZonePropsMixin",
		reflect.TypeOf((*CfnHostedZonePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHostedZonePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZonePropsMixin.HostedZoneConfigProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_HostedZoneConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZonePropsMixin.HostedZoneTagProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_HostedZoneTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZonePropsMixin.QueryLoggingConfigProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_QueryLoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnHostedZonePropsMixin.VPCProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_VPCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyMixinProps",
		reflect.TypeOf((*CfnKeySigningKeyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnKeySigningKeyPropsMixin",
		reflect.TypeOf((*CfnKeySigningKeyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeySigningKeyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupMixinProps",
		reflect.TypeOf((*CfnRecordSetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecordSetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.AliasTargetProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_AliasTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.CidrRoutingConfigProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_CidrRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.CoordinatesProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_CoordinatesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.GeoLocationProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_GeoLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.GeoProximityLocationProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_GeoProximityLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetGroupPropsMixin.RecordSetProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_RecordSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetMixinProps",
		reflect.TypeOf((*CfnRecordSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin",
		reflect.TypeOf((*CfnRecordSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecordSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin.AliasTargetProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_AliasTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin.CidrRoutingConfigProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_CidrRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin.CoordinatesProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_CoordinatesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin.GeoLocationProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_GeoLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_route53.mixins.CfnRecordSetPropsMixin.GeoProximityLocationProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_GeoProximityLocationProperty)(nil)).Elem(),
	)
}
