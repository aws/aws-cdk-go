package previewawspinpointmixins


// Specifies how recently segment members were active.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   recencyProperty := &RecencyProperty{
//   	Duration: jsii.String("duration"),
//   	RecencyType: jsii.String("recencyType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-recency.html
//
type CfnSegmentPropsMixin_RecencyProperty struct {
	// The duration to use when determining which users have been active or inactive with your app.
	//
	// Possible values: `HR_24` | `DAY_7` | `DAY_14` | `DAY_30` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-recency.html#cfn-pinpoint-segment-recency-duration
	//
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The type of recency dimension to use for the segment.
	//
	// Valid values are: `ACTIVE` and `INACTIVE` . If the value is `ACTIVE` , the segment includes users who have used your app within the specified duration are included in the segment. If the value is `INACTIVE` , the segment includes users who haven't used your app within the specified duration are included in the segment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-segment-recency.html#cfn-pinpoint-segment-recency-recencytype
	//
	RecencyType *string `field:"optional" json:"recencyType" yaml:"recencyType"`
}

