package previewawssyntheticsmixins


// Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary.
//
// If you omit this parameter, no changes are made to any baseline screenshots that the canary might be using already.
//
// Visual monitoring is supported only on canaries running the *syn-puppeteer-node-3.2* runtime or later. For more information, see [Visual monitoring](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html) and [Visual monitoring blueprint](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   visualReferenceProperty := &VisualReferenceProperty{
//   	BaseCanaryRunId: jsii.String("baseCanaryRunId"),
//   	BaseScreenshots: []interface{}{
//   		&BaseScreenshotProperty{
//   			IgnoreCoordinates: []*string{
//   				jsii.String("ignoreCoordinates"),
//   			},
//   			ScreenshotName: jsii.String("screenshotName"),
//   		},
//   	},
//   	BrowserType: jsii.String("browserType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-visualreference.html
//
type CfnCanaryPropsMixin_VisualReferenceProperty struct {
	// Specifies which canary run to use the screenshots from as the baseline for future visual monitoring with this canary.
	//
	// Valid values are `nextrun` to use the screenshots from the next run after this update is made, `lastrun` to use the screenshots from the most recent run before this update was made, or the value of `Id` in the [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) from any past run of this canary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-visualreference.html#cfn-synthetics-canary-visualreference-basecanaryrunid
	//
	BaseCanaryRunId *string `field:"optional" json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// An array of screenshots that are used as the baseline for comparisons during visual monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-visualreference.html#cfn-synthetics-canary-visualreference-basescreenshots
	//
	BaseScreenshots interface{} `field:"optional" json:"baseScreenshots" yaml:"baseScreenshots"`
	// The browser type associated with this visual reference configuration.
	//
	// Valid values are `CHROME` and `FIREFOX` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-visualreference.html#cfn-synthetics-canary-visualreference-browsertype
	//
	BrowserType *string `field:"optional" json:"browserType" yaml:"browserType"`
}

