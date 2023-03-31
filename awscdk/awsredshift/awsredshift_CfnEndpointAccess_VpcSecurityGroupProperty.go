package awsredshift


// The security groups associated with the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSecurityGroupProperty := &vpcSecurityGroupProperty{
//   	status: jsii.String("status"),
//   	vpcSecurityGroupId: jsii.String("vpcSecurityGroupId"),
//   }
//
type CfnEndpointAccess_VpcSecurityGroupProperty struct {
	// The status of the endpoint.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The identifier of the VPC security group.
	VpcSecurityGroupId *string `field:"optional" json:"vpcSecurityGroupId" yaml:"vpcSecurityGroupId"`
}

