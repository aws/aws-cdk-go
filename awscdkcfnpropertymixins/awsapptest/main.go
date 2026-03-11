package awsapptest

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCaseMixinProps",
		reflect.TypeOf((*CfnTestCaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin",
		reflect.TypeOf((*CfnTestCasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTestCasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.BatchProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_BatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.CloudFormationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_CloudFormationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.CompareActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_CompareActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.DataSetProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_DataSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.DatabaseCDCProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_DatabaseCDCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.FileMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_FileMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.InputFileProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_InputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.InputProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.M2ManagedActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2ManagedActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.M2ManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2ManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.M2NonManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2NonManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.MainframeActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.MainframeActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.MainframeActionTypeProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.OutputFileProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_OutputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.OutputProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.ResourceActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_ResourceActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.ScriptProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_ScriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.SourceDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_SourceDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.StepActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_StepActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.StepProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_StepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.TN3270Property",
		reflect.TypeOf((*CfnTestCasePropsMixin_TN3270Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.TargetDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_TargetDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apptest.CfnTestCasePropsMixin.TestCaseLatestVersionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_TestCaseLatestVersionProperty)(nil)).Elem(),
	)
}
