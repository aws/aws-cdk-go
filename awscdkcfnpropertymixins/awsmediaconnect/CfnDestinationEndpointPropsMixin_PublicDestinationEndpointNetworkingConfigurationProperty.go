package awsmediaconnect


// The networking configuration for a publicly accessible destination endpoint, including the CIDR blocks that are allowed to connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   publicDestinationEndpointNetworkingConfigurationProperty := &PublicDestinationEndpointNetworkingConfigurationProperty{
//   	DestinationAllowRules: []interface{}{
//   		&PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfiguration.html
//
type CfnDestinationEndpointPropsMixin_PublicDestinationEndpointNetworkingConfigurationProperty struct {
	// The rules that specify which CIDR blocks can connect to the destination endpoint.
	//
	// You can specify up to 10 rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfiguration.html#cfn-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfiguration-destinationallowrules
	//
	DestinationAllowRules interface{} `field:"optional" json:"destinationAllowRules" yaml:"destinationAllowRules"`
}

