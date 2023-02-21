package awss3outposts


// Contains the virtual private cloud (VPC) configuration for the specified access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnAccessPoint_VpcConfigurationProperty struct {
	// The ID of the VPC configuration.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

