package awsgreengrassv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		reflect.TypeOf((*CfnComponentVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrComponentName", GoGetter: "AttrComponentName"},
			_jsii_.MemberProperty{JsiiProperty: "attrComponentVersion", GoGetter: "AttrComponentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inlineRecipe", GoGetter: "InlineRecipe"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComponentVersion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.ComponentDependencyRequirementProperty",
		reflect.TypeOf((*CfnComponentVersion_ComponentDependencyRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.ComponentPlatformProperty",
		reflect.TypeOf((*CfnComponentVersion_ComponentPlatformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaContainerParamsProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaContainerParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaDeviceMountProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaDeviceMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaEventSourceProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaEventSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaExecutionParametersProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaExecutionParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaFunctionRecipeSourceProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaFunctionRecipeSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaLinuxProcessParamsProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaLinuxProcessParamsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersion.LambdaVolumeMountProperty",
		reflect.TypeOf((*CfnComponentVersion_LambdaVolumeMountProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_greengrassv2.CfnComponentVersionProps",
		reflect.TypeOf((*CfnComponentVersionProps)(nil)).Elem(),
	)
}
