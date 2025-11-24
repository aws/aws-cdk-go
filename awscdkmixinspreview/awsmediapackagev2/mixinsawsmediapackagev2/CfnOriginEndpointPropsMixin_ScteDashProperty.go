package mixinsawsmediapackagev2


// The SCTE configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scteDashProperty := &ScteDashProperty{
//   	AdMarkerDash: jsii.String("adMarkerDash"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctedash.html
//
type CfnOriginEndpointPropsMixin_ScteDashProperty struct {
	// Choose how ad markers are included in the packaged content.
	//
	// If you include ad markers in the content stream in your upstream encoders, then you need to inform MediaPackage what to do with the ad markers in the output.
	//
	// Value description:
	//
	// - `Binary` - The SCTE-35 marker is expressed as a hex-string (Base64 string) rather than full XML.
	// - `XML` - The SCTE marker is expressed fully in XML.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctedash.html#cfn-mediapackagev2-originendpoint-sctedash-admarkerdash
	//
	AdMarkerDash *string `field:"optional" json:"adMarkerDash" yaml:"adMarkerDash"`
}

