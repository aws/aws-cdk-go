package awsbedrockalphamixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseEmbeddingsModel",
		reflect.TypeOf((*KnowledgeBaseEmbeddingsModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBaseEmbeddingsModel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseEmbeddingsModelProps",
		reflect.TypeOf((*KnowledgeBaseEmbeddingsModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseOpenSearchServerlessStorage",
		reflect.TypeOf((*KnowledgeBaseOpenSearchServerlessStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBaseOpenSearchServerlessStorage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseOpenSearchServerlessStorageProps",
		reflect.TypeOf((*KnowledgeBaseOpenSearchServerlessStorageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseSupplementalDataStorage",
		reflect.TypeOf((*KnowledgeBaseSupplementalDataStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBaseSupplementalDataStorage{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseSupplementalDataStorageProps",
		reflect.TypeOf((*KnowledgeBaseSupplementalDataStorageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-bedrock-alpha.mixins.KnowledgeBaseTypeConfiguration",
		reflect.TypeOf((*KnowledgeBaseTypeConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBaseTypeConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			return &j
		},
	)
}
