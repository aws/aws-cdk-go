package awsemrcontainers

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointMixinProps",
		reflect.TypeOf((*CfnEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin",
		reflect.TypeOf((*CfnEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.CertificateProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CertificateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.CloudWatchMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_CloudWatchMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.ConfigurationOverridesProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_ConfigurationOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.ContainerLogRotationConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_ContainerLogRotationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.EMREKSConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_EMREKSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnEndpointPropsMixin.S3MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnEndpointPropsMixin_S3MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationMixinProps",
		reflect.TypeOf((*CfnSecurityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSecurityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.AtRestEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AtRestEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.AuthorizationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_AuthorizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.ContainerInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_ContainerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.ContainerProviderProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_ContainerProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.EksInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EksInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.IAMConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_IAMConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.IdentityCenterConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_IdentityCenterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.InTransitEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_InTransitEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.LakeFormationConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_LakeFormationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.LocalDiskEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_LocalDiskEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.S3EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_S3EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.SecureNamespaceInfoProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_SecureNamespaceInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.SecurityConfigurationDataProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_SecurityConfigurationDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnSecurityConfigurationPropsMixin.TLSCertificateConfigurationProperty",
		reflect.TypeOf((*CfnSecurityConfigurationPropsMixin_TLSCertificateConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnVirtualClusterMixinProps",
		reflect.TypeOf((*CfnVirtualClusterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnVirtualClusterPropsMixin",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVirtualClusterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnVirtualClusterPropsMixin.ContainerInfoProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_ContainerInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnVirtualClusterPropsMixin.ContainerProviderProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_ContainerProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_emrcontainers.CfnVirtualClusterPropsMixin.EksInfoProperty",
		reflect.TypeOf((*CfnVirtualClusterPropsMixin_EksInfoProperty)(nil)).Elem(),
	)
}
