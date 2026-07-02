package awsbedrockagentcore


// AWS Skills baked into the Harness's underlying Runtime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSkillAwsSkillsSourceProperty := &HarnessSkillAwsSkillsSourceProperty{
//   	Paths: []*string{
//   		jsii.String("paths"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillawsskillssource.html
//
type CfnHarness_HarnessSkillAwsSkillsSourceProperty struct {
	// Optionally filter allowed skills with glob syntax, e.g., ['core-skills/*'].
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillawsskillssource.html#cfn-bedrockagentcore-harness-harnessskillawsskillssource-paths
	//
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

