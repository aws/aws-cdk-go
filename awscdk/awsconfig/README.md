# AWS Config Construct Library

[AWS Config](https://docs.aws.amazon.com/config/latest/developerguide/WhatIsConfig.html) provides a detailed view of the configuration of AWS resources in your AWS account.
This includes how the resources are related to one another and how they were configured in the
past so that you can see how the configurations and relationships change over time.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## Initial Setup

Before using the constructs provided in this module, you need to set up AWS Config
in the region in which it will be used. This setup includes the one-time creation of the
following resources per region:

* `ConfigurationRecorder`: Configure which resources will be recorded for config changes.
* `DeliveryChannel`: Configure where to store the recorded data.

The following guides provide the steps for getting started with AWS Config:

* [Using the AWS Console](https://docs.aws.amazon.com/config/latest/developerguide/gs-console.html)
* [Using the AWS CLI](https://docs.aws.amazon.com/config/latest/developerguide/gs-cli.html)

## Rules

AWS Config can evaluate the configuration settings of your AWS resources by creating AWS Config rules,
which represent your ideal configuration settings.

See [Evaluating Resources with AWS Config Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html) to learn more about AWS Config rules.

### AWS Managed Rules

AWS Config provides AWS managed rules, which are predefined, customizable rules that AWS Config
uses to evaluate whether your AWS resources comply with common best practices.

For example, you could create a managed rule that checks whether active access keys are rotated
within the number of days specified.

```go
// https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
// https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html
config.NewManagedRule(this, jsii.String("AccessKeysRotated"), &ManagedRuleProps{
	Identifier: config.ManagedRuleIdentifiers_ACCESS_KEYS_ROTATED(),
	InputParameters: map[string]interface{}{
		"maxAccessKeyAge": jsii.Number(60),
	},

	// default is 24 hours
	MaximumExecutionFrequency: config.MaximumExecutionFrequency_TWELVE_HOURS,
})
```

Identifiers for AWS managed rules are available through static constants in the `ManagedRuleIdentifiers` class.
You can find supported input parameters in the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).

The following higher level constructs for AWS managed rules are available.

#### Access Key rotation

Checks whether your active access keys are rotated within the number of days specified.

```go
// compliant if access keys have been rotated within the last 90 days
// compliant if access keys have been rotated within the last 90 days
config.NewAccessKeysRotated(this, jsii.String("AccessKeyRotated"))
```

#### CloudFormation Stack drift detection

Checks whether your CloudFormation stack's actual configuration differs, or has drifted,
from it's expected configuration.

```go
// compliant if stack's status is 'IN_SYNC'
// non-compliant if the stack's drift status is 'DRIFTED'
// compliant if stack's status is 'IN_SYNC'
// non-compliant if the stack's drift status is 'DRIFTED'
config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"), &CloudFormationStackDriftDetectionCheckProps{
	OwnStackOnly: jsii.Boolean(true),
})
```

#### CloudFormation Stack notifications

Checks whether your CloudFormation stacks are sending event notifications to a SNS topic.

```go
// topics to which CloudFormation stacks may send event notifications
topic1 := sns.NewTopic(this, jsii.String("AllowedTopic1"))
topic2 := sns.NewTopic(this, jsii.String("AllowedTopic2"))

// non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
// non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
config.NewCloudFormationStackNotificationCheck(this, jsii.String("NotificationCheck"), &CloudFormationStackNotificationCheckProps{
	Topics: []iTopic{
		topic1,
		topic2,
	},
})
```

### Custom rules

You can develop custom rules and add them to AWS Config. You associate each custom rule with an
AWS Lambda function and Guard.

#### Custom Lambda Rules

Lambda function which contains the logic that evaluates whether your AWS resources comply with the rule.

```go
// Lambda function containing logic that evaluates compliance with the rule.
evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
})

// A custom rule that runs on configuration changes of EC2 instances
customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
	ConfigurationChanges: jsii.Boolean(true),
	LambdaFunction: evalComplianceFn,
	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
})
```

#### Custom Policy Rules

Guard which contains the logic that evaluates whether your AWS resources comply with the rule.

```go
samplePolicyText := `
# This rule checks if point in time recovery (PITR) is enabled on active Amazon DynamoDB tables
let status = ['ACTIVE']

rule tableisactive when
    resourceType == "AWS::DynamoDB::Table" {
    configuration.tableStatus == %status
}

rule checkcompliance when
    resourceType == "AWS::DynamoDB::Table"
    tableisactive {
        let pitr = supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus
        %pitr == "ENABLED"
}
`

config.NewCustomPolicy(this, jsii.String("Custom"), &CustomPolicyProps{
	PolicyText: samplePolicyText,
	EnableDebugLog: jsii.Boolean(true),
	RuleScope: config.RuleScope_FromResources([]resourceType{
		config.*resourceType_DYNAMODB_TABLE(),
	}),
})
```

### Triggers

AWS Lambda executes functions in response to events that are published by AWS Services.
The function for a custom Config rule receives an event that is published by AWS Config,
and is responsible for evaluating the compliance of the rule.

Evaluations can be triggered by configuration changes, periodically, or both.
To create a custom rule, define a `CustomRule` and specify the Lambda Function
to run and the trigger types.

```go
var evalComplianceFn function


config.NewCustomRule(this, jsii.String("CustomRule"), &CustomRuleProps{
	LambdaFunction: evalComplianceFn,
	ConfigurationChanges: jsii.Boolean(true),
	Periodic: jsii.Boolean(true),

	// default is 24 hours
	MaximumExecutionFrequency: config.MaximumExecutionFrequency_SIX_HOURS,
})
```

When the trigger for a rule occurs, the Lambda function is invoked by publishing an event.
See [example events for AWS Config Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules_example-events.html)

The AWS documentation has examples of Lambda functions for evaluations that are
[triggered by configuration changes](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules_nodejs-sample.html#event-based-example-rule) and [triggered periodically](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules_nodejs-sample.html#periodic-example-rule)

### Scope

By default rules are triggered by changes to all [resources](https://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources).

Use the `RuleScope` APIs (`fromResource()`, `fromResources()` or `fromTag()`) to restrict
the scope of both managed and custom rules:

```go
var evalComplianceFn function
sshRule := config.NewManagedRule(this, jsii.String("SSH"), &ManagedRuleProps{
	Identifier: config.ManagedRuleIdentifiers_EC2_SECURITY_GROUPS_INCOMING_SSH_DISABLED(),
	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_SECURITY_GROUP(), jsii.String("sg-1234567890abcdefgh")),
})
customRule := config.NewCustomRule(this, jsii.String("Lambda"), &CustomRuleProps{
	LambdaFunction: evalComplianceFn,
	ConfigurationChanges: jsii.Boolean(true),
	RuleScope: config.RuleScope_FromResources([]resourceType{
		config.*resourceType_CLOUDFORMATION_STACK(),
		config.*resourceType_S3_BUCKET(),
	}),
})

tagRule := config.NewCustomRule(this, jsii.String("CostCenterTagRule"), &CustomRuleProps{
	LambdaFunction: evalComplianceFn,
	ConfigurationChanges: jsii.Boolean(true),
	RuleScope: config.RuleScope_FromTag(jsii.String("Cost Center"), jsii.String("MyApp")),
})
```

### Evaluation Mode

You can specify the [evaluation mode](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config-rules.html) for a rule to determine when AWS Config runs evaluations.

Use the `evaluationModes` property to specify the evaluation mode:

```go
var fn function
var samplePolicyText string


config.NewManagedRule(this, jsii.String("ManagedRule"), &ManagedRuleProps{
	Identifier: config.ManagedRuleIdentifiers_API_GW_XRAY_ENABLED(),
	EvaluationModes: config.EvaluationMode_DETECTIVE_AND_PROACTIVE(),
})

config.NewCustomRule(this, jsii.String("CustomRule"), &CustomRuleProps{
	LambdaFunction: fn,
	EvaluationModes: config.EvaluationMode_PROACTIVE(),
})

config.NewCustomPolicy(this, jsii.String("CustomPolicy"), &CustomPolicyProps{
	PolicyText: samplePolicyText,
	EvaluationModes: config.EvaluationMode_DETECTIVE(),
})
```

**Note**: Proactive evaluation mode is not supported for all rules. See [AWS Config documentation](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config-rules.html#aws-config-rules-evaluation-modes) for more information.

### Events

You can define Amazon EventBridge event rules which trigger when a compliance check fails
or when a rule is re-evaluated.

Use the `onComplianceChange()` APIs to trigger an EventBridge event when a compliance check
of your AWS Config Rule fails:

```go
// Topic to which compliance notification events will be published
complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))

rule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
rule.onComplianceChange(jsii.String("TopicEvent"), &OnEventOptions{
	Target: targets.NewSnsTopic(complianceTopic),
})
```

Use the `onReEvaluationStatus()` status to trigger an EventBridge event when an AWS Config
rule is re-evaluated.

```go
// Topic to which re-evaluation notification events will be published
reEvaluationTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))

rule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))
rule.onReEvaluationStatus(jsii.String("ReEvaluationEvent"), &OnEventOptions{
	Target: targets.NewSnsTopic(reEvaluationTopic),
})
```

### Example

The following example creates a custom rule that evaluates whether EC2 instances are compliant.
Compliance events are published to an SNS topic.

```go
// Lambda function containing logic that evaluates compliance with the rule.
evalComplianceFn := lambda.NewFunction(this, jsii.String("CustomFunction"), &FunctionProps{
	Code: lambda.AssetCode_FromInline(jsii.String("exports.handler = (event) => console.log(event);")),
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_18_X(),
})

// A custom rule that runs on configuration changes of EC2 instances
customRule := config.NewCustomRule(this, jsii.String("Custom"), &CustomRuleProps{
	ConfigurationChanges: jsii.Boolean(true),
	LambdaFunction: evalComplianceFn,
	RuleScope: config.RuleScope_FromResource(config.ResourceType_EC2_INSTANCE()),
})

// A rule to detect stack drifts
driftRule := config.NewCloudFormationStackDriftDetectionCheck(this, jsii.String("Drift"))

// Topic to which compliance notification events will be published
complianceTopic := sns.NewTopic(this, jsii.String("ComplianceTopic"))

// Send notification on compliance change events
driftRule.onComplianceChange(jsii.String("ComplianceChange"), &OnEventOptions{
	Target: targets.NewSnsTopic(complianceTopic),
})
```
