package awspipes


// The parameters required to set up a target for your pipe.
//
// For more information about pipe target parameters, including how to use dynamic path parameters, see [Target parameters](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetParametersProperty := &pipeTargetParametersProperty{
//   	batchJobParameters: &pipeTargetBatchJobParametersProperty{
//   		jobDefinition: jsii.String("jobDefinition"),
//   		jobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		arrayProperties: &batchArrayPropertiesProperty{
//   			size: jsii.Number(123),
//   		},
//   		containerOverrides: &batchContainerOverridesProperty{
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			environment: []interface{}{
//   				&batchEnvironmentVariableProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			instanceType: jsii.String("instanceType"),
//   			resourceRequirements: []interface{}{
//   				&batchResourceRequirementProperty{
//   					type: jsii.String("type"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		dependsOn: []interface{}{
//   			&batchJobDependencyProperty{
//   				jobId: jsii.String("jobId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		retryStrategy: &batchRetryStrategyProperty{
//   			attempts: jsii.Number(123),
//   		},
//   	},
//   	cloudWatchLogsParameters: &pipeTargetCloudWatchLogsParametersProperty{
//   		logStreamName: jsii.String("logStreamName"),
//   		timestamp: jsii.String("timestamp"),
//   	},
//   	ecsTaskParameters: &pipeTargetEcsTaskParametersProperty{
//   		taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		capacityProviderStrategy: []interface{}{
//   			&capacityProviderStrategyItemProperty{
//   				capacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				base: jsii.Number(123),
//   				weight: jsii.Number(123),
//   			},
//   		},
//   		enableEcsManagedTags: jsii.Boolean(false),
//   		enableExecuteCommand: jsii.Boolean(false),
//   		group: jsii.String("group"),
//   		launchType: jsii.String("launchType"),
//   		networkConfiguration: &networkConfigurationProperty{
//   			awsvpcConfiguration: &awsVpcConfigurationProperty{
//   				subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				assignPublicIp: jsii.String("assignPublicIp"),
//   				securityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   			},
//   		},
//   		overrides: &ecsTaskOverrideProperty{
//   			containerOverrides: []interface{}{
//   				&ecsContainerOverrideProperty{
//   					command: []*string{
//   						jsii.String("command"),
//   					},
//   					cpu: jsii.Number(123),
//   					environment: []interface{}{
//   						&ecsEnvironmentVariableProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					environmentFiles: []interface{}{
//   						&ecsEnvironmentFileProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					memory: jsii.Number(123),
//   					memoryReservation: jsii.Number(123),
//   					name: jsii.String("name"),
//   					resourceRequirements: []interface{}{
//   						&ecsResourceRequirementProperty{
//   							type: jsii.String("type"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   			cpu: jsii.String("cpu"),
//   			ephemeralStorage: &ecsEphemeralStorageProperty{
//   				sizeInGiB: jsii.Number(123),
//   			},
//   			executionRoleArn: jsii.String("executionRoleArn"),
//   			inferenceAcceleratorOverrides: []interface{}{
//   				&ecsInferenceAcceleratorOverrideProperty{
//   					deviceName: jsii.String("deviceName"),
//   					deviceType: jsii.String("deviceType"),
//   				},
//   			},
//   			memory: jsii.String("memory"),
//   			taskRoleArn: jsii.String("taskRoleArn"),
//   		},
//   		placementConstraints: []interface{}{
//   			&placementConstraintProperty{
//   				expression: jsii.String("expression"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		placementStrategy: []interface{}{
//   			&placementStrategyProperty{
//   				field: jsii.String("field"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		platformVersion: jsii.String("platformVersion"),
//   		propagateTags: jsii.String("propagateTags"),
//   		referenceId: jsii.String("referenceId"),
//   		tags: []cfnTag{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		taskCount: jsii.Number(123),
//   	},
//   	eventBridgeEventBusParameters: &pipeTargetEventBridgeEventBusParametersProperty{
//   		detailType: jsii.String("detailType"),
//   		endpointId: jsii.String("endpointId"),
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: jsii.String("source"),
//   		time: jsii.String("time"),
//   	},
//   	httpParameters: &pipeTargetHttpParametersProperty{
//   		headerParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		pathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		queryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	inputTemplate: jsii.String("inputTemplate"),
//   	kinesisStreamParameters: &pipeTargetKinesisStreamParametersProperty{
//   		partitionKey: jsii.String("partitionKey"),
//   	},
//   	lambdaFunctionParameters: &pipeTargetLambdaFunctionParametersProperty{
//   		invocationType: jsii.String("invocationType"),
//   	},
//   	redshiftDataParameters: &pipeTargetRedshiftDataParametersProperty{
//   		database: jsii.String("database"),
//   		sqls: []*string{
//   			jsii.String("sqls"),
//   		},
//
//   		// the properties below are optional
//   		dbUser: jsii.String("dbUser"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		statementName: jsii.String("statementName"),
//   		withEvent: jsii.Boolean(false),
//   	},
//   	sageMakerPipelineParameters: &pipeTargetSageMakerPipelineParametersProperty{
//   		pipelineParameterList: []interface{}{
//   			&sageMakerPipelineParameterProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	sqsQueueParameters: &pipeTargetSqsQueueParametersProperty{
//   		messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   		messageGroupId: jsii.String("messageGroupId"),
//   	},
//   	stepFunctionStateMachineParameters: &pipeTargetStateMachineParametersProperty{
//   		invocationType: jsii.String("invocationType"),
//   	},
//   }
//
type CfnPipe_PipeTargetParametersProperty struct {
	// The parameters for using an AWS Batch job as a target.
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// The parameters for using an CloudWatch Logs log stream as a target.
	CloudWatchLogsParameters interface{} `field:"optional" json:"cloudWatchLogsParameters" yaml:"cloudWatchLogsParameters"`
	// The parameters for using an Amazon ECS task as a target.
	EcsTaskParameters interface{} `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// The parameters for using an EventBridge event bus as a target.
	EventBridgeEventBusParameters interface{} `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
	// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	//
	// To remove an input template, specify an empty string.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// The parameters for using a Kinesis stream as a source.
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// The parameters for using a Lambda function as a target.
	LambdaFunctionParameters interface{} `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API BatchExecuteStatement.
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The parameters for using a SageMaker pipeline as a target.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// The parameters for using a Amazon SQS stream as a source.
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// The parameters for using a Step Functions state machine as a target.
	StepFunctionStateMachineParameters interface{} `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

