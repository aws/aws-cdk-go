# Amazon EventBridge Construct Library

Amazon EventBridge delivers a near real-time stream of system events that
describe changes in AWS resources. For example, an AWS CodePipeline emits the
[State
Change](https://docs.aws.amazon.com/eventbridge/latest/userguide/event-types.html#codepipeline-event-type)
event when the pipeline changes its state.

* **Events**: An event indicates a change in your AWS environment. AWS resources
  can generate events when their state changes. For example, Amazon EC2
  generates an event when the state of an EC2 instance changes from pending to
  running, and Amazon EC2 Auto Scaling generates events when it launches or
  terminates instances. AWS CloudTrail publishes events when you make API calls.
  You can generate custom application-level events and publish them to
  EventBridge. You can also set up scheduled events that are generated on
  a periodic basis. For a list of services that generate events, and sample
  events from each service, see [EventBridge Event Examples From Each
  Supported
  Service](https://docs.aws.amazon.com/eventbridge/latest/userguide/event-types.html).
* **Targets**: A target processes events. Targets can include Amazon EC2
  instances, AWS Lambda functions, Kinesis streams, Amazon ECS tasks, Step
  Functions state machines, Amazon SNS topics, Amazon SQS queues, Amazon CloudWatch LogGroups, and built-in
  targets. A target receives events in JSON format.
* **Rules**: A rule matches incoming events and routes them to targets for
  processing. A single rule can route to multiple targets, all of which are
  processed in parallel. Rules are not processed in a particular order. This
  enables different parts of an organization to look for and process the events
  that are of interest to them. A rule can customize the JSON sent to the
  target, by passing only certain parts or by overwriting it with a constant.
* **EventBuses**: An event bus can receive events from your own custom applications
  or it can receive events from applications and services created by AWS SaaS partners.
  See [Creating an Event Bus](https://docs.aws.amazon.com/eventbridge/latest/userguide/create-event-bus.html).

## Rule

The `Rule` construct defines an EventBridge rule which monitors an
event based on an [event
pattern](https://docs.aws.amazon.com/eventbridge/latest/userguide/filtering-examples-structure.html)
and invoke **event targets** when the pattern is matched against a triggered
event. Event targets are objects that implement the `IRuleTarget` interface.

Normally, you will use one of the `source.onXxx(name[, target[, options]]) -> Rule` methods on the event source to define an event rule associated with
the specific activity. You can targets either via props, or add targets using
`rule.addTarget`.

For example, to define an rule that triggers a CodeBuild project build when a
commit is pushed to the "master" branch of a CodeCommit repository:

```go
var repo repository
var project project


onCommitRule := repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
	target: targets.NewCodeBuildProject(project),
	branches: []*string{
		jsii.String("master"),
	},
})
```

You can add additional targets, with optional [input
transformer](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_InputTransformer.html)
using `eventRule.addTarget(target[, input])`. For example, we can add a SNS
topic target which formats a human-readable message for the commit.

For example, this adds an SNS topic as a target:

```go
var onCommitRule rule
var topic topic


onCommitRule.addTarget(targets.NewSnsTopic(topic, &snsTopicProps{
	message: events.ruleTargetInput.fromText(
	fmt.Sprintf("A commit was pushed to the repository %v on branch %v", codecommit.referenceEvent.repositoryName, codecommit.*referenceEvent.referenceName)),
}))
```

Or using an Object:

```go
var onCommitRule rule
var topic topic


onCommitRule.addTarget(targets.NewSnsTopic(topic, &snsTopicProps{
	message: events.ruleTargetInput.fromObject(map[string]*string{
		"DataType": fmt.Sprintf("custom_%v", events.EventField.fromPath(jsii.String("$.detail-type"))),
	}),
}))
```

To define a pattern, use the matcher API, which provides a number of factory methods to declare different logical predicates. For example, to match all S3 events for objects larger than 1024 bytes, stored using one of the storage classes Glacier, Glacier IR or Deep Archive and coming from any region other than the AWS GovCloud ones:

```go
// Example automatically generated from non-compiling source. May contain errors.
rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
	eventPattern: &eventPattern{
		detail: map[string]interface{}{
			"object": map[string][]*string{
				// Matchers may appear at any level
				"size": events.Match.greaterThan(jsii.Number(1024)),
			},

			// 'OR' condition
			"source-storage-class": events.Match.anyOf(events.Match.prefix(jsii.String("GLACIER")), events.Match.exactString(jsii.String("DEEP_ARCHIVE"))),
		},
		detailtype: events.match.equalsIgnoreCase(jsii.String("object created")),

		// If you prefer, you can use a low level array of strings, as directly consumed by EventBridge
		source: []*string{
			jsii.String("aws.s3"),
		},

		region: events.*match.anythingButPrefix(jsii.String("us-gov")),
	},
})
```

## Scheduling

You can configure a Rule to run on a schedule (cron or rate).
Rate must be specified in minutes, hours or days.

The following example runs a task every day at 4am:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"

var cluster cluster
var taskDefinition taskDefinition
var role role


ecsTaskTarget := awscdk.NewEcsTask(&ecsTaskProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	role: role,
})

awscdk.NewRule(this, jsii.String("ScheduleRule"), &ruleProps{
	schedule: awscdk.Schedule.cron(&cronOptions{
		minute: jsii.String("0"),
		hour: jsii.String("4"),
	}),
	targets: []iRuleTarget{
		ecsTaskTarget,
	},
})
```

If you want to specify Fargate platform version, set `platformVersion` in EcsTask's props like the following example:

```go
var cluster cluster
var taskDefinition taskDefinition
var role role


platformVersion := ecs.fargatePlatformVersion_VERSION1_4
ecsTaskTarget := targets.NewEcsTask(&ecsTaskProps{
	cluster: cluster,
	taskDefinition: taskDefinition,
	role: role,
	platformVersion: platformVersion,
})
```

## Event Targets

The `@aws-cdk/aws-events-targets` module includes classes that implement the `IRuleTarget`
interface for various AWS services.

The following targets are supported:

* `targets.CodeBuildProject`: Start an AWS CodeBuild build
* `targets.CodePipeline`: Start an AWS CodePipeline pipeline execution
* `targets.EcsTask`: Start a task on an Amazon ECS cluster
* `targets.LambdaFunction`: Invoke an AWS Lambda function
* `targets.SnsTopic`: Publish into an SNS topic
* `targets.SqsQueue`: Send a message to an Amazon SQS Queue
* `targets.SfnStateMachine`: Trigger an AWS Step Functions state machine
* `targets.BatchJob`: Queue an AWS Batch Job
* `targets.AwsApi`: Make an AWS API call
* `targets.ApiGateway`: Invoke an AWS API Gateway
* `targets.ApiDestination`: Make an call to an external destination

### Cross-account and cross-region targets

It's possible to have the source of the event and a target in separate AWS accounts and regions:

```go
import "github.com/aws/aws-cdk-go/awscdk"
import codebuild "github.com/aws/aws-cdk-go/awscdk"
import codecommit "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

app := awscdk.NewApp()

account1 := "11111111111"
account2 := "22222222222"

stack1 := awscdk.NewStack(app, jsii.String("Stack1"), &stackProps{
	env: &environment{
		account: account1,
		region: jsii.String("us-west-1"),
	},
})
repo := codecommit.NewRepository(stack1, jsii.String("Repository"), &repositoryProps{
	repositoryName: jsii.String("myrepository"),
})

stack2 := awscdk.NewStack(app, jsii.String("Stack2"), &stackProps{
	env: &environment{
		account: account2,
		region: jsii.String("us-east-1"),
	},
})
project := codebuild.NewProject(stack2, jsii.String("Project"), &projectProps{
})

repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
	target: targets.NewCodeBuildProject(project),
})
```

In this situation, the CDK will wire the 2 accounts together:

* It will generate a rule in the source stack with the event bus of the target account as the target
* It will generate a rule in the target stack, with the provided target
* It will generate a separate stack that gives the source account permissions to publish events
  to the event bus of the target account in the given region,
  and make sure its deployed before the source stack

For more information, see the
[AWS documentation on cross-account events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html).

## Archiving

It is possible to archive all or some events sent to an event bus. It is then possible to [replay these events](https://aws.amazon.com/blogs/aws/new-archive-and-replay-events-with-amazon-eventbridge/).

```go
bus := events.NewEventBus(this, jsii.String("bus"), &eventBusProps{
	eventBusName: jsii.String("MyCustomEventBus"),
})

bus.archive(jsii.String("MyArchive"), &baseArchiveProps{
	archiveName: jsii.String("MyCustomEventBusArchive"),
	description: jsii.String("MyCustomerEventBus Archive"),
	eventPattern: &eventPattern{
		account: []*string{
			awscdk.*stack.of(this).account,
		},
	},
	retention: awscdk.Duration.days(jsii.Number(365)),
})
```

## Granting PutEvents to an existing EventBus

To import an existing EventBus into your CDK application, use `EventBus.fromEventBusArn`, `EventBus.fromEventBusAttributes`
or `EventBus.fromEventBusName` factory method.

Then, you can use the `grantPutEventsTo` method to grant `event:PutEvents` to the eventBus.

```go
var lambdaFunction function


eventBus := events.eventBus.fromEventBusArn(this, jsii.String("ImportedEventBus"), jsii.String("arn:aws:events:us-east-1:111111111:event-bus/my-event-bus"))

// now you can just call methods on the eventbus
eventBus.grantPutEventsTo(lambdaFunction)
```
