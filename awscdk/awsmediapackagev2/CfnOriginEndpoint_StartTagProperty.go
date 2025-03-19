package awsmediapackagev2


// To insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset.
//
// When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   startTagProperty := &StartTagProperty{
//   	TimeOffset: jsii.Number(123),
//
//   	// the properties below are optional
//   	Precise: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-starttag.html
//
type CfnOriginEndpoint_StartTagProperty struct {
	// Specify the value for TIME-OFFSET within your EXT-X-START tag.
	//
	// Enter a signed floating point value which, if positive, must be less than the configured manifest duration minus three times the configured segment target duration. If negative, the absolute value must be larger than three times the configured segment target duration, and the absolute value must be smaller than the configured manifest duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-starttag.html#cfn-mediapackagev2-originendpoint-starttag-timeoffset
	//
	TimeOffset *float64 `field:"required" json:"timeOffset" yaml:"timeOffset"`
	// Specify the value for PRECISE within your EXT-X-START tag.
	//
	// Leave blank, or choose false, to use the default value NO. Choose yes to use the value YES.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-starttag.html#cfn-mediapackagev2-originendpoint-starttag-precise
	//
	Precise interface{} `field:"optional" json:"precise" yaml:"precise"`
}

