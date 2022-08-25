package awsevents


// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventPattern interface{}
//
//   cfnRuleProps := &cfnRuleProps{
//   	description: jsii.String("description"),
//   	eventBusName: jsii.String("eventBusName"),
//   	eventPattern: eventPattern,
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   	state: jsii.String("state"),
//   	targets: []interface{}{
//   		&targetProperty{
//   			arn: jsii.String("arn"),
//   			id: jsii.String("id"),
//
//   			// the properties below are optional
//   			batchParameters: &batchParametersProperty{
//   				jobDefinition: jsii.String("jobDefinition"),
//   				jobName: jsii.String("jobName"),
//
//   				// the properties below are optional
//   				arrayProperties: &batchArrayPropertiesProperty{
//   					size: jsii.Number(123),
//   				},
//   				retryStrategy: &batchRetryStrategyProperty{
//   					attempts: jsii.Number(123),
//   				},
//   			},
//   			deadLetterConfig: &deadLetterConfigProperty{
//   				arn: jsii.String("arn"),
//   			},
//   			ecsParameters: &ecsParametersProperty{
//   				taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   				// the properties below are optional
//   				capacityProviderStrategy: []interface{}{
//   					&capacityProviderStrategyItemProperty{
//   						capacityProvider: jsii.String("capacityProvider"),
//
//   						// the properties below are optional
//   						base: jsii.Number(123),
//   						weight: jsii.Number(123),
//   					},
//   				},
//   				enableEcsManagedTags: jsii.Boolean(false),
//   				enableExecuteCommand: jsii.Boolean(false),
//   				group: jsii.String("group"),
//   				launchType: jsii.String("launchType"),
//   				networkConfiguration: &networkConfigurationProperty{
//   					awsVpcConfiguration: &awsVpcConfigurationProperty{
//   						subnets: []*string{
//   							jsii.String("subnets"),
//   						},
//
//   						// the properties below are optional
//   						assignPublicIp: jsii.String("assignPublicIp"),
//   						securityGroups: []*string{
//   							jsii.String("securityGroups"),
//   						},
//   					},
//   				},
//   				placementConstraints: []interface{}{
//   					&placementConstraintProperty{
//   						expression: jsii.String("expression"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				placementStrategies: []interface{}{
//   					&placementStrategyProperty{
//   						field: jsii.String("field"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   				platformVersion: jsii.String("platformVersion"),
//   				propagateTags: jsii.String("propagateTags"),
//   				referenceId: jsii.String("referenceId"),
//   				tagList: []interface{}{
//   					&cfnTag{
//   						key: jsii.String("key"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				taskCount: jsii.Number(123),
//   			},
//   			httpParameters: &httpParametersProperty{
//   				headerParameters: map[string]*string{
//   					"headerParametersKey": jsii.String("headerParameters"),
//   				},
//   				pathParameterValues: []*string{
//   					jsii.String("pathParameterValues"),
//   				},
//   				queryStringParameters: map[string]*string{
//   					"queryStringParametersKey": jsii.String("queryStringParameters"),
//   				},
//   			},
//   			input: jsii.String("input"),
//   			inputPath: jsii.String("inputPath"),
//   			inputTransformer: &inputTransformerProperty{
//   				inputTemplate: jsii.String("inputTemplate"),
//
//   				// the properties below are optional
//   				inputPathsMap: map[string]*string{
//   					"inputPathsMapKey": jsii.String("inputPathsMap"),
//   				},
//   			},
//   			kinesisParameters: &kinesisParametersProperty{
//   				partitionKeyPath: jsii.String("partitionKeyPath"),
//   			},
//   			redshiftDataParameters: &redshiftDataParametersProperty{
//   				database: jsii.String("database"),
//   				sql: jsii.String("sql"),
//
//   				// the properties below are optional
//   				dbUser: jsii.String("dbUser"),
//   				secretManagerArn: jsii.String("secretManagerArn"),
//   				statementName: jsii.String("statementName"),
//   				withEvent: jsii.Boolean(false),
//   			},
//   			retryPolicy: &retryPolicyProperty{
//   				maximumEventAgeInSeconds: jsii.Number(123),
//   				maximumRetryAttempts: jsii.Number(123),
//   			},
//   			roleArn: jsii.String("roleArn"),
//   			runCommandParameters: &runCommandParametersProperty{
//   				runCommandTargets: []interface{}{
//   					&runCommandTargetProperty{
//   						key: jsii.String("key"),
//   						values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			sageMakerPipelineParameters: &sageMakerPipelineParametersProperty{
//   				pipelineParameterList: []interface{}{
//   					&sageMakerPipelineParameterProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			sqsParameters: &sqsParametersProperty{
//   				messageGroupId: jsii.String("messageGroupId"),
//   			},
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// The description of the rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or ARN of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The event pattern of the rule.
	//
	// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
	EventPattern interface{} `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The name of the rule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	//
	// If you're setting an event bus in another account as the target and that account granted permission to your account through an organization instead of directly by the account ID, you must specify a `RoleArn` with proper permissions in the `Target` structure, instead of here in this parameter.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The scheduling expression.
	//
	// For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see [Creating an Amazon EventBridge rule that runs on a schedule](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-create-rule-schedule.html) .
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The state of the rule.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.
	//
	// Targets are the resources that are invoked when a rule is triggered.
	//
	// > Each rule can have up to five (5) targets associated with it at one time.
	//
	// You can configure the following as targets for Events:
	//
	// - [API destination](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-destinations.html)
	// - [API Gateway](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-gateway-target.html)
	// - Batch job queue
	// - CloudWatch group
	// - CodeBuild project
	// - CodePipeline
	// - EC2 `CreateSnapshot` API call
	// - EC2 Image Builder
	// - EC2 `RebootInstances` API call
	// - EC2 `StopInstances` API call
	// - EC2 `TerminateInstances` API call
	// - ECS task
	// - [Event bus in a different account or Region](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cross-account.html)
	// - [Event bus in the same account and Region](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-bus-to-bus.html)
	// - Firehose delivery stream
	// - Glue workflow
	// - [Incident Manager response plan](https://docs.aws.amazon.com//incident-manager/latest/userguide/incident-creation.html#incident-tracking-auto-eventbridge)
	// - Inspector assessment template
	// - Kinesis stream
	// - Lambda function
	// - Redshift cluster
	// - SageMaker Pipeline
	// - SNS topic
	// - SQS queue
	// - Step Functions state machine
	// - Systems Manager Automation
	// - Systems Manager OpsItem
	// - Systems Manager Run Command
	//
	// Creating rules with built-in targets is supported only in the AWS Management Console . The built-in targets are `EC2 CreateSnapshot API call` , `EC2 RebootInstances API call` , `EC2 StopInstances API call` , and `EC2 TerminateInstances API call` .
	//
	// For some target types, `PutTargets` provides target-specific parameters. If the target is a Kinesis data stream, you can optionally specify which shard the event goes to by using the `KinesisParameters` argument. To invoke a command on multiple EC2 instances with one rule, you can use the `RunCommandParameters` field.
	//
	// To be able to make API calls against the resources that you own, Amazon EventBridge needs the appropriate permissions. For AWS Lambda and Amazon SNS resources, EventBridge relies on resource-based policies. For EC2 instances, Kinesis Data Streams, AWS Step Functions state machines and API Gateway REST APIs, EventBridge relies on IAM roles that you specify in the `RoleARN` argument in `PutTargets` . For more information, see [Authentication and Access Control](https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html) in the *Amazon EventBridge User Guide* .
	//
	// If another AWS account is in the same region and has granted you permission (using `PutPermission` ), you can send events to that account. Set that account's event bus as a target of the rules in your account. To send the matched events to the other account, specify that account's event bus as the `Arn` value when you run `PutTargets` . If your account sends events to another account, your account is charged for each sent event. Each event sent to another account is charged as a custom event. The account receiving the event is not charged. For more information, see [Amazon EventBridge Pricing](https://docs.aws.amazon.com/eventbridge/pricing/) .
	//
	// > `Input` , `InputPath` , and `InputTransformer` are not available with `PutTarget` if the target is an event bus of a different AWS account.
	//
	// If you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
	//
	// For more information about enabling cross-account events, see [PutPermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html) .
	//
	// *Input* , *InputPath* , and *InputTransformer* are mutually exclusive and optional parameters of a target. When a rule is triggered due to a matched event:
	//
	// - If none of the following arguments are specified for a target, then the entire event is passed to the target in JSON format (unless the target is Amazon EC2 Run Command or Amazon ECS task, in which case nothing from the event is passed to the target).
	// - If *Input* is specified in the form of valid JSON, then the matched event is overridden with this constant.
	// - If *InputPath* is specified in the form of JSONPath (for example, `$.detail` ), then only the part of the event specified in the path is passed to the target (for example, only the detail part of the event is passed).
	// - If *InputTransformer* is specified, then one or more specified JSONPaths are extracted from the event and used as values in a template that you specify as the input to the target.
	//
	// When you specify `InputPath` or `InputTransformer` , you must use JSON dot notation, not bracket notation.
	//
	// When you add targets to a rule and the associated rule triggers soon after, new or updated targets might not be immediately invoked. Allow a short period of time for changes to take effect.
	//
	// This action can partially fail if too many requests are made at the same time. If that happens, `FailedEntryCount` is non-zero in the response and each entry in `FailedEntries` provides the ID of the failed target and the error code.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

