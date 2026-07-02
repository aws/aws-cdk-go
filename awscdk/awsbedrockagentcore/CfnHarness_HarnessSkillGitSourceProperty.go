package awsbedrockagentcore


// A git repository containing the skill, cloned over HTTPS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSkillGitSourceProperty := &HarnessSkillGitSourceProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	Auth: &HarnessSkillGitAuthProperty{
//   		CredentialArn: jsii.String("credentialArn"),
//
//   		// the properties below are optional
//   		Username: jsii.String("username"),
//   	},
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html
//
type CfnHarness_HarnessSkillGitSourceProperty struct {
	// The HTTPS URL of the git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// Authentication configuration for accessing a private git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// Subdirectory within the repository containing the skill.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}

