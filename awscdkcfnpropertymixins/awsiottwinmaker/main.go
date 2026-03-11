package awsiottwinmaker

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypeMixinProps",
		reflect.TypeOf((*CfnComponentTypeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin",
		reflect.TypeOf((*CfnComponentTypePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentTypePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.CompositeComponentTypeProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_CompositeComponentTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.DataConnectorProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataConnectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.DataTypeProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.DataValueProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_DataValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.ErrorProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_ErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.FunctionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_FunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.LambdaFunctionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_LambdaFunctionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.PropertyDefinitionProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_PropertyDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.RelationshipProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_RelationshipProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnComponentTypePropsMixin.StatusProperty",
		reflect.TypeOf((*CfnComponentTypePropsMixin_StatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityMixinProps",
		reflect.TypeOf((*CfnEntityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin",
		reflect.TypeOf((*CfnEntityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEntityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.ComponentProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_ComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.CompositeComponentProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_CompositeComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.DataValueProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_DataValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.ErrorProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_ErrorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.PropertyGroupProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_PropertyGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.PropertyProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_PropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnEntityPropsMixin.StatusProperty",
		reflect.TypeOf((*CfnEntityPropsMixin_StatusProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnSceneMixinProps",
		reflect.TypeOf((*CfnSceneMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnScenePropsMixin",
		reflect.TypeOf((*CfnScenePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnScenePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnSyncJobMixinProps",
		reflect.TypeOf((*CfnSyncJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnSyncJobPropsMixin",
		reflect.TypeOf((*CfnSyncJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSyncJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnWorkspaceMixinProps",
		reflect.TypeOf((*CfnWorkspaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iottwinmaker.CfnWorkspacePropsMixin",
		reflect.TypeOf((*CfnWorkspacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkspacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
