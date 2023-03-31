package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Use an AWS Lambda function as an event rule target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
//   	deadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	maxEventAge: cdk.duration.hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	retryAttempts: jsii.Number(2),
//   }))
//
// Experimental.
type LambdaFunction interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Lambda as a result from an EventBridge event.
	// Experimental.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__awseventsIRuleTarget
}

// Experimental.
func NewLambdaFunction(handler awslambda.IFunction, props *LambdaFunctionProps) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"monocdk.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunction_Override(l LambdaFunction, handler awslambda.IFunction, props *LambdaFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		l,
	)
}

func (l *jsiiProxy_LambdaFunction) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := l.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

