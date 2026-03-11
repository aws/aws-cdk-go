package awscomprehend

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierMixinProps",
		reflect.TypeOf((*CfnDocumentClassifierMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDocumentClassifierPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.AugmentedManifestsListItemProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_AugmentedManifestsListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.DocumentClassifierDocumentsProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierDocumentsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.DocumentClassifierInputDataConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierInputDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.DocumentClassifierOutputDataConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentClassifierOutputDataConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.DocumentReaderConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_DocumentReaderConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnDocumentClassifierPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnDocumentClassifierPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelMixinProps",
		reflect.TypeOf((*CfnFlywheelMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin",
		reflect.TypeOf((*CfnFlywheelPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlywheelPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.DataSecurityConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_DataSecurityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.DocumentClassificationConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_DocumentClassificationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.EntityRecognitionConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_EntityRecognitionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.EntityTypesListItemProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_EntityTypesListItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.TaskConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_TaskConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_comprehend.CfnFlywheelPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnFlywheelPropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
}
