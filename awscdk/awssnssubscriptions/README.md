# CDK Construct Library for Amazon Simple Notification Service Subscriptions

This library provides constructs for adding subscriptions to an Amazon SNS topic.
Subscriptions can be added by calling the `.addSubscription(...)` method on the topic.

## Subscriptions

Subscriptions can be added to the following endpoints:

* HTTPS
* Amazon SQS
* AWS Lambda
* Email
* SMS

Subscriptions to Amazon SQS and AWS Lambda can be added on topics across regions.

Create an Amazon SNS Topic to add subscriptions.

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
```

### HTTPS

Add an HTTP or HTTPS Subscription to your topic:

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.addSubscription(subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/")))
```

The URL being subscribed can also be [tokens](https://docs.aws.amazon.com/cdk/latest/guide/tokens.html), that resolve
to a URL during deployment. A typical use case is when the URL is passed in as a [CloudFormation
parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html). The
following code defines a CloudFormation parameter and uses it in a URL subscription.

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
url := awscdk.NewCfnParameter(this, jsii.String("url-param"))

myTopic.addSubscription(subscriptions.NewUrlSubscription(url.valueAsString))
```

### Amazon SQS

Subscribe a queue to your topic:

```go
myQueue := sqs.NewQueue(this, jsii.String("MyQueue"))
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))

myTopic.addSubscription(subscriptions.NewSqsSubscription(myQueue))
```

KMS key permissions will automatically be granted to SNS when a subscription is made to
an encrypted queue.

Note that subscriptions of queues in different accounts need to be manually confirmed by
reading the initial message from the queue and visiting the link found in it.

### AWS Lambda

Subscribe an AWS Lambda function to your topic:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
var myFunction function


myTopic := sns.NewTopic(this, jsii.String("myTopic"))
myTopic.addSubscription(subscriptions.NewLambdaSubscription(myFunction))
```

### Email

Subscribe an email address to your topic:

```go
myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
myTopic.addSubscription(subscriptions.NewEmailSubscription(jsii.String("foo@bar.com")))
```

The email being subscribed can also be [tokens](https://docs.aws.amazon.com/cdk/latest/guide/tokens.html), that resolve
to an email address during deployment. A typical use case is when the email address is passed in as a [CloudFormation
parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html). The
following code defines a CloudFormation parameter and uses it in an email subscription.

```go
myTopic := sns.NewTopic(this, jsii.String("Topic"))
emailAddress := awscdk.NewCfnParameter(this, jsii.String("email-param"))

myTopic.addSubscription(subscriptions.NewEmailSubscription(emailAddress.valueAsString))
```

Note that email subscriptions require confirmation by visiting the link sent to the
email address.

### SMS

Subscribe an sms number to your topic:

```go
myTopic := sns.NewTopic(this, jsii.String("Topic"))

myTopic.addSubscription(subscriptions.NewSmsSubscription(jsii.String("+15551231234")))
```

The number being subscribed can also be [tokens](https://docs.aws.amazon.com/cdk/latest/guide/tokens.html), that resolve
to a number during deployment. A typical use case is when the number is passed in as a [CloudFormation
parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html). The
following code defines a CloudFormation parameter and uses it in an sms subscription.

```go
myTopic := sns.NewTopic(this, jsii.String("Topic"))
smsNumber := awscdk.NewCfnParameter(this, jsii.String("sms-param"))

myTopic.addSubscription(subscriptions.NewSmsSubscription(smsNumber.valueAsString))
```
