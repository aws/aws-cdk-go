# Amazon CloudWatch Logs Construct Library

This library supplies constructs for working with CloudWatch Logs.

## Log Groups/Streams

The basic unit of CloudWatch is a *Log Group*. Every log group typically has the
same kind of data logged to it, in the same format. If there are multiple
applications or services logging into the Log Group, each of them creates a new
*Log Stream*.

Every log operation creates a "log event", which can consist of a simple string
or a single-line JSON object. JSON objects have the advantage that they afford
more filtering abilities (see below).

The only configurable attribute for log streams is the retention period, which
configures after how much time the events in the log stream expire and are
deleted.

The default retention period if not supplied is 2 years, but it can be set to
one of the values in the `RetentionDays` enum to configure a different
retention period (including infinite retention).

```go
// Configure log group for short retention
logGroup := awscdk.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
	Retention: awscdk.RetentionDays_ONE_WEEK,
})// Configure log group for infinite retention
logGroup := awscdk.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
	Retention: infinity,
})
```

## LogRetention

The `LogRetention` construct is a way to control the retention period of log groups that are created outside of the CDK. The construct is usually
used on log groups that are auto created by AWS services, such as [AWS
lambda](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-cloudwatchlogs.html).

This is implemented using a [CloudFormation custom
resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html)
which pre-creates the log group if it doesn't exist, and sets the specified log retention period (never expire, by default).

By default, the log group will be created in the same region as the stack. The `logGroupRegion` property can be used to configure
log groups in other regions. This is typically useful when controlling retention for log groups auto-created by global services that
publish their log group to a specific region, such as AWS Chatbot creating a log group in `us-east-1`.

By default, the log group created by LogRetention will be retained after the stack is deleted. If the RemovalPolicy is set to DESTROY, then the log group will be deleted when the stack is deleted.

## Resource Policy

CloudWatch Resource Policies allow other AWS services or IAM Principals to put log events into the log groups.
A resource policy is automatically created when `addToResourcePolicy` is called on the LogGroup for the first time:

```go
logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"))
logGroup.addToResourcePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("logs:CreateLogStream"),
		jsii.String("logs:PutLogEvents"),
	},
	Principals: []iPrincipal{
		iam.NewServicePrincipal(jsii.String("es.amazonaws.com")),
	},
	Resources: []*string{
		logGroup.LogGroupArn,
	},
}))
```

Or more conveniently, write permissions to the log group can be granted as follows which gives same result as in the above example.

```go
logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"))
logGroup.grantWrite(iam.NewServicePrincipal(jsii.String("es.amazonaws.com")))
```

Similarily, read permissions can be granted to the log group as follows.

```go
logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"))
logGroup.grantRead(iam.NewServicePrincipal(jsii.String("es.amazonaws.com")))
```

Be aware that any ARNs or tokenized values passed to the resource policy will be converted into AWS Account IDs.
This is because CloudWatch Logs Resource Policies do not accept ARNs as principals, but they do accept
Account ID strings. Non-ARN principals, like Service principals or Any principals, are accepted by CloudWatch.

## Encrypting Log Groups

By default, log group data is always encrypted in CloudWatch Logs. You have the
option to encrypt log group data using a AWS KMS customer master key (CMK) should
you not wish to use the default AWS encryption. Keep in mind that if you decide to
encrypt a log group, any service or IAM identity that needs to read the encrypted
log streams in the future will require the same CMK to decrypt the data.

Here's a simple example of creating an encrypted Log Group using a KMS CMK.

```go
import kms "github.com/aws/aws-cdk-go/awscdk"


logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
	EncryptionKey: kms.NewKey(this, jsii.String("Key")),
})
```

See the AWS documentation for more detailed information about [encrypting CloudWatch
Logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/encrypt-log-data-kms.html).

## Subscriptions and Destinations

Log events matching a particular filter can be sent to either a Lambda function
or a Kinesis stream.

