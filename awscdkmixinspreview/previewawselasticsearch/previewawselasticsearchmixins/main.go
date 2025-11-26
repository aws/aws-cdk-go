package previewawselasticsearchmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainMixinProps",
		reflect.TypeOf((*CfnDomainMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.AdvancedSecurityOptionsInputProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_AdvancedSecurityOptionsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.CognitoOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_CognitoOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.ColdStorageOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ColdStorageOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.DomainEndpointOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_DomainEndpointOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.EBSOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EBSOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.ElasticsearchClusterConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ElasticsearchClusterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.EncryptionAtRestOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_EncryptionAtRestOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.LogPublishingOptionProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_LogPublishingOptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.MasterUserOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_MasterUserOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.NodeToNodeEncryptionOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_NodeToNodeEncryptionOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.SnapshotOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_SnapshotOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.VPCOptionsProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_VPCOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticsearch.mixins.CfnDomainPropsMixin.ZoneAwarenessConfigProperty",
		reflect.TypeOf((*CfnDomainPropsMixin_ZoneAwarenessConfigProperty)(nil)).Elem(),
	)
}
