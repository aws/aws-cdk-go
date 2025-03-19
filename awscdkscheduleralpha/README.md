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

1. **Schedule**: A schedule is the main resource you create, configure, and manage using Amazon EventBridge Scheduler. Every schedule has a schedule expression that determines when, and with what frequency, the schedule runs. EventBridge Scheduler supports three types of schedules: rate, cron, and one-time schedules. When you create a schedule, you configure a target for the schedule to invoke.
2. **Target**: A target is an API operation that EventBridge Scheduler calls on your behalf every time your schedule runs. EventBridge Scheduler
   supports two types of targets: templated targets and universal targets. Templated targets invoke common API operations across a core groups of
   services. For example, EventBridge Scheduler supports templated targets for invoking AWS Lambda Function or starting execution of Step Functions state
   machine. For API operations that are not supported by templated targets you can use customizable universal targets. Universal targets support calling
   more than 6,000 API operations across over 270 AWS services.
3. **Schedule Group**: A schedule group is an Amazon EventBridge Scheduler resource that you use to organize your schedules. Your AWS account comes
   with a default scheduler group. A new schedule will always be added to a scheduling group. If you do not provide a scheduling group to add to, it
   will be added to the default scheduling group. You can create up to 500 schedule groups in your AWS account. Groups can be used to organize the
   schedules logically, access the schedule metrics and manage permissions at group granularity (see details below). Schedule groups support tagging.
   With EventBridge Scheduler, you apply tags to schedule groups, not to individual schedules to organize your resources.

## Defining a schedule

```go
var fn function


target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"payload": jsii.String("useful"),
	}),
})

schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Target: Target,
	Description: jsii.String("This is a test schedule that invokes a lambda function every 10 minutes."),
})
```

### Schedule Expressions

You can choose from three schedule types when configuring your schedule: rate-based, cron-based, and one-time schedules.

Both rate-based and cron-based schedules are recurring schedules. You can configure each recurring schedule type using a schedule expression.

For
cron-based schedules you can specify a time zone in which EventBridge Scheduler evaluates the expression.

```go
var target lambdaInvoke


rateBasedSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Target: Target,
	Description: jsii.String("This is a test rate-based schedule"),
})

cronBasedSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Cron(&CronOptionsWithTimezone{
		Minute: jsii.String("0"),
		Hour: jsii.String("23"),
		Day: jsii.String("20"),
		Month: jsii.String("11"),
		TimeZone: awscdk.TimeZone_AMERICA_NEW_YORK(),
	}),
	Target: Target,
	Description: jsii.String("This is a test cron-based schedule that will run at 11:00 PM, on day 20 of the month, only in November in New York timezone"),
})
```

A one-time schedule is a schedule that invokes a target only once. You configure a one-time schedule by specifying the time of day, date,
and time zone in which EventBridge Scheduler evaluates the schedule.

```go
var target lambdaInvoke


oneTimeSchedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_At(
	NewDate(jsii.Number(2022), jsii.Number(10), jsii.Number(20), jsii.Number(19), jsii.Number(20), jsii.Number(23)), awscdk.TimeZone_AMERICA_NEW_YORK()),
	Target: Target,
	Description: jsii.String("This is a one-time schedule in New York timezone"),
})
```

### Grouping Schedules

Your AWS account comes with a default scheduler group. You can access the default group in CDK with:

```go
defaultScheduleGroup := awscdkscheduleralpha.ScheduleGroup_FromDefaultScheduleGroup(this, jsii.String("DefaultGroup"))
```

You can add a schedule to a custom scheduling group managed by you. If a custom group is not specified, the schedule is added to the default group.

```go
var target lambdaInvoke


scheduleGroup := awscdkscheduleralpha.NewScheduleGroup(this, jsii.String("ScheduleGroup"), &ScheduleGroupProps{
	ScheduleGroupName: jsii.String("MyScheduleGroup"),
})

awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Target: Target,
	ScheduleGroup: ScheduleGroup,
})
```

### Disabling Schedules

By default, a schedule will be enabled. You can disable a schedule by setting the `enabled` property to false:

```go
var target lambdaInvoke


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Target: target,
	Enabled: jsii.Boolean(false),
})
```

### Configuring a start and end time of the Schedule

If you choose a recurring schedule, you can set the start and end time of the Schedule by specifying the `start` and `end`.

```go
var target lambdaInvoke


awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(cdk.Duration_Hours(jsii.Number(12))),
	Target: target,
	Start: NewDate(jsii.String("2023-01-01T00:00:00.000Z")),
	End: NewDate(jsii.String("2023-02-01T00:00:00.000Z")),
})
```

## Scheduler Targets

The `@aws-cdk/aws-scheduler-targets-alpha` module includes classes that implement the `IScheduleTarget` interface for
various AWS services. EventBridge Scheduler supports two types of targets:

1. **Templated targets** which invoke common API
   operations across a core groups of services, and
2. **Universal targets** that you can customize to call more
   than 6,000 operations across over 270 services.

A list of supported targets can be found at `@aws-cdk/aws-scheduler-targets-alpha`.

### Input

Targets can be invoked with a custom input. The `ScheduleTargetInput` class supports free-form text input and JSON-formatted object input:

```go
input := awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
	"QueueName": jsii.String("MyQueue"),
})
```

You can include context attributes in your target payload. EventBridge Scheduler will replace each keyword with
its respective value and deliver it to the target. See
[full list of supported context attributes](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-context-attributes.html):

1. `ContextAttribute.scheduleArn()` – The ARN of the schedule.
2. `ContextAttribute.scheduledTime()` – The time you specified for the schedule to invoke its target, e.g., 2022-03-22T18:59:43Z.
3. `ContextAttribute.executionId()` – The unique ID that EventBridge Scheduler assigns for each attempted invocation of a target, e.g., d32c5kddcf5bb8c3.
4. `ContextAttribute.attemptNumber()` – A counter that identifies the attempt number for the current invocation, e.g., 1.

