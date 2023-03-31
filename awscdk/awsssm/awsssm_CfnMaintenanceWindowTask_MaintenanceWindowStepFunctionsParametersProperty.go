package awsssm


// The `MaintenanceWindowStepFunctionsParameters` property type specifies the parameters for the execution of a `STEP_FUNCTIONS` task in a Systems Manager maintenance window.
//
// `MaintenanceWindowStepFunctionsParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceWindowStepFunctionsParametersProperty := &maintenanceWindowStepFunctionsParametersProperty{
//   	input: jsii.String("input"),
//   	name: jsii.String("name"),
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowStepFunctionsParametersProperty struct {
	// The inputs for the `STEP_FUNCTIONS` task.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The name of the `STEP_FUNCTIONS` task.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

