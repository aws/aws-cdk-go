package awsbedrockagentcore


// The workload identity details for the agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workloadIdentityDetailsProperty := &WorkloadIdentityDetailsProperty{
//   	WorkloadIdentityArn: jsii.String("workloadIdentityArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-workloadidentitydetails.html
//
type CfnRuntime_WorkloadIdentityDetailsProperty struct {
	// The Amazon Resource Name (ARN) for the workload identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-workloadidentitydetails.html#cfn-bedrockagentcore-runtime-workloadidentitydetails-workloadidentityarn
	//
	WorkloadIdentityArn *string `field:"required" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
}

