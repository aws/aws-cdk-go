package mixinsawsapptest

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCaseMixinProps",
		reflect.TypeOf((*CfnTestCaseMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin",
		reflect.TypeOf((*CfnTestCasePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTestCasePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.BatchProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_BatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.CloudFormationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_CloudFormationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.CompareActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_CompareActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.DataSetProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_DataSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.DatabaseCDCProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_DatabaseCDCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.FileMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_FileMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.InputFileProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_InputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.InputProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.M2ManagedActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2ManagedActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.M2ManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2ManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.M2NonManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_M2NonManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.MainframeActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.MainframeActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.MainframeActionTypeProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_MainframeActionTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.OutputFileProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_OutputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.OutputProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.ResourceActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_ResourceActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.ScriptProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_ScriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.SourceDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_SourceDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.StepActionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_StepActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.StepProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_StepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.TN3270Property",
		reflect.TypeOf((*CfnTestCasePropsMixin_TN3270Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.TargetDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_TargetDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apptest.mixins.CfnTestCasePropsMixin.TestCaseLatestVersionProperty",
		reflect.TypeOf((*CfnTestCasePropsMixin_TestCaseLatestVersionProperty)(nil)).Elem(),
	)
}
