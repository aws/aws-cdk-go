# Event Targets for Amazon EventBridge

This library contains integration classes to send Amazon EventBridge to any
number of supported AWS Services. Instances of these classes should be passed
to the `rule.addTarget()` method.

Currently supported are:

* [Event Targets for Amazon EventBridge](#event-targets-for-amazon-eventbridge)

  * [Event retry policy and using dead-letter queues](#event-retry-policy-and-using-dead-letter-queues)
  * [Invoke a Lambda function](#invoke-a-lambda-function)
  * [Log an event into a LogGroup](#log-an-event-into-a-loggroup)
  * [Start a CodeBuild build](#start-a-codebuild-build)
  * [Start a CodePipeline pipeline](#start-a-codepipeline-pipeline)
  * [Start a StepFunctions state machine](#start-a-stepfunctions-state-machine)
  * [Queue a Batch job](#queue-a-batch-job)
  * [Invoke an API Gateway REST API](#invoke-an-api-gateway-rest-api)
  * [Invoke an AWS API](#invoke-an-aws-api)
  * [Invoke an API Destination](#invoke-an-api-destination)
  * [Invoke an AppSync GraphQL API](#invoke-an-appsync-graphql-api)
  * [Put an event on an EventBridge bus](#put-an-event-on-an-eventbridge-bus)
  * [Run an ECS Task](#run-an-ecs-task)

    * [Tagging Tasks](#tagging-tasks)
    * [Launch type for ECS Task](#launch-type-for-ecs-task)
    * [Assign public IP addresses to tasks](#assign-public-ip-addresses-to-tasks)
    * [Enable Amazon ECS Exec for ECS Task](#enable-amazon-ecs-exec-for-ecs-task)
  * [Run a Redshift query](#schedule-a-redshift-query-serverless-or-cluster)
  * [Publish to an SNS topic](#publish-to-an-sns-topic)

See the README of the `aws-cdk-lib/aws-events` library for more information on
EventBridge.

## Event retry policy and using dead-letter queues

The Codebuild, CodePipeline, Lambda, Kinesis Data Streams, StepFunctions, LogGroup, SQSQueue, SNSTopic and ECSTask targets support attaching a [dead letter queue and setting retry policies](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html). See the [lambda example](#invoke-a-lambda-function).
Use [escape hatches](https://docs.aws.amazon.com/cdk/latest/guide/cfn_layer.html) for the other target types.

## Invoke a Lambda function

Use the `LambdaFunction` target to invoke a lambda function.

The code snippet below creates an event rule with a Lambda function as a target
triggered for every events from `aws.ec2` source. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
})

rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
	EventPattern: &EventPattern{
		Source: []*string{
			jsii.String("aws.ec2"),
		},
	},
})

queue := sqs.NewQueue(this, jsii.String("Queue"))

rule.AddTarget(targets.NewLambdaFunction(fn, &LambdaFunctionProps{
	DeadLetterQueue: queue,
	 // Optional: add a dead letter queue
	MaxEventAge: awscdk.Duration_Hours(jsii.Number(2)),
	 // Optional: set the maxEventAge retry policy
	RetryAttempts: jsii.Number(2),
}))
```

## Log an event into a LogGroup

Use the `LogGroup` target to log your events in a CloudWatch LogGroup.

For example, the following code snippet creates an event rule with a CloudWatch LogGroup as a target.
Every events sent from the `aws.ec2` source will be sent to the CloudWatch LogGroup.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"), &LogGroupProps{
	LogGroupName: jsii.String("MyLogGroup"),
})

rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
	EventPattern: &EventPattern{
		Source: []*string{
			jsii.String("aws.ec2"),
		},
	},
})

rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup))
```

A rule target input can also be specified to modify the event that is sent to the log group.
Unlike other event targets, CloudWatchLogs requires a specific input template format.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"
var logGroup logGroup
var rule rule


rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
	LogEvent: targets.LogGroupTargetInput_FromObjectV2(&LogGroupTargetInputOptions{
		Timestamp: events.EventField_FromPath(jsii.String("$.time")),
		Message: events.EventField_*FromPath(jsii.String("$.detail-type")),
	}),
}))
```

If you want to use static values to overwrite the `message` make sure that you provide a `string`
value.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"
var logGroup logGroup
var rule rule


rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
	LogEvent: targets.LogGroupTargetInput_FromObjectV2(&LogGroupTargetInputOptions{
		Message: jSON.stringify(map[string]*string{
			"CustomField": jsii.String("CustomValue"),
		}),
	}),
}))
```

The cloudwatch log event target will create an AWS custom resource internally which will default
to set `installLatestAwsSdk` to `true`. This may be problematic for CN partition deployment. To
workaround this issue, set `installLatestAwsSdk` to `false`.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"
var logGroup logGroup
var rule rule


rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
	InstallLatestAwsSdk: jsii.Boolean(false),
}))
```

## Start a CodeBuild build

Use the `CodeBuildProject` target to trigger a CodeBuild project.

The code snippet below creates a CodeCommit repository that triggers a CodeBuild project
on commit to the master branch. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import "github.com/aws/aws-cdk-go/awscdk"
import codecommit "github.com/aws/aws-cdk-go/awscdk"


repo := codecommit.NewRepository(this, jsii.String("MyRepo"), &RepositoryProps{
	RepositoryName: jsii.String("aws-cdk-codebuild-events"),
})

project := codebuild.NewProject(this, jsii.String("MyProject"), &ProjectProps{
	Source: codebuild.Source_CodeCommit(&CodeCommitSourceProps{
		Repository: repo,
	}),
})

deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

// trigger a build when a commit is pushed to the repo
onCommitRule := repo.onCommit(jsii.String("OnCommit"), &OnCommitOptions{
	Target: targets.NewCodeBuildProject(project, &CodeBuildProjectProps{
		DeadLetterQueue: deadLetterQueue,
	}),
	Branches: []*string{
		jsii.String("master"),
	},
})
```

## Start a CodePipeline pipeline

Use the `CodePipeline` target to trigger a CodePipeline pipeline.

The code snippet below creates a CodePipeline pipeline that is triggered every hour

```go
import codepipeline "github.com/aws/aws-cdk-go/awscdk"


pipeline := codepipeline.NewPipeline(this, jsii.String("Pipeline"))

rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Expression(jsii.String("rate(1 hour)")),
})

rule.AddTarget(targets.NewCodePipeline(pipeline))
```

## Start a StepFunctions state machine

Use the `SfnStateMachine` target to trigger a State Machine.

The code snippet below creates a Simple StateMachine that is triggered every minute with a
dummy object as input.
You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html)
to the target.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
})

dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &StateMachineProps{
	Definition: sfn.NewWait(this, jsii.String("Hello"), &WaitProps{
		Time: sfn.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(10))),
	}),
})

rule.AddTarget(targets.NewSfnStateMachine(stateMachine, &SfnStateMachineProps{
	Input: events.RuleTargetInput_FromObject(map[string]*string{
		"SomeParam": jsii.String("SomeValue"),
	}),
	DeadLetterQueue: dlq,
	Role: role,
}))
```

## Queue a Batch job

Use the `BatchJob` target to queue a Batch job.

The code snippet below creates a Simple JobQueue that is triggered every hour with a
dummy object as input.
You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html)
to the target.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
import ecs "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc


computeEnvironment := batch.NewFargateComputeEnvironment(this, jsii.String("ComputeEnv"), &FargateComputeEnvironmentProps{
	Vpc: Vpc,
})

jobQueue := batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
	Priority: jsii.Number(1),
	ComputeEnvironments: []orderedComputeEnvironment{
		&orderedComputeEnvironment{
			ComputeEnvironment: *ComputeEnvironment,
			Order: jsii.Number(1),
		},
	},
})

jobDefinition := batch.NewEcsJobDefinition(this, jsii.String("MyJob"), &EcsJobDefinitionProps{
	Container: batch.NewEcsEc2ContainerDefinition(this, jsii.String("Container"), &EcsEc2ContainerDefinitionProps{
		Image: ecs.ContainerImage_FromRegistry(jsii.String("test-repo")),
		Memory: cdk.Size_Mebibytes(jsii.Number(2048)),
		Cpu: jsii.Number(256),
	}),
})

queue := sqs.NewQueue(this, jsii.String("Queue"))

rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewBatchJob(jobQueue.JobQueueArn, jobQueue, jobDefinition.JobDefinitionArn, jobDefinition, &BatchJobProps{
	DeadLetterQueue: queue,
	Event: events.RuleTargetInput_FromObject(map[string]*string{
		"SomeParam": jsii.String("SomeValue"),
	}),
	RetryAttempts: jsii.Number(2),
	MaxEventAge: awscdk.Duration_*Hours(jsii.Number(2)),
}))
```

## Invoke an API Gateway REST API

Use the `ApiGateway` target to trigger a REST API.

The code snippet below creates a Api Gateway REST API that is invoked every hour.

```go
import api "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
})

fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = e => {}")),
})

restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &LambdaRestApiProps{
	Handler: fn,
})

dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

rule.AddTarget(
targets.NewApiGateway(restApi, &ApiGatewayProps{
	Path: jsii.String("/*/test"),
	Method: jsii.String("GET"),
	Stage: jsii.String("prod"),
	PathParameterValues: []*string{
		jsii.String("path-value"),
	},
	HeaderParameters: map[string]*string{
		"Header1": jsii.String("header1"),
	},
	QueryStringParameters: map[string]*string{
		"QueryParam1": jsii.String("query-param-1"),
	},
	DeadLetterQueue: dlq,
}))
```

## Invoke an API Gateway V2 HTTP API

Use the `ApiGatewayV2` target to trigger a HTTP API.

```go
import apigwv2 "github.com/aws/aws-cdk-go/awscdk"

var httpApi httpApi
var rule rule


rule.AddTarget(targets.NewApiGatewayV2(httpApi))
```

## Invoke an AWS API

Use the `AwsApi` target to make direct AWS API calls from EventBridge rules. This is useful for invoking AWS services that don't have a dedicated EventBridge target.

### Basic Usage

The following example shows how to update an ECS service when a rule is triggered:

```go
rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewAwsApi(&AwsApiProps{
	Service: jsii.String("ECS"),
	Action: jsii.String("updateService"),
	Parameters: map[string]interface{}{
		"service": jsii.String("my-service"),
		"forceNewDeployment": jsii.Boolean(true),
	},
}))
```

### IAM Permissions

By default, the AwsApi target automatically creates the necessary IAM permissions based on the service and action you specify. The permission format follows the pattern: `service:Action`.

For example:

* `ECS` service with `updateService` action → `ecs:UpdateService` permission
* `RDS` service with `createDBSnapshot` action → `rds:CreateDBSnapshot` permission

### Custom IAM Policy

In some cases, you may need to provide a custom IAM policy statement, especially when:

* You need to restrict permissions to specific resources (instead of `*`)
* The service requires additional permissions beyond the main action
* You want more granular control over the permissions

```go
import "github.com/aws/aws-cdk-go/awscdk"
import s3 "github.com/aws/aws-cdk-go/awscdk"

var rule rule
var bucket bucket


rule.AddTarget(targets.NewAwsApi(&AwsApiProps{
	Service: jsii.String("s3"),
	Action: jsii.String("GetBucketEncryption"),
	Parameters: map[string]*string{
		"Bucket": bucket.bucketName,
	},
	PolicyStatement: iam.NewPolicyStatement(&PolicyStatementProps{
		Effect: iam.Effect_ALLOW,
		Actions: []*string{
			jsii.String("s3:GetEncryptionConfiguration"),
		},
		Resources: []*string{
			bucket.BucketArn,
		},
	}),
}))
```

## Invoke an API Destination

Use the `targets.ApiDestination` target to trigger an external API. You need to
create an `events.Connection` and `events.ApiDestination` as well.

The code snippet below creates an external destination that is invoked every hour.

```go
connection := events.NewConnection(this, jsii.String("Connection"), &ConnectionProps{
	Authorization: events.Authorization_ApiKey(jsii.String("x-api-key"), awscdk.SecretValue_SecretsManager(jsii.String("ApiSecretName"))),
	Description: jsii.String("Connection with API Key x-api-key"),
})

destination := events.NewApiDestination(this, jsii.String("Destination"), &ApiDestinationProps{
	Connection: Connection,
	Endpoint: jsii.String("https://example.com"),
	Description: jsii.String("Calling example.com with API key x-api-key"),
})

rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
	Targets: []iRuleTarget{
		targets.NewApiDestination(destination),
	},
})
```

You can also import an existing connection and destination
to create additional rules:

```go
connection := events.Connection_FromEventBusArn(this, jsii.String("Connection"), jsii.String("arn:aws:events:us-east-1:123456789012:event-bus/EventBusName"), jsii.String("arn:aws:secretsmanager:us-east-1:123456789012:secret:SecretName-f3gDy9"))

apiDestinationArn := "arn:aws:events:us-east-1:123456789012:api-destination/DestinationName/11111111-1111-1111-1111-111111111111"
apiDestinationArnForPolicy := "arn:aws:events:us-east-1:123456789012:api-destination/DestinationName"
destination := events.ApiDestination_FromApiDestinationAttributes(this, jsii.String("Destination"), &ApiDestinationAttributes{
	ApiDestinationArn: jsii.String(ApiDestinationArn),
	Connection: Connection,
	ApiDestinationArnForPolicy: jsii.String(ApiDestinationArnForPolicy),
})

rule := events.NewRule(this, jsii.String("OtherRule"), &RuleProps{
	Schedule: events.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Targets: []iRuleTarget{
		targets.NewApiDestination(destination),
	},
})
```

## Invoke an AppSync GraphQL API

Use the `AppSync` target to trigger an AppSync GraphQL API. You need to
create an `AppSync.GraphqlApi` configured with `AWS_IAM` authorization mode.

The code snippet below creates an AppSync GraphQL API target that is invoked every hour, calling the `publish` mutation.

```go
import "github.com/aws/aws-cdk-go/awscdk"


api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
	Name: jsii.String("api"),
	Definition: appsync.Definition_FromFile(jsii.String("schema.graphql")),
	AuthorizationConfig: &AuthorizationConfig{
		DefaultAuthorization: &AuthorizationMode{
			AuthorizationType: appsync.AuthorizationType_IAM,
		},
	},
})

rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
	Variables: events.RuleTargetInput_FromObject(map[string]*string{
		"message": jsii.String("hello world"),
	}),
}))
```

You can pass an existing role with the proper permissions to be used for the target when the rule is triggered. The code snippet below uses an existing role and grants permissions to use the publish Mutation on the GraphQL API.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


api := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedAPI"), &GraphqlApiAttributes{
	GraphqlApiId: jsii.String("<api-id>"),
	GraphqlApiArn: jsii.String("<api-arn>"),
	GraphQLEndpointArn: jsii.String("<api-endpoint-arn>"),
	Visibility: appsync.Visibility_GLOBAL,
	Modes: []authorizationType{
		appsync.*authorizationType_IAM,
	},
})

rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Minutes(jsii.Number(1))),
})
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
})

