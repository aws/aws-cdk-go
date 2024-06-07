package awsmediapackagev2


// <p>The SCTE configuration.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scteDashProperty := &ScteDashProperty{
//   	AdMarkerDash: jsii.String("adMarkerDash"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctedash.html
//
type CfnOriginEndpoint_ScteDashProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-sctedash.html#cfn-mediapackagev2-originendpoint-sctedash-admarkerdash
	//
	AdMarkerDash *string `field:"optional" json:"adMarkerDash" yaml:"adMarkerDash"`
}

