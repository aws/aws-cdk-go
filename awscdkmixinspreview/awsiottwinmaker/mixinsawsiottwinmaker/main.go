package mixinsawsiottwinmaker

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypeMixinProps",
		reflect.TypeOf((*CfnComponentTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin",
		reflect.TypeOf((*CfnComponentTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.CompositeComponentTypeProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_CompositeComponentTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.DataConnectorProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataConnectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.DataTypeProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.DataValueProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.ErrorProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_ErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.FunctionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_FunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.LambdaFunctionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_LambdaFunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.PropertyDefinitionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_PropertyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.RelationshipProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_RelationshipProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnComponentTypePropsMixin.StatusProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_StatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityMixinProps",
		reflect.TypeOf((*CfnEntityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin",
		reflect.TypeOf((*CfnEntityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEntityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.ComponentProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_ComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.CompositeComponentProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_CompositeComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.DataValueProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_DataValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.ErrorProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_ErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.PropertyProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_PropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnEntityPropsMixin.StatusProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_StatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnSceneMixinProps",
		reflect.TypeOf((*CfnSceneMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnScenePropsMixin",
		reflect.TypeOf((*CfnScenePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScenePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnSyncJobMixinProps",
		reflect.TypeOf((*CfnSyncJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnSyncJobPropsMixin",
		reflect.TypeOf((*CfnSyncJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSyncJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iottwinmaker.mixins.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