If the Kinesis stream lives in a different account, a `CrossAccountDestination`
object needs to be added in the destination account which will act as a proxy
for the remote Kinesis stream. This object is automatically created for you
if you use the CDK Kinesis library.

Create a `SubscriptionFilter`, initialize it with an appropriate `Pattern` (see
below) and supply the intended destination:

```go
import destinations "github.com/aws/aws-cdk-go/awscdk"

var fn function
var logGroup logGroup


logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
	LogGroup: LogGroup,
	Destination: destinations.NewLambdaDestination(fn),
	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
	FilterName: jsii.String("ErrorInMainThread"),
})
```

## Metric Filters

CloudWatch Logs can extract and emit metrics based on a textual log stream.
Depending on your needs, this may be a more convenient way of generating metrics
for you application than making calls to CloudWatch Metrics yourself.

A `MetricFilter` either emits a fixed number every time it sees a log event
matching a particular pattern (see below), or extracts a number from the log
event and uses that as the metric value.

Example:

```go
awscdk.NewMetricFilter(this, jsii.String("MetricFilter"), &MetricFilterProps{
	LogGroup: LogGroup,
	MetricNamespace: jsii.String("MyApp"),
	MetricName: jsii.String("Latency"),
	FilterPattern: awscdk.FilterPattern_Exists(jsii.String("$.latency")),
	MetricValue: jsii.String("$.latency"),
})
```

Remember that if you want to use a value from the log event as the metric value,
you must mention it in your pattern somewhere.

A very simple MetricFilter can be created by using the `logGroup.extractMetric()`
helper function:

```go
var logGroup logGroup

logGroup.extractMetric(jsii.String("$.jsonField"), jsii.String("Namespace"), jsii.String("MetricName"))
```

Will extract the value of `jsonField` wherever it occurs in JSON-structured
log records in the LogGroup, and emit them to CloudWatch Metrics under
the name `Namespace/MetricName`.

### Exposing Metric on a Metric Filter

You can expose a metric on a metric filter by calling the `MetricFilter.metric()` API.
This has a default of `statistic = 'avg'` if the statistic is not set in the `props`.

```go
var logGroup logGroup

mf := logs.NewMetricFilter(this, jsii.String("MetricFilter"), &MetricFilterProps{
	LogGroup: LogGroup,
	MetricNamespace: jsii.String("MyApp"),
	MetricName: jsii.String("Latency"),
	FilterPattern: logs.FilterPattern_Exists(jsii.String("$.latency")),
	MetricValue: jsii.String("$.latency"),
	Dimensions: map[string]*string{
		"ErrorCode": jsii.String("$.errorCode"),
	},
	Unit: cloudwatch.Unit_MILLISECONDS,
})

//expose a metric from the metric filter
metric := mf.Metric()

//you can use the metric to create a new alarm
//you can use the metric to create a new alarm
cloudwatch.NewAlarm(this, jsii.String("alarm from metric filter"), &AlarmProps{
	Metric: Metric,
	Threshold: jsii.Number(100),
	EvaluationPeriods: jsii.Number(2),
})
```

## Patterns

Patterns describe which log events match a subscription or metric filter. There
are three types of patterns:

* Text patterns
* JSON patterns
* Space-delimited table patterns

All patterns are constructed by using static functions on the `FilterPattern`
class.

In addition to the patterns above, the following special patterns exist:

* `FilterPattern.allEvents()`: matches all log events.
* `FilterPattern.literal(string)`: if you already know what pattern expression to
  use, this function takes a string and will use that as the log pattern. For
  more information, see the [Filter and Pattern
  Syntax](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).

### Text Patterns

Text patterns match if the literal strings appear in the text form of the log
line.

* `FilterPattern.allTerms(term, term, ...)`: matches if all of the given terms
  (substrings) appear in the log event.
* `FilterPattern.anyTerm(term, term, ...)`: matches if all of the given terms
  (substrings) appear in the log event.
