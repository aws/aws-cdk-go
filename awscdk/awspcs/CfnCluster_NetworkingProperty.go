package awspcs


// The networking configuration for the cluster's control plane.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkingProperty := &NetworkingProperty{
//   	NetworkType: jsii.String("networkType"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html
//
type CfnCluster_NetworkingProperty struct {
	// The IP address version the cluster uses.
	//
	// The default is `IPV4` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html#cfn-pcs-cluster-networking-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The list of security group IDs associated with the Elastic Network Interface (ENI) created in subnets.
	//
	// The following rules are required:
	//
	// - Inbound rule 1
	//
	// - Protocol: All
	// - Ports: All
	// - Source: Self
	// - Outbound rule 1
	//
	// - Protocol: All
	// - Ports: All
	// - Destination: 0.0.0.0/0 (IPv4) or ::/0 (IPv6)
	// - Outbound rule 2
	//
	// - Protocol: All
	// - Ports: All
	// - Destination: Self.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html#cfn-pcs-cluster-networking-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources.
	//
	// The subnet must have an available IP address, cannot reside in AWS Outposts , AWS Wavelength , or an AWS Local Zone.
	//
	// Example: `subnet-abcd1234`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html#cfn-pcs-cluster-networking-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

