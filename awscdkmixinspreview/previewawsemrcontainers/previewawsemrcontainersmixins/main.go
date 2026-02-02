package previewawsemrcontainersmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointMixinProps",
		reflect.TypeOf((*CfnEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin",
		reflect.TypeOf((*CfnEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.CloudWatchMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CloudWatchMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.ConfigurationOverridesProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_ConfigurationOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.ContainerLogRotationConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_ContainerLogRotationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.EMREKSConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_EMREKSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnEndpointPropsMixin.S3MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_S3MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.AtRestEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AtRestEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.AuthorizationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AuthorizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.ContainerInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_ContainerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.ContainerProviderProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_ContainerProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.EksInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EksInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.IAMConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_IAMConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.IdentityCenterConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_IdentityCenterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.InTransitEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_InTransitEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.LocalDiskEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_LocalDiskEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.S3EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_S3EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.SecureNamespaceInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_SecureNamespaceInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.SecurityConfigurationDataProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_SecurityConfigurationDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnSecurityConfigurationPropsMixin.TLSCertificateConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_TLSCertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterMixinProps",
		reflect.TypeOf((*CfnVirtualClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin.ContainerInfoProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_ContainerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin.ContainerProviderProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_ContainerProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrcontainers.mixins.CfnVirtualClusterPropsMixin.EksInfoProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_EksInfoProperty)(nil)).Elem(),
	)
}
