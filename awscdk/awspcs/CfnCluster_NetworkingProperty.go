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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html#cfn-pcs-cluster-networking-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of subnet IDs where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources.
	//
	// The subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or an AWS Local Zone. AWS PCS currently supports only 1 subnet in this list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pcs-cluster-networking.html#cfn-pcs-cluster-networking-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

