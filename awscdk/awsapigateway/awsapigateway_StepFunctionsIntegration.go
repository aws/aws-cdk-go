package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Options to integrate with various StepFunction API.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
type StepFunctionsIntegration interface {
}

// The jsii proxy struct for StepFunctionsIntegration
type jsiiProxy_StepFunctionsIntegration struct {
	_ byte // padding
}

func NewStepFunctionsIntegration() StepFunctionsIntegration {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewStepFunctionsIntegration_Override(s StepFunctionsIntegration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		nil, // no parameters
		s,
	)
}

// Integrates a Synchronous Express State Machine from AWS Step Functions to an API Gateway method.
//
// Example:
//   stateMachine := stepfunctions.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   	stateMachineType: stepfunctions.stateMachineType_EXPRESS,
//   	definition: stepfunctions.chain.start(stepfunctions.NewPass(this, jsii.String("Pass"))),
//   })
//
//   api := apigateway.NewRestApi(this, jsii.String("Api"), &restApiProps{
//   	restApiName: jsii.String("MyApi"),
//   })
//   api.root.addMethod(jsii.String("GET"), apigateway.stepFunctionsIntegration.startExecution(stateMachine))
//
func StepFunctionsIntegration_StartExecution(stateMachine awsstepfunctions.IStateMachine, options *StepFunctionsExecutionIntegrationOptions) AwsIntegration {
	_init_.Initialize()

	var returns AwsIntegration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.StepFunctionsIntegration",
		"startExecution",
		[]interface{}{stateMachine, options},
		&returns,
	)

	return returns
}