```go
text := fmt.Sprintf("Attempt number: %v", awscdkscheduleralpha.ContextAttribute_AttemptNumber())
input := awscdkscheduleralpha.ScheduleTargetInput_FromText(text)
```

### Specifying an execution role

An execution role is an IAM role that EventBridge Scheduler assumes in order to interact with other AWS services on your behalf.

The classes for templated schedule targets automatically create an IAM role with all the minimum necessary
permissions to interact with the templated target. If you wish you may specify your own IAM role, then the templated targets
will grant minimal required permissions. For example, the `LambdaInvoke` target will grant the
IAM execution role `lambda:InvokeFunction` permission to invoke the Lambda function.

```go
var fn function


role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("scheduler.amazonaws.com")),
})

target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"payload": jsii.String("useful"),
	}),
	Role: Role,
})
```

### Specifying an encryption key

EventBridge Scheduler integrates with AWS Key Management Service (AWS KMS) to encrypt and decrypt your data using an AWS KMS key.
EventBridge Scheduler supports two types of KMS keys: AWS owned keys, and customer managed keys.

By default, all events in Scheduler are encrypted with a key that AWS owns and manages.
If you wish you can also provide a customer managed key to encrypt and decrypt the payload that your schedule delivers to its target using the `key` property.
Target classes will automatically add AWS `kms:Decrypt` permission to your schedule's execution role permissions policy.

```go
var key key
var fn function


target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
		"payload": jsii.String("useful"),
	}),
})

schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(10))),
	Target: Target,
	Key: Key,
})
```

> See [Data protection in Amazon EventBridge Scheduler](https://docs.aws.amazon.com/scheduler/latest/UserGuide/data-protection.html) for more details.

## Configuring flexible time windows

You can configure flexible time windows by specifying the `timeWindow` property.
Flexible time windows are disabled by default.

```go
var target lambdaInvoke


schedule := awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(12))),
	Target: Target,
	TimeWindow: awscdkscheduleralpha.TimeWindow_Flexible(awscdk.Duration_*Hours(jsii.Number(10))),
})
```

> See [Configuring flexible time windows](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-flexible-time-windows.html) for more details.

## Error-handling

You can configure how your schedule handles failures, when EventBridge Scheduler is unable to deliver an event
successfully to a target, by using two primary mechanisms: a retry policy, and a dead-letter queue (DLQ).

A retry policy determines the number of times EventBridge Scheduler must retry a failed event, and how long
to keep an unprocessed event.

A DLQ is a standard Amazon SQS queue EventBridge Scheduler uses to deliver failed events to, after the retry
policy has been exhausted. You can use a DLQ to troubleshoot issues with your schedule or its downstream target.
If you've configured a retry policy for your schedule, EventBridge Scheduler delivers the dead-letter event after
exhausting the maximum number of retries you set in the retry policy.

```go
var fn function


dlq := sqs.NewQueue(this, jsii.String("DLQ"), &QueueProps{
	QueueName: jsii.String("MyDLQ"),
})

target := targets.NewLambdaInvoke(fn, &ScheduleTargetBaseProps{
	DeadLetterQueue: dlq,
	MaxEventAge: awscdk.Duration_Minutes(jsii.Number(1)),
	RetryAttempts: jsii.Number(3),
})
```

## Monitoring

You can monitor Amazon EventBridge Scheduler using CloudWatch, which collects raw data
and processes it into readable, near real-time metrics. EventBridge Scheduler emits
a set of metrics for all schedules, and an additional set of metrics for schedules that
have an associated dead-letter queue (DLQ). If you configure a DLQ for your schedule,
EventBridge Scheduler publishes additional metrics when your schedule exhausts its retry policy.

### Metrics for all schedules

The `Schedule` class provides static methods for accessing all schedules metrics with default configuration, such as `metricAllErrors` for viewing errors when executing targets.

```go
cloudwatch.NewAlarm(this, jsii.String("SchedulesErrorAlarm"), &AlarmProps{
	Metric: awscdkscheduleralpha.Schedule_MetricAllErrors(),
	Threshold: jsii.Number(0),
	EvaluationPeriods: jsii.Number(1),
})
```

### Metrics for a Schedule Group

To view metrics for a specific group you can use methods on class `ScheduleGroup`:

```go
scheduleGroup := awscdkscheduleralpha.NewScheduleGroup(this, jsii.String("ScheduleGroup"), &ScheduleGroupProps{
	ScheduleGroupName: jsii.String("MyScheduleGroup"),
})

cloudwatch.NewAlarm(this, jsii.String("MyGroupErrorAlarm"), &AlarmProps{
	Metric: scheduleGroup.metricTargetErrors(),
	EvaluationPeriods: jsii.Number(1),
	Threshold: jsii.Number(0),
})

// Or use default group
defaultScheduleGroup := awscdkscheduleralpha.ScheduleGroup_FromDefaultScheduleGroup(this, jsii.String("DefaultScheduleGroup"))
cloudwatch.NewAlarm(this, jsii.String("DefaultScheduleGroupErrorAlarm"), &AlarmProps{
	Metric: defaultScheduleGroup.MetricTargetErrors(),
	EvaluationPeriods: jsii.Number(1),
	Threshold: jsii.Number(0),
})
```

See full list of metrics and their description at
[Monitoring Using CloudWatch Metrics](https://docs.aws.amazon.com/scheduler/latest/UserGuide/monitoring-cloudwatch.html)
in the *AWS EventBridge Scheduler User Guide*.
