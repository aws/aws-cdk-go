package awsmsk


// Specifies information about subnets and security groups for the VPC that your clients will use to connect with the serverless cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnServerlessCluster_VpcConfigProperty struct {
	// A list of subnets in at least two different Availability Zones that host your client applications.
	//
	// We recommend that you specify a backup subnet in a different Availability Zone for failover in case of an outage.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Specifies up to five security groups that control inbound and outbound traffic for the serverless cluster.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

