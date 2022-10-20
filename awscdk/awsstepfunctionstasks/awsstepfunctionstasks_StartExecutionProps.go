package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for StartExecution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var input interface{}
//
//   startExecutionProps := &startExecutionProps{
//   	input: map[string]interface{}{
//   		"inputKey": input,
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	name: jsii.String("name"),
//   }
//
// Deprecated: - use 'StepFunctionsStartExecution'.
type StartExecutionProps struct {
	// The JSON input for the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'.
	Input *map[string]interface{} `field:"optional" json:"input" yaml:"input"`
	// The service integration pattern indicates different ways to call StartExecution to Step Functions.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// The name of the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

