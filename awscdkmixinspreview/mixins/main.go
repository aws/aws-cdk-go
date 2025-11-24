package mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.mixins.CfnPropertyMixinOptions",
		reflect.TypeOf((*CfnPropertyMixinOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.mixins.PropertyMergeStrategy",
		reflect.TypeOf((*PropertyMergeStrategy)(nil)).Elem(),
		map[string]interface{}{
			"OVERRIDE": PropertyMergeStrategy_OVERRIDE,
			"MERGE": PropertyMergeStrategy_MERGE,
		},
	)
}
