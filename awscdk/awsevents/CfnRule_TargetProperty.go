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
//   targetProperty := &TargetProperty{
//   	Arn: jsii.String("arn"),
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	AppSyncParameters: &AppSyncParametersProperty{
//   		GraphQlOperation: jsii.String("graphQlOperation"),
//   	},
//   	BatchParameters: &BatchParametersProperty{
//   		JobDefinition: jsii.String("jobDefinition"),
//   		JobName: jsii.String("jobName"),
//
//   		// the properties below are optional
//   		ArrayProperties: &BatchArrayPropertiesProperty{
//   			Size: jsii.Number(123),
//   		},
//   		RetryStrategy: &BatchRetryStrategyProperty{
//   			Attempts: jsii.Number(123),
//   		},
//   	},
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	EcsParameters: &EcsParametersProperty{
//   		TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   		// the properties below are optional
//   		CapacityProviderStrategy: []interface{}{
//   			&CapacityProviderStrategyItemProperty{
//   				CapacityProvider: jsii.String("capacityProvider"),
//
//   				// the properties below are optional
//   				Base: jsii.Number(123),
//   				Weight: jsii.Number(123),
//   			},
//   		},
//   		EnableEcsManagedTags: jsii.Boolean(false),
//   		EnableExecuteCommand: jsii.Boolean(false),
//   		Group: jsii.String("group"),
//   		LaunchType: jsii.String("launchType"),
//   		NetworkConfiguration: &NetworkConfigurationProperty{
//   			AwsVpcConfiguration: &AwsVpcConfigurationProperty{
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//
//   				// the properties below are optional
//   				AssignPublicIp: jsii.String("assignPublicIp"),
//   				SecurityGroups: []*string{
//   					jsii.String("securityGroups"),
//   				},
//   			},
//   		},
//   		PlacementConstraints: []interface{}{
//   			&PlacementConstraintProperty{
//   				Expression: jsii.String("expression"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlacementStrategies: []interface{}{
//   			&PlacementStrategyProperty{
//   				Field: jsii.String("field"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		PlatformVersion: jsii.String("platformVersion"),
//   		PropagateTags: jsii.String("propagateTags"),
//   		ReferenceId: jsii.String("referenceId"),
//   		TagList: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		TaskCount: jsii.Number(123),
//   	},
//   	HttpParameters: &HttpParametersProperty{
//   		HeaderParameters: map[string]*string{
//   			"headerParametersKey": jsii.String("headerParameters"),
//   		},
//   		PathParameterValues: []*string{
//   			jsii.String("pathParameterValues"),
//   		},
//   		QueryStringParameters: map[string]*string{
//   			"queryStringParametersKey": jsii.String("queryStringParameters"),
//   		},
//   	},
//   	Input: jsii.String("input"),
//   	InputPath: jsii.String("inputPath"),
//   	InputTransformer: &InputTransformerProperty{
//   		InputTemplate: jsii.String("inputTemplate"),
//
//   		// the properties below are optional
//   		InputPathsMap: map[string]*string{
//   			"inputPathsMapKey": jsii.String("inputPathsMap"),
//   		},
//   	},
//   	KinesisParameters: &KinesisParametersProperty{
//   		PartitionKeyPath: jsii.String("partitionKeyPath"),
//   	},
//   	RedshiftDataParameters: &RedshiftDataParametersProperty{
//   		Database: jsii.String("database"),
//
//   		// the properties below are optional
//   		DbUser: jsii.String("dbUser"),
//   		SecretManagerArn: jsii.String("secretManagerArn"),
//   		Sql: jsii.String("sql"),
//   		Sqls: []*string{
//   			jsii.String("sqls"),
//   		},
//   		StatementName: jsii.String("statementName"),
//   		WithEvent: jsii.Boolean(false),
//   	},
//   	RetryPolicy: &RetryPolicyProperty{
//   		MaximumEventAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	RunCommandParameters: &RunCommandParametersProperty{
//   		RunCommandTargets: []interface{}{
//   			&RunCommandTargetProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	SageMakerPipelineParameters: &SageMakerPipelineParametersProperty{
//   		PipelineParameterList: []interface{}{
//   			&SageMakerPipelineParameterProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SqsParameters: &SqsParametersProperty{
//   		MessageGroupId: jsii.String("messageGroupId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html
//
type CfnRule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The ID of the target within the specified rule.
	//
	// Use this ID to reference the target when updating the rule. We recommend using a memorable and unique string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Contains the GraphQL operation to be parsed and executed, if the event target is an AWS AppSync API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-appsyncparameters
	//
	AppSyncParameters interface{} `field:"optional" json:"appSyncParameters" yaml:"appSyncParameters"`
	// If the event target is an AWS Batch job, this contains the job definition, job name, and other parameters.
	//
	// For more information, see [Jobs](https://docs.aws.amazon.com/batch/latest/userguide/jobs.html) in the *AWS Batch User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-batchparameters
	//
	BatchParameters interface{} `field:"optional" json:"batchParameters" yaml:"batchParameters"`
	// The `DeadLetterConfig` that defines the target queue to send dead-letter queue events to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-deadletterconfig
	//
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Contains the Amazon ECS task definition and task count to be used, if the event target is an Amazon ECS task.
	//
	// For more information about Amazon ECS tasks, see [Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon EC2 Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-ecsparameters
	//
	EcsParameters interface{} `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway endpoint or EventBridge ApiDestination.
	//
	// If you specify an API Gateway API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-httpparameters
	//
	HttpParameters interface{} `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-input
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The value of the JSONPath that is used for extracting part of the matched event when passing it to the target.
	//
	// You may use JSON dot notation or bracket notation. For more information about JSON paths, see [JSONPath](https://docs.aws.amazon.com/http://goessner.net/articles/JsonPath/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-inputpath
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Settings to enable you to provide custom input to a target based on certain event data.
	//
	// You can extract one or more key-value pairs from the event and then use that data to send customized input to the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-inputtransformer
	//
	InputTransformer interface{} `field:"optional" json:"inputTransformer" yaml:"inputTransformer"`
	// The custom parameter you can use to control the shard assignment, when the target is a Kinesis data stream.
	//
	// If you do not include this parameter, the default is to use the `eventId` as the partition key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-kinesisparameters
	//
	KinesisParameters interface{} `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// Contains the Amazon Redshift Data API parameters to use when the target is a Amazon Redshift cluster.
	//
	// If you specify a Amazon Redshift Cluster as a Target, you can use this to specify parameters to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-redshiftdataparameters
	//
	RedshiftDataParameters interface{} `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The `RetryPolicy` object that contains the retry policy configuration to use for the dead-letter queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-retrypolicy
	//
	RetryPolicy interface{} `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered.
	//
	// If one rule triggers multiple targets, you can use a different IAM role for each target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-runcommandparameters
	//
	RunCommandParameters interface{} `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// Contains the SageMaker Model Building Pipeline parameters to start execution of a SageMaker Model Building Pipeline.
	//
	// If you specify a SageMaker Model Building Pipeline as a target, you can use this to specify parameters to start a pipeline execution based on EventBridge events.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-sagemakerpipelineparameters
	//
	SageMakerPipelineParameters interface{} `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Contains the message group ID to use when the target is a FIFO queue.
	//
	// If you specify an SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-target.html#cfn-events-rule-target-sqsparameters
	//
	SqsParameters interface{} `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

