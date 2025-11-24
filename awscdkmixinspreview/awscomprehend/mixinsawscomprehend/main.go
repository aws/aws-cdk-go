package mixinsawscomprehend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierMixinProps",
		reflect.TypeOf((*CfnDocumentClassifierMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDocumentClassifierPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.AugmentedManifestsListItemProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_AugmentedManifestsListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.DocumentClassifierDocumentsProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierDocumentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.DocumentClassifierInputDataConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierInputDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.DocumentClassifierOutputDataConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierOutputDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.DocumentReaderConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentReaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnDocumentClassifierPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelMixinProps",
		reflect.TypeOf((*CfnFlywheelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin",
		reflect.TypeOf((*CfnFlywheelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlywheelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.DataSecurityConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_DataSecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.DocumentClassificationConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_DocumentClassificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.EntityRecognitionConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_EntityRecognitionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.EntityTypesListItemProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_EntityTypesListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.TaskConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_TaskConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_comprehend.mixins.CfnFlywheelPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
}
