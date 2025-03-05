# AWS CloudTrail Construct Library

## Trail

AWS CloudTrail enables governance, compliance, and operational and risk auditing of your AWS account. Actions taken by
a user, role, or an AWS service are recorded as events in CloudTrail. Learn more at the [CloudTrail
documentation](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-user-guide.html).

The `Trail` construct enables ongoing delivery of events as log files to an Amazon S3 bucket. Learn more about [Creating
a Trail for Your AWS Account](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.html).
The following code creates a simple CloudTrail for your account -

```go
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"))
```

By default, this will create a new S3 Bucket that CloudTrail will write to, and choose a few other reasonable defaults
such as turning on multi-region and global service events.
The defaults for each property and how to override them are all documented on the `TrailProps` interface.

## Log File Validation

In order to validate that the CloudTrail log file was not modified after CloudTrail delivered it, CloudTrail provides a
digital signature for each file. Learn more at [Validating CloudTrail Log File
Integrity](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-log-file-validation-intro.html).

This is enabled on the `Trail` construct by default, but can be turned off by setting `enableFileValidation` to `false`.

```go
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &TrailProps{
	EnableFileValidation: jsii.Boolean(false),
})
```

## Notifications

Amazon SNS notifications can be configured upon new log files containing Trail events are delivered to S3.
Learn more at [Configuring Amazon SNS Notifications for
CloudTrail](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/configure-sns-notifications-for-cloudtrail.html).
The following code configures an SNS topic to be notified -

```go
topic := sns.NewTopic(this, jsii.String("TrailTopic"))
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &TrailProps{
	SnsTopic: topic,
})
```

## Service Integrations

Besides sending trail events to S3, they can also be configured to notify other AWS services -

### Amazon CloudWatch Logs

CloudTrail events can be delivered to a CloudWatch Logs LogGroup. By default, a new LogGroup is created with a
default retention setting. The following code enables sending CloudWatch logs but specifies a particular retention
period for the created Log Group.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &TrailProps{
	SendToCloudWatchLogs: jsii.Boolean(true),
	CloudWatchLogsRetention: logs.RetentionDays_FOUR_MONTHS,
})
```

If you would like to use a specific log group instead, this can be configured via `cloudwatchLogGroup`.

### Amazon EventBridge

Amazon EventBridge rules can be configured to be triggered when CloudTrail events occur using the `Trail.onEvent()` API.
Using APIs available in `aws-events`, these events can be filtered to match to those that are of interest, either from
a specific service, account or time range. See [Events delivered via
CloudTrail](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/EventTypes.html#events-for-services-not-listed)
to learn more about the event structure for events from CloudTrail.

The following code filters events for S3 from a specific AWS account and triggers a lambda function.

```go
myFunctionHandler := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Code: lambda.Code_FromAsset(jsii.String("resource/myfunction")),
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("index.handler"),
})

eventRule := cloudtrail.Trail_OnEvent(this, jsii.String("MyCloudWatchEvent"), &OnEventOptions{
	Target: targets.NewLambdaFunction(myFunctionHandler),
})

eventRule.AddEventPattern(&EventPattern{
	Account: []*string{
		jsii.String("123456789012"),
	},
	Source: []*string{
		jsii.String("aws.s3"),
	},
})
```

## Multi-Region & Global Service Events

By default, a `Trail` is configured to deliver log files from multiple regions to a single S3 bucket for a given
account. This creates shadow trails (replication of the trails) in all of the other regions. Learn more about [How
CloudTrail Behaves Regionally](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-regional-and-global-services)
and about the [`IsMultiRegion`
property](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail).

For most services, events are recorded in the region where the action occurred. For global services such as AWS IAM,
AWS STS, Amazon CloudFront, Route 53, etc., events are delivered to any trail that includes global services. Learn more
[About Global Service Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-global-service-events).

Events for global services are turned on by default for `Trail` constructs in the CDK.

The following code disables multi-region trail delivery and trail delivery for global services for a specific `Trail` -

```go
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &TrailProps{
	// ...
	IsMultiRegionTrail: jsii.Boolean(false),
	IncludeGlobalServiceEvents: jsii.Boolean(false),
})
```

## Events Types

**Management events** provide information about management operations that are performed on resources in your AWS
account. These are also known as control plane operations. Learn more about [Management
Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-events).

By default, a `Trail` logs all management events. However, they can be configured to either be turned off, or to only
log 'Read' or 'Write' events.

The following code configures the `Trail` to only track management events that are of type 'Read'.

```go
trail := cloudtrail.NewTrail(this, jsii.String("CloudTrail"), &TrailProps{
	// ...
	ManagementEvents: cloudtrail.ReadWriteType_READ_ONLY,
})
```

**Data events** provide information about the resource operations performed on or in a resource. These are also known
as data plane operations. Learn more about [Data
Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html#cloudtrail-concepts-events).
By default, no data events are logged for a `Trail`.

AWS CloudTrail supports data event logging for Amazon S3 objects and AWS Lambda functions.

The `logAllS3DataEvents()` API configures the trail to log all S3 data events while the `addS3EventSelector()` API can
be used to configure logging of S3 data events for specific buckets and specific object prefix. The following code
configures logging of S3 data events for `fooBucket` and with object prefix `bar/`.

```go
import s3 "github.com/aws/aws-cdk-go/awscdk"
var bucket bucket


trail := cloudtrail.NewTrail(this, jsii.String("MyAmazingCloudTrail"))

// Adds an event selector to the bucket foo
trail.AddS3EventSelector([]s3EventSelector{
	&s3EventSelector{
		Bucket: Bucket,
		ObjectPrefix: jsii.String("bar/"),
	},
})
```

Similarly, the `logAllLambdaDataEvents()` configures the trail to log all Lambda data events while the
`addLambdaEventSelector()` API can be used to configure logging for specific Lambda functions. The following code
configures logging of Lambda data events for a specific Function.

```go
trail := cloudtrail.NewTrail(this, jsii.String("MyAmazingCloudTrail"))
amazingFunction := lambda.NewFunction(this, jsii.String("AnAmazingFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_LATEST(),
	Handler: jsii.String("hello.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda")),
})

// Add an event selector to log data events for the provided Lambda functions.
trail.AddLambdaEventSelector([]iFunction{
	amazingFunction,
})
```

## Organization Trail

It is possible to create a trail that will be applied to all accounts in an organization if the current account manages an organization.
To enable this, the property `isOrganizationTrail` must be set. If this property is set and the current account does not manage an organization, the stack will fail to deploy.

```go
cloudtrail.NewTrail(this, jsii.String("OrganizationTrail"), &TrailProps{
	IsOrganizationTrail: jsii.Boolean(true),
})
```

## CloudTrail Insights

Set `InsightSelector` to enable Insight.
Insights selector values can be `ApiCallRateInsight`, `ApiErrorRateInsight`, or both.

```go
cloudtrail.NewTrail(this, jsii.String("Insights"), &TrailProps{
	InsightTypes: []insightType{
		cloudtrail.*insightType_API_CALL_RATE(),
		cloudtrail.*insightType_API_ERROR_RATE(),
	},
})
```
