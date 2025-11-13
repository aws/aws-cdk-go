package interfacesawsquicksight


// A reference to a VPCConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCConnectionReference := &VPCConnectionReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   	VpcConnectionId: jsii.String("vpcConnectionId"),
//   }
//
type VPCConnectionReference struct {
	// The AwsAccountId of the VPCConnection resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the VPCConnection resource.
	VpcConnectionArn *string `field:"required" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
	// The VPCConnectionId of the VPCConnection resource.
	VpcConnectionId *string `field:"required" json:"vpcConnectionId" yaml:"vpcConnectionId"`
}

