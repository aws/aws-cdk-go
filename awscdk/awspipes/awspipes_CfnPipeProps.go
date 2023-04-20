package awspipes


// Properties for defining a `CfnPipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipeProps := &CfnPipeProps{
//   	RoleArn: jsii.String("roleArn"),
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DesiredState: jsii.String("desiredState"),
//   	Enrichment: jsii.String("enrichment"),
//   	EnrichmentParameters: &PipeEnrichmentParametersProperty{
//   		HttpParameters: &PipeEnrichmentHttpParametersProperty{
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
//   	},
//   	Name: jsii.String("name"),
//   	SourceParameters: &PipeSourceParametersProperty{
//   		ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			QueueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   		DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   			StartingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   		},
//   		FilterCriteria: &FilterCriteriaProperty{
//   			Filters: []interface{}{
//   				&FilterProperty{
//   					Pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   		KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   			StartingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   			TopicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &MSKAccessCredentialsProperty{
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   		},
//   		RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			QueueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			VirtualHost: jsii.String("virtualHost"),
//   		},
//   		SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
//   			TopicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			AdditionalBootstrapServers: []*string{
//   				jsii.String("additionalBootstrapServers"),
//   			},
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &SelfManagedKafkaAccessConfigurationCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram256Auth: jsii.String("saslScram256Auth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			ServerRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   			StartingPosition: jsii.String("startingPosition"),
//   			Vpc: &SelfManagedKafkaAccessConfigurationVpcProperty{
//   				SecurityGroup: []*string{
//   					jsii.String("securityGroup"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		SqsQueueParameters: &PipeSourceSqsQueueParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
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
//   			Tags: []cfnTag{
//   				&cfnTag{
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
//   	},
//   }
//
type CfnPipeProps struct {
	// The ARN of the role that allows the pipe to send data to the target.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the source resource.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The ARN of the target resource.
	Target *string `field:"required" json:"target" yaml:"target"`
	// A description of the pipe.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state the pipe should be in.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The ARN of the enrichment resource.
	Enrichment *string `field:"optional" json:"enrichment" yaml:"enrichment"`
	// The parameters required to set up enrichment on your pipe.
	EnrichmentParameters interface{} `field:"optional" json:"enrichmentParameters" yaml:"enrichmentParameters"`
	// The name of the pipe.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The parameters required to set up a source for your pipe.
	SourceParameters interface{} `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
	// The list of key-value pairs to associate with the pipe.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The parameters required to set up a target for your pipe.
	//
	// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
	TargetParameters interface{} `field:"optional" json:"targetParameters" yaml:"targetParameters"`
}

