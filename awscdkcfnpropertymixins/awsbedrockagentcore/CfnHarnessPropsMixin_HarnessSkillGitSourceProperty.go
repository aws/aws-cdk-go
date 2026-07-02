package awsbedrockagentcore


// A git repository containing the skill, cloned over HTTPS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessSkillGitSourceProperty := &HarnessSkillGitSourceProperty{
//   	Auth: &HarnessSkillGitAuthProperty{
//   		CredentialArn: jsii.String("credentialArn"),
//   		Username: jsii.String("username"),
//   	},
//   	Path: jsii.String("path"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html
//
type CfnHarnessPropsMixin_HarnessSkillGitSourceProperty struct {
	// Authentication configuration for accessing a private git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-auth
	//
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// Subdirectory within the repository containing the skill.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTPS URL of the git repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitsource.html#cfn-bedrockagentcore-harness-harnessskillgitsource-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}

