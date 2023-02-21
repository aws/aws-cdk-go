package awsmwaa


// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
//
// To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkConfigurationProperty := &networkConfigurationProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnEnvironment_NetworkConfigurationProperty struct {
	// A list of one or more security group IDs.
	//
	// Accepts up to 5 security group IDs. A security group must be attached to the same VPC as the subnets. To learn more, see [Security in your VPC on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/vpc-security.html) .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs.
	//
	// *Required* to create an environment. Must be private subnets in two different availability zones. A subnet must be attached to the same VPC as the security group. To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

