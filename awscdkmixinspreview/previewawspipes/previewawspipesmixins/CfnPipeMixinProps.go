package previewawspipesmixins


// Properties for CfnPipePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPipeMixinProps := &CfnPipeMixinProps{
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
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	LogConfiguration: &PipeLogConfigurationProperty{
//   		CloudwatchLogsLogDestination: &CloudwatchLogsLogDestinationProperty{
//   			LogGroupArn: jsii.String("logGroupArn"),
//   		},
//   		FirehoseLogDestination: &FirehoseLogDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		},
//   		IncludeExecutionData: []*string{
//   			jsii.String("includeExecutionData"),
//   		},
//   		Level: jsii.String("level"),
//   		S3LogDestination: &S3LogDestinationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			OutputFormat: jsii.String("outputFormat"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Source: jsii.String("source"),
//   	SourceParameters: &PipeSourceParametersProperty{
//   		ActiveMqBrokerParameters: &PipeSourceActiveMQBrokerParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			QueueName: jsii.String("queueName"),
//   		},
//   		DynamoDbStreamParameters: &PipeSourceDynamoDBStreamParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   		},
//   		FilterCriteria: &FilterCriteriaProperty{
//   			Filters: []interface{}{
//   				&FilterProperty{
//   					Pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   		KinesisStreamParameters: &PipeSourceKinesisStreamParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			MaximumRecordAgeInSeconds: jsii.Number(123),
//   			MaximumRetryAttempts: jsii.Number(123),
//   			OnPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			ParallelizationFactor: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   			StartingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		ManagedStreamingKafkaParameters: &PipeSourceManagedStreamingKafkaParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			ConsumerGroupId: jsii.String("consumerGroupId"),
//   			Credentials: &MSKAccessCredentialsProperty{
//   				ClientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				SaslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			StartingPosition: jsii.String("startingPosition"),
//   			TopicName: jsii.String("topicName"),
//   		},
//   		RabbitMqBrokerParameters: &PipeSourceRabbitMQBrokerParametersProperty{
//   			BatchSize: jsii.Number(123),
//   			Credentials: &MQBrokerAccessCredentialsProperty{
//   				BasicAuth: jsii.String("basicAuth"),
//   			},
//   			MaximumBatchingWindowInSeconds: jsii.Number(123),
//   			QueueName: jsii.String("queueName"),
//   			VirtualHost: jsii.String("virtualHost"),
//   		},
//   		SelfManagedKafkaParameters: &PipeSourceSelfManagedKafkaParametersProperty{
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
//   			TopicName: jsii.String("topicName"),
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
//   	Target: jsii.String("target"),
//   	TargetParameters: &PipeTargetParametersProperty{
//   		BatchJobParameters: &PipeTargetBatchJobParametersProperty{
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
//   			JobDefinition: jsii.String("jobDefinition"),
//   			JobName: jsii.String("jobName"),
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
//   			CapacityProviderStrategy: []interface{}{
//   				&CapacityProviderStrategyItemProperty{
//   					Base: jsii.Number(123),
//   					CapacityProvider: jsii.String("capacityProvider"),
//   					Weight: jsii.Number(123),
//   				},
//   			},
//   			EnableEcsManagedTags: jsii.Boolean(false),
//   			EnableExecuteCommand: jsii.Boolean(false),
//   			Group: jsii.String("group"),
//   			LaunchType: jsii.String("launchType"),
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   					AssignPublicIp: jsii.String("assignPublicIp"),
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   					Subnets: []*string{
//   						jsii.String("subnets"),
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
//   			TaskDefinitionArn: jsii.String("taskDefinitionArn"),
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
//   			DbUser: jsii.String("dbUser"),
//   			SecretManagerArn: jsii.String("secretManagerArn"),
//   			Sqls: []*string{
//   				jsii.String("sqls"),
//   			},
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
//   			TimeValue: jsii.String("timeValue"),
//   			VersionValue: jsii.String("versionValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html
//
type CfnPipeMixinProps struct {
	// A description of the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The state the pipe should be in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-desiredstate
	//
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The ARN of the enrichment resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-enrichment
	//
	Enrichment *string `field:"optional" json:"enrichment" yaml:"enrichment"`
	// The parameters required to set up enrichment on your pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-enrichmentparameters
	//
	EnrichmentParameters interface{} `field:"optional" json:"enrichmentParameters" yaml:"enrichmentParameters"`
	// The identifier of the AWS  customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.
	//
	// To update a pipe that is using the default AWS owned key to use a customer managed key instead, or update a pipe that is using a customer managed key to use a different customer managed key, specify a customer managed key identifier.
	//
	// To update a pipe that is using a customer managed key to use the default AWS owned key , specify an empty string.
	//
	// For more information, see [Managing keys](https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html) in the *AWS Key Management Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The logging configuration settings for the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// The name of the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the role that allows the pipe to send data to the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the source resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The parameters required to set up a source for your pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-sourceparameters
	//
	SourceParameters interface{} `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
	// The list of key-value pairs to associate with the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the target resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The parameters required to set up a target for your pipe.
	//
	// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-targetparameters
	//
	TargetParameters interface{} `field:"optional" json:"targetParameters" yaml:"targetParameters"`
}

