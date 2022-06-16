# Actions for AWS IoT Rule

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This library contains integration classes to send data to any number of
supported AWS Services. Instances of these classes should be passed to
`TopicRule` defined in `@aws-cdk/aws-iot`.

Currently supported are:

* Republish a message to another MQTT topic
* Invoke a Lambda function
* Put objects to a S3 bucket
* Put logs to CloudWatch Logs
* Capture CloudWatch metrics
* Change state for a CloudWatch alarm
* Put records to Kinesis Data stream
* Put records to Kinesis Data Firehose stream
* Send messages to SQS queues
* Publish messages on SNS topics

## Republish a message to another MQTT topic

The code snippet below creates an AWS IoT Rule that republish a message to
another MQTT topic when it is triggered.

```go
iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &iotRepublishMqttActionProps{
			qualityOfService: actions.mqttQualityOfService_AT_LEAST_ONCE,
		}),
	},
})
```

## Invoke a Lambda function

The code snippet below creates an AWS IoT Rule that invoke a Lambda function
when it is triggered.

```go
func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromInline(jsii.String("\n    exports.handler = (event) => {\n      console.log(\"It is test for lambda action of AWS IoT Rule.\", event);\n    };")),
})

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewLambdaFunctionAction(func),
	},
})
```

## Put objects to a S3 bucket

The code snippet below creates an AWS IoT Rule that put objects to a S3 bucket
when it is triggered.

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewS3PutObjectAction(bucket),
	},
})
```

The property `key` of `S3PutObjectAction` is given the value `${topic()}/${timestamp()}` by default. This `${topic()}`
and `${timestamp()}` is called Substitution templates. For more information see
[this documentation](https://docs.aws.amazon.com/iot/latest/developerguide/iot-substitution-templates.html).
In above sample, `${topic()}` is replaced by a given MQTT topic as `device/001/data`. And `${timestamp()}` is replaced
by the number of the current timestamp in milliseconds as `1636289461203`. So if the MQTT broker receives an MQTT topic
`device/001/data` on `2021-11-07T00:00:00.000Z`, the S3 bucket object will be put to `device/001/data/1636243200000`.

You can also set specific `key` as following:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
			key: jsii.String("${year}/${month}/${day}/${topic(2)}"),
		}),
	},
})
```

If you wanna set access control to the S3 bucket object, you can specify `accessControl` as following:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewS3PutObjectAction(bucket, &s3PutObjectActionProps{
			accessControl: s3.bucketAccessControl_PUBLIC_READ,
		}),
	},
})
```

## Put logs to CloudWatch Logs

The code snippet below creates an AWS IoT Rule that put logs to CloudWatch Logs
when it is triggered.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewCloudWatchLogsAction(logGroup),
	},
})
```

## Capture CloudWatch metrics

The code snippet below creates an AWS IoT Rule that capture CloudWatch metrics
when it is triggered.

```go
topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, namespace, unit, value, timestamp FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewCloudWatchPutMetricAction(&cloudWatchPutMetricActionProps{
			metricName: jsii.String("${topic(2)}"),
			metricNamespace: jsii.String("${namespace}"),
			metricUnit: jsii.String("${unit}"),
			metricValue: jsii.String("${value}"),
			metricTimestamp: jsii.String("${timestamp}"),
		}),
	},
})
```

## Change the state of an Amazon CloudWatch alarm

The code snippet below creates an AWS IoT Rule that changes the state of an Amazon CloudWatch alarm when it is triggered:

```go
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


metric := cloudwatch.NewMetric(&metricProps{
	namespace: jsii.String("MyNamespace"),
	metricName: jsii.String("MyMetric"),
	dimensions: map[string]interface{}{
		"MyDimension": jsii.String("MyDimensionValue"),
	},
})
alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &alarmProps{
	metric: metric,
	threshold: jsii.Number(100),
	evaluationPeriods: jsii.Number(3),
	datapointsToAlarm: jsii.Number(2),
})

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewCloudWatchSetAlarmStateAction(alarm, &cloudWatchSetAlarmStateActionProps{
			reason: jsii.String("AWS Iot Rule action is triggered"),
			alarmStateToSet: cloudwatch.alarmState_ALARM,
		}),
	},
})
```

## Put records to Kinesis Data stream

The code snippet below creates an AWS IoT Rule that put records to Kinesis Data
stream when it is triggered.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"


stream := kinesis.NewStream(this, jsii.String("MyStream"))

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewKinesisPutRecordAction(stream, &kinesisPutRecordActionProps{
			partitionKey: jsii.String("${newuuid()}"),
		}),
	},
})
```

## Put records to Kinesis Data Firehose stream

The code snippet below creates an AWS IoT Rule that put records to Put records
to Kinesis Data Firehose stream when it is triggered.

```go
import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"


bucket := s3.NewBucket(this, jsii.String("MyBucket"))
stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &deliveryStreamProps{
	destinations: []iDestination{
		destinations.NewS3Bucket(bucket),
	},
})

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewFirehosePutRecordAction(stream, &firehosePutRecordActionProps{
			batchMode: jsii.Boolean(true),
			recordSeparator: actions.firehoseRecordSeparator_NEWLINE,
		}),
	},
})
```

## Send messages to an SQS queue

The code snippet below creates an AWS IoT Rule that send messages
to an SQS queue when it is triggered:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"


queue := sqs.NewQueue(this, jsii.String("MyQueue"))

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewSqsQueueAction(queue, &sqsQueueActionProps{
			useBase64: jsii.Boolean(true),
		}),
	},
})
```

## Publish messages on an SNS topic

The code snippet below creates and AWS IoT Rule that publishes messages to an SNS topic when it is triggered:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


topic := sns.NewTopic(this, jsii.String("MyTopic"))

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	actions: []iAction{
		actions.NewSnsTopicAction(topic, &snsTopicActionProps{
			messageFormat: actions.snsActionMessageFormat_JSON,
		}),
	},
})
```
