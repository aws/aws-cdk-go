package mixinsawsbcmdataexports

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportMixinProps",
		reflect.TypeOf((*CfnExportMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin",
		reflect.TypeOf((*CfnExportPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExportPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.DataQueryProperty",
		reflect.TypeOf((*CfnExportPropsMixin_DataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.DestinationConfigurationsProperty",
		reflect.TypeOf((*CfnExportPropsMixin_DestinationConfigurationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.ExportProperty",
		reflect.TypeOf((*CfnExportPropsMixin_ExportProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.RefreshCadenceProperty",
		reflect.TypeOf((*CfnExportPropsMixin_RefreshCadenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnExportPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnExportPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bcmdataexports.mixins.CfnExportPropsMixin.S3OutputConfigurationsProperty",
		reflect.TypeOf((*CfnExportPropsMixin_S3OutputConfigurationsProperty)(nil)).Elem(),
	)
}
