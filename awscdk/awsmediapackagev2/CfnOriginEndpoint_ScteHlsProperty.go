package awsmediapackagev2


// The SCTE-35 HLS configuration associated with the origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scteHlsProperty := &ScteHlsProperty{
//   	AdMarkerHls: jsii.String("adMarkerHls"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctehls.html
//
type CfnOriginEndpoint_ScteHlsProperty struct {
	// The SCTE-35 HLS ad-marker configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctehls.html#cfn-mediapackagev2-originendpoint-sctehls-admarkerhls
	//
	AdMarkerHls *string `field:"optional" json:"adMarkerHls" yaml:"adMarkerHls"`
}

