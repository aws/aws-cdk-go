package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   standardDestinationEndpointNetworkingConfigurationProperty := &StandardDestinationEndpointNetworkingConfigurationProperty{
//   	Public: &PublicDestinationEndpointNetworkingConfigurationProperty{
//   		DestinationAllowRules: []interface{}{
//   			&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   	},
//   	Vpc: &VpcDestinationEndpointNetworkingConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointnetworkingconfiguration.html
//
type CfnDestinationEndpointPropsMixin_StandardDestinationEndpointNetworkingConfigurationProperty struct {
	// The networking configuration for a publicly accessible destination endpoint, including the CIDR blocks that are allowed to connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointnetworkingconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointnetworkingconfiguration-public
	//
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// The networking configuration for a destination endpoint in a virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-standarddestinationendpointnetworkingconfiguration.html#cfn-mediaconnect-destinationendpoint-standarddestinationendpointnetworkingconfiguration-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