// allow EventBridge to use the `publish` mutation
api.GrantMutation(role, jsii.String("publish"))

rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
	Variables: events.RuleTargetInput_FromObject(map[string]*string{
		"message": jsii.String("hello world"),
	}),
	EventRole: role,
}))
```

## Put an event on an EventBridge bus

Use the `EventBus` target to route event to a different EventBus.

The code snippet below creates the scheduled event rule that route events to an imported event bus.

```go
rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Expression(jsii.String("rate(1 minute)")),
})

rule.AddTarget(targets.NewEventBus(events.EventBus_FromEventBusArn(this, jsii.String("External"), jsii.String("arn:aws:events:eu-west-1:999999999999:event-bus/test-bus"))))
```

## Run an ECS Task

Use the `EcsTask` target to run an ECS Task.

The code snippet below creates a scheduled event rule that will run the task described in `taskDefinition` every hour.

### Tagging Tasks

By default, ECS tasks run from EventBridge targets will not have tags applied to
them. You can set the `propagateTags` field to propagate the tags set on the task
definition to the task initialized by the event trigger.

If you want to set tags independent of those applied to the TaskDefinition, you
can use the `tags` array. Both of these fields can be used together or separately
to set tags on the triggered task.

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"

var cluster iCluster
var taskDefinition taskDefinition


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(
targets.NewEcsTask(&EcsTaskProps{
	Cluster: cluster,
	TaskDefinition: taskDefinition,
	PropagateTags: ecs.PropagatedTagSource_TASK_DEFINITION,
	Tags: []tag{
		&tag{
			Key: jsii.String("my-tag"),
			Value: jsii.String("my-tag-value"),
		},
	},
}))
```

