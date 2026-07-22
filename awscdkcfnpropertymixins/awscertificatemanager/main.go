package awscertificatemanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAccountMixinProps",
		reflect.TypeOf((*CfnAccountMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAccountPropsMixin",
		reflect.TypeOf((*CfnAccountPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccountPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAccountPropsMixin.ExpiryEventsConfigurationProperty",
		reflect.TypeOf((*CfnAccountPropsMixin_ExpiryEventsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationMixinProps",
		reflect.TypeOf((*CfnAcmeDomainValidationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin",
		reflect.TypeOf((*CfnAcmeDomainValidationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAcmeDomainValidationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin.DnsPrevalidationOptionsProperty",
		reflect.TypeOf((*CfnAcmeDomainValidationPropsMixin_DnsPrevalidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin.DomainScopeProperty",
		reflect.TypeOf((*CfnAcmeDomainValidationPropsMixin_DomainScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin.PrevalidationOptionsProperty",
		reflect.TypeOf((*CfnAcmeDomainValidationPropsMixin_PrevalidationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeDomainValidationPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnAcmeDomainValidationPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeEndpointMixinProps",
		reflect.TypeOf((*CfnAcmeEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeEndpointPropsMixin",
		reflect.TypeOf((*CfnAcmeEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAcmeEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeEndpointPropsMixin.CertificateAuthorityProperty",
		reflect.TypeOf((*CfnAcmeEndpointPropsMixin_CertificateAuthorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeEndpointPropsMixin.PublicCertificateAuthorityProperty",
		reflect.TypeOf((*CfnAcmeEndpointPropsMixin_PublicCertificateAuthorityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeEndpointPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnAcmeEndpointPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingMixinProps",
		reflect.TypeOf((*CfnAcmeExternalAccountBindingMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin",
		reflect.TypeOf((*CfnAcmeExternalAccountBindingPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAcmeExternalAccountBindingPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin.ExpirationProperty",
		reflect.TypeOf((*CfnAcmeExternalAccountBindingPropsMixin_ExpirationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnAcmeExternalAccountBindingPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnAcmeExternalAccountBindingPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnCertificateMixinProps",
		reflect.TypeOf((*CfnCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnCertificatePropsMixin",
		reflect.TypeOf((*CfnCertificatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_certificatemanager.CfnCertificatePropsMixin.DomainValidationOptionProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_DomainValidationOptionProperty)(nil)).Elem(),
	)
}
