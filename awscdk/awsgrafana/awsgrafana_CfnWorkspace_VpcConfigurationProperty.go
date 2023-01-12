package awsgrafana


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
type CfnWorkspace_VpcConfigurationProperty struct {
	// `CfnWorkspace.VpcConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnWorkspace.VpcConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

