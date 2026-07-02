package awsbedrockagentcore


// Authentication configuration for accessing a private git repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessSkillGitAuthProperty := &HarnessSkillGitAuthProperty{
//   	CredentialArn: jsii.String("credentialArn"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitauth.html
//
type CfnHarnessPropsMixin_HarnessSkillGitAuthProperty struct {
	// The ARN of the credential in AgentCore Identity containing the password or personal access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitauth.html#cfn-bedrockagentcore-harness-harnessskillgitauth-credentialarn
	//
	CredentialArn *string `field:"optional" json:"credentialArn" yaml:"credentialArn"`
	// Username for authentication.
	//
	// Defaults to 'oauth2' if not specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskillgitauth.html#cfn-bedrockagentcore-harness-harnessskillgitauth-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}

