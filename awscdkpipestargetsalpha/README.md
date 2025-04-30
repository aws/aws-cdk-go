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

EventBridge Pipes Targets let you create a target for an EventBridge Pipe.

For more details see the [service documentation](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html).

## Targets

Pipe targets are the end point of an EventBridge Pipe. The following targets are supported:

* `targets.ApiDestinationTarget`: [Send event source to an EventBridge API destination](#amazon-eventbridge-api-destination)
* `targets.ApiGatewayTarget`: [Send event source to an API Gateway REST API](#amazon-api-gateway-rest-api)
* `targets.CloudWatchLogsTarget`: [Send event source to a CloudWatch Logs log group](#amazon-cloudwatch-logs-log-group)
* `targets.EventBridgeTarget`: [Send event source to an EventBridge event bus](#amazon-eventbridge-event-bus)
* `targets.FirehoseTarget`: [Send event source to an Amazon Data Firehose delivery stream](#amazon-data-firehose-delivery-stream)
* `targets.KinesisTarget`: [Send event source to a Kinesis data stream](#amazon-kinesis-data-stream)
* `targets.LambdaFunction`: [Send event source to a Lambda function](#aws-lambda-function)
* `targets.SageMakerTarget`: [Send event source to a SageMaker pipeline](#amazon-sagemaker-pipeline)
* `targets.SfnStateMachine`: [Invoke a Step Functions state machine from an event source](#aws-step-functions-state-machine)
* `targets.SqsTarget`: [Send event source to an SQS queue](#amazon-sqs)

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

### Amazon API Gateway Rest API

A REST API can be used as a target for a pipe.
The REST API will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = e => {}")),
})

restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &LambdaRestApiProps{
	Handler: fn,
})
apiTarget := targets.NewApiGatewayTarget(restApi)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: apiTarget,
})
```

The input to the target REST API can be transformed:

```go
var sourceQueue queue


fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
	Handler: jsii.String("index.handler"),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Code: lambda.Code_FromInline(jsii.String("exports.handler = e => {}")),
})

restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &LambdaRestApiProps{
	Handler: fn,
})
apiTarget := targets.NewApiGatewayTarget(restApi, &ApiGatewayTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: apiTarget,
})
```

### Amazon CloudWatch Logs Log Group

A CloudWatch Logs log group can be used as a target for a pipe.
The log group will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetLogGroup logGroup


logGroupTarget := targets.NewCloudWatchLogsTarget(targetLogGroup)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: logGroupTarget,
})
```

The input to the target log group can be transformed:

```go
var sourceQueue queue
var targetLogGroup logGroup


logGroupTarget := targets.NewCloudWatchLogsTarget(targetLogGroup, &CloudWatchLogsTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: logGroupTarget,
})
```

### Amazon EventBridge Event Bus

An EventBridge event bus can be used as a target for a pipe.
The event bus will receive the (enriched/filtered) source payload.

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

### Amazon Data Firehose Delivery Stream

An Amazon Data Firehose delivery stream can be used as a target for a pipe.
The delivery stream will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetDeliveryStream deliveryStream


deliveryStreamTarget := targets.NewFirehoseTarget(targetDeliveryStream)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: deliveryStreamTarget,
})
```

The input to the target delivery stream can be transformed:

```go
var sourceQueue queue
var targetDeliveryStream deliveryStream


deliveryStreamTarget := targets.NewFirehoseTarget(targetDeliveryStream, &FirehoseTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: deliveryStreamTarget,
})
```

### Amazon Kinesis Data Stream

A Kinesis data stream can be used as a target for a pipe.
The data stream will receive the (enriched/filtered) source payload.

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

### AWS Lambda Function

A Lambda function can be used as a target for a pipe.
The Lambda function will be invoked with the (enriched/filtered) source payload.

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

The target Lambda function is invoked synchronously by default. You can also choose to invoke the Lambda Function asynchronously by setting `invocationType` property to `FIRE_AND_FORGET`.

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

### Amazon SageMaker Pipeline

A SageMaker pipeline can be used as a target for a pipe.
The pipeline will receive the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetPipeline iPipeline


pipelineTarget := targets.NewSageMakerTarget(targetPipeline)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipelineTarget,
})
```

The input to the target pipeline can be transformed:

```go
var sourceQueue queue
var targetPipeline iPipeline


pipelineTarget := targets.NewSageMakerTarget(targetPipeline, &SageMakerTargetParameters{
	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
		"body": jsii.String("👀"),
	}),
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
	Target: pipelineTarget,
})
```

### AWS Step Functions State Machine

A Step Functions state machine can be used as a target for a pipe.
The state machine will be invoked with the (enriched/filtered) source payload.

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

You can specify the invocation type when the target state machine is invoked:

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

The input to the target state machine can be transformed:

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

### Amazon SQS Queue

An SQS queue can be used as a target for a pipe.
The queue will receive the (enriched/filtered) source payload.

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
