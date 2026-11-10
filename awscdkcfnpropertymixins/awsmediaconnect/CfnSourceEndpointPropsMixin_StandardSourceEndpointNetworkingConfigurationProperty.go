package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   standardSourceEndpointNetworkingConfigurationProperty := &StandardSourceEndpointNetworkingConfigurationProperty{
//   	Public: &PublicSourceEndpointNetworkingConfigurationProperty{
//   		SourceAllowRules: []interface{}{
//   			&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   	},
//   	Vpc: &VpcSourceEndpointNetworkingConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointnetworkingconfiguration.html
//
type CfnSourceEndpointPropsMixin_StandardSourceEndpointNetworkingConfigurationProperty struct {
	// The networking configuration for a publicly accessible source endpoint, including the CIDR blocks that are allowed to connect.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointnetworkingconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointnetworkingconfiguration-public
	//
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// The networking configuration for a source endpoint in a virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-standardsourceendpointnetworkingconfiguration.html#cfn-mediaconnect-sourceendpoint-standardsourceendpointnetworkingconfiguration-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

