package awskinesisanalyticsv2


// Describes the parameters of a VPC used by the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnApplication_VpcConfigurationProperty struct {
	// The array of [SecurityGroup](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html) IDs used by the VPC configuration.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The array of [Subnet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Subnet.html) IDs used by the VPC configuration.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

