// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		reflect.TypeOf((*AcceleratorType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_AcceleratorType{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ContainerDefinition",
		reflect.TypeOf((*ContainerDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImage",
		reflect.TypeOf((*ContainerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ContainerImageConfig",
		reflect.TypeOf((*ContainerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "endpointRef", GoGetter: "EndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "findInstanceProductionVariant", GoMethod: "FindInstanceProductionVariant"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProductionVariants", GoGetter: "InstanceProductionVariants"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Endpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.EndpointAttributes",
		reflect.TypeOf((*EndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.EndpointConfig",
		reflect.TypeOf((*EndpointConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addInstanceProductionVariant", GoMethod: "AddInstanceProductionVariant"},
			_jsii_.MemberMethod{JsiiMethod: "addServerlessProductionVariant", GoMethod: "AddServerlessProductionVariant"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigArn", GoGetter: "EndpointConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigName", GoGetter: "EndpointConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_EndpointConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEndpointConfig)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.EndpointConfigProps",
		reflect.TypeOf((*EndpointConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.EndpointProps",
		reflect.TypeOf((*EndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-sagemaker-alpha.IEndpoint",
		reflect.TypeOf((*IEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "endpointName", GoGetter: "EndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "endpointRef", GoGetter: "EndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awssagemakerIEndpoint)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-sagemaker-alpha.IEndpointConfig",
		reflect.TypeOf((*IEndpointConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigArn", GoGetter: "EndpointConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigName", GoGetter: "EndpointConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IEndpointConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-sagemaker-alpha.IEndpointInstanceProductionVariant",
		reflect.TypeOf((*IEndpointInstanceProductionVariant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "autoScaleInstanceCount", GoMethod: "AutoScaleInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCpuUtilization", GoMethod: "MetricCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricDiskUtilization", GoMethod: "MetricDiskUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricGpuMemoryUtilization", GoMethod: "MetricGpuMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricGpuUtilization", GoMethod: "MetricGpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationResponseCode", GoMethod: "MetricInvocationResponseCode"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsPerInstance", GoMethod: "MetricInvocationsPerInstance"},
			_jsii_.MemberMethod{JsiiMethod: "metricMemoryUtilization", GoMethod: "MetricMemoryUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricModelLatency", GoMethod: "MetricModelLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricOverheadLatency", GoMethod: "MetricOverheadLatency"},
			_jsii_.MemberProperty{JsiiProperty: "variantName", GoGetter: "VariantName"},
		},
		func() interface{} {
			return &jsiiProxy_IEndpointInstanceProductionVariant{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-sagemaker-alpha.IModel",
		reflect.TypeOf((*IModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelName", GoGetter: "ModelName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IModel{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.InstanceProductionVariantProps",
		reflect.TypeOf((*InstanceProductionVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		reflect.TypeOf((*InstanceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_InstanceType{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-sagemaker-alpha.InvocationHttpResponseCode",
		reflect.TypeOf((*InvocationHttpResponseCode)(nil)).Elem(),
		map[string]interface{}{
			"INVOCATION_4XX_ERRORS": InvocationHttpResponseCode_INVOCATION_4XX_ERRORS,
			"INVOCATION_5XX_ERRORS": InvocationHttpResponseCode_INVOCATION_5XX_ERRORS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.InvocationsScalingProps",
		reflect.TypeOf((*InvocationsScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.Model",
		reflect.TypeOf((*Model)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainer", GoMethod: "AddContainer"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelName", GoGetter: "ModelName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Model{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IModel)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ModelAttributes",
		reflect.TypeOf((*ModelAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.ModelData",
		reflect.TypeOf((*ModelData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ModelData{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ModelDataConfig",
		reflect.TypeOf((*ModelDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ModelProps",
		reflect.TypeOf((*ModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.Pipeline",
		reflect.TypeOf((*Pipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantStartPipelineExecution", GoMethod: "GrantStartPipelineExecution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineArn", GoGetter: "PipelineArn"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineName", GoGetter: "PipelineName"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineRef", GoGetter: "PipelineRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Pipeline{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awssagemakerIPipeline)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.PipelineAttributes",
		reflect.TypeOf((*PipelineAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.PipelineProps",
		reflect.TypeOf((*PipelineProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-sagemaker-alpha.ScalableInstanceCount",
		reflect.TypeOf((*ScalableInstanceCount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "doScaleOnMetric", GoMethod: "DoScaleOnMetric"},
			_jsii_.MemberMethod{JsiiMethod: "doScaleOnSchedule", GoMethod: "DoScaleOnSchedule"},
			_jsii_.MemberMethod{JsiiMethod: "doScaleToTrackMetric", GoMethod: "DoScaleToTrackMetric"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "scalableTargetRef", GoGetter: "ScalableTargetRef"},
			_jsii_.MemberMethod{JsiiMethod: "scaleOnInvocations", GoMethod: "ScaleOnInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ScalableInstanceCount{}
			_jsii_.InitJsiiProxy(&j.Type__awsapplicationautoscalingBaseScalableAttribute)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ScalableInstanceCountProps",
		reflect.TypeOf((*ScalableInstanceCountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-sagemaker-alpha.ServerlessProductionVariantProps",
		reflect.TypeOf((*ServerlessProductionVariantProps)(nil)).Elem(),
	)
}
