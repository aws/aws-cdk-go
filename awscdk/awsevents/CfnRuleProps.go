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
//   cfnRuleProps := &CfnRuleProps{
//   	Description: jsii.String("description"),
//   	EventBusName: jsii.String("eventBusName"),
//   	EventPattern: eventPattern,
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	State: jsii.String("state"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			Arn: jsii.String("arn"),
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AppSyncParameters: &AppSyncParametersProperty{
//   				GraphQlOperation: jsii.String("graphQlOperation"),
//   			},
//   			BatchParameters: &BatchParametersProperty{
//   				JobDefinition: jsii.String("jobDefinition"),
//   				JobName: jsii.String("jobName"),
//
//   				// the properties below are optional
//   				ArrayProperties: &BatchArrayPropertiesProperty{
//   					Size: jsii.Number(123),
//   				},
//   				RetryStrategy: &BatchRetryStrategyProperty{
//   					Attempts: jsii.Number(123),
//   				},
//   			},
//   			DeadLetterConfig: &DeadLetterConfigProperty{
//   				Arn: jsii.String("arn"),
//   			},
//   			EcsParameters: &EcsParametersProperty{
//   				TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   				// the properties below are optional
//   				CapacityProviderStrategy: []interface{}{
//   					&CapacityProviderStrategyItemProperty{
//   						CapacityProvider: jsii.String("capacityProvider"),
//
//   						// the properties below are optional
//   						Base: jsii.Number(123),
//   						Weight: jsii.Number(123),
//   					},
//   				},
//   				EnableEcsManagedTags: jsii.Boolean(false),
//   				EnableExecuteCommand: jsii.Boolean(false),
//   				Group: jsii.String("group"),
//   				LaunchType: jsii.String("launchType"),
//   				NetworkConfiguration: &NetworkConfigurationProperty{
//   					AwsVpcConfiguration: &AwsVpcConfigurationProperty{
//   						Subnets: []*string{
//   							jsii.String("subnets"),
//   						},
//
//   						// the properties below are optional
//   						AssignPublicIp: jsii.String("assignPublicIp"),
//   						SecurityGroups: []*string{
//   							jsii.String("securityGroups"),
//   						},
//   					},
//   				},
//   				PlacementConstraints: []interface{}{
//   					&PlacementConstraintProperty{
//   						Expression: jsii.String("expression"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				PlacementStrategies: []interface{}{
//   					&PlacementStrategyProperty{
//   						Field: jsii.String("field"),
//   						Type: jsii.String("type"),
//   					},
//   				},
//   				PlatformVersion: jsii.String("platformVersion"),
//   				PropagateTags: jsii.String("propagateTags"),
//   				ReferenceId: jsii.String("referenceId"),
//   				TagList: []interface{}{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				TaskCount: jsii.Number(123),
//   			},
//   			HttpParameters: &HttpParametersProperty{
//   				HeaderParameters: map[string]*string{
//   					"headerParametersKey": jsii.String("headerParameters"),
//   				},
//   				PathParameterValues: []*string{
//   					jsii.String("pathParameterValues"),
//   				},
//   				QueryStringParameters: map[string]*string{
//   					"queryStringParametersKey": jsii.String("queryStringParameters"),
//   				},
//   			},
//   			Input: jsii.String("input"),
//   			InputPath: jsii.String("inputPath"),
//   			InputTransformer: &InputTransformerProperty{
//   				InputTemplate: jsii.String("inputTemplate"),
//
//   				// the properties below are optional
//   				InputPathsMap: map[string]*string{
//   					"inputPathsMapKey": jsii.String("inputPathsMap"),
//   				},
//   			},
//   			KinesisParameters: &KinesisParametersProperty{
//   				PartitionKeyPath: jsii.String("partitionKeyPath"),
//   			},
//   			RedshiftDataParameters: &RedshiftDataParametersProperty{
//   				Database: jsii.String("database"),
//
//   				// the properties below are optional
//   				DbUser: jsii.String("dbUser"),
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   				Sql: jsii.String("sql"),
//   				Sqls: []*string{
//   					jsii.String("sqls"),
//   				},
//   				StatementName: jsii.String("statementName"),
//   				WithEvent: jsii.Boolean(false),
//   			},
//   			RetryPolicy: &RetryPolicyProperty{
//   				MaximumEventAgeInSeconds: jsii.Number(123),
//   				MaximumRetryAttempts: jsii.Number(123),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   			RunCommandParameters: &RunCommandParametersProperty{
//   				RunCommandTargets: []interface{}{
//   					&RunCommandTargetProperty{
//   						Key: jsii.String("key"),
//   						Values: []*string{
//   							jsii.String("values"),
//   						},
//   					},
//   				},
//   			},
//   			SageMakerPipelineParameters: &SageMakerPipelineParametersProperty{
//   				PipelineParameterList: []interface{}{
//   					&SageMakerPipelineParameterProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			SqsParameters: &SqsParametersProperty{
//   				MessageGroupId: jsii.String("messageGroupId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html
//
type CfnRuleProps struct {
	// The description of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name or ARN of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventbusname
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The event pattern of the rule.
	//
	// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern
	//
	EventPattern interface{} `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The name of the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	//
	// If you're setting an event bus in another account as the target and that account granted permission to your account through an organization instead of directly by the account ID, you must specify a `RoleArn` with proper permissions in the `Target` structure, instead of here in this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The scheduling expression.
	//
	// For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see [Creating an Amazon EventBridge rule that runs on a schedule](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-create-rule-schedule.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The state of the rule.
	//
	// Valid values include:
	//
	// - `DISABLED` : The rule is disabled. EventBridge does not match any events against the rule.
	// - `ENABLED` : The rule is enabled. EventBridge matches events against the rule, *except* for AWS management events delivered through CloudTrail.
	// - `ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS` : The rule is enabled for all events, including AWS management events delivered through CloudTrail.
	//
	// Management events provide visibility into management operations that are performed on resources in your AWS account. These are also known as control plane operations. For more information, see [Logging management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html#logging-management-events) in the *CloudTrail User Guide* , and [Filtering management events from AWS services](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-service-event.html#eb-service-event-cloudtrail) in the *Amazon EventBridge User Guide* .
	//
	// This value is only valid for rules on the [default](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is-how-it-works-concepts.html#eb-bus-concepts-buses) event bus or [custom event buses](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-create-event-bus.html) . It does not apply to [partner event buses](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.
	//
	// Targets are the resources that are invoked when a rule is triggered.
	//
	// The maximum number of entries per request is 10.
	//
	// > Each rule can have up to five (5) targets associated with it at one time.
	//
	// For a list of services you can configure as targets for events, see [EventBridge targets](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html) in the *Amazon EventBridge User Guide* .
	//
	// Creating rules with built-in targets is supported only in the AWS Management Console . The built-in targets are:
	//
	// - `Amazon EBS CreateSnapshot API call`
	// - `Amazon EC2 RebootInstances API call`
	// - `Amazon EC2 StopInstances API call`
	// - `Amazon EC2 TerminateInstances API call`
	//
	// For some target types, `PutTargets` provides target-specific parameters. If the target is a Kinesis data stream, you can optionally specify which shard the event goes to by using the `KinesisParameters` argument. To invoke a command on multiple EC2 instances with one rule, you can use the `RunCommandParameters` field.
	//
	// To be able to make API calls against the resources that you own, Amazon EventBridge needs the appropriate permissions:
	//
	// - For AWS Lambda and Amazon SNS resources, EventBridge relies on resource-based policies.
	// - For EC2 instances, Kinesis Data Streams, AWS Step Functions state machines and API Gateway APIs, EventBridge relies on IAM roles that you specify in the `RoleARN` argument in `PutTargets` .
	//
	// For more information, see [Authentication and Access Control](https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html) in the *Amazon EventBridge User Guide* .
	//
	// If another AWS account is in the same region and has granted you permission (using `PutPermission` ), you can send events to that account. Set that account's event bus as a target of the rules in your account. To send the matched events to the other account, specify that account's event bus as the `Arn` value when you run `PutTargets` . If your account sends events to another account, your account is charged for each sent event. Each event sent to another account is charged as a custom event. The account receiving the event is not charged. For more information, see [Amazon EventBridge Pricing](https://docs.aws.amazon.com/eventbridge/pricing/) .
	//
	// > `Input` , `InputPath` , and `InputTransformer` are not available with `PutTarget` if the target is an event bus of a different AWS account.
	//
	// If you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
	//
	// > If you have an IAM role on a cross-account event bus target, a `PutTargets` call without a role on the same target (same `Id` and `Arn` ) will not remove the role.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

