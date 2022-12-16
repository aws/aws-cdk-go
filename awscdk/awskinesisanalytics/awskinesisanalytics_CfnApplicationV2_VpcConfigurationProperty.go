package awskinesisanalytics


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
type CfnApplicationV2_VpcConfigurationProperty struct {
	// `CfnApplicationV2.VpcConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnApplicationV2.VpcConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

