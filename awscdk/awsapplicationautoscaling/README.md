# AWS Auto Scaling Construct Library

**Application AutoScaling** is used to configure autoscaling for all
services other than scaling EC2 instances. For example, you will use this to
scale ECS tasks, DynamoDB capacity, Spot Fleet sizes, Comprehend document classification endpoints, Lambda function provisioned concurrency and more.

As a CDK user, you will probably not have to interact with this library
directly; instead, it will be used by other construct libraries to
offer AutoScaling features for their own constructs.

This document will describe the general autoscaling features and concepts;
your particular service may offer only a subset of these.

## AutoScaling basics

Resources can offer one or more **attributes** to autoscale, typically
representing some capacity dimension of the underlying service. For example,
a DynamoDB Table offers autoscaling of the read and write capacity of the
table proper and its Global Secondary Indexes, an ECS Service offers
autoscaling of its task count, an RDS Aurora cluster offers scaling of its
replica count, and so on.

When you enable autoscaling for an attribute, you specify a minimum and a
maximum value for the capacity. AutoScaling policies that respond to metrics
will never go higher or lower than the indicated capacity (but scheduled
scaling actions might, see below).

There are three ways to scale your capacity:

* **In response to a metric** (also known as step scaling); for example, you
  might want to scale out if the CPU usage across your cluster starts to rise,
  and scale in when it drops again.
* **By trying to keep a certain metric around a given value** (also known as
  target tracking scaling); you might want to automatically scale out an in to
  keep your CPU usage around 50%.
* **On a schedule**; you might want to organize your scaling around traffic
  flows you expect, by scaling out in the morning and scaling in in the
  evening.

The general pattern of autoscaling will look like this:

```go
var resource someScalableResource


capacity := resource.autoScaleCapacity(&caps{
	minCapacity: jsii.Number(5),
	maxCapacity: jsii.Number(100),
})
```

## Step Scaling

This type of scaling scales in and out in deterministic steps that you
configure, in response to metric values. For example, your scaling strategy
to scale in response to CPU usage might look like this:

```plaintext
 Scaling        -1          (no change)          +1       +3
            │        │                       │        │        │
            ├────────┼───────────────────────┼────────┼────────┤
            │        │                       │        │        │
CPU usage   0%      10%                     50%       70%     100%
```

(Note that this is not necessarily a recommended scaling strategy, but it's
a possible one. You will have to determine what thresholds are right for you).

You would configure it like this:

```go
var capacity scalableAttribute
var cpuUtilization metric


capacity.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
	Metric: cpuUtilization,
	ScalingSteps: []scalingInterval{
		&scalingInterval{
			Upper: jsii.Number(10),
			Change: -jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(50),
			Change: +jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(70),
			Change: +jsii.Number(3),
		},
	},

	// Change this to AdjustmentType.PercentChangeInCapacity to interpret the
	// 'change' numbers before as percentages instead of capacity counts.
	AdjustmentType: appscaling.AdjustmentType_CHANGE_IN_CAPACITY,
})
```

The AutoScaling construct library will create the required CloudWatch alarms and
AutoScaling policies for you.

### Scaling based on multiple datapoints

The Step Scaling configuration above will initiate a scaling event when a single
datapoint of the scaling metric is breaching a scaling step breakpoint. In cases
where you might want to initiate scaling actions on a larger number of datapoints
(ie in order to smooth out randomness in the metric data), you can use the
optional `evaluationPeriods` and `datapointsToAlarm` properties:

```go
var capacity scalableAttribute
var cpuUtilization metric


capacity.scaleOnMetric(jsii.String("ScaleToCPUWithMultipleDatapoints"), &basicStepScalingPolicyProps{
	Metric: cpuUtilization,
	ScalingSteps: []scalingInterval{
		&scalingInterval{
			Upper: jsii.Number(10),
			Change: -jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(50),
			Change: +jsii.Number(1),
		},
		&scalingInterval{
			Lower: jsii.Number(70),
			Change: +jsii.Number(3),
		},
	},

	// if the cpuUtilization metric has a period of 1 minute, then data points
	// in the last 10 minutes will be evaluated
	EvaluationPeriods: jsii.Number(10),

	// Only trigger a scaling action when 6 datapoints out of the last 10 are
	// breaching. If this is left unspecified, then ALL datapoints in the
	// evaluation period must be breaching to trigger a scaling action
	DatapointsToAlarm: jsii.Number(6),
})
```

## Target Tracking Scaling

This type of scaling scales in and out in order to keep a metric (typically
representing utilization) around a value you prefer. This type of scaling is
typically heavily service-dependent in what metric you can use, and so
different services will have different methods here to set up target tracking
scaling.

The following example configures the read capacity of a DynamoDB table
to be around 60% utilization:

