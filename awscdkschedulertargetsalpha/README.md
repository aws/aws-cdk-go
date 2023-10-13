# Amazon EventBridge Scheduler Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

[Amazon EventBridge Scheduler](https://aws.amazon.com/blogs/compute/introducing-amazon-eventbridge-scheduler/) is a feature from Amazon EventBridge
that allows you to create, run, and manage scheduled tasks at scale. With EventBridge Scheduler, you can schedule one-time or recurrently tens
of millions of tasks across many AWS services without provisioning or managing underlying infrastructure.

This library contains integration classes for Amazon EventBridge Scheduler to call any
number of supported AWS Services.

The following targets are supported:

1. `targets.LambdaInvoke`: [Invoke an AWS Lambda function](#invoke-a-lambda-function))
2. `targets.StepFunctionsStartExecution`: [Start an AWS Step Function](#start-an-aws-step-function)

## Invoke a Lambda function

Use the `LambdaInvoke` target to invoke a lambda function.

The code snippet below creates an event rule with a Lambda function as a target
called every hour by Event Bridge Scheduler with custom payload. You can optionally attach a
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
called every hour by Event Bridge Scheduler with a custom payload.

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
