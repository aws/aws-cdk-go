# S3 Bucket Notifications Destinations

This module includes integration classes for using Topics, Queues or Lambdas
as S3 Notification Destinations.

## Examples

The following example shows how to send a notification to an SNS
topic when an object is created in an S3 bucket:

```go
import sns "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Bucket"))
topic := sns.NewTopic(this, jsii.String("Topic"))

bucket.addEventNotification(s3.eventType_OBJECT_CREATED_PUT, s3n.NewSnsDestination(topic))
```

The following example shows how to send a notification to an SQS queue
when an object is created in an S3 bucket:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Bucket"))
queue := sqs.NewQueue(this, jsii.String("Queue"))

bucket.addEventNotification(s3.eventType_OBJECT_CREATED_PUT, s3n.NewSqsDestination(queue))
```

The following example shows how to send a notification to a Lambda function when an object is created in an S3 bucket:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"


bucket := s3.NewBucket(this, jsii.String("Bucket"))
fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
	runtime: lambda.runtime_NODEJS_14_X(),
	handler: jsii.String("index.handler"),
	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
})

bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(fn))
```
