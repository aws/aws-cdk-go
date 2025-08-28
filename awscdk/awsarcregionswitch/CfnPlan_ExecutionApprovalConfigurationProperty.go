package awsarcregionswitch


// Configuration for approval steps in a Region switch plan execution.
//
// Approval steps require manual intervention before the execution can proceed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionApprovalConfigurationProperty := &ExecutionApprovalConfigurationProperty{
//   	ApprovalRole: jsii.String("approvalRole"),
//
//   	// the properties below are optional
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionapprovalconfiguration.html
//
type CfnPlan_ExecutionApprovalConfigurationProperty struct {
	// The IAM approval role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionapprovalconfiguration.html#cfn-arcregionswitch-plan-executionapprovalconfiguration-approvalrole
	//
	ApprovalRole *string `field:"required" json:"approvalRole" yaml:"approvalRole"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-executionapprovalconfiguration.html#cfn-arcregionswitch-plan-executionapprovalconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