### Launch type for ECS Task

By default, if `isEc2Compatible` for the `taskDefinition` is true, the EC2 type is used as
the launch type for the task, otherwise the FARGATE type.
If you want to override the default launch type, you can set the `launchType` property.

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"

var cluster iCluster
var taskDefinition taskDefinition


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	LaunchType: ecs.LaunchType_FARGATE,
}))
```

### Assign public IP addresses to tasks

You can set the `assignPublicIp` flag to assign public IP addresses to tasks.
If you want to detach the public IP address from the task, you have to set the flag `false`.
You can specify the flag `true` only when the launch type is set to FARGATE.

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"
import ec2 "github.com/aws/aws-cdk-go/awscdk"

var cluster iCluster
var taskDefinition taskDefinition


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(
targets.NewEcsTask(&EcsTaskProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	AssignPublicIp: jsii.Boolean(true),
	SubnetSelection: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
}))
```

### Enable Amazon ECS Exec for ECS Task

If you use Amazon ECS Exec, you can run commands in or get a shell to a container running on an Amazon EC2 instance or on AWS Fargate.

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"

var cluster iCluster
var taskDefinition taskDefinition


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	TaskCount: jsii.Number(1),
	ContainerOverrides: []containerOverride{
		&containerOverride{
			ContainerName: jsii.String("TheContainer"),
			Command: []*string{
				jsii.String("echo"),
				events.EventField_FromPath(jsii.String("$.detail.event")),
			},
		},
	},
	EnableExecuteCommand: jsii.Boolean(true),
}))
```

### Overriding Values in the Task Definition

You can override values in the task definition by setting the corresponding properties in the `EcsTaskProps`. All
values in the [`TaskOverrides` API](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskOverride.html) are
supported.

```go
import ecs "github.com/aws/aws-cdk-go/awscdk"

