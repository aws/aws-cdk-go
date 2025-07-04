package awssynthetics


// A structure representing a screenshot that is used as a baseline during visual monitoring comparisons made by the canary.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseScreenshotProperty := &BaseScreenshotProperty{
//   	ScreenshotName: jsii.String("screenshotName"),
//
//   	// the properties below are optional
//   	IgnoreCoordinates: []*string{
//   		jsii.String("ignoreCoordinates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-basescreenshot.html
//
type CfnCanary_BaseScreenshotProperty struct {
	// The name of the screenshot.
	//
	// This is generated the first time the canary is run after the `UpdateCanary` operation that specified for this canary to perform visual monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-basescreenshot.html#cfn-synthetics-canary-basescreenshot-screenshotname
	//
	ScreenshotName *string `field:"required" json:"screenshotName" yaml:"screenshotName"`
	// Coordinates that define the part of a screen to ignore during screenshot comparisons.
	//
	// To obtain the coordinates to use here, use the CloudWatch console to draw the boundaries on the screen. For more information, see [Edit or delete a canary](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/synthetics_canaries_deletion.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-basescreenshot.html#cfn-synthetics-canary-basescreenshot-ignorecoordinates
	//
	IgnoreCoordinates *[]*string `field:"optional" json:"ignoreCoordinates" yaml:"ignoreCoordinates"`
}