```go
import dynamodb "github.com/aws/aws-cdk-go/awscdk"

var table table


readCapacity := table.AutoScaleReadCapacity(&EnableScalingProps{
	MinCapacity: jsii.Number(10),
	MaxCapacity: jsii.Number(1000),
})
readCapacity.ScaleOnUtilization(&UtilizationScalingProps{
	TargetUtilizationPercent: jsii.Number(60),
})
```

## Scheduled Scaling

This type of scaling is used to change capacities based on time. It works
by changing the `minCapacity` and `maxCapacity` of the attribute, and so
can be used for two purposes:

* Scale in and out on a schedule by setting the `minCapacity` high or
  the `maxCapacity` low.
* Still allow the regular scaling actions to do their job, but restrict
  the range they can scale over (by setting both `minCapacity` and
  `maxCapacity` but changing their range over time).

The following schedule expressions can be used:

* `at(yyyy-mm-ddThh:mm:ss)` -- scale at a particular moment in time
* `rate(value unit)` -- scale every minute/hour/day
* `cron(mm hh dd mm dow)` -- scale on arbitrary schedules

Of these, the cron expression is the most useful but also the most
complicated. A schedule is expressed as a cron expression. The `Schedule` class has a `cron` method to help build cron expressions.

The following example scales the fleet out in the morning, and lets natural
scaling take over at night:

```go
import "github.com/aws/aws-cdk-go/awscdk"
var resource someScalableResource


capacity := resource.autoScaleCapacity(&caps{
	minCapacity: jsii.Number(1),
	maxCapacity: jsii.Number(50),
})

capacity.scaleOnSchedule(jsii.String("PrescaleInTheMorning"), &scalingSchedule{
	Schedule: appscaling.Schedule_Cron(&CronOptions{
		Hour: jsii.String("8"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(20),
	TimeZone: awscdk.TimeZone_AMERICA_DENVER(),
})

capacity.scaleOnSchedule(jsii.String("AllowDownscalingAtNight"), &scalingSchedule{
	Schedule: appscaling.Schedule_*Cron(&CronOptions{
		Hour: jsii.String("20"),
		Minute: jsii.String("0"),
	}),
	MinCapacity: jsii.Number(1),
	TimeZone: awscdk.TimeZone_AMERICA_DENVER(),
})
```

## Examples

### Lambda Provisioned Concurrency Auto Scaling

```go
import "github.com/aws/aws-cdk-go/awscdk"

var code code


handler := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: Code,

	ReservedConcurrentExecutions: jsii.Number(2),
})

fnVer := handler.currentVersion

target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &ScalableTargetProps{
	ServiceNamespace: appscaling.ServiceNamespace_LAMBDA,
	MaxCapacity: jsii.Number(100),
	MinCapacity: jsii.Number(10),
	ResourceId: fmt.Sprintf("function:%v:%v", handler.FunctionName, fnVer.Version),
	ScalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
})

target.ScaleToTrackMetric(jsii.String("PceTracking"), &BasicTargetTrackingScalingPolicyProps{
	TargetValue: jsii.Number(0.9),
	PredefinedMetric: appscaling.PredefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
})
```

### ElastiCache Redis shards scaling with target value

```go
shardsScalableTarget := appscaling.NewScalableTarget(this, jsii.String("ElastiCacheRedisShardsScalableTarget"), &ScalableTargetProps{
	ServiceNamespace: appscaling.ServiceNamespace_ELASTICACHE,
	ScalableDimension: jsii.String("elasticache:replication-group:NodeGroups"),
	MinCapacity: jsii.Number(2),
	MaxCapacity: jsii.Number(10),
	ResourceId: jsii.String("replication-group/main-cluster"),
})

shardsScalableTarget.ScaleToTrackMetric(jsii.String("ElastiCacheRedisShardsCPUUtilization"), &BasicTargetTrackingScalingPolicyProps{
	TargetValue: jsii.Number(20),
	PredefinedMetric: appscaling.PredefinedMetric_ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION,
})
```

### SageMaker variant provisioned concurrency utilization with target value

```go
target := appscaling.NewScalableTarget(this, jsii.String("SageMakerVariantScalableTarget"), &ScalableTargetProps{
	ServiceNamespace: appscaling.ServiceNamespace_SAGEMAKER,
	ScalableDimension: jsii.String("sagemaker:variant:DesiredProvisionedConcurrency"),
	MinCapacity: jsii.Number(2),
	MaxCapacity: jsii.Number(10),
	ResourceId: jsii.String("endpoint/MyEndpoint/variant/MyVariant"),
})

target.ScaleToTrackMetric(jsii.String("SageMakerVariantProvisionedConcurrencyUtilization"), &BasicTargetTrackingScalingPolicyProps{
	TargetValue: jsii.Number(0.9),
	PredefinedMetric: appscaling.PredefinedMetric_SAGEMAKER_VARIANT_PROVISIONED_CONCURRENCY_UTILIZATION,
})
```
