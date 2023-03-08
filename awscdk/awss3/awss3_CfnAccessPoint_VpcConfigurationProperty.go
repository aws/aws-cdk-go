package awss3


// The Virtual Private Cloud (VPC) configuration for this access point.
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
	// If this field is specified, the access point will only allow connections from the specified VPC ID.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

