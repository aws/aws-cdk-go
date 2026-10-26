package awslicensemanagerusersubscriptions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointMixinProps",
		reflect.TypeOf((*CfnLicenseServerEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLicenseServerEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.CredentialsProviderProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_CredentialsProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.LicenseServerProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_LicenseServerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.LicenseServerSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_LicenseServerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.RdsSalSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_RdsSalSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.SecretsManagerCredentialsProviderProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_SecretsManagerCredentialsProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.ServerEndpointProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_ServerEndpointProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_licensemanagerusersubscriptions.CfnLicenseServerEndpointPropsMixin.ServerSettingsProperty",
		reflect.TypeOf((*CfnLicenseServerEndpointPropsMixin_ServerSettingsProperty)(nil)).Elem(),
	)
}
