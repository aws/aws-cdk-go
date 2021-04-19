package awssynthetics

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.ArtifactsBucketLocation",
		reflect.TypeOf((*ArtifactsBucketLocation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.AssetCode",
		reflect.TypeOf((*AssetCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_AssetCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.Canary",
		reflect.TypeOf((*Canary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsBucket", GoGetter: "ArtifactsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "canaryId", GoGetter: "CanaryId"},
			_jsii_.MemberProperty{JsiiProperty: "canaryName", GoGetter: "CanaryName"},
			_jsii_.MemberProperty{JsiiProperty: "canaryState", GoGetter: "CanaryState"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metricDuration", GoMethod: "MetricDuration"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricSuccessPercent", GoMethod: "MetricSuccessPercent"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Canary{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CanaryProps",
		reflect.TypeOf((*CanaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.CfnCanary",
		reflect.TypeOf((*CfnCanary)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "artifactS3Location", GoGetter: "ArtifactS3Location"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "failureRetentionPeriod", GoGetter: "FailureRetentionPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "runConfig", GoGetter: "RunConfig"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeVersion", GoGetter: "RuntimeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "startCanaryAfterCreation", GoGetter: "StartCanaryAfterCreation"},
			_jsii_.MemberProperty{JsiiProperty: "successRetentionPeriod", GoGetter: "SuccessRetentionPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfig", GoGetter: "VpcConfig"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCanary{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CfnCanary.CodeProperty",
		reflect.TypeOf((*CfnCanary_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CfnCanary.RunConfigProperty",
		reflect.TypeOf((*CfnCanary_RunConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CfnCanary.ScheduleProperty",
		reflect.TypeOf((*CfnCanary_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CfnCanary.VPCConfigProperty",
		reflect.TypeOf((*CfnCanary_VPCConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CfnCanaryProps",
		reflect.TypeOf((*CfnCanaryProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.Code",
		reflect.TypeOf((*Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_Code{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CodeConfig",
		reflect.TypeOf((*CodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_synthetics.CustomTestOptions",
		reflect.TypeOf((*CustomTestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.InlineCode",
		reflect.TypeOf((*InlineCode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InlineCode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.Runtime",
		reflect.TypeOf((*Runtime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Runtime{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.S3Code",
		reflect.TypeOf((*S3Code)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3Code{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Code)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.Schedule",
		reflect.TypeOf((*Schedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "expressionString", GoGetter: "ExpressionString"},
		},
		func() interface{} {
			return &jsiiProxy_Schedule{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_synthetics.Test",
		reflect.TypeOf((*Test)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
		},
		func() interface{} {
			return &jsiiProxy_Test{}
		},
	)
}
