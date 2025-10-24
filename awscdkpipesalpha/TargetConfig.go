package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspipes"
)

// Target config properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   targetConfig := &TargetConfig{
//   	TargetParameters: &PipeTargetParametersProperty{
//   		BatchJobParameters: &PipeTargetBatchJobParametersProperty{
//   			JobDefinition: jsii.String("jobDefinition"),
//   			JobName: jsii.String("jobName"),
//
//   			// the properties below are optional
//   			ArrayProperties: &BatchArrayPropertiesProperty{
//   				Size: jsii.Number(123),
//   			},
//   			ContainerOverrides: &BatchContainerOverridesProperty{
//   				Command: []*string{
//   					jsii.String("command"),
//   				},
//   				Environment: []interface{}{
//   					&BatchEnvironmentVariableProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				InstanceType: jsii.String("instanceType"),
//   				ResourceRequirements: []interface{}{
//   					&BatchResourceRequirementProperty{
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			DependsOn: []interface{}{
//   				&BatchJobDependencyProperty{
//   					JobId: jsii.String("jobId"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			RetryStrategy: &BatchRetryStrategyProperty{
//   				Attempts: jsii.Number(123),
//   			},
//   		},
//   		CloudWatchLogsParameters: &PipeTargetCloudWatchLogsParametersProperty{
//   			LogStreamName: jsii.String("logStreamName"),
//   			Timestamp: jsii.String("timestamp"),
//   		},
//   		EcsTaskParameters: &PipeTargetEcsTaskParametersProperty{
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					CapacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					Base: jsii.Number(123),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   				},
//   			},
//   			Overrides: &EcsTaskOverrideProperty{
//   				ContainerOverrides: []interface{}{
//   					&EcsContainerOverrideProperty{
//   						Command: []*string{
//   							jsii.String("command"),
//   						},
//   						Cpu: jsii.Number(123),
//   						Environment: []interface{}{
//   							&EcsEnvironmentVariableProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EnvironmentFiles: []interface{}{
//   							&EcsEnvironmentFileProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Memory: jsii.Number(123),
//   						MemoryReservation: jsii.Number(123),
//   						Name: jsii.String("name"),
//   						ResourceRequirements: []interface{}{
//   							&EcsResourceRequirementProperty{
//   								Type: jsii.String("type"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Cpu: jsii.String("cpu"),
//   				EphemeralStorage: &EcsEphemeralStorageProperty{
//   					SizeInGiB: jsii.Number(123),
//   				},
//   				ExecutionRoleArn: jsii.String("executionRoleArn"),
//   				InferenceAcceleratorOverrides: []interface{}{
//   					&EcsInferenceAcceleratorOverrideProperty{
//   						DeviceName: jsii.String("deviceName"),
//   						DeviceType: jsii.String("deviceType"),
//   					},
//   				},
//   				Memory: jsii.String("memory"),
//   				TaskRoleArn: jsii.String("taskRoleArn"),
//   			},
//   			PlacementConstraints: []interface{}{
//   				&PlacementConstraintProperty{
//   					Expression: jsii.String("expression"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlacementStrategy: []interface{}{
//   				&PlacementStrategyProperty{
//   					Field: jsii.String("field"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			PlatformVersion: jsii.String("platformVersion"),
//   			PropagateTags: jsii.String("propagateTags"),
//   			ReferenceId: jsii.String("referenceId"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			TaskCount: jsii.Number(123),
//   		},
//   		EventBridgeEventBusParameters: &PipeTargetEventBridgeEventBusParametersProperty{
//   			DetailType: jsii.String("detailType"),
//   			EndpointId: jsii.String("endpointId"),
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			Source: jsii.String("source"),
//   			Time: jsii.String("time"),
//   		},
//   		HttpParameters: &PipeTargetHttpParametersProperty{
//   			HeaderParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			PathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			QueryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		InputTemplate: jsii.String("inputTemplate"),
//   		KinesisStreamParameters: &PipeTargetKinesisStreamParametersProperty{
//   			PartitionKey: jsii.String("partitionKey"),
//   		},
//   		LambdaFunctionParameters: &PipeTargetLambdaFunctionParametersProperty{
//   			InvocationType: jsii.String("invocationType"),
//   		},
//   		RedshiftDataParameters: &PipeTargetRedshiftDataParametersProperty{
//   			Database: jsii.String("database"),
//   			Sqls: []*string{
//   				jsii.String("sqls"),
//   			},
//
//   			// the properties below are optional
//   			DbUser: jsii.String("dbUser"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   			StatementName: jsii.String("statementName"),
//   			WithEvent: jsii.Boolean(false),
//   		},
//   		SageMakerPipelineParameters: &PipeTargetSageMakerPipelineParametersProperty{
//   			PipelineParameterList: []interface{}{
//   				&SageMakerPipelineParameterProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		SqsQueueParameters: &PipeTargetSqsQueueParametersProperty{
//   			MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   			MessageGroupId: jsii.String("messageGroupId"),
//   		},
//   		StepFunctionStateMachineParameters: &PipeTargetStateMachineParametersProperty{
//   			InvocationType: jsii.String("invocationType"),
//   		},
//   		TimestreamParameters: &PipeTargetTimestreamParametersProperty{
//   			DimensionMappings: []interface{}{
//   				&DimensionMappingProperty{
//   					DimensionName: jsii.String("dimensionName"),
//   					DimensionValue: jsii.String("dimensionValue"),
//   					DimensionValueType: jsii.String("dimensionValueType"),
//   				},
//   			},
//   			TimeValue: jsii.String("timeValue"),
//   			VersionValue: jsii.String("versionValue"),
//
//   			// the properties below are optional
//   			EpochTimeUnit: jsii.String("epochTimeUnit"),
//   			MultiMeasureMappings: []interface{}{
//   				&MultiMeasureMappingProperty{
//   					MultiMeasureAttributeMappings: []interface{}{
//   						&MultiMeasureAttributeMappingProperty{
//   							MeasureValue: jsii.String("measureValue"),
//   							MeasureValueType: jsii.String("measureValueType"),
//   							MultiMeasureAttributeName: jsii.String("multiMeasureAttributeName"),
//   						},
//   					},
//   					MultiMeasureName: jsii.String("multiMeasureName"),
//   				},
//   			},
//   			SingleMeasureMappings: []interface{}{
//   				&SingleMeasureMappingProperty{
//   					MeasureName: jsii.String("measureName"),
//   					MeasureValue: jsii.String("measureValue"),
//   					MeasureValueType: jsii.String("measureValueType"),
//   				},
//   			},
//   			TimeFieldType: jsii.String("timeFieldType"),
//   			TimestampFormat: jsii.String("timestampFormat"),
//   		},
//   	},
//   }
//
// Experimental.
type TargetConfig struct {
	// The parameters required to set up a target for your pipe.
	// Experimental.
	TargetParameters *awspipes.CfnPipe_PipeTargetParametersProperty `field:"required" json:"targetParameters" yaml:"targetParameters"`
}

