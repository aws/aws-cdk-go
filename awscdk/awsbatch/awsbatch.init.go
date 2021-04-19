package awsbatch

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_batch.AllocationStrategy",
		reflect.TypeOf((*AllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"BEST_FIT": AllocationStrategy_BEST_FIT,
			"BEST_FIT_PROGRESSIVE": AllocationStrategy_BEST_FIT_PROGRESSIVE,
			"SPOT_CAPACITY_OPTIMIZED": AllocationStrategy_SPOT_CAPACITY_OPTIMIZED,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment",
		reflect.TypeOf((*CfnComputeEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentName", GoGetter: "ComputeEnvironmentName"},
			_jsii_.MemberProperty{JsiiProperty: "computeResources", GoGetter: "ComputeResources"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnComputeEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment.ComputeResourcesProperty",
		reflect.TypeOf((*CfnComputeEnvironment_ComputeResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment.Ec2ConfigurationObjectProperty",
		reflect.TypeOf((*CfnComputeEnvironment_Ec2ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironment.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnComputeEnvironment_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnComputeEnvironmentProps",
		reflect.TypeOf((*CfnComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.CfnJobDefinition",
		reflect.TypeOf((*CfnJobDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "containerProperties", GoGetter: "ContainerProperties"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "jobDefinitionName", GoGetter: "JobDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeProperties", GoGetter: "NodeProperties"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "platformCapabilities", GoGetter: "PlatformCapabilities"},
			_jsii_.MemberProperty{JsiiProperty: "propagateTags", GoGetter: "PropagateTags"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "retryStrategy", GoGetter: "RetryStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.ContainerPropertiesProperty",
		reflect.TypeOf((*CfnJobDefinition_ContainerPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.DeviceProperty",
		reflect.TypeOf((*CfnJobDefinition_DeviceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.EnvironmentProperty",
		reflect.TypeOf((*CfnJobDefinition_EnvironmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.EvaluateOnExitProperty",
		reflect.TypeOf((*CfnJobDefinition_EvaluateOnExitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.FargatePlatformConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinition_FargatePlatformConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.LinuxParametersProperty",
		reflect.TypeOf((*CfnJobDefinition_LinuxParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.LogConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinition_LogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.MountPointsProperty",
		reflect.TypeOf((*CfnJobDefinition_MountPointsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnJobDefinition_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.NodePropertiesProperty",
		reflect.TypeOf((*CfnJobDefinition_NodePropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.NodeRangePropertyProperty",
		reflect.TypeOf((*CfnJobDefinition_NodeRangePropertyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.ResourceRequirementProperty",
		reflect.TypeOf((*CfnJobDefinition_ResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.RetryStrategyProperty",
		reflect.TypeOf((*CfnJobDefinition_RetryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.SecretProperty",
		reflect.TypeOf((*CfnJobDefinition_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.TimeoutProperty",
		reflect.TypeOf((*CfnJobDefinition_TimeoutProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.TmpfsProperty",
		reflect.TypeOf((*CfnJobDefinition_TmpfsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.UlimitProperty",
		reflect.TypeOf((*CfnJobDefinition_UlimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.VolumesHostProperty",
		reflect.TypeOf((*CfnJobDefinition_VolumesHostProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinition.VolumesProperty",
		reflect.TypeOf((*CfnJobDefinition_VolumesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobDefinitionProps",
		reflect.TypeOf((*CfnJobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.CfnJobQueue",
		reflect.TypeOf((*CfnJobQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "computeEnvironmentOrder", GoGetter: "ComputeEnvironmentOrder"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueueName", GoGetter: "JobQueueName"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobQueue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobQueue.ComputeEnvironmentOrderProperty",
		reflect.TypeOf((*CfnJobQueue_ComputeEnvironmentOrderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.CfnJobQueueProps",
		reflect.TypeOf((*CfnJobQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.ComputeEnvironment",
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
		"aws-cdk-lib.aws_batch.ComputeEnvironmentProps",
		reflect.TypeOf((*ComputeEnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_batch.ComputeResourceType",
		reflect.TypeOf((*ComputeResourceType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": ComputeResourceType_ON_DEMAND,
			"SPOT": ComputeResourceType_SPOT,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.ComputeResources",
		reflect.TypeOf((*ComputeResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.ExposedSecret",
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
		"aws-cdk-lib.aws_batch.IComputeEnvironment",
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
		"aws-cdk-lib.aws_batch.IJobDefinition",
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
		"aws-cdk-lib.aws_batch.IJobQueue",
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
		"aws-cdk-lib.aws_batch.IMultiNodeProps",
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
		"aws-cdk-lib.aws_batch.INodeRangeProps",
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
		"aws-cdk-lib.aws_batch.JobDefinition",
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
		"aws-cdk-lib.aws_batch.JobDefinitionContainer",
		reflect.TypeOf((*JobDefinitionContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.JobDefinitionProps",
		reflect.TypeOf((*JobDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_batch.JobQueue",
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
		"aws-cdk-lib.aws_batch.JobQueueComputeEnvironment",
		reflect.TypeOf((*JobQueueComputeEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.JobQueueProps",
		reflect.TypeOf((*JobQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.LaunchTemplateSpecification",
		reflect.TypeOf((*LaunchTemplateSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_batch.LogConfiguration",
		reflect.TypeOf((*LogConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_batch.LogDriver",
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
}
