package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSkillProperty := &HarnessSkillProperty{
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html
//
type CfnHarness_HarnessSkillProperty struct {
	// The filesystem path to the skill definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html#cfn-bedrockagentcore-harness-harnessskill-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
}

