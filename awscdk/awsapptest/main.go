package awsapptest

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_apptest.CfnTestCase",
		reflect.TypeOf((*CfnTestCase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastUpdateTime", GoGetter: "AttrLastUpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrLatestVersion", GoGetter: "AttrLatestVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrTestCaseArn", GoGetter: "AttrTestCaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrTestCaseId", GoGetter: "AttrTestCaseId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTestCaseVersion", GoGetter: "AttrTestCaseVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "steps", GoGetter: "Steps"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "testCaseRef", GoGetter: "TestCaseRef"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTestCase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapptestITestCaseRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.BatchProperty",
		reflect.TypeOf((*CfnTestCase_BatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.CloudFormationActionProperty",
		reflect.TypeOf((*CfnTestCase_CloudFormationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.CompareActionProperty",
		reflect.TypeOf((*CfnTestCase_CompareActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.DataSetProperty",
		reflect.TypeOf((*CfnTestCase_DataSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.DatabaseCDCProperty",
		reflect.TypeOf((*CfnTestCase_DatabaseCDCProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.FileMetadataProperty",
		reflect.TypeOf((*CfnTestCase_FileMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.InputFileProperty",
		reflect.TypeOf((*CfnTestCase_InputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.InputProperty",
		reflect.TypeOf((*CfnTestCase_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.M2ManagedActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCase_M2ManagedActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.M2ManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCase_M2ManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.M2NonManagedApplicationActionProperty",
		reflect.TypeOf((*CfnTestCase_M2NonManagedApplicationActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.MainframeActionPropertiesProperty",
		reflect.TypeOf((*CfnTestCase_MainframeActionPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.MainframeActionProperty",
		reflect.TypeOf((*CfnTestCase_MainframeActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.MainframeActionTypeProperty",
		reflect.TypeOf((*CfnTestCase_MainframeActionTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.OutputFileProperty",
		reflect.TypeOf((*CfnTestCase_OutputFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.OutputProperty",
		reflect.TypeOf((*CfnTestCase_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.ResourceActionProperty",
		reflect.TypeOf((*CfnTestCase_ResourceActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.ScriptProperty",
		reflect.TypeOf((*CfnTestCase_ScriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.SourceDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCase_SourceDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.StepActionProperty",
		reflect.TypeOf((*CfnTestCase_StepActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.StepProperty",
		reflect.TypeOf((*CfnTestCase_StepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.TN3270Property",
		reflect.TypeOf((*CfnTestCase_TN3270Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.TargetDatabaseMetadataProperty",
		reflect.TypeOf((*CfnTestCase_TargetDatabaseMetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCase.TestCaseLatestVersionProperty",
		reflect.TypeOf((*CfnTestCase_TestCaseLatestVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_apptest.CfnTestCaseProps",
		reflect.TypeOf((*CfnTestCaseProps)(nil)).Elem(),
	)
}