* `FilterPattern.anyTermGroup([term, term, ...], [term, term, ...], ...)`: matches if
  all of the terms in any of the groups (specified as arrays) matches. This is
  an OR match.

Examples:

```go
// Search for lines that contain both "ERROR" and "MainThread"
pattern1 := logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread"))

// Search for lines that either contain both "ERROR" and "MainThread", or
// both "WARN" and "Deadlock".
pattern2 := logs.FilterPattern_AnyTermGroup([]*string{
	jsii.String("ERROR"),
	jsii.String("MainThread"),
}, []*string{
	jsii.String("WARN"),
	jsii.String("Deadlock"),
})
```

## JSON Patterns

JSON patterns apply if the log event is the JSON representation of an object
(without any other characters, so it cannot include a prefix such as timestamp
or log level). JSON patterns can make comparisons on the values inside the
fields.

* **Strings**: the comparison operators allowed for strings are `=` and `!=`.
  String values can start or end with a `*` wildcard.
* **Numbers**: the comparison operators allowed for numbers are `=`, `!=`,
  `<`, `<=`, `>`, `>=`.

Fields in the JSON structure are identified by identifier the complete object as `$`
and then descending into it, such as `$.field` or `$.list[0].field`.

* `FilterPattern.stringValue(field, comparison, string)`: matches if the given
  field compares as indicated with the given string value.
* `FilterPattern.numberValue(field, comparison, number)`: matches if the given
  field compares as indicated with the given numerical value.
* `FilterPattern.isNull(field)`: matches if the given field exists and has the
  value `null`.
* `FilterPattern.notExists(field)`: matches if the given field is not in the JSON
  structure.
* `FilterPattern.exists(field)`: matches if the given field is in the JSON
  structure.
* `FilterPattern.booleanValue(field, boolean)`: matches if the given field
  is exactly the given boolean value.
* `FilterPattern.all(jsonPattern, jsonPattern, ...)`: matches if all of the
  given JSON patterns match. This makes an AND combination of the given
  patterns.
* `FilterPattern.any(jsonPattern, jsonPattern, ...)`: matches if any of the
  given JSON patterns match. This makes an OR combination of the given
  patterns.

Example:

```go
// Search for all events where the component field is equal to
// "HttpServer" and either error is true or the latency is higher
// than 1000.
pattern := logs.FilterPattern_All(logs.FilterPattern_StringValue(jsii.String("$.component"), jsii.String("="), jsii.String("HttpServer")), logs.FilterPattern_Any(logs.FilterPattern_BooleanValue(jsii.String("$.error"), jsii.Boolean(true)), logs.FilterPattern_NumberValue(jsii.String("$.latency"), jsii.String(">"), jsii.Number(1000))))
```

## Space-delimited table patterns

If the log events are rows of a space-delimited table, this pattern can be used
to identify the columns in that structure and add conditions on any of them. The
canonical example where you would apply this type of pattern is Apache server
logs.

Text that is surrounded by `"..."` quotes or `[...]` square brackets will
be treated as one column.

* `FilterPattern.spaceDelimited(column, column, ...)`: construct a
  `SpaceDelimitedTextPattern` object with the indicated columns. The columns
  map one-by-one the columns found in the log event. The string `"..."` may
  be used to specify an arbitrary number of unnamed columns anywhere in the
  name list (but may only be specified once).

After constructing a `SpaceDelimitedTextPattern`, you can use the following
two members to add restrictions:

* `pattern.whereString(field, comparison, string)`: add a string condition.
  The rules are the same as for JSON patterns.
* `pattern.whereNumber(field, comparison, number)`: add a numerical condition.
  The rules are the same as for JSON patterns.

Multiple restrictions can be added on the same column; they must all apply.

Example:

