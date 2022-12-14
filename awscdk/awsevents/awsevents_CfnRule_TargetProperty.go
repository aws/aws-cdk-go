package awsevents


// Targets are the resources to be invoked when a rule is triggered.
//
// For a complete list of services and resources that can be set as a target, see [PutTargets](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutTargets.html) .
//
// If you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetProperty := &targetProperty{
//   	arn: jsii.String("arn"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	batchParameters: &batchParametersProperty{
//   		jobDefinition: jsii.String("jobDefinition"),
//   		jobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		arrayProperties: &batchArrayPropertiesProperty{
//   			size: jsii.Number(123),
//   		},
//   		retryStrategy: &batchRetryStrategyProperty{
//   			attempts: jsii.Number(123),
//   		},
//   	},
//   	deadLetterConfig: &deadLetterConfigProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	ecsParameters: &ecsParametersProperty{
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
//   			awsVpcConfiguration: &awsVpcConfigurationProperty{
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
//   		placementConstraints: []interface{}{
//   			&placementConstraintProperty{
//   				expression: jsii.String("expression"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		placementStrategies: []interface{}{
//   			&placementStrategyProperty{
//   				field: jsii.String("field"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		platformVersion: jsii.String("platformVersion"),
//   		propagateTags: jsii.String("propagateTags"),
//   		referenceId: jsii.String("referenceId"),
//   		tagList: []interface{}{
//   			&cfnTag{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		taskCount: jsii.Number(123),
//   	},
//   	httpParameters: &httpParametersProperty{
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
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   	inputTransformer: &inputTransformerProperty{
//   		inputTemplate: jsii.String("inputTemplate"),
//
//   		// the properties below are optional
//   		inputPathsMap: map[string]*string{
//   			"inputPathsMapKey": jsii.String("inputPathsMap"),
//   		},
//   	},
//   	kinesisParameters: &kinesisParametersProperty{
//   		partitionKeyPath: jsii.String("partitionKeyPath"),
//   	},
//   	redshiftDataParameters: &redshiftDataParametersProperty{
//   		database: jsii.String("database"),
//   		sql: jsii.String("sql"),
//
//   		// the properties below are optional
//   		dbUser: jsii.String("dbUser"),
//   		secretManagerArn: jsii.String("secretManagerArn"),
//   		statementName: jsii.String("statementName"),
//   		withEvent: jsii.Boolean(false),
//   	},
//   	retryPolicy: &retryPolicyProperty{
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   	runCommandParameters: &runCommandParametersProperty{
//   		runCommandTargets: []interface{}{
//   			&runCommandTargetProperty{
//   				key: jsii.String("key"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	sageMakerPipelineParameters: &sageMakerPipelineParametersProperty{
//   		pipelineParameterList: []interface{}{
//   			&sageMakerPipelineParameterProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	sqsParameters: &sqsParametersProperty{
//   		messageGroupId: jsii.String("messageGroupId"),
//   	},
//   }
//
type CfnRule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the target.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The ID of the target within the specified rule.
	//
	// Use this ID to reference the target when updating the rule. We recommend using a memorable and unique string.
	Id *string `field:"required" json:"id" yaml:"id"`
	// If the event target is an AWS Batch job, this contains the job definition, job name, and other parameters.
	//
	// For more information, see [Jobs](https://docs.aws.amazon.com/batch/latest/userguide/jobs.html) in the *AWS Batch User Guide* .
	BatchParameters interface{} `field:"optional" json:"batchParameters" yaml:"batchParameters"`
	// The `DeadLetterConfig` that defines the target queue to send dead-letter queue events to.
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Contains the Amazon ECS task definition and task count to be used, if the event target is an Amazon ECS task.
	//
	// For more information about Amazon ECS tasks, see [Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon EC2 Container Service Developer Guide* .
	EcsParameters interface{} `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination.
	//
	// If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence.
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The value of the JSONPath that is used for extracting part of the matched event when passing it to the target.
	//
	// You must use JSON dot notation, not bracket notation. For more information about JSON paths, see [JSONPath](https://docs.aws.amazon.com/http://goessner.net/articles/JsonPath/) .
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Settings to enable you to provide custom input to a target based on certain event data.
	//
	// You can extract one or more key-value pairs from the event and then use that data to send customized input to the target.
	InputTransformer interface{} `field:"optional" json:"inputTransformer" yaml:"inputTransformer"`
	// The custom parameter you can use to control the shard assignment, when the target is a Kinesis data stream.
	//
	// If you do not include this parameter, the default is to use the `eventId` as the partition key.
	KinesisParameters interface{} `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// Contains the Amazon Redshift Data API parameters to use when the target is a Amazon Redshift cluster.
	//
	// If you specify a Amazon Redshift Cluster as a Target, you can use this to specify parameters to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The `RetryPolicy` object that contains the retry policy configuration to use for the dead-letter queue.
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered.
	//
	// If one rule triggers multiple targets, you can use a different IAM role for each target.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command.
	RunCommandParameters interface{} `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// Contains the SageMaker Model Building Pipeline parameters to start execution of a SageMaker Model Building Pipeline.
	//
	// If you specify a SageMaker Model Building Pipeline as a target, you can use this to specify parameters to start a pipeline execution based on EventBridge events.
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Contains the message group ID to use when the target is a FIFO queue.
	//
	// If you specify an SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
	SqsParameters interface{} `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