var cluster iCluster
var taskDefinition taskDefinition


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewEcsTask(&EcsTaskProps{
	Cluster: Cluster,
	TaskDefinition: TaskDefinition,
	TaskCount: jsii.Number(1),

	// Overrides the cpu and memory values in the task definition
	Cpu: jsii.String("512"),
	Memory: jsii.String("512"),
}))
```

## Schedule a Redshift query (serverless or cluster)

Use the `RedshiftQuery` target to schedule an Amazon Redshift Query.

The code snippet below creates the scheduled event rule that route events to an Amazon Redshift Query

```go
import redshiftserverless "github.com/aws/aws-cdk-go/awscdk"

var workgroup cfnWorkgroup


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

rule.AddTarget(targets.NewRedshiftQuery(workgroup.AttrWorkgroupWorkgroupArn, &RedshiftQueryProps{
	Database: jsii.String("dev"),
	DeadLetterQueue: dlq,
	Sql: []*string{
		jsii.String("SELECT * FROM foo"),
		jsii.String("SELECT * FROM baz"),
	},
}))
```

## Publish to an SNS Topic

Use the `SnsTopic` target to publish to an SNS Topic.

The code snippet below creates the scheduled event rule that publishes to an SNS Topic using a resource policy.

```go
import sns "github.com/aws/aws-cdk-go/awscdk"

var topic iTopic


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewSnsTopic(topic))
```

Alternatively, a role can be attached to the target when the rule is triggered.

```go
import sns "github.com/aws/aws-cdk-go/awscdk"

var topic iTopic


rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
})

rule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
	AuthorizeUsingRole: jsii.Boolean(true),
}))
```
