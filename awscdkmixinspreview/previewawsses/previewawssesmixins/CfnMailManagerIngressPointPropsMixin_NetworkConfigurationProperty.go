package previewawssesmixins


// The network type (IPv4-only, Dual-Stack, PrivateLink) of the ingress endpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkConfigurationProperty := &NetworkConfigurationProperty{
//   	PrivateNetworkConfiguration: &PrivateNetworkConfigurationProperty{
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   	},
//   	PublicNetworkConfiguration: &PublicNetworkConfigurationProperty{
//   		IpType: jsii.String("ipType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-networkconfiguration.html
//
type CfnMailManagerIngressPointPropsMixin_NetworkConfigurationProperty struct {
	// Specifies the network configuration for the private ingress point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-networkconfiguration.html#cfn-ses-mailmanageringresspoint-networkconfiguration-privatenetworkconfiguration
	//
	PrivateNetworkConfiguration interface{} `field:"optional" json:"privateNetworkConfiguration" yaml:"privateNetworkConfiguration"`
	// Specifies the network configuration for the public ingress point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-networkconfiguration.html#cfn-ses-mailmanageringresspoint-networkconfiguration-publicnetworkconfiguration
	//
	PublicNetworkConfiguration interface{} `field:"optional" json:"publicNetworkConfiguration" yaml:"publicNetworkConfiguration"`
}

