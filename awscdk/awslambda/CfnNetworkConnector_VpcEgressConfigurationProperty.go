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
//   	NetworkProtocol: jsii.String("networkProtocol"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	ElasticIpConfiguration: &ElasticIpConfigurationProperty{
//   		Type: jsii.String("type"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html
//
type CfnNetworkConnector_VpcEgressConfigurationProperty struct {
	// The types of Lambda compute resources that can use this connector.
	//
	// Supported values are MicroVm and WebFunction.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-associatedcomputeresourcetypes
	//
	AssociatedComputeResourceTypes *[]*string `field:"required" json:"associatedComputeResourceTypes" yaml:"associatedComputeResourceTypes"`
	// The network protocol for the connector.
	//
	// Specify IPv4 for IPv4-only networking, or DualStack for both IPv4 and IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-networkprotocol
	//
	NetworkProtocol *string `field:"required" json:"networkProtocol" yaml:"networkProtocol"`
	// The IDs of the VPC security groups to attach to the ENIs.
	//
	// All security groups must be in the same VPC as the subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the VPC subnets where Lambda provisions elastic network interfaces (ENIs).
	//
	// Specify 1 to 16 subnets. All subnets must be in the same VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Specifies whether Lambda assigns an Elastic IP address to each of the network connector's elastic network interfaces (ENIs).
	//
	// If you don't specify an Elastic IP configuration, no Elastic IP address is assigned. When you update this configuration, the change applies to the ENIs of the new connector version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-vpcegressconfiguration.html#cfn-lambda-networkconnector-vpcegressconfiguration-elasticipconfiguration
	//
	ElasticIpConfiguration interface{} `field:"optional" json:"elasticIpConfiguration" yaml:"elasticIpConfiguration"`
}

