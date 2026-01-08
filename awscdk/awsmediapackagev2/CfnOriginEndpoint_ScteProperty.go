package awsmediapackagev2


// The SCTE-35 configuration associated with the origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scteProperty := &ScteProperty{
//   	ScteFilter: []*string{
//   		jsii.String("scteFilter"),
//   	},
//   	ScteInSegments: jsii.String("scteInSegments"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-scte.html
//
type CfnOriginEndpoint_ScteProperty struct {
	// The filter associated with the SCTE-35 configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-scte.html#cfn-mediapackagev2-originendpoint-scte-sctefilter
	//
	ScteFilter *[]*string `field:"optional" json:"scteFilter" yaml:"scteFilter"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-scte.html#cfn-mediapackagev2-originendpoint-scte-scteinsegments
	//
	ScteInSegments *string `field:"optional" json:"scteInSegments" yaml:"scteInSegments"`
}

