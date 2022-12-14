package awspipes


// Properties for defining a `CfnPipe`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipeProps := &cfnPipeProps{
//   	roleArn: jsii.String("roleArn"),
//   	source: jsii.String("source"),
//   	target: jsii.String("target"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	desiredState: jsii.String("desiredState"),
//   	enrichment: jsii.String("enrichment"),
//   	enrichmentParameters: &pipeEnrichmentParametersProperty{
//   		httpParameters: &pipeEnrichmentHttpParametersProperty{
//   			headerParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			pathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			queryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		inputTemplate: jsii.String("inputTemplate"),
//   	},
//   	name: jsii.String("name"),
//   	sourceParameters: &pipeSourceParametersProperty{
//   		activeMqBrokerParameters: &pipeSourceActiveMQBrokerParametersProperty{
//   			credentials: &mQBrokerAccessCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   			},
//   			queueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   		dynamoDbStreamParameters: &pipeSourceDynamoDBStreamParametersProperty{
//   			startingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			deadLetterConfig: &deadLetterConfigProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			maximumRecordAgeInSeconds: jsii.Number(123),
//   			maximumRetryAttempts: jsii.Number(123),
//   			onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			parallelizationFactor: jsii.Number(123),
//   		},
//   		filterCriteria: &filterCriteriaProperty{
//   			filters: []interface{}{
//   				&filterProperty{
//   					pattern: jsii.String("pattern"),
//   				},
//   			},
//   		},
//   		kinesisStreamParameters: &pipeSourceKinesisStreamParametersProperty{
//   			startingPosition: jsii.String("startingPosition"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			deadLetterConfig: &deadLetterConfigProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			maximumRecordAgeInSeconds: jsii.Number(123),
//   			maximumRetryAttempts: jsii.Number(123),
//   			onPartialBatchItemFailure: jsii.String("onPartialBatchItemFailure"),
//   			parallelizationFactor: jsii.Number(123),
//   			startingPositionTimestamp: jsii.String("startingPositionTimestamp"),
//   		},
//   		managedStreamingKafkaParameters: &pipeSourceManagedStreamingKafkaParametersProperty{
//   			topicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			consumerGroupId: jsii.String("consumerGroupId"),
//   			credentials: &mSKAccessCredentialsProperty{
//   				clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				saslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			startingPosition: jsii.String("startingPosition"),
//   		},
//   		rabbitMqBrokerParameters: &pipeSourceRabbitMQBrokerParametersProperty{
//   			credentials: &mQBrokerAccessCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   			},
//   			queueName: jsii.String("queueName"),
//
//   			// the properties below are optional
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			virtualHost: jsii.String("virtualHost"),
//   		},
//   		selfManagedKafkaParameters: &pipeSourceSelfManagedKafkaParametersProperty{
//   			topicName: jsii.String("topicName"),
//
//   			// the properties below are optional
//   			additionalBootstrapServers: []*string{
//   				jsii.String("additionalBootstrapServers"),
//   			},
//   			batchSize: jsii.Number(123),
//   			consumerGroupId: jsii.String("consumerGroupId"),
//   			credentials: &selfManagedKafkaAccessConfigurationCredentialsProperty{
//   				basicAuth: jsii.String("basicAuth"),
//   				clientCertificateTlsAuth: jsii.String("clientCertificateTlsAuth"),
//   				saslScram256Auth: jsii.String("saslScram256Auth"),
//   				saslScram512Auth: jsii.String("saslScram512Auth"),
//   			},
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   			serverRootCaCertificate: jsii.String("serverRootCaCertificate"),
//   			startingPosition: jsii.String("startingPosition"),
//   			vpc: &selfManagedKafkaAccessConfigurationVpcProperty{
//   				securityGroup: []*string{
//   					jsii.String("securityGroup"),
//   				},
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   		},
//   		sqsQueueParameters: &pipeSourceSqsQueueParametersProperty{
//   			batchSize: jsii.Number(123),
//   			maximumBatchingWindowInSeconds: jsii.Number(123),
//   		},
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	targetParameters: &pipeTargetParametersProperty{
//   		batchJobParameters: &pipeTargetBatchJobParametersProperty{
//   			jobDefinition: jsii.String("jobDefinition"),
//   			jobName: jsii.String("jobName"),
//
//   			// the properties below are optional
//   			arrayProperties: &batchArrayPropertiesProperty{
//   				size: jsii.Number(123),
//   			},
//   			containerOverrides: &batchContainerOverridesProperty{
//   				command: []*string{
//   					jsii.String("command"),
//   				},
//   				environment: []interface{}{
//   					&batchEnvironmentVariableProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				instanceType: jsii.String("instanceType"),
//   				resourceRequirements: []interface{}{
//   					&batchResourceRequirementProperty{
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			dependsOn: []interface{}{
//   				&batchJobDependencyProperty{
//   					jobId: jsii.String("jobId"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			retryStrategy: &batchRetryStrategyProperty{
//   				attempts: jsii.Number(123),
//   			},
//   		},
//   		cloudWatchLogsParameters: &pipeTargetCloudWatchLogsParametersProperty{
//   			logStreamName: jsii.String("logStreamName"),
//   			timestamp: jsii.String("timestamp"),
//   		},
//   		ecsTaskParameters: &pipeTargetEcsTaskParametersProperty{
//   			taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   			// the properties below are optional
//   			capacityProviderStrategy: []interface{}{
//   				&capacityProviderStrategyItemProperty{
//   					capacityProvider: jsii.String("capacityProvider"),
//
//   					// the properties below are optional
//   					base: jsii.Number(123),
//   					weight: jsii.Number(123),
//   				},
//   			},
//   			enableEcsManagedTags: jsii.Boolean(false),
//   			enableExecuteCommand: jsii.Boolean(false),
//   			group: jsii.String("group"),
//   			launchType: jsii.String("launchType"),
//   			networkConfiguration: &networkConfigurationProperty{
//   				awsvpcConfiguration: &awsVpcConfigurationProperty{
//   					subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//
//   					// the properties below are optional
//   					assignPublicIp: jsii.String("assignPublicIp"),
//   					securityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   				},
//   			},
//   			overrides: &ecsTaskOverrideProperty{
//   				containerOverrides: []interface{}{
//   					&ecsContainerOverrideProperty{
//   						command: []*string{
//   							jsii.String("command"),
//   						},
//   						cpu: jsii.Number(123),
//   						environment: []interface{}{
//   							&ecsEnvironmentVariableProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						environmentFiles: []interface{}{
//   							&ecsEnvironmentFileProperty{
//   								type: jsii.String("type"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						memory: jsii.Number(123),
//   						memoryReservation: jsii.Number(123),
//   						name: jsii.String("name"),
//   						resourceRequirements: []interface{}{
//   							&ecsResourceRequirementProperty{
//   								type: jsii.String("type"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				cpu: jsii.String("cpu"),
//   				ephemeralStorage: &ecsEphemeralStorageProperty{
//   					sizeInGiB: jsii.Number(123),
//   				},
//   				executionRoleArn: jsii.String("executionRoleArn"),
//   				inferenceAcceleratorOverrides: []interface{}{
//   					&ecsInferenceAcceleratorOverrideProperty{
//   						deviceName: jsii.String("deviceName"),
//   						deviceType: jsii.String("deviceType"),
//   					},
//   				},
//   				memory: jsii.String("memory"),
//   				taskRoleArn: jsii.String("taskRoleArn"),
//   			},
//   			placementConstraints: []interface{}{
//   				&placementConstraintProperty{
//   					expression: jsii.String("expression"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			placementStrategy: []interface{}{
//   				&placementStrategyProperty{
//   					field: jsii.String("field"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			platformVersion: jsii.String("platformVersion"),
//   			propagateTags: jsii.String("propagateTags"),
//   			referenceId: jsii.String("referenceId"),
//   			tags: []cfnTag{
//   				&cfnTag{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			taskCount: jsii.Number(123),
//   		},
//   		eventBridgeEventBusParameters: &pipeTargetEventBridgeEventBusParametersProperty{
//   			detailType: jsii.String("detailType"),
//   			endpointId: jsii.String("endpointId"),
//   			resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			source: jsii.String("source"),
//   			time: jsii.String("time"),
//   		},
//   		httpParameters: &pipeTargetHttpParametersProperty{
//   			headerParameters: map[string]*string{
//   				"headerParametersKey": jsii.String("headerParameters"),
//   			},
//   			pathParameterValues: []*string{
//   				jsii.String("pathParameterValues"),
//   			},
//   			queryStringParameters: map[string]*string{
//   				"queryStringParametersKey": jsii.String("queryStringParameters"),
//   			},
//   		},
//   		inputTemplate: jsii.String("inputTemplate"),
//   		kinesisStreamParameters: &pipeTargetKinesisStreamParametersProperty{
//   			partitionKey: jsii.String("partitionKey"),
//   		},
//   		lambdaFunctionParameters: &pipeTargetLambdaFunctionParametersProperty{
//   			invocationType: jsii.String("invocationType"),
//   		},
//   		redshiftDataParameters: &pipeTargetRedshiftDataParametersProperty{
//   			database: jsii.String("database"),
//   			sqls: []*string{
//   				jsii.String("sqls"),
//   			},
//
//   			// the properties below are optional
//   			dbUser: jsii.String("dbUser"),
//   			secretManagerArn: jsii.String("secretManagerArn"),
//   			statementName: jsii.String("statementName"),
//   			withEvent: jsii.Boolean(false),
//   		},
//   		sageMakerPipelineParameters: &pipeTargetSageMakerPipelineParametersProperty{
//   			pipelineParameterList: []interface{}{
//   				&sageMakerPipelineParameterProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		sqsQueueParameters: &pipeTargetSqsQueueParametersProperty{
//   			messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   			messageGroupId: jsii.String("messageGroupId"),
//   		},
//   		stepFunctionStateMachineParameters: &pipeTargetStateMachineParametersProperty{
//   			invocationType: jsii.String("invocationType"),
//   		},
//   	},
//   }
//
type CfnPipeProps struct {
	// `AWS::Pipes::Pipe.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `AWS::Pipes::Pipe.Source`.
	Source *string `field:"required" json:"source" yaml:"source"`
	// `AWS::Pipes::Pipe.Target`.
	Target *string `field:"required" json:"target" yaml:"target"`
	// `AWS::Pipes::Pipe.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Pipes::Pipe.DesiredState`.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// `AWS::Pipes::Pipe.Enrichment`.
	Enrichment *string `field:"optional" json:"enrichment" yaml:"enrichment"`
	// `AWS::Pipes::Pipe.EnrichmentParameters`.
	EnrichmentParameters interface{} `field:"optional" json:"enrichmentParameters" yaml:"enrichmentParameters"`
	// `AWS::Pipes::Pipe.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Pipes::Pipe.SourceParameters`.
	SourceParameters interface{} `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
	// `AWS::Pipes::Pipe.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Pipes::Pipe.TargetParameters`.
	TargetParameters interface{} `field:"optional" json:"targetParameters" yaml:"targetParameters"`
}

