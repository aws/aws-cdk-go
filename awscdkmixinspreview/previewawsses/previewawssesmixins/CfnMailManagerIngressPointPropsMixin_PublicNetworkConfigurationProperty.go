package previewawssesmixins


// Specifies the network configuration for the public ingress point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicNetworkConfigurationProperty := &PublicNetworkConfigurationProperty{
//   	IpType: jsii.String("ipType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-publicnetworkconfiguration.html
//
type CfnMailManagerIngressPointPropsMixin_PublicNetworkConfigurationProperty struct {
	// The IP address type for the public ingress point.
	//
	// Valid values are IPV4 and DUAL_STACK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-publicnetworkconfiguration.html#cfn-ses-mailmanageringresspoint-publicnetworkconfiguration-iptype
	//
	IpType *string `field:"optional" json:"ipType" yaml:"ipType"`
}

