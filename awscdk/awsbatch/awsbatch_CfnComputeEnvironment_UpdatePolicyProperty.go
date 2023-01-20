package awsbatch


// Specifies the infrastructure update policy for the compute environment.
//
// For more information about infrastructure updates, see [Infrastructure updates](https://docs.aws.amazon.com/batch/latest/userguide/infrastructure-updates.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updatePolicyProperty := &updatePolicyProperty{
//   	jobExecutionTimeoutMinutes: jsii.Number(123),
//   	terminateJobsOnUpdate: jsii.Boolean(false),
//   }
//
type CfnComputeEnvironment_UpdatePolicyProperty struct {
	// Specifies the job timeout, in minutes, when the compute environment infrastructure is updated.
	//
	// The default value is 30.
	JobExecutionTimeoutMinutes *float64 `field:"optional" json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// Specifies whether jobs are automatically terminated when the computer environment infrastructure is updated.
	//
	// The default value is `false` .
	TerminateJobsOnUpdate interface{} `field:"optional" json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
}

