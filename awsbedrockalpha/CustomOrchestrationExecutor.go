package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Contains details about the Lambda function containing the orchestration logic carried out upon invoking the custom orchestration.
//
// Example:
//   orchestrationFunction := lambda.NewFunction(this, jsii.String("OrchestrationFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_10(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(jsii.String("lambda/orchestration")),
//   })
//
//   agent := bedrock.NewAgent(this, jsii.String("CustomOrchestrationAgent"), &AgentProps{
//   	Instruction: jsii.String("You are a helpful assistant with custom orchestration logic."),
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	CustomOrchestrationExecutor: bedrock.CustomOrchestrationExecutor_FromLambda(orchestrationFunction),
//   })
//
// Experimental.
type CustomOrchestrationExecutor interface {
	// The Lambda function that contains the custom orchestration logic.
	//
	// This function is called when the agent needs to make decisions about action execution.
	// Experimental.
	LambdaFunction() awslambda.IFunction
	// The type of orchestration this executor performs.
	// Experimental.
	Type() OrchestrationType
}

// The jsii proxy struct for CustomOrchestrationExecutor
type jsiiProxy_CustomOrchestrationExecutor struct {
	_ byte // padding
}

func (j *jsiiProxy_CustomOrchestrationExecutor) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomOrchestrationExecutor) Type() OrchestrationType {
	var returns OrchestrationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Defines an orchestration executor with a Lambda function containing the business logic.
// Experimental.
func CustomOrchestrationExecutor_FromLambda(lambdaFunction awslambda.IFunction) CustomOrchestrationExecutor {
	_init_.Initialize()

	if err := validateCustomOrchestrationExecutor_FromLambdaParameters(lambdaFunction); err != nil {
		panic(err)
	}
	var returns CustomOrchestrationExecutor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.CustomOrchestrationExecutor",
		"fromLambda",
		[]interface{}{lambdaFunction},
		&returns,
	)

	return returns
}

