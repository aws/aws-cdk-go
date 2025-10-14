package awsbedrockagentcore


// The workload identity details for the gateway.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-workloadidentitydetails.html
//
type CfnGateway_WorkloadIdentityDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-workloadidentitydetails.html#cfn-bedrockagentcore-gateway-workloadidentitydetails-workloadidentityarn
	//
	WorkloadIdentityArn *string `field:"required" json:"workloadIdentityArn" yaml:"workloadIdentityArn"`
}

