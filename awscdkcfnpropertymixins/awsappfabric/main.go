package awsappfabric

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationMixinProps",
		reflect.TypeOf((*CfnAppAuthorizationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin",
		reflect.TypeOf((*CfnAppAuthorizationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppAuthorizationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin.ApiKeyCredentialProperty",
		reflect.TypeOf((*CfnAppAuthorizationPropsMixin_ApiKeyCredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin.CredentialProperty",
		reflect.TypeOf((*CfnAppAuthorizationPropsMixin_CredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin.Oauth2CredentialProperty",
		reflect.TypeOf((*CfnAppAuthorizationPropsMixin_Oauth2CredentialProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnAppAuthorizationPropsMixin.TenantProperty",
		reflect.TypeOf((*CfnAppAuthorizationPropsMixin_TenantProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationMixinProps",
		reflect.TypeOf((*CfnIngestionDestinationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnIngestionDestinationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.AuditLogDestinationConfigurationProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_AuditLogDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.AuditLogProcessingConfigurationProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_AuditLogProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.DestinationConfigurationProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_DestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.FirehoseStreamProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_FirehoseStreamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.ProcessingConfigurationProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_ProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_appfabric.CfnIngestionDestinationPropsMixin.S3BucketProperty",
		reflect.TypeOf((*CfnIngestionDestinationPropsMixin_S3BucketProperty)(nil)).Elem(),
	)
}
