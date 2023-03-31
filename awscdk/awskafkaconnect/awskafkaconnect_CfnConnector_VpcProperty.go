package awskafkaconnect


// Information about the VPC in which the connector resides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcProperty := &vpcProperty{
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
type CfnConnector_VpcProperty struct {
	// The security groups for the connector.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The subnets for the connector.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

