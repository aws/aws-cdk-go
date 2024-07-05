# Amazon EventBridge Pipes Enrichments Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

EventBridge Pipes Enrichments let you create enrichments for an EventBridge Pipe.

For more details see the service documentation:

[Documentation](https://docs.aws.amazon.com/eventbridge/latest/userguide/pipes-enrichment.html)

## Pipe sources

Pipe enrichments are invoked prior to sending the events to a target of a EventBridge Pipe.

### Lambda function

A Lambda function can be used to enrich events of a pipe.

```go
var sourceQueue queue
var targetQueue queue

var enrichmentFunction function


enrichment := enrichments.NewLambdaEnrichment(enrichmentFunction)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: NewSomeSource(sourceQueue),
	Enrichment: Enrichment,
	Target: NewSomeTarget(targetQueue),
})
```

### Step Functions state machine

Step Functions state machine can be used to enrich events of a pipe.

**Note:** EventBridge Pipes only supports Express workflows invoked synchronously.

> Visit [Amazon EventBridge Pipes event enrichment](https://docs.aws.amazon.com/eventbridge/latest/userguide/pipes-enrichment.html) for more details.

```go
var sourceQueue queue
var targetQueue queue

var enrichmentStateMachine stateMachine


enrichment := enrichments.NewStepFunctionsEnrichment(enrichmentStateMachine)

pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
	Source: NewSomeSource(sourceQueue),
	Enrichment: Enrichment,
	Target: NewSomeTarget(targetQueue),
})
```
