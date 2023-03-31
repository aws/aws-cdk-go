package awsssm


// The `MaintenanceWindowAutomationParameters` property type specifies the parameters for an `AUTOMATION` task type for a maintenance window task in AWS Systems Manager .
//
// `MaintenanceWindowAutomationParameters` is a property of the [TaskInvocationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html) property type.
//
// For information about available parameters in Automation runbooks, you can view the content of the runbook itself in the Systems Manager console. For information, see [View runbook content](https://docs.aws.amazon.com/systems-manager/latest/userguide/automation-documents-reference-details.html#view-automation-json) in the *AWS Systems Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   maintenanceWindowAutomationParametersProperty := &maintenanceWindowAutomationParametersProperty{
//   	documentVersion: jsii.String("documentVersion"),
//   	parameters: parameters,
//   }
//
type CfnMaintenanceWindowTask_MaintenanceWindowAutomationParametersProperty struct {
	// The version of an Automation runbook to use during task execution.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// The parameters for the AUTOMATION task.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

