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

1. `targets.SqsTarget`: [Send event source to a Queue](#amazon-sqs)
2. `targets.SfnStateMachine`: [Invoke a State Machine from an event source](#aws-step-functions)

### Amazon SQS

A SQS message queue can be used as a target for a pipe. Messages will be pushed to the queue.

```go
var sourceQueue queue
var targetQueue queue


pipeTarget := targets.NewSqsTarget(targetQueue)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: NewSomeSource(sourceQueue),
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
	Source: NewSomeSource(sourceQueue),
	Target: pipeTarget,
})
```

### AWS Step Functions State Machine

A State Machine can be used as a target for a pipe. The State Machine will be invoked with the (enriched/filtered) source payload.

```go
var sourceQueue queue
var targetStateMachine iStateMachine


pipeTarget := targets.NewSfnStateMachine(targetStateMachine, &SfnStateMachineParameters{
})

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: NewSomeSource(sourceQueue),
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
	Source: NewSomeSource(sourceQueue),
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
	Source: NewSomeSource(sourceQueue),
	Target: pipeTarget,
})
```
