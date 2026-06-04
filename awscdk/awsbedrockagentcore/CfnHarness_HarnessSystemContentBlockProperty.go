package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSystemContentBlockProperty := &HarnessSystemContentBlockProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssystemcontentblock.html
//
type CfnHarness_HarnessSystemContentBlockProperty struct {
	// The text content of the system prompt block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssystemcontentblock.html#cfn-bedrockagentcore-harness-harnesssystemcontentblock-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

