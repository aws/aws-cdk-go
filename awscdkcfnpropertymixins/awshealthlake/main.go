package awshealthlake

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfileMixinProps",
		reflect.TypeOf((*CfnDataTransformationProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin",
		reflect.TypeOf((*CfnDataTransformationProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataTransformationProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin.ExistingVersionedProfileSourceProperty",
		reflect.TypeOf((*CfnDataTransformationProfilePropsMixin_ExistingVersionedProfileSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin.ProfileMappingSourceProperty",
		reflect.TypeOf((*CfnDataTransformationProfilePropsMixin_ProfileMappingSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin.SourceProperty",
		reflect.TypeOf((*CfnDataTransformationProfilePropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnDataTransformationProfilePropsMixin.StarterProfileSourceProperty",
		reflect.TypeOf((*CfnDataTransformationProfilePropsMixin_StarterProfileSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastoreMixinProps",
		reflect.TypeOf((*CfnFHIRDatastoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFHIRDatastorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin.CreatedAtProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_CreatedAtProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin.KmsEncryptionConfigProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_KmsEncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin.PreloadDataConfigProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_PreloadDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_healthlake.CfnFHIRDatastorePropsMixin.SseConfigurationProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_SseConfigurationProperty)(nil)).Elem(),
	)
}
