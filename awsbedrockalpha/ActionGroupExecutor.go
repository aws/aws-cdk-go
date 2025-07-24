package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Defines how fulfillment of the action group is handled after the necessary information has been elicited from the user.
//
// Valid executors are:
// - Lambda function
// - Return Control.
//
// Example:
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3Schema := bedrock.ApiSchema_FromS3File(bucket, jsii.String("schemas/action-group.yaml"))
//
//   actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
//   })
//
//   actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
//   	Name: jsii.String("query-library"),
//   	Description: jsii.String("Use these functions to get information about the books in the library."),
//   	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
//   	Enabled: jsii.Boolean(true),
//   	ApiSchema: s3Schema,
//   })
//
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
//   	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
//   })
//
//   agent.AddActionGroup(actionGroup)
//
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/action-handle.html
//
// Experimental.
type ActionGroupExecutor interface {
	// The custom control type for the action group executor.
	//
	// Currently only supports 'RETURN_CONTROL' which returns results directly in the InvokeAgent response.
	// Experimental.
	CustomControl() CustomControl
	// The Lambda function that will be called by the action group.
	//
	// Contains the business logic for handling the action group's invocation.
	// Experimental.
	LambdaFunction() awslambda.IFunction
}

// The jsii proxy struct for ActionGroupExecutor
type jsiiProxy_ActionGroupExecutor struct {
	_ byte // padding
}

func (j *jsiiProxy_ActionGroupExecutor) CustomControl() CustomControl {
	var returns CustomControl
	_jsii_.Get(
		j,
		"customControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionGroupExecutor) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}


// Defines an action group with a Lambda function containing the business logic.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-lambda.html
//
// Experimental.
func ActionGroupExecutor_FromLambda(lambdaFunction awslambda.IFunction) ActionGroupExecutor {
	_init_.Initialize()

	if err := validateActionGroupExecutor_FromLambdaParameters(lambdaFunction); err != nil {
		panic(err)
	}
	var returns ActionGroupExecutor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ActionGroupExecutor",
		"fromLambda",
		[]interface{}{lambdaFunction},
		&returns,
	)

	return returns
}

func ActionGroupExecutor_RETURN_CONTROL() ActionGroupExecutor {
	_init_.Initialize()
	var returns ActionGroupExecutor
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.ActionGroupExecutor",
		"RETURN_CONTROL",
		&returns,
	)
	return returns
}

