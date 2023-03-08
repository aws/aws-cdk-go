package awssynthetics


// Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary.
//
// If you omit this parameter, no changes are made to any baseline screenshots that the canary might be using already.
//
// Visual monitoring is supported only on canaries running the *syn-puppeteer-node-3.2* runtime or later. For more information, see [Visual monitoring](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html) and [Visual monitoring blueprint](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   visualReferenceProperty := &visualReferenceProperty{
//   	baseCanaryRunId: jsii.String("baseCanaryRunId"),
//
//   	// the properties below are optional
//   	baseScreenshots: []interface{}{
//   		&baseScreenshotProperty{
//   			screenshotName: jsii.String("screenshotName"),
//
//   			// the properties below are optional
//   			ignoreCoordinates: []*string{
//   				jsii.String("ignoreCoordinates"),
//   			},
//   		},
//   	},
//   }
//
type CfnCanary_VisualReferenceProperty struct {
	// Specifies which canary run to use the screenshots from as the baseline for future visual monitoring with this canary.
	//
	// Valid values are `nextrun` to use the screenshots from the next run after this update is made, `lastrun` to use the screenshots from the most recent run before this update was made, or the value of `Id` in the [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) from any past run of this canary.
	BaseCanaryRunId *string `field:"required" json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// An array of screenshots that are used as the baseline for comparisons during visual monitoring.
	BaseScreenshots interface{} `field:"optional" json:"baseScreenshots" yaml:"baseScreenshots"`
}

