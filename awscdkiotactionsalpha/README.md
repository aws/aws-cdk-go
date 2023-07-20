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
* Write messages into columns of DynamoDB
* Put messages IoT Events input
* Send messages to HTTPS endpoints

## Republish a message to another MQTT topic

The code snippet below creates an AWS IoT Rule that republish a message to
another MQTT topic when it is triggered.

```go
iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewIotRepublishMqttAction(jsii.String("${topic()}/republish"), &IotRepublishMqttActionProps{
			QualityOfService: actions.MqttQualityOfService_AT_LEAST_ONCE,
		}),
	},
})
```

## Invoke a Lambda function

The code snippet below creates an AWS IoT Rule that invoke a Lambda function
when it is triggered.

```go
func := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_NODEJS_14_X(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String(`
	    exports.handler = (event) => {
	      console.log("It is test for lambda action of AWS IoT Rule.", event);
	    };`)),
})

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewLambdaFunctionAction(func),
	},
})
```

## Put objects to a S3 bucket

The code snippet below creates an AWS IoT Rule that puts objects to a S3 bucket
when it is triggered.

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	Actions: []iAction{
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

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewS3PutObjectAction(bucket, &S3PutObjectActionProps{
			Key: jsii.String("${year}/${month}/${day}/${topic(2)}"),
		}),
	},
})
```

If you wanna set access control to the S3 bucket object, you can specify `accessControl` as following:

```go
bucket := s3.NewBucket(this, jsii.String("MyBucket"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewS3PutObjectAction(bucket, &S3PutObjectActionProps{
			AccessControl: s3.BucketAccessControl_PUBLIC_READ,
		}),
	},
})
```

## Put logs to CloudWatch Logs

The code snippet below creates an AWS IoT Rule that puts logs to CloudWatch Logs
when it is triggered.

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewCloudWatchLogsAction(logGroup),
	},
})
```

## Capture CloudWatch metrics

The code snippet below creates an AWS IoT Rule that capture CloudWatch metrics
when it is triggered.

```go
topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, namespace, unit, value, timestamp FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewCloudWatchPutMetricAction(&CloudWatchPutMetricActionProps{
			MetricName: jsii.String("${topic(2)}"),
			MetricNamespace: jsii.String("${namespace}"),
			MetricUnit: jsii.String("${unit}"),
			MetricValue: jsii.String("${value}"),
			MetricTimestamp: jsii.String("${timestamp}"),
		}),
	},
})
```

## Start Step Functions State Machine

The code snippet below creates an AWS IoT Rule that starts a Step Functions State Machine
when it is triggered.

```go
stateMachine := stepfunctions.NewStateMachine(this, jsii.String("SM"), &StateMachineProps{
	DefinitionBody: stepfunctions.DefinitionBody_FromChainable(stepfunctions.NewWait(this, jsii.String("Hello"), &WaitProps{
		Time: stepfunctions.WaitTime_Duration(awscdk.Duration_Seconds(jsii.Number(10))),
	})),
})

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewStepFunctionsStateMachineAction(stateMachine),
	},
})
```

## Change the state of an Amazon CloudWatch alarm

The code snippet below creates an AWS IoT Rule that changes the state of an Amazon CloudWatch alarm when it is triggered:

```go
import "github.com/aws/aws-cdk-go/awscdk"


metric := cloudwatch.NewMetric(&MetricProps{
	Namespace: jsii.String("MyNamespace"),
	MetricName: jsii.String("MyMetric"),
	Dimensions: map[string]interface{}{
		"MyDimension": jsii.String("MyDimensionValue"),
	},
})
alarm := cloudwatch.NewAlarm(this, jsii.String("MyAlarm"), &AlarmProps{
	Metric: metric,
	Threshold: jsii.Number(100),
	EvaluationPeriods: jsii.Number(3),
	DatapointsToAlarm: jsii.Number(2),
})

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewCloudWatchSetAlarmStateAction(alarm, &CloudWatchSetAlarmStateActionProps{
			Reason: jsii.String("AWS Iot Rule action is triggered"),
			AlarmStateToSet: cloudwatch.AlarmState_ALARM,
		}),
	},
})
```

