# Metrics Facade Alpha Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## CloudWatch Metrics

This library automatically generates typed CloudWatch metric helpers for AWS resources. These metrics come in two flavors: **resource-scoped** and **unscoped**.

### Resource-Scoped Metrics

Resource-scoped metrics are created from a resource reference (e.g. a Lambda function). The resource dimensions are automatically injected, so metric methods return metrics already filtered to that specific resource.

#### Metrics Basic Usage

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/awscdkmetricsfacadealpha/awslambda"

// Works with L2 constructs
var fn Function

// Also works with L1 constructs
var cfnFunction CfnFunction

fnMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics_FromFunction(fn)

cloudwatch.NewAlarm(scope, jsii.String("ErrorAlarm"), &AlarmProps{
	Metric: fnMetrics.errors(),
	Threshold: jsii.Number(5),
	EvaluationPeriods: jsii.Number(1),
	ComparisonOperator: cloudwatch.ComparisonOperator_GREATER_THAN_THRESHOLD,
})
cfnMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics_FromFunction(cfnFunction)
```

#### Metric Features

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/awscdkmetricsfacadealpha/awslambda"

var fn Function

fnMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics_FromFunction(fn)

// FunctionName dimension is automatically injected from the function reference
metric := fnMetrics.invocations()
```

**Statistic Override**: Override default statistics via `MetricOptions`

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws-samples/dummy/awscdkmetricsfacadealpha/awslambda"

var fn Function

fnMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics_FromFunction(fn)

// Default statistic for Duration is Average
defaultMetric := fnMetrics.duration()

// Override statistic via MetricOptions
p99Metric := fnMetrics.duration(&MetricOptions{
	Statistic: jsii.String("p99"),
})
```

### Unscoped Metrics

Unscoped metrics are not tied to a specific resource. They can be created with explicit dimensions or with no dimensions for account-wide metrics.

#### Unscoped Metrics Basic Usage

```go
import "github.com/aws-samples/dummy/awscdkmetricsfacadealpha/awslambda"
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


// Explicit dimensions
fnMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics.NewFunctionMetrics(&FunctionMetricsProps{
	functionName: jsii.String("my-function"),
})
fnMetrics.errors()

// Per-resource dimensions (e.g. alias or version)
resourceMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics.NewResourceMetrics(&ResourceMetricsProps{
	functionName: jsii.String("my-function"),
	resource: jsii.String("prod"),
})
resourceMetrics.provisionedConcurrencyUtilization()

// Account-wide metrics (no dimensions)
accountMetrics := awscdkmetricsfacadealphaawslambda.LambdaMetrics.NewAccountMetrics()
accountMetrics.concurrentExecutions()
accountMetrics.unreservedConcurrentExecutions()
```

### Available Metrics

Metrics are generated from the AWS CloudWatch metrics spec database. Common examples:

**Lambda Metrics**:

* `invocations()` - Function invocation count
* `errors()` - Function error count
* `duration()` - Function execution duration
* `throttles()` - Function throttle count
* `concurrentExecutions()` - Concurrent execution count

Import metrics from service-specific modules:

```go
import "github.com/aws-samples/dummy/awscdkmetricsfacadealpha/awslambda"
```
