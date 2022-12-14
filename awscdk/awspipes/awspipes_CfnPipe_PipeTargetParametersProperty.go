package awspipes


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
	// `CfnPipe.PipeTargetParametersProperty.BatchJobParameters`.
	BatchJobParameters interface{} `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// `CfnPipe.PipeTargetParametersProperty.CloudWatchLogsParameters`.
	CloudWatchLogsParameters interface{} `field:"optional" json:"cloudWatchLogsParameters" yaml:"cloudWatchLogsParameters"`
	// `CfnPipe.PipeTargetParametersProperty.EcsTaskParameters`.
	EcsTaskParameters interface{} `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// `CfnPipe.PipeTargetParametersProperty.EventBridgeEventBusParameters`.
	EventBridgeEventBusParameters interface{} `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
	// `CfnPipe.PipeTargetParametersProperty.HttpParameters`.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// `CfnPipe.PipeTargetParametersProperty.InputTemplate`.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// `CfnPipe.PipeTargetParametersProperty.KinesisStreamParameters`.
	KinesisStreamParameters interface{} `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// `CfnPipe.PipeTargetParametersProperty.LambdaFunctionParameters`.
	LambdaFunctionParameters interface{} `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// `CfnPipe.PipeTargetParametersProperty.RedshiftDataParameters`.
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// `CfnPipe.PipeTargetParametersProperty.SageMakerPipelineParameters`.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// `CfnPipe.PipeTargetParametersProperty.SqsQueueParameters`.
	SqsQueueParameters interface{} `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// `CfnPipe.PipeTargetParametersProperty.StepFunctionStateMachineParameters`.
	StepFunctionStateMachineParameters interface{} `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}

