package previewawshealthlakemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastoreMixinProps",
		reflect.TypeOf((*CfnFHIRDatastoreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFHIRDatastorePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin.CreatedAtProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_CreatedAtProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin.IdentityProviderConfigurationProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_IdentityProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin.KmsEncryptionConfigProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_KmsEncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin.PreloadDataConfigProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_PreloadDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_healthlake.mixins.CfnFHIRDatastorePropsMixin.SseConfigurationProperty",
		reflect.TypeOf((*CfnFHIRDatastorePropsMixin_SseConfigurationProperty)(nil)).Elem(),
	)
}
