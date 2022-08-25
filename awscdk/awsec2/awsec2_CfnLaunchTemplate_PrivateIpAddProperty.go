package awsec2


// Specifies a secondary private IPv4 address for a network interface.
//
// `PrivateIpAdd` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddProperty := &privateIpAddProperty{
//   	primary: jsii.Boolean(false),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   }
//
type CfnLaunchTemplate_PrivateIpAddProperty struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// The private IPv4 addresses.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

