package awsroute53

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionMixinProps",
		reflect.TypeOf((*CfnCidrCollectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin",
		reflect.TypeOf((*CfnCidrCollectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCidrCollectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnCidrCollectionPropsMixin.LocationProperty",
		reflect.TypeOf((*CfnCidrCollectionPropsMixin_LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnDNSSECMixinProps",
		reflect.TypeOf((*CfnDNSSECMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnDNSSECPropsMixin",
		reflect.TypeOf((*CfnDNSSECPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDNSSECPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckMixinProps",
		reflect.TypeOf((*CfnHealthCheckMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckPropsMixin",
		reflect.TypeOf((*CfnHealthCheckPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHealthCheckPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckPropsMixin.AlarmIdentifierProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_AlarmIdentifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckPropsMixin.HealthCheckConfigProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_HealthCheckConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHealthCheckPropsMixin.HealthCheckTagProperty",
		reflect.TypeOf((*CfnHealthCheckPropsMixin_HealthCheckTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZoneMixinProps",
		reflect.TypeOf((*CfnHostedZoneMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin",
		reflect.TypeOf((*CfnHostedZonePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHostedZonePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin.HostedZoneConfigProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_HostedZoneConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin.HostedZoneFeaturesProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_HostedZoneFeaturesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin.HostedZoneTagProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_HostedZoneTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin.QueryLoggingConfigProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_QueryLoggingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnHostedZonePropsMixin.VPCProperty",
		reflect.TypeOf((*CfnHostedZonePropsMixin_VPCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnKeySigningKeyMixinProps",
		reflect.TypeOf((*CfnKeySigningKeyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnKeySigningKeyPropsMixin",
		reflect.TypeOf((*CfnKeySigningKeyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKeySigningKeyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupMixinProps",
		reflect.TypeOf((*CfnRecordSetGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecordSetGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.AliasTargetProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_AliasTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.CidrRoutingConfigProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_CidrRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.CoordinatesProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_CoordinatesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.GeoLocationProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_GeoLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.GeoProximityLocationProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_GeoProximityLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetGroupPropsMixin.RecordSetProperty",
		reflect.TypeOf((*CfnRecordSetGroupPropsMixin_RecordSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetMixinProps",
		reflect.TypeOf((*CfnRecordSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin",
		reflect.TypeOf((*CfnRecordSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecordSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin.AliasTargetProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_AliasTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin.CidrRoutingConfigProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_CidrRoutingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin.CoordinatesProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_CoordinatesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin.GeoLocationProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_GeoLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnRecordSetPropsMixin.GeoProximityLocationProperty",
		reflect.TypeOf((*CfnRecordSetPropsMixin_GeoProximityLocationProperty)(nil)).Elem(),
	)
}
