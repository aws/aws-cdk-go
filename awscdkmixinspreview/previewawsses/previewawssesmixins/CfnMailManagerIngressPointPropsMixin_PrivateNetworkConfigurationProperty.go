package previewawssesmixins


// Specifies the network configuration for the private ingress point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateNetworkConfigurationProperty := &PrivateNetworkConfigurationProperty{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-privatenetworkconfiguration.html
//
type CfnMailManagerIngressPointPropsMixin_PrivateNetworkConfigurationProperty struct {
	// The identifier of the VPC endpoint to associate with this private ingress point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-privatenetworkconfiguration.html#cfn-ses-mailmanageringresspoint-privatenetworkconfiguration-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

