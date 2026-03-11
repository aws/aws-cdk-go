package awselasticsearch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.AdvancedSecurityOptionsInputProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AdvancedSecurityOptionsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.CognitoOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CognitoOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.ColdStorageOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ColdStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.DomainEndpointOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEndpointOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.EBSOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EBSOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.ElasticsearchClusterConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ElasticsearchClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.LogPublishingOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LogPublishingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.MasterUserOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MasterUserOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.NodeToNodeEncryptionOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeToNodeEncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.SnapshotOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SnapshotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.VPCOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_VPCOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_elasticsearch.CfnDomainPropsMixin.ZoneAwarenessConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ZoneAwarenessConfigProperty)(nil)).Elem(),
	)
}
