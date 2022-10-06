package awscdkintegtestsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		reflect.TypeOf((*ActualResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
		},
		func() interface{} {
			return &jsiiProxy_ActualResult{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.ApiCallBase",
		reflect.TypeOf((*ApiCallBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiCallResource", GoGetter: "ApiCallResource"},
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberProperty{JsiiProperty: "expectedResult", GoGetter: "ExpectedResult"},
			_jsii_.MemberProperty{JsiiProperty: "flattenResponse", GoGetter: "FlattenResponse"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiCallBase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IApiCall)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AssertionRequest",
		reflect.TypeOf((*AssertionRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AssertionResult",
		reflect.TypeOf((*AssertionResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AssertionResultData",
		reflect.TypeOf((*AssertionResultData)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/integ-tests-alpha.AssertionType",
		reflect.TypeOf((*AssertionType)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": AssertionType_EQUALS,
			"OBJECT_LIKE": AssertionType_OBJECT_LIKE,
			"ARRAY_WITH": AssertionType_ARRAY_WITH,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
		reflect.TypeOf((*AssertionsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPolicyStatementFromSdkCall", GoMethod: "AddPolicyStatementFromSdkCall"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "encode", GoMethod: "Encode"},
			_jsii_.MemberProperty{JsiiProperty: "handlerRoleArn", GoGetter: "HandlerRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AssertionsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.AwsApiCall",
		reflect.TypeOf((*AwsApiCall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiCallResource", GoGetter: "ApiCallResource"},
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberProperty{JsiiProperty: "expectedResult", GoGetter: "ExpectedResult"},
			_jsii_.MemberProperty{JsiiProperty: "flattenResponse", GoGetter: "FlattenResponse"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AwsApiCall{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiCallBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AwsApiCallOptions",
		reflect.TypeOf((*AwsApiCallOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AwsApiCallProps",
		reflect.TypeOf((*AwsApiCallProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AwsApiCallRequest",
		reflect.TypeOf((*AwsApiCallRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.AwsApiCallResult",
		reflect.TypeOf((*AwsApiCallResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.EqualsAssertion",
		reflect.TypeOf((*EqualsAssertion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EqualsAssertion{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.EqualsAssertionProps",
		reflect.TypeOf((*EqualsAssertionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		reflect.TypeOf((*ExpectedResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
		},
		func() interface{} {
			return &jsiiProxy_ExpectedResult{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/integ-tests-alpha.IApiCall",
		reflect.TypeOf((*IApiCall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
		},
		func() interface{} {
			j := jsiiProxy_IApiCall{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/integ-tests-alpha.IDeployAssert",
		reflect.TypeOf((*IDeployAssert)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsApiCall", GoMethod: "AwsApiCall"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "invokeFunction", GoMethod: "InvokeFunction"},
		},
		func() interface{} {
			return &jsiiProxy_IDeployAssert{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.IntegTest",
		reflect.TypeOf((*IntegTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		reflect.TypeOf((*IntegTestCase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTestCase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseProps",
		reflect.TypeOf((*IntegTestCaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		reflect.TypeOf((*IntegTestCaseStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTestCaseStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStackProps",
		reflect.TypeOf((*IntegTestCaseStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.IntegTestProps",
		reflect.TypeOf((*IntegTestProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/integ-tests-alpha.InvocationType",
		reflect.TypeOf((*InvocationType)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": InvocationType_EVENT,
			"REQUEST_RESPONE": InvocationType_REQUEST_RESPONE,
			"DRY_RUN": InvocationType_DRY_RUN,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.LambdaInvokeFunction",
		reflect.TypeOf((*LambdaInvokeFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiCallResource", GoGetter: "ApiCallResource"},
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberProperty{JsiiProperty: "expectedResult", GoGetter: "ExpectedResult"},
			_jsii_.MemberProperty{JsiiProperty: "flattenResponse", GoGetter: "FlattenResponse"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvokeFunction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsApiCall)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/integ-tests-alpha.LambdaInvokeFunctionProps",
		reflect.TypeOf((*LambdaInvokeFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/integ-tests-alpha.LogType",
		reflect.TypeOf((*LogType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": LogType_NONE,
			"TAIL": LogType_TAIL,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/integ-tests-alpha.Match",
		reflect.TypeOf((*Match)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Match{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/integ-tests-alpha.Status",
		reflect.TypeOf((*Status)(nil)).Elem(),
		map[string]interface{}{
			"PASS": Status_PASS,
			"FAIL": Status_FAIL,
		},
	)
}
