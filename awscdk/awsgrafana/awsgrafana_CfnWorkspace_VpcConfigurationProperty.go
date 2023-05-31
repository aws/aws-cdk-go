package awsgrafana


// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
//
// > Provided `securityGroupIds` and `subnetIds` must be part of the same VPC.
// >
// > Connecting to a private VPC is not yet available in the Asia Pacific (Seoul) Region (ap-northeast-2).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &VpcConfigurationProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnWorkspace_VpcConfigurationProperty struct {
	// The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.
	//
	// Duplicates not allowed.
	//
	// *Array Members* : Minimum number of 1 items. Maximum number of 5 items.
	//
	// *Length* : Minimum length of 0. Maximum length of 255.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.
	//
	// Duplicates not allowed.
	//
	// *Array Members* : Minimum number of 2 items. Maximum number of 6 items.
	//
	// *Length* : Minimum length of 0. Maximum length of 255.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

