package mixinsawsbatch


// Specifies the infrastructure update policy for the Amazon EC2 compute environment.
//
// For more information about infrastructure updates, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   updatePolicyProperty := &UpdatePolicyProperty{
//   	JobExecutionTimeoutMinutes: jsii.Number(123),
//   	TerminateJobsOnUpdate: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-updatepolicy.html
//
type CfnComputeEnvironmentPropsMixin_UpdatePolicyProperty struct {
	// Specifies the job timeout (in minutes) when the compute environment infrastructure is updated.
	//
	// The default value is 30.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-updatepolicy.html#cfn-batch-computeenvironment-updatepolicy-jobexecutiontimeoutminutes
	//
	// Default: - 30.
	//
	JobExecutionTimeoutMinutes *float64 `field:"optional" json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// Specifies whether jobs are automatically terminated when the compute environment infrastructure is updated.
	//
	// The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-updatepolicy.html#cfn-batch-computeenvironment-updatepolicy-terminatejobsonupdate
	//
	// Default: - false.
	//
	TerminateJobsOnUpdate interface{} `field:"optional" json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
}

