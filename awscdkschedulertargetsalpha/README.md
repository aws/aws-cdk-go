# Amazon EventBridge Scheduler Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Developer Preview](https://img.shields.io/badge/cdk--constructs-developer--preview-informational.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are in **developer preview** before they
> become stable. We will only make breaking changes to address unforeseen API issues. Therefore,
> these APIs are not subject to [Semantic Versioning](https://semver.org/), and breaking changes
> will be announced in release notes. This means that while you may use them, you may need to
> update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[Amazon EventBridge Scheduler](https://aws.amazon.com/blogs/compute/introducing-amazon-eventbridge-scheduler/) is a feature from Amazon EventBridge
that allows you to create, run, and manage scheduled tasks at scale. With EventBridge Scheduler, you can schedule millions of one-time or recurring tasks across various AWS services without provisioning or managing underlying infrastructure.

This library contains integration classes for Amazon EventBridge Scheduler to call any
number of supported AWS Services.

The following targets are supported:

1. `targets.LambdaInvoke`: [Invoke an AWS Lambda function](#invoke-a-lambda-function)
2. `targets.StepFunctionsStartExecution`: [Start an AWS Step Function](#start-an-aws-step-function)
3. `targets.CodeBuildStartBuild`: [Start a CodeBuild job](#start-a-codebuild-job)
4. `targets.SqsSendMessage`: [Send a Message to an Amazon SQS Queue](#send-a-message-to-an-sqs-queue)
5. `targets.SnsPublish`: [Publish messages to an Amazon SNS topic](#publish-messages-to-an-amazon-sns-topic)
6. `targets.EventBridgePutEvents`: [Put Events on EventBridge](#send-events-to-an-eventbridge-event-bus)
7. `targets.InspectorStartAssessmentRun`: [Start an Amazon Inspector assessment run](#start-an-amazon-inspector-assessment-run)
8. `targets.KinesisStreamPutRecord`: [Put a record to an Amazon Kinesis Data Stream](#put-a-record-to-an-amazon-kinesis-data-stream)
9. `targets.FirehosePutRecord`: [Put a record to an Amazon Data Firehose](#put-a-record-to-an-amazon-data-firehose)
10. `targets.CodePipelineStartPipelineExecution`: [Start a CodePipeline execution](#start-a-codepipeline-execution)
11. `targets.SageMakerStartPipelineExecution`: [Start a SageMaker pipeline execution](#start-a-sagemaker-pipeline-execution)
12. `targets.Universal`: [Invoke a wider set of AWS API](#invoke-a-wider-set-of-aws-api)

## Invoke a Lambda function

Use the `LambdaInvoke` target to invoke a lambda function.

The code snippet below creates an event rule with a Lambda function as a target
called every hour by EventBridge Scheduler with a custom payload. You can optionally attach a
[dead letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html).

```go
import "github.com/aws/aws-cdk-go/awscdk"


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
})

dlq := sqs.NewQueue(this, jsii.String("DLQ"), &QueueProps{
	QueueName: jsii.String("MyDLQ"),
})

target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	DeadLetterQueue: dlq,
	MaxEventAge: awscdk.Duration_Minutes(jsii.Number(1)),
	RetryAttempts: jsii.Number(3),
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"payload": jsii.String("useful"),
	}),
})

schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: Target,
})
```

## Start an AWS Step Function

Use the `StepFunctionsStartExecution` target to start a new execution on a StepFunction.

The code snippet below creates an event rule with a Step Function as a target
called every hour by EventBridge Scheduler with a custom payload.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import tasks "github.com/aws/aws-cdk-go/awscdk"


payload := map[string]*string{
	"Name": jsii.String("MyParameter"),
	"Value": jsii.String("🌥️"),
}

putParameterStep := tasks.NewCallAwsService(this, jsii.String("PutParameter"), &CallAwsServiceProps{
	Service: jsii.String("ssm"),
	Action: jsii.String("putParameter"),
	IamResources: []*string{
		jsii.String("*"),
	},
	Parameters: map[string]interface{}{
		"Name.$": jsii.String("$.Name"),
		"Value.$": jsii.String("$.Value"),
		"Type": jsii.String("String"),
		"Overwrite": jsii.Boolean(true),
	},
})

stateMachine := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
	DefinitionBody: sfn.DefinitionBody_FromChainable(putParameterStep),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: targets.NewStepFunctionsStartExecution(stateMachine, &ScheduleTargetBaseProps{
		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
	}),
})
```

## Start a CodeBuild job

Use the `CodeBuildStartBuild` target to start a new build run on a CodeBuild project.

The code snippet below creates an event rule with a CodeBuild project as target which is
called every hour by EventBridge Scheduler.

```go
import codebuild "github.com/aws/aws-cdk-go/awscdk"

var project project


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewCodeBuildStartBuild(project),
})
```

## Send a Message To an SQS Queue

Use the `SqsSendMessage` target to send a message to an SQS Queue.

The code snippet below creates an event rule with an SQS Queue as a target
called every hour by EventBridge Scheduler with a custom payload.

Contains the `messageGroupId` to use when the target is a FIFO queue. If you specify
a FIFO queue as a target, the queue must have content-based deduplication enabled.

```go
payload := "test"
messageGroupId := "id"
queue := sqs.NewQueue(this, jsii.String("MyQueue"), &QueueProps{
	Fifo: jsii.Boolean(true),
	ContentBasedDeduplication: jsii.Boolean(true),
})

target := targets.NewSqsSendMessage(queue, &SqsSendMessageProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromText(payload),
	MessageGroupId: jsii.String(MessageGroupId),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
	Target: Target,
})
```

## Publish messages to an Amazon SNS topic

Use the `SnsPublish` target to publish messages to an Amazon SNS topic.

The code snippets below create an event rule with a Amazon SNS topic as a target.
It's called every hour by Amazon EventBridge Scheduler with a custom payload.

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


topic := sns.NewTopic(this, jsii.String("Topic"))

payload := map[string]*string{
	"message": jsii.String("Hello scheduler!"),
}

target := targets.NewSnsPublish(topic, &ScheduleTargetBaseProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: Target,
})
```

## Send events to an EventBridge event bus

Use the `EventBridgePutEvents` target to send events to an EventBridge event bus.

The code snippet below creates an event rule with an EventBridge event bus as a target
called every hour by EventBridge Scheduler with a custom event payload.

```go
import events "github.com/aws/aws-cdk-go/awscdk"


eventBus := events.NewEventBus(this, jsii.String("EventBus"), &EventBusProps{
	EventBusName: jsii.String("DomainEvents"),
})

eventEntry := &EventBridgePutEventsEntry{
	EventBus: EventBus,
	Source: jsii.String("PetService"),
	Detail: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"Name": jsii.String("Fluffy"),
	}),
	DetailType: jsii.String("🐶"),
}

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
	Target: targets.NewEventBridgePutEvents(eventEntry),
})
```

## Start an Amazon Inspector assessment run

Use the `InspectorStartAssessmentRun` target to start an Inspector assessment run.

The code snippet below creates an event rule with an assessment template as the target which is
called every hour by EventBridge Scheduler.

```go
import inspector "github.com/aws/aws-cdk-go/awscdk"

var cfnAssessmentTemplate cfnAssessmentTemplate


assessmentTemplate := inspector.AssessmentTemplate_FromCfnAssessmentTemplate(this, jsii.String("MyAssessmentTemplate"), cfnAssessmentTemplate)

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewInspectorStartAssessmentRun(assessmentTemplate),
})
```

## Put a record to an Amazon Kinesis Data Stream

Use the `KinesisStreamPutRecord` target to put a record to an Amazon Kinesis Data Stream.

The code snippet below creates an event rule with a stream as the target which is
called every hour by EventBridge Scheduler.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"


stream := kinesis.NewStream(this, jsii.String("MyStream"))

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewKinesisStreamPutRecord(stream, &KinesisStreamPutRecordProps{
		PartitionKey: jsii.String("key"),
	}),
})
```

## Put a record to an Amazon Data Firehose

Use the `FirehosePutRecord` target to put a record to an Amazon Data Firehose delivery stream.

The code snippet below creates an event rule with a delivery stream as a target
called every hour by EventBridge Scheduler with a custom payload.

```go
import firehose "github.com/aws/aws-cdk-go/awscdk"
var deliveryStream iDeliveryStream


payload := map[string]*string{
	"Data": jsii.String("record"),
}

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewFirehosePutRecord(deliveryStream, &ScheduleTargetBaseProps{
		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
	}),
})
```

## Start a CodePipeline execution

Use the `CodePipelineStartPipelineExecution` target to start a new execution for a CodePipeline pipeline.

The code snippet below creates an event rule with a CodePipeline pipeline as the target which is
called every hour by EventBridge Scheduler.

```go
import codepipeline "github.com/aws/aws-cdk-go/awscdk"

var pipeline pipeline


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewCodePipelineStartPipelineExecution(pipeline),
})
```

## Start a SageMaker pipeline execution

Use the `SageMakerStartPipelineExecution` target to start a new execution for a SageMaker pipeline.

The code snippet below creates an event rule with a SageMaker pipeline as the target which is
called every hour by EventBridge Scheduler.

```go
import sagemaker "github.com/aws/aws-cdk-go/awscdk"

var pipeline iPipeline


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewSageMakerStartPipelineExecution(pipeline, &SageMakerStartPipelineExecutionProps{
		PipelineParameterList: []sageMakerPipelineParameter{
			&sageMakerPipelineParameter{
				Name: jsii.String("parameter-name"),
				Value: jsii.String("parameter-value"),
			},
		},
	}),
})
```

## Invoke a wider set of AWS API

Use the `Universal` target to invoke AWS API. See [https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html)

The code snippet below creates an event rule with AWS API as the target which is
called at midnight every day by EventBridge Scheduler.

```go
awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Cron(&CronOptionsWithTimezone{
		Minute: jsii.String("0"),
		Hour: jsii.String("0"),
	}),
	Target: targets.NewUniversal(&UniversalTargetProps{
		Service: jsii.String("rds"),
		Action: jsii.String("stopDBCluster"),
		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
			"DbClusterIdentifier": jsii.String("my-db"),
		}),
	}),
})
```

The `service` must be in lowercase and the `action` must be in camelCase.

By default, an IAM policy for the Scheduler is extracted from the API call. The action in the policy is constructed using the `service` and `action` prop.
Re-using the example above, the action will be `rds:stopDBCluster`. Note that not all IAM actions follow the same pattern. In such scenario, please use the
`policyStatements` prop to override the policy:

```go
awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
	Target: targets.NewUniversal(&UniversalTargetProps{
		Service: jsii.String("sqs"),
		Action: jsii.String("sendMessage"),
		PolicyStatements: []policyStatement{
			iam.NewPolicyStatement(&PolicyStatementProps{
				Actions: []*string{
					jsii.String("sqs:SendMessage"),
				},
				Resources: []*string{
					jsii.String("arn:aws:sqs:us-east-1:123456789012:my_queue"),
				},
			}),
			iam.NewPolicyStatement(&PolicyStatementProps{
				Actions: []*string{
					jsii.String("kms:Decrypt"),
					jsii.String("kms:GenerateDataKey*"),
				},
				Resources: []*string{
					jsii.String("arn:aws:kms:us-east-1:123456789012:key/0987dcba-09fe-87dc-65ba-ab0987654321"),
				},
			}),
		},
	}),
})
```

> Note: The default policy uses `*` in the resources field as CDK does not have a straight forward way to auto-discover the resources permission required.
> It is recommended that you scope the field down to specific resources to have a better security posture.
