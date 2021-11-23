package awscdkbatchalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.AllocationStrategy",
		reflect.TypeOf((*AllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"BEST_FIT": AllocationStrategy_BEST_FIT,
			"BEST_FIT_PROGRESSIVE": AllocationStrategy_BEST_FIT_PROGRESSIVE,
			"SPOT_CAPACITY_OPTIMIZED": AllocationStrategy_SPOT_CAPACITY_OPTIMIZED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.ComputeEnvironment",
		reflect.TypeOf((*ComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IComputeEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ComputeEnvironmentProps",
		reflect.TypeOf((*ComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.ComputeResourceType",
		reflect.TypeOf((*ComputeResourceType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": ComputeResourceType_ON_DEMAND,
			"SPOT": ComputeResourceType_SPOT,
			"FARGATE": ComputeResourceType_FARGATE,
			"FARGATE_SPOT": ComputeResourceType_FARGATE_SPOT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.ComputeResources",
		reflect.TypeOf((*ComputeResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.ExposedSecret",
		reflect.TypeOf((*ExposedSecret)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "optionName", GoGetter: "OptionName"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
		},
		func() interface{} {
			return &jsiiProxy_ExposedSecret{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IComputeEnvironment",
		reflect.TypeOf((*IComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentArn", GoGetter: "ComputeEnvironmentArn"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IJobDefinition",
		reflect.TypeOf((*IJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IJobQueue",
		reflect.TypeOf((*IJobQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueArn", GoGetter: "JobQueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueName", GoGetter: "JobQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IJobQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.IMultiNodeProps",
		reflect.TypeOf((*IMultiNodeProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "mainNode", GoGetter: "MainNode"},
			_jsii_.MemberProperty{JsiiProperty: "rangeProps", GoGetter: "RangeProps"},
		},
		func() interface{} {
			return &jsiiProxy_IMultiNodeProps{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-batch-alpha.INodeRangeProps",
		reflect.TypeOf((*INodeRangeProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "fromNodeIndex", GoGetter: "FromNodeIndex"},
			_jsii_.MemberProperty{JsiiProperty: "toNodeIndex", GoGetter: "ToNodeIndex"},
		},
		func() interface{} {
			return &jsiiProxy_INodeRangeProps{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.JobDefinition",
		reflect.TypeOf((*JobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionArn", GoGetter: "JobDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobDefinitionContainer",
		reflect.TypeOf((*JobDefinitionContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobDefinitionProps",
		reflect.TypeOf((*JobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-batch-alpha.JobQueue",
		reflect.TypeOf((*JobQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueArn", GoGetter: "JobQueueArn"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueName", GoGetter: "JobQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_JobQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJobQueue)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobQueueComputeEnvironment",
		reflect.TypeOf((*JobQueueComputeEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.JobQueueProps",
		reflect.TypeOf((*JobQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.LaunchTemplateSpecification",
		reflect.TypeOf((*LaunchTemplateSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-batch-alpha.LogConfiguration",
		reflect.TypeOf((*LogConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.LogDriver",
		reflect.TypeOf((*LogDriver)(nil)).Elem(),
		map[string]interface{}{
			"AWSLOGS": LogDriver_AWSLOGS,
			"FLUENTD": LogDriver_FLUENTD,
			"GELF": LogDriver_GELF,
			"JOURNALD": LogDriver_JOURNALD,
			"LOGENTRIES": LogDriver_LOGENTRIES,
			"JSON_FILE": LogDriver_JSON_FILE,
			"SPLUNK": LogDriver_SPLUNK,
			"SYSLOG": LogDriver_SYSLOG,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-batch-alpha.PlatformCapabilities",
		reflect.TypeOf((*PlatformCapabilities)(nil)).Elem(),
		map[string]interface{}{
			"EC2": PlatformCapabilities_EC2,
			"FARGATE": PlatformCapabilities_FARGATE,
		},
	)
}
