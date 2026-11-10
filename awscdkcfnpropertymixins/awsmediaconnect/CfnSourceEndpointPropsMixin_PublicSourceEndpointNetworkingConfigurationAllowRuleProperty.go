package awsmediaconnect


// A rule that allows a specific CIDR block to connect to the public source endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   publicSourceEndpointNetworkingConfigurationAllowRuleProperty := &PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfigurationallowrule.html
//
type CfnSourceEndpointPropsMixin_PublicSourceEndpointNetworkingConfigurationAllowRuleProperty struct {
	// The CIDR block that is allowed to connect to the source endpoint, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfigurationallowrule.html#cfn-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfigurationallowrule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

