package previewawsopensearchservicemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnApplicationPropsMixin.AppConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AppConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnApplicationPropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnApplicationPropsMixin.IamIdentityCenterOptionsProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IamIdentityCenterOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.AIMLOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AIMLOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.AdvancedSecurityOptionsInputProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AdvancedSecurityOptionsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.CognitoOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CognitoOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.ColdStorageOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ColdStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.DomainEndpointOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEndpointOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.EBSOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EBSOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.IAMFederationOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IAMFederationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.IdentityCenterOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdentityCenterOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.IdpProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.JWTOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JWTOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.LogPublishingOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LogPublishingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.MasterUserOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MasterUserOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.NodeConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.NodeOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.NodeToNodeEncryptionOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeToNodeEncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.OffPeakWindowOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_OffPeakWindowOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.OffPeakWindowProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_OffPeakWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.S3VectorsEngineProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_S3VectorsEngineProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.SAMLOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SAMLOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.ServerlessVectorAccelerationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ServerlessVectorAccelerationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.ServiceSoftwareOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ServiceSoftwareOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.SnapshotOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SnapshotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.SoftwareUpdateOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SoftwareUpdateOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.VPCOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_VPCOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.WindowStartTimeProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_WindowStartTimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_opensearchservice.mixins.CfnDomainPropsMixin.ZoneAwarenessConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ZoneAwarenessConfigProperty)(nil)).Elem(),
	)
}
