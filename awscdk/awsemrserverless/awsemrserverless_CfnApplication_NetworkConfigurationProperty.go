package awsemrserverless


// The network configuration for customer VPC connectivity.
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
type CfnApplication_NetworkConfigurationProperty struct {
	// The array of security group Ids for customer VPC connectivity.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 32
	//
	// *Pattern* : `^[-0-9a-zA-Z]+`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The array of subnet Ids for customer VPC connectivity.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 32
	//
	// *Pattern* : `^[-0-9a-zA-Z]+`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

