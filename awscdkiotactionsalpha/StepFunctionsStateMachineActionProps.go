package awscdkiotactionsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configuration properties of an action for the Step Functions State Machine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iot_actions_alpha "github.com/aws/aws-cdk-go/awscdkiotactionsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   stepFunctionsStateMachineActionProps := &StepFunctionsStateMachineActionProps{
//   	ExecutionNamePrefix: jsii.String("executionNamePrefix"),
//   	Role: role,
//   }
//
// Experimental.
type StepFunctionsStateMachineActionProps struct {
	// The IAM role that allows access to AWS service.
	// Default: a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Name of the state machine execution prefix.
	//
	// The name given to the state machine execution consists of this prefix followed by a UUID. Step Functions creates a unique name for each state machine execution if one is not provided.
	// See: https://docs.aws.amazon.com/iot/latest/developerguide/stepfunctions-rule-action.html#stepfunctions-rule-action-parameters
	//
	// Default: : None - Step Functions creates a unique name for each state machine execution if one is not provided.
	//
	// Experimental.
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
}

