package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for RunLambdaTask.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var taskInput taskInput
//
//   runLambdaTaskProps := &runLambdaTaskProps{
//   	clientContext: jsii.String("clientContext"),
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	invocationType: awscdk.Aws_stepfunctions_tasks.invocationType_REQUEST_RESPONSE,
//   	payload: taskInput,
//   	qualifier: jsii.String("qualifier"),
//   }
//
// Deprecated: Use `LambdaInvoke`.
type RunLambdaTaskProps struct {
	// Client context to pass to the function.
	// Deprecated: Use `LambdaInvoke`.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// The service integration pattern indicates different ways to invoke Lambda function.
	//
	// The valid value for Lambda is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN,
	// it determines whether to pause the workflow until a task token is returned.
	//
	// If this is set to WAIT_FOR_TASK_TOKEN, the JsonPath.taskToken value must be included
	// somewhere in the payload and the Lambda must call
	// `SendTaskSuccess/SendTaskFailure` using that token.
	// Deprecated: Use `LambdaInvoke`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Invocation type of the Lambda function.
	// Deprecated: Use `LambdaInvoke`.
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The JSON that you want to provide to your Lambda function as input.
	// Deprecated: Use `LambdaInvoke`.
	Payload awsstepfunctions.TaskInput `field:"optional" json:"payload" yaml:"payload"`
	// Version or alias of the function to be invoked.
	// Deprecated: pass a Version or Alias object as lambdaFunction instead.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

