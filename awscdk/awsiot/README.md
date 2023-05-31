# AWS IoT Construct Library

AWS IoT Core lets you connect billions of IoT devices and route trillions of
messages to AWS services without managing infrastructure.

## Installation

Install the module:

```console
$ npm i @aws-cdk/aws-iot
```

Import it into your code:

```go
import iot "github.com/aws/aws-cdk-go/awscdk"
import actions "github.com/aws/aws-cdk-go/awscdk"
```

## `TopicRule`

Create a topic rule that give your devices the ability to interact with AWS services.
You can create a topic rule with an action that invoke the Lambda action as following:

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
	TopicRuleName: jsii.String("MyTopicRule"),
	 // optional
	Description: jsii.String("invokes the lambda function"),
	 // optional
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	Actions: []iAction{
		actions.NewLambdaFunctionAction(func),
	},
})
```

Or, you can add an action after constructing the `TopicRule` instance as following:

```go
var func function


topicRule := iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
})
topicRule.AddAction(actions.NewLambdaFunctionAction(func))
```

You can also supply `errorAction` as following,
and the IoT Rule will trigger it if a rule's action is unable to perform:

```go
import logs "github.com/aws/aws-cdk-go/awscdk"


logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"))

iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	ErrorAction: actions.NewCloudWatchLogsAction(logGroup),
})
```

If you wanna make the topic rule disable, add property `enabled: false` as following:

```go
iot.NewTopicRule(this, jsii.String("TopicRule"), &TopicRuleProps{
	Sql: iot.IotSql_FromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp FROM 'device/+/data'")),
	Enabled: jsii.Boolean(false),
})
```

See also [@aws-cdk/aws-iot-actions](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-iot-actions-readme.html) for other actions.
