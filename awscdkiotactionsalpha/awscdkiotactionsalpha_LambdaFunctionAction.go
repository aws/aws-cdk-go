// Receipt rule actions for AWS IoT
package awscdkiotactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkiotactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkiotalpha/v2"
)

// The action to invoke an AWS Lambda function, passing in an MQTT message.
//
// Example:
//   func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("\n    exports.handler = (event) => {\n      console.log(\"It is test for lambda action of AWS IoT Rule.\", event);\n    };")),
//   })
//
//   iot.NewTopicRule(this, jsii.String("TopicRule"), &topicRuleProps{
//   	sql: iot.iotSql.fromStringAsVer20160323(jsii.String("SELECT topic(2) as device_id, timestamp() as timestamp, temperature FROM 'device/+/data'")),
//   	actions: []iAction{
//   		actions.NewLambdaFunctionAction(func),
//   	},
//   })
//
// Experimental.
type LambdaFunctionAction interface {
	awscdkiotalpha.IAction
}

// The jsii proxy struct for LambdaFunctionAction
type jsiiProxy_LambdaFunctionAction struct {
	internal.Type__awscdkiotalphaIAction
}

// Experimental.
func NewLambdaFunctionAction(func_ awslambda.IFunction) LambdaFunctionAction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionActionParameters(func_); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.LambdaFunctionAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionAction_Override(l LambdaFunctionAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iot-actions-alpha.LambdaFunctionAction",
		[]interface{}{func_},
		l,
	)
}

