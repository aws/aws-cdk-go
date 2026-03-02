package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Use an AWS Lambda function as an event rule target.
//
// Example:
//   myFunctionHandler := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Code: lambda.Code_FromAsset(jsii.String("resource/myfunction")),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   })
//
//   eventRule := cloudtrail.Trail_OnEvent(this, jsii.String("MyCloudWatchEvent"), &OnEventOptions{
//   	Target: targets.NewLambdaFunction(myFunctionHandler),
//   })
//
//   eventRule.AddEventPattern(&EventPattern{
//   	Account: []*string{
//   		jsii.String("123456789012"),
//   	},
//   	Source: []*string{
//   		jsii.String("aws.s3"),
//   	},
//   })
//
type LambdaFunction interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this Lambda as a result from an EventBridge event.
	Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__awseventsIRuleTarget
}

func NewLambdaFunction(handler awslambda.IFunction, props *LambdaFunctionProps) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		&j,
	)

	return &j
}

func NewLambdaFunction_Override(l LambdaFunction, handler awslambda.IFunction, props *LambdaFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LambdaFunction",
		[]interface{}{handler, props},
		l,
	)
}

func (l *jsiiProxy_LambdaFunction) Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig {
	if err := l.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

