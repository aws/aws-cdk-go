package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNetworkConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnNetworkConnectorMixinProps := &CfnNetworkConnectorMixinProps{
//   	Configuration: &ConfigProperty{
//   		VpcEgressConfiguration: &VpcEgressConfigurationProperty{
//   			AssociatedComputeResourceTypes: []*string{
//   				jsii.String("associatedComputeResourceTypes"),
//   			},
//   			NetworkProtocol: jsii.String("networkProtocol"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OperatorRole: jsii.String("operatorRole"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-networkconnector.html
//
type CfnNetworkConnectorMixinProps struct {
	// The network configuration for the connector.
	//
	// Specify a VpcEgressConfiguration to enable outbound traffic routing through your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-networkconnector.html#cfn-lambda-networkconnector-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// A unique name for the network connector within your account and Region.
	//
	// Must be 1 to 64 alphanumeric characters, hyphens, or underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-networkconnector.html#cfn-lambda-networkconnector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the IAM role that Lambda assumes to manage elastic network interfaces in your VPC.
	//
	// This role must have permissions for ec2:CreateNetworkInterface and related describe operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-networkconnector.html#cfn-lambda-networkconnector-operatorrole
	//
	OperatorRole *string `field:"optional" json:"operatorRole" yaml:"operatorRole"`
	// A list of tags to apply to the network connector.
	//
	// Use tags to categorize network connectors for cost allocation, access control, or operational management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-networkconnector.html#cfn-lambda-networkconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

