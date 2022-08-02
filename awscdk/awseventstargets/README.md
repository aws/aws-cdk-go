# Event Targets for Amazon EventBridge

This library contains integration classes to send Amazon EventBridge to any
number of supported AWS Services. Instances of these classes should be passed
to the `rule.addTarget()` method.

Currently supported are:

* [Start a CodeBuild build](#start-a-codebuild-build)
* [Start a CodePipeline pipeline](#start-a-codepipeline-pipeline)
* Run an ECS task
* [Invoke a Lambda function](#invoke-a-lambda-function)
* [Invoke a API Gateway REST API](#invoke-an-api-gateway-rest-api)
* Publish a message to an SNS topic
* Send a message to an SQS queue
* [Start a StepFunctions state machine](#start-a-stepfunctions-state-machine)
* [Queue a Batch job](#queue-a-batch-job)
* Make an AWS API call
* Put a record to a Kinesis stream
* [Log an event into a LogGroup](#log-an-event-into-a-loggroup)
* Put a record to a Kinesis Data Firehose stream
* [Put an event on an EventBridge bus](#put-an-event-on-an-eventbridge-bus)
* [Send an event to EventBridge API Destination](#invoke-an-api-destination)

See the README of the `@aws-cdk/aws-events` library for more information on
EventBridge.

## Event retry policy and using dead-letter queues

The Codebuild, CodePipeline, Lambda, StepFunctions, LogGroup, SQSQueue, SNSTopic and ECSTask targets support attaching a [dead letter queue and setting retry policies](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html). See the [lambda example](#invoke-a-lambda-function).
Use [escape hatches](https://docs.aws.amazon.com/cdk/latest/guide/cfn_layer.html) for the other target types.

## Invoke a Lambda function

Use the `LambdaFunction` target to invoke a lambda function.

The code snippet below creates an event rule with a Lambda function as a target
triggered for every events from `aws.ec2` source. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
})

rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
	eventPattern: &eventPattern{
		source: []*string{
			jsii.String("aws.ec2"),
		},
	},
})

queue := sqs.NewQueue(this, jsii.String("Queue"))

rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
	deadLetterQueue: queue,
	 // Optional: add a dead letter queue
	maxEventAge: cdk.duration.hours(jsii.Number(2)),
	 // Optional: set the maxEventAge retry policy
	retryAttempts: jsii.Number(2),
}))
```

## Log an event into a LogGroup

Use the `LogGroup` target to log your events in a CloudWatch LogGroup.

For example, the following code snippet creates an event rule with a CloudWatch LogGroup as a target.
Every events sent from the `aws.ec2` source will be sent to the CloudWatch LogGroup.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"), &logGroupProps{
	logGroupName: jsii.String("MyLogGroup"),
})

rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
	eventPattern: &eventPattern{
		source: []*string{
			jsii.String("aws.ec2"),
		},
	},
})

rule.addTarget(targets.NewCloudWatchLogGroup(logGroup))
```

A rule target input can also be specified to modify the event that is sent to the log group.
Unlike other event targets, CloudWatchLogs requires a specific input template format.

```go
// Example automatically generated from non-compiling source. May contain errors.
import logs "github.com/aws/aws-cdk-go/awscdk"
var logGroup logGroup
var rule rule


rule.addTarget(targets.NewCloudWatchLogGroup(logGroup, &logGroupProps{
	logEvent: targets.logGroupTargetInput(map[string]interface{}{
		"timestamp": events.EventField.from(jsii.String("$.time")),
		"message": events.EventField.from(jsii.String("$.detail-type")),
	}),
}))
```

If you want to use static values to overwrite the `message` make sure that you provide a `string`
value.

```go
// Example automatically generated from non-compiling source. May contain errors.
import logs "github.com/aws/aws-cdk-go/awscdk"
var logGroup logGroup
var rule rule


rule.addTarget(targets.NewCloudWatchLogGroup(logGroup, &logGroupProps{
	logEvent: targets.logGroupTargetInput(map[string]*string{
		"message": JSON.stringify(map[string]*string{
			"CustomField": jsii.String("CustomValue"),
		}),
	}),
}))
```

## Start a CodeBuild build

Use the `CodeBuildProject` target to trigger a CodeBuild project.

The code snippet below creates a CodeCommit repository that triggers a CodeBuild project
on commit to the master branch. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import codebuild "github.com/aws/aws-cdk-go/awscdk"
import codecommit "github.com/aws/aws-cdk-go/awscdk"


repo := codecommit.NewRepository(this, jsii.String("MyRepo"), &repositoryProps{
	repositoryName: jsii.String("aws-cdk-codebuild-events"),
})

project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
	source: codebuild.source.codeCommit(&codeCommitSourceProps{
		repository: repo,
	}),
})

deadLetterQueue := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

// trigger a build when a commit is pushed to the repo
onCommitRule := repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
	target: targets.NewCodeBuildProject(project, &codeBuildProjectProps{
		deadLetterQueue: deadLetterQueue,
	}),
	branches: []*string{
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

rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.expression(jsii.String("rate(1 hour)")),
})

rule.addTarget(targets.NewCodePipeline(pipeline))
```

## Start a StepFunctions state machine

Use the `SfnStateMachine` target to trigger a State Machine.

The code snippet below creates a Simple StateMachine that is triggered every minute with a
dummy object as input.
You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html)
to the target.

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import sfn "github.com/aws/aws-cdk-go/awscdk"


rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
})

dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

role := iam.NewRole(this, jsii.String("Role"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
})
stateMachine := sfn.NewStateMachine(this, jsii.String("SM"), &stateMachineProps{
	definition: sfn.NewWait(this, jsii.String("Hello"), &waitProps{
		time: sfn.waitTime.duration(cdk.*duration.seconds(jsii.Number(10))),
	}),
})

rule.addTarget(targets.NewSfnStateMachine(stateMachine, &sfnStateMachineProps{
	input: events.ruleTargetInput.fromObject(map[string]*string{
		"SomeParam": jsii.String("SomeValue"),
	}),
	deadLetterQueue: dlq,
	role: role,
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
// Example automatically generated from non-compiling source. May contain errors.
import batch "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"


jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
	"computeEnvironments": []map[string]interface{}{
		map[string]interface{}{
			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
				"managed": jsii.Boolean(false),
			}),
			"order": jsii.Number(1),
		},
	},
})

jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
	"container": map[string]repositoryImage{
		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
	},
})

queue := sqs.NewQueue(this, jsii.String("Queue"))

rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
})

rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
	deadLetterQueue: queue,
	event: events.ruleTargetInput.fromObject(map[string]*string{
		"SomeParam": jsii.String("SomeValue"),
	}),
	retryAttempts: jsii.Number(2),
	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
}))
```

## Invoke an API Gateway REST API

Use the `ApiGateway` target to trigger a REST API.

The code snippet below creates a Api Gateway REST API that is invoked every hour.

```go
import api "github.com/aws/aws-cdk-go/awscdk"
import lambda "github.com/aws/aws-cdk-go/awscdk"


rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
})

fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
	handler: jsii.String("index.handler"),
	runtime: lambda.runtime_NODEJS_14_X(),
	code: lambda.code.fromInline(jsii.String("exports.handler = e => {}")),
})

restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &lambdaRestApiProps{
	handler: fn,
})

dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))

rule.addTarget(
targets.NewApiGateway(restApi, &apiGatewayProps{
	path: jsii.String("/*/test"),
	method: jsii.String("GET"),
	stage: jsii.String("prod"),
	pathParameterValues: []*string{
		jsii.String("path-value"),
	},
	headerParameters: map[string]*string{
		"Header1": jsii.String("header1"),
	},
	queryStringParameters: map[string]*string{
		"QueryParam1": jsii.String("query-param-1"),
	},
	deadLetterQueue: dlq,
}))
```

## Invoke an API Destination

Use the `targets.ApiDestination` target to trigger an external API. You need to
create an `events.Connection` and `events.ApiDestination` as well.

The code snippet below creates an external destination that is invoked every hour.

```go
connection := events.NewConnection(this, jsii.String("Connection"), &connectionProps{
	authorization: events.authorization.apiKey(jsii.String("x-api-key"), awscdk.SecretValue.secretsManager(jsii.String("ApiSecretName"))),
	description: jsii.String("Connection with API Key x-api-key"),
})

destination := events.NewApiDestination(this, jsii.String("Destination"), &apiDestinationProps{
	connection: connection,
	endpoint: jsii.String("https://example.com"),
	description: jsii.String("Calling example.com with API key x-api-key"),
})

rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
	targets: []iRuleTarget{
		targets.NewApiDestination(destination),
	},
})
```

## Put an event on an EventBridge bus

Use the `EventBus` target to route event to a different EventBus.

The code snippet below creates the scheduled event rule that route events to an imported event bus.

```go
rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
	schedule: events.schedule.expression(jsii.String("rate(1 minute)")),
})

rule.addTarget(targets.NewEventBus(events.eventBus.fromEventBusArn(this, jsii.String("External"), jsii.String("arn:aws:events:eu-west-1:999999999999:event-bus/test-bus"))))
```
