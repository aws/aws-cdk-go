package awsmediapackagev2


// The base URLs to use for retrieving segments.
//
// You can specify multiple locations and indicate the priority and weight for when each should be used, for use in mutli-CDN workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dashBaseUrlProperty := &DashBaseUrlProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	DvbPriority: jsii.Number(123),
//   	DvbWeight: jsii.Number(123),
//   	ServiceLocation: jsii.String("serviceLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html
//
type CfnOriginEndpoint_DashBaseUrlProperty struct {
	// A source location for segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html#cfn-mediapackagev2-originendpoint-dashbaseurl-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// For use with DVB-DASH profiles only.
	//
	// The priority of this location for servings segments. The lower the number, the higher the priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html#cfn-mediapackagev2-originendpoint-dashbaseurl-dvbpriority
	//
	DvbPriority *float64 `field:"optional" json:"dvbPriority" yaml:"dvbPriority"`
	// For use with DVB-DASH profiles only.
	//
	// The weighting for source locations that have the same priority.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html#cfn-mediapackagev2-originendpoint-dashbaseurl-dvbweight
	//
	DvbWeight *float64 `field:"optional" json:"dvbWeight" yaml:"dvbWeight"`
	// The name of the source location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-dashbaseurl.html#cfn-mediapackagev2-originendpoint-dashbaseurl-servicelocation
	//
	ServiceLocation *string `field:"optional" json:"serviceLocation" yaml:"serviceLocation"`
}

