# Amazon EventBridge Pipes Targets Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

EventBridge Pipes Targets let you create a target for a EventBridge Pipe.

For more details see the service documentation:

[Documentation](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html)

## Targets

Pipe targets are the end point of a EventBridge Pipe.

The following targets are supported:

1. `targets.SqsTarget`: [Send event source to an SQS queue](#amazon-sqs)
2. `targets.SfnStateMachine`: [Invoke a State Machine from an event source](#aws-step-functions-state-machine)
3. `targets.LambdaFunction`: [Send event source to a Lambda function](#aws-lambda-function)
4. `targets.ApiDestinationTarget`: [Send event source to an EventBridge API destination](#amazon-eventbridge-api-destination)
5. `targets.KinesisTarget`: [Send event source to a Kinesis data stream](#amazon-kinesis-data-stream)
6. `targets.EventBridgeTarget`: [Send event source to an EventBridge event bus](#amazon-eventbridge-event-bus)

### Amazon SQS

A SQS message queue can be used as a target for a pipe. Messages will be pushed to the queue.

```go
var sourceQueue queue
var targetQueue queue


pipeTarget := targets.NewSqsTarget(targetQueue)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

The target input can be transformed:

```go
var sourceQueue queue
var targetQueue queue


pipeTarget := targets.NewSqsTarget(targetQueue, &SqsTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"SomeKey": pipes.DynamicInput_fromEventPath(jsii.String("$.body")),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

### AWS Step Functions State Machine

A State Machine can be used as a target for a pipe. The State Machine will be invoked with the (enriched) source payload.

```go
var sourceQueue queue
var targetStateMachine iStateMachine


pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

Specifying the Invocation Type when the target State Machine is invoked:

```go
var sourceQueue queue
var targetStateMachine iStateMachine


pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
	InvocationType: targets.StateMachineInvocationType_FIRE_AND_FORGET,
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

The input to the target State Machine can be transformed:

```go
var sourceQueue queue
var targetStateMachine iStateMachine


pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("<$.body>"),
	}),
	InvocationType: targets.StateMachineInvocationType_FIRE_AND_FORGET,
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

### AWS Lambda Function

A Lambda Function can be used as a target for a pipe. The Lambda Function will be invoked with the (enriched) source payload.

```go
var sourceQueue queue
var targetFunction iFunction


pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

The target Lambda Function is invoked synchronously by default. You can also choose to invoke the Lambda Function asynchronously by setting `invocationType` property to `FIRE_AND_FORGET`.

```go
var sourceQueue queue
var targetFunction iFunction


pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
	InvocationType: targets.LambdaFunctionInvocationType_FIRE_AND_FORGET,
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

The input to the target Lambda Function can be transformed:

```go
var sourceQueue queue
var targetFunction iFunction


pipeTarget := targets.NewLambdaFunction(targetFunction, &LambdaFunctionParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipeTarget,
})
```

### Amazon EventBridge API Destination

An EventBridge API destination can be used as a target for a pipe.
The API destination will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var dest apiDestination


apiTarget := targets.NewApiDestinationTarget(dest)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: apiTarget,
})
```

The input to the target API destination can be transformed:

```go
var sourceQueue queue
var dest apiDestination


apiTarget := targets.NewApiDestinationTarget(dest, &ApiDestinationTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: apiTarget,
})
```

### Amazon Kinesis Data Stream

A data stream can be used as a target for a pipe. The data stream will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetStream stream


streamTarget := targets.NewKinesisTarget(targetStream, &KinesisTargetParameters{
	PartitionKey: jsii.String("pk"),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: streamTarget,
})
```

The input to the target data stream can be transformed:

```go
var sourceQueue queue
var targetStream stream


streamTarget := targets.NewKinesisTarget(targetStream, &KinesisTargetParameters{
	PartitionKey: jsii.String("pk"),
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: streamTarget,
})
```

### Amazon EventBridge Event Bus

An event bus can be used as a target for a pipe. The event bus will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetEventBus eventBus


eventBusTarget := targets.NewEventBridgeTarget(targetEventBus)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: eventBusTarget,
})
```

The input to the target event bus can be transformed:

```go
var sourceQueue queue
var targetEventBus eventBus


eventBusTarget := targets.NewEventBridgeTarget(targetEventBus, &EventBridgeTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: eventBusTarget,
})
```