## Put records to Kinesis Data stream

The code snippet below creates an AWS IoT Rule that puts records to Kinesis Data
stream when it is triggered.

```go
import kinesis "github.com/aws/aws-cdk-go/awscdk"


stream := kinesis.NewStream(this, jsii.String("MyStream"))

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewKinesisPutRecordAction(stream, &KinesisPutRecordActionProps{
			PartitionKey: jsii.String("${newuuid()}"),
		}),
	},
})
```

## Put records to Kinesis Data Firehose stream

The code snippet below creates an AWS IoT Rule that puts records to Put records
to Kinesis Data Firehose stream when it is triggered.

```go
import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"


bucket := s3.NewBucket(this, jsii.String("MyBucket"))
stream := firehose.NewDeliveryStream(this, jsii.String("MyStream"), &DeliveryStreamProps{
	Destinations: []iDestination{
		destinations.NewS3Bucket(bucket),
	},
})

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewFirehosePutRecordAction(stream, &FirehosePutRecordActionProps{
			BatchMode: jsii.Boolean(true),
			RecordSeparator: actions.FirehoseRecordSeparator_NEWLINE,
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

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewSqsQueueAction(queue, &SqsQueueActionProps{
			UseBase64: jsii.Boolean(true),
		}),
	},
})
```

## Publish messages on an SNS topic

The code snippet below creates and AWS IoT Rule that publishes messages to an SNS topic when it is triggered:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


topic := sns.NewTopic(this, jsii.String("MyTopic"))

topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewSnsTopicAction(topic, &SnsTopicActionProps{
			MessageFormat: actions.SnsActionMessageFormat_JSON,
		}),
	},
})
```

## Write attributes of a message to DynamoDB

The code snippet below creates an AWS IoT rule that writes all or part of an
MQTT message to DynamoDB using the DynamoDBv2 action.

```go
import dynamodb "github.com/aws/aws-cdk-go/awscdk"

var table table


topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewDynamoDBv2PutItemAction(table),
	},
})
```

## Put messages IoT Events input

The code snippet below creates an AWS IoT Rule that puts messages
to an IoT Events input when it is triggered:

```go
import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
import iam "github.com/aws/aws-cdk-go/awscdk"

var role iRole


input := iotevents.NewInput(this, jsii.String("MyInput"), &InputProps{
	AttributeJsonPaths: []*string{
		jsii.String("payload.temperature"),
		jsii.String("payload.transactionId"),
	},
})
topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT * FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewIotEventsPutMessageAction(input, &IotEventsPutMessageActionProps{
			BatchMode: jsii.Boolean(true),
			 // optional property, default is 'false'
			MessageId: jsii.String("${payload.transactionId}"),
			 // optional property, default is a new UUID
			Role: role,
		}),
	},
})
```

## Send Messages to HTTPS Endpoints

The code snippet below creates an AWS IoT Rule that sends messages
to an HTTPS endpoint when it is triggered:

```go
topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, year, month, day FROM 'device/+/data'")),
})

topicRule.AddAction(
actions.NewHttpsAction(jsii.String("https://example.com/endpoint"), &HttpsActionProps{
	ConfirmationUrl: jsii.String("https://example.com"),
	Headers: []httpActionHeader{
		&httpActionHeader{
			Key: jsii.String("key0"),
			Value: jsii.String("value0"),
		},
		&httpActionHeader{
			Key: jsii.String("key1"),
			Value: jsii.String("value1"),
		},
	},
	Auth: &HttpActionSigV4Auth{
		ServiceName: jsii.String("serviceName"),
		SigningRegion: jsii.String("us-east-1"),
	},
}))
```
