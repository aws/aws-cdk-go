package awslambda


// The VPC egress configuration for the network connector.
//
// Specifies the subnets, security groups, and network protocol for routing outbound traffic through your VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEgressConfigurationProperty := &VpcEgressConfigurationProperty{
//   	AssociatedComputeResourceTypes: []*string{
//   		jsii.String("associatedComputeResourceTypes"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	NetworkProtocol: jsii.String("networkProtocol"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html
//
type CfnNetworkConnector_VpcEgressConfigurationProperty struct {
	// The types of Lambda compute resources that can use this connector.
	//
	// Currently, only MicroVm is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-associatedcomputeresourcetypes
	//
	AssociatedComputeResourceTypes *[]*string `field:"required" json:"associatedComputeResourceTypes" yaml:"associatedComputeResourceTypes"`
	// The IDs of the VPC subnets where Lambda provisions elastic network interfaces (ENIs).
	//
	// Specify 1 to 16 subnets. All subnets must be in the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The network protocol for the connector.
	//
	// Specify IPv4 for IPv4-only networking, or DualStack for both IPv4 and IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-networkprotocol
	//
	NetworkProtocol *string `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// The IDs of the VPC security groups to attach to the ENIs.
	//
	// Specify 0 to 5 security groups. All security groups must be in the same VPC as the subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

