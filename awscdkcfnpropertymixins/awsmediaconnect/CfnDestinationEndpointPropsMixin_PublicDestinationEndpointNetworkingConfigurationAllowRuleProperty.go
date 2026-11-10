package awsmediaconnect


// A rule that allows a specific CIDR block to connect to the public destination endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   publicDestinationEndpointNetworkingConfigurationAllowRuleProperty := &PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfigurationallowrule.html
//
type CfnDestinationEndpointPropsMixin_PublicDestinationEndpointNetworkingConfigurationAllowRuleProperty struct {
	// The CIDR block that is allowed to connect to the destination endpoint, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfigurationallowrule.html#cfn-mediaconnect-destinationendpoint-publicdestinationendpointnetworkingconfigurationallowrule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

