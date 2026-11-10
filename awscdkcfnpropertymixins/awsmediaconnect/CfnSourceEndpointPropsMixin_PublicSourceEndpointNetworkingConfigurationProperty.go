package awsmediaconnect


// The networking configuration for a publicly accessible source endpoint, including the CIDR blocks that are allowed to connect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   publicSourceEndpointNetworkingConfigurationProperty := &PublicSourceEndpointNetworkingConfigurationProperty{
//   	SourceAllowRules: []interface{}{
//   		&PublicSourceEndpointNetworkingConfigurationAllowRuleProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfiguration.html
//
type CfnSourceEndpointPropsMixin_PublicSourceEndpointNetworkingConfigurationProperty struct {
	// The rules that specify which CIDR blocks can connect to the source endpoint.
	//
	// You can specify up to 10 rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfiguration.html#cfn-mediaconnect-sourceendpoint-publicsourceendpointnetworkingconfiguration-sourceallowrules
	//
	SourceAllowRules interface{} `field:"optional" json:"sourceAllowRules" yaml:"sourceAllowRules"`
}

