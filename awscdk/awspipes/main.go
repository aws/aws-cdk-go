package awspipes

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_pipes.CfnPipe",
		reflect.TypeOf((*CfnPipe)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreationTime", GoGetter: "AttrCreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrCurrentState", GoGetter: "AttrCurrentState"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastModifiedTime", GoGetter: "AttrLastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateReason", GoGetter: "AttrStateReason"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "desiredState", GoGetter: "DesiredState"},
			_jsii_.MemberProperty{JsiiProperty: "enrichment", GoGetter: "Enrichment"},
			_jsii_.MemberProperty{JsiiProperty: "enrichmentParameters", GoGetter: "EnrichmentParameters"},
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
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceParameters", GoGetter: "SourceParameters"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetParameters", GoGetter: "TargetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipe{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.AwsVpcConfigurationProperty",
		reflect.TypeOf((*CfnPipe_AwsVpcConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchArrayPropertiesProperty",
		reflect.TypeOf((*CfnPipe_BatchArrayPropertiesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchContainerOverridesProperty",
		reflect.TypeOf((*CfnPipe_BatchContainerOverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchEnvironmentVariableProperty",
		reflect.TypeOf((*CfnPipe_BatchEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchJobDependencyProperty",
		reflect.TypeOf((*CfnPipe_BatchJobDependencyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchResourceRequirementProperty",
		reflect.TypeOf((*CfnPipe_BatchResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.BatchRetryStrategyProperty",
		reflect.TypeOf((*CfnPipe_BatchRetryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.CapacityProviderStrategyItemProperty",
		reflect.TypeOf((*CfnPipe_CapacityProviderStrategyItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.DeadLetterConfigProperty",
		reflect.TypeOf((*CfnPipe_DeadLetterConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsContainerOverrideProperty",
		reflect.TypeOf((*CfnPipe_EcsContainerOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsEnvironmentFileProperty",
		reflect.TypeOf((*CfnPipe_EcsEnvironmentFileProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsEnvironmentVariableProperty",
		reflect.TypeOf((*CfnPipe_EcsEnvironmentVariableProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsEphemeralStorageProperty",
		reflect.TypeOf((*CfnPipe_EcsEphemeralStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsInferenceAcceleratorOverrideProperty",
		reflect.TypeOf((*CfnPipe_EcsInferenceAcceleratorOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsResourceRequirementProperty",
		reflect.TypeOf((*CfnPipe_EcsResourceRequirementProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.EcsTaskOverrideProperty",
		reflect.TypeOf((*CfnPipe_EcsTaskOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.FilterCriteriaProperty",
		reflect.TypeOf((*CfnPipe_FilterCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.FilterProperty",
		reflect.TypeOf((*CfnPipe_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.MQBrokerAccessCredentialsProperty",
		reflect.TypeOf((*CfnPipe_MQBrokerAccessCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.MSKAccessCredentialsProperty",
		reflect.TypeOf((*CfnPipe_MSKAccessCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnPipe_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeEnrichmentHttpParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeEnrichmentHttpParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeEnrichmentParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeEnrichmentParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceActiveMQBrokerParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceActiveMQBrokerParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceDynamoDBStreamParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceDynamoDBStreamParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceKinesisStreamParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceKinesisStreamParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceManagedStreamingKafkaParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceManagedStreamingKafkaParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceRabbitMQBrokerParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceRabbitMQBrokerParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceSelfManagedKafkaParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceSelfManagedKafkaParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeSourceSqsQueueParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeSourceSqsQueueParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetBatchJobParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetBatchJobParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetCloudWatchLogsParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetCloudWatchLogsParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetEcsTaskParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetEcsTaskParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetEventBridgeEventBusParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetEventBridgeEventBusParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetHttpParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetHttpParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetKinesisStreamParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetKinesisStreamParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetLambdaFunctionParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetLambdaFunctionParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetRedshiftDataParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetRedshiftDataParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetSageMakerPipelineParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetSageMakerPipelineParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetSqsQueueParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetSqsQueueParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PipeTargetStateMachineParametersProperty",
		reflect.TypeOf((*CfnPipe_PipeTargetStateMachineParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PlacementConstraintProperty",
		reflect.TypeOf((*CfnPipe_PlacementConstraintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.PlacementStrategyProperty",
		reflect.TypeOf((*CfnPipe_PlacementStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.SageMakerPipelineParameterProperty",
		reflect.TypeOf((*CfnPipe_SageMakerPipelineParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.SelfManagedKafkaAccessConfigurationCredentialsProperty",
		reflect.TypeOf((*CfnPipe_SelfManagedKafkaAccessConfigurationCredentialsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipe.SelfManagedKafkaAccessConfigurationVpcProperty",
		reflect.TypeOf((*CfnPipe_SelfManagedKafkaAccessConfigurationVpcProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_pipes.CfnPipeProps",
		reflect.TypeOf((*CfnPipeProps)(nil)).Elem(),
	)
}
