package awslambda


// The network configuration for the connector.
//
// Specify a VpcEgressConfiguration to enable outbound traffic routing through your VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   configProperty := &ConfigProperty{
//   	VpcEgressConfiguration: &VpcEgressConfigurationProperty{
//   		AssociatedComputeResourceTypes: []*string{
//   			jsii.String("associatedComputeResourceTypes"),
//   		},
//   		NetworkProtocol: jsii.String("networkProtocol"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-config.html
//
type CfnNetworkConnectorPropsMixin_ConfigProperty struct {
	// The VPC egress configuration for the network connector.
	//
	// Specifies the subnets, security groups, and network protocol for routing outbound traffic through your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-networkconnector-config.html#cfn-lambda-networkconnector-config-vpcegressconfiguration
	//
	VpcEgressConfiguration interface{} `field:"optional" json:"vpcEgressConfiguration" yaml:"vpcEgressConfiguration"`
}

