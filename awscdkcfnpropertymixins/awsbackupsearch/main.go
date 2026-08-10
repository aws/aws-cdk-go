package awsbackupsearch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backupsearch.CfnSearchJobMixinProps",
		reflect.TypeOf((*CfnSearchJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_backupsearch.CfnSearchJobPropsMixin",
		reflect.TypeOf((*CfnSearchJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSearchJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backupsearch.CfnSearchJobPropsMixin.SearchScopeProperty",
		reflect.TypeOf((*CfnSearchJobPropsMixin_SearchScopeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_backupsearch.CfnSearchJobPropsMixin.TagsItemsProperty",
		reflect.TypeOf((*CfnSearchJobPropsMixin_TagsItemsProperty)(nil)).Elem(),
	)
}