```go
// Search for all events where the component is "HttpServer" and the
// result code is not equal to 200.
pattern := logs.FilterPattern_SpaceDelimited(jsii.String("time"), jsii.String("component"), jsii.String("..."), jsii.String("result_code"), jsii.String("latency")).WhereString(jsii.String("component"), jsii.String("="), jsii.String("HttpServer")).WhereNumber(jsii.String("result_code"), jsii.String("!="), jsii.Number(200))
```

## Logs Insights Query Definition

Creates a query definition for CloudWatch Logs Insights.

Example:

```go
logs.NewQueryDefinition(this, jsii.String("QueryDefinition"), &QueryDefinitionProps{
	QueryDefinitionName: jsii.String("MyQuery"),
	QueryString: logs.NewQueryString(&QueryStringProps{
		Fields: []*string{
			jsii.String("@timestamp"),
			jsii.String("@message"),
		},
		ParseStatements: []*string{
			jsii.String("@message \"[*] *\" as loggingType, loggingMessage"),
			jsii.String("@message \"<*>: *\" as differentLoggingType, differentLoggingMessage"),
		},
		FilterStatements: []*string{
			jsii.String("loggingType = \"ERROR\""),
			jsii.String("loggingMessage = \"A very strange error occurred!\""),
		},
		Sort: jsii.String("@timestamp desc"),
		Limit: jsii.Number(20),
	}),
})
```

## Data Protection Policy

Creates a data protection policy and assigns it to the log group. A data protection policy can help safeguard sensitive data that's ingested by the log group by auditing and masking the sensitive log data. When a user who does not have permission to view masked data views a log event that includes masked data, the sensitive data is replaced by asterisks.

For more information, see [Protect sensitive log data with masking](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html).

For a list of types of managed identifiers that can be audited and masked, see [Types of data that you can protect](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/protect-sensitive-log-data-types.html).

If a new identifier is supported but not yet in the `DataIdentifiers` enum, the name of the identifier can be supplied as `name` in the constructor instead.

To add a custom data identifier, supply a custom `name` and `regex` to the `CustomDataIdentifiers` constructor.
For more information on custom data identifiers, see [Custom data identifiers](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL-custom-data-identifiers.html).

Each policy may consist of a log group, S3 bucket, and/or Firehose delivery stream audit destination.

Example:

```go
import kinesisfirehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"


logGroupDestination := logs.NewLogGroup(this, jsii.String("LogGroupLambdaAudit"), &LogGroupProps{
	LogGroupName: jsii.String("auditDestinationForCDK"),
})

bucket := s3.NewBucket(this, jsii.String("audit-bucket"))
s3Destination := destinations.NewS3Bucket(bucket)

deliveryStream := kinesisfirehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
	Destinations: []iDestination{
		s3Destination,
	},
})

dataProtectionPolicy := logs.NewDataProtectionPolicy(&DataProtectionPolicyProps{
	Name: jsii.String("data protection policy"),
	Description: jsii.String("policy description"),
	Identifiers: []dataIdentifier{
		logs.*dataIdentifier_DRIVERSLICENSE_US(),
		 // managed data identifier
		logs.NewDataIdentifier(jsii.String("EmailAddress")),
		 // forward compatibility for new managed data identifiers
		logs.NewCustomDataIdentifier(jsii.String("EmployeeId"), jsii.String("EmployeeId-\\d{9}")),
	},
	 // custom data identifier
	LogGroupAuditDestination: logGroupDestination,
	S3BucketAuditDestination: bucket,
	DeliveryStreamNameAuditDestination: deliveryStream.DeliveryStreamName,
})

logs.NewLogGroup(this, jsii.String("LogGroupLambda"), &LogGroupProps{
	LogGroupName: jsii.String("cdkIntegLogGroup"),
	DataProtectionPolicy: dataProtectionPolicy,
})
```

## Notes

Be aware that Log Group ARNs will always have the string `:*` appended to
them, to match the behavior of [the CloudFormation `AWS::Logs::LogGroup`
resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#aws-resource-logs-loggroup-return-values).
