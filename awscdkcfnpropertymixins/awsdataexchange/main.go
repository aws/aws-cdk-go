package awsdataexchange

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnDataSetMixinProps",
		reflect.TypeOf((*CfnDataSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnDataSetPropsMixin",
		reflect.TypeOf((*CfnDataSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionMixinProps",
		reflect.TypeOf((*CfnEventActionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin",
		reflect.TypeOf((*CfnEventActionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventActionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.AutoExportRevisionDestinationEntryProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_AutoExportRevisionDestinationEntryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.AutoExportRevisionToS3RequestDetailsProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_AutoExportRevisionToS3RequestDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.EventProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_EventProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.ExportServerSideEncryptionProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_ExportServerSideEncryptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_dataexchange.CfnEventActionPropsMixin.RevisionPublishedProperty",
		reflect.TypeOf((*CfnEventActionPropsMixin_RevisionPublishedProperty)(nil)).Elem(),
	)
}
