package awsacmpca

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityActivationMixinProps",
		reflect.TypeOf((*CfnCertificateAuthorityActivationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityActivationPropsMixin",
		reflect.TypeOf((*CfnCertificateAuthorityActivationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificateAuthorityActivationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityMixinProps",
		reflect.TypeOf((*CfnCertificateAuthorityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCertificateAuthorityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.AccessDescriptionProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_AccessDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.AccessMethodProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_AccessMethodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.CrlConfigurationProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_CrlConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.CrlDistributionPointExtensionConfigurationProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_CrlDistributionPointExtensionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.CsrExtensionsProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_CsrExtensionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.CustomAttributeProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_CustomAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.EdiPartyNameProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_EdiPartyNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.GeneralNameProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_GeneralNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.KeyUsageProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_KeyUsageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.OcspConfigurationProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_OcspConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.OtherNameProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_OtherNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.RevocationConfigurationProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_RevocationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateAuthorityPropsMixin.SubjectProperty",
		reflect.TypeOf((*CfnCertificateAuthorityPropsMixin_SubjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificateMixinProps",
		reflect.TypeOf((*CfnCertificateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.ApiPassthroughProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_ApiPassthroughProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.CustomAttributeProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_CustomAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.CustomExtensionProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_CustomExtensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.EdiPartyNameProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_EdiPartyNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.ExtendedKeyUsageProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_ExtendedKeyUsageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.ExtensionsProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_ExtensionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.GeneralNameProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_GeneralNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.KeyUsageProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_KeyUsageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.OtherNameProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_OtherNameProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.PolicyInformationProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_PolicyInformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.PolicyQualifierInfoProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_PolicyQualifierInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.QualifierProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_QualifierProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.SubjectProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_SubjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnCertificatePropsMixin.ValidityProperty",
		reflect.TypeOf((*CfnCertificatePropsMixin_ValidityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnPermissionMixinProps",
		reflect.TypeOf((*CfnPermissionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_acmpca.CfnPermissionPropsMixin",
		reflect.TypeOf((*CfnPermissionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPermissionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
