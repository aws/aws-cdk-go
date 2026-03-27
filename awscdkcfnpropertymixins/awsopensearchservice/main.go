package awsopensearchservice

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnApplicationPropsMixin.AppConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AppConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnApplicationPropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnApplicationPropsMixin.IamIdentityCenterOptionsProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IamIdentityCenterOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin",
		reflect.TypeOf((*CfnDomainPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDomainPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.AIMLOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AIMLOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.AdvancedSecurityOptionsInputProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AdvancedSecurityOptionsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.ClusterConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.CognitoOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CognitoOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.ColdStorageOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ColdStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.DeploymentStrategyOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DeploymentStrategyOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.DomainEndpointOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEndpointOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.EBSOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EBSOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.IAMFederationOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IAMFederationOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.IdentityCenterOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdentityCenterOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.IdpProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_IdpProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.JWTOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_JWTOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.LogPublishingOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LogPublishingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.MasterUserOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MasterUserOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.NodeConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.NodeOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.NodeToNodeEncryptionOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeToNodeEncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.OffPeakWindowOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_OffPeakWindowOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.OffPeakWindowProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_OffPeakWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.S3VectorsEngineProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_S3VectorsEngineProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.SAMLOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SAMLOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.ServerlessVectorAccelerationProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ServerlessVectorAccelerationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.ServiceSoftwareOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ServiceSoftwareOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.SnapshotOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SnapshotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.SoftwareUpdateOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SoftwareUpdateOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.VPCOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_VPCOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.WindowStartTimeProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_WindowStartTimeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_opensearchservice.CfnDomainPropsMixin.ZoneAwarenessConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ZoneAwarenessConfigProperty)(nil)).Elem(),
	)
}
