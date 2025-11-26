package previewawsresourcegroupsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupMixinProps",
		reflect.TypeOf((*CfnGroupMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin",
		reflect.TypeOf((*CfnGroupPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGroupPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin.ConfigurationItemProperty",
		reflect.TypeOf((*CfnGroupPropsMixin_ConfigurationItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin.ConfigurationParameterProperty",
		reflect.TypeOf((*CfnGroupPropsMixin_ConfigurationParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin.QueryProperty",
		reflect.TypeOf((*CfnGroupPropsMixin_QueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin.ResourceQueryProperty",
		reflect.TypeOf((*CfnGroupPropsMixin_ResourceQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnGroupPropsMixin.TagFilterProperty",
		reflect.TypeOf((*CfnGroupPropsMixin_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskMixinProps",
		reflect.TypeOf((*CfnTagSyncTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_resourcegroups.mixins.CfnTagSyncTaskPropsMixin",
		reflect.TypeOf((*CfnTagSyncTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTagSyncTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}
