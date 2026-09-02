package awsbcmdataexports

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportMixinProps",
		reflect.TypeOf((*CfnExportMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin",
		reflect.TypeOf((*CfnExportPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExportPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.DataQueryProperty",
		reflect.TypeOf((*CfnExportPropsMixin_DataQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.DestinationConfigurationsProperty",
		reflect.TypeOf((*CfnExportPropsMixin_DestinationConfigurationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.ExportProperty",
		reflect.TypeOf((*CfnExportPropsMixin_ExportProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.RefreshCadenceProperty",
		reflect.TypeOf((*CfnExportPropsMixin_RefreshCadenceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnExportPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.S3DestinationProperty",
		reflect.TypeOf((*CfnExportPropsMixin_S3DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcmdataexports.CfnExportPropsMixin.S3OutputConfigurationsProperty",
		reflect.TypeOf((*CfnExportPropsMixin_S3OutputConfigurationsProperty)(nil)).Elem(),
	)
}
