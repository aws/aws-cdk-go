package awskendra


// Provides the configuration information to connect to an Amazon VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataSourceVpcConfigurationProperty := &dataSourceVpcConfigurationProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnDataSource_DataSourceVpcConfigurationProperty struct {
	// A list of identifiers of security groups within your Amazon VPC.
	//
	// The security groups should enable Amazon Kendra to connect to the data source.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of identifiers for subnets within your Amazon VPC.
	//
	// The subnets should be able to connect to each other in the VPC, and they should have outgoing access to the Internet through a NAT device.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
}

