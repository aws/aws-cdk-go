package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSkillProperty := &HarnessSkillProperty{
//   	AwsSkills: &HarnessSkillAwsSkillsSourceProperty{
//   		Paths: []*string{
//   			jsii.String("paths"),
//   		},
//   	},
//   	Git: &HarnessSkillGitSourceProperty{
//   		Url: jsii.String("url"),
//
//   		// the properties below are optional
//   		Auth: &HarnessSkillGitAuthProperty{
//   			CredentialArn: jsii.String("credentialArn"),
//
//   			// the properties below are optional
//   			Username: jsii.String("username"),
//   		},
//   		Path: jsii.String("path"),
//   	},
//   	Path: jsii.String("path"),
//   	S3: &HarnessSkillS3SourceProperty{
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html
//
type CfnHarness_HarnessSkillProperty struct {
	// AWS Skills baked into the Harness's underlying Runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html#cfn-bedrockagentcore-harness-harnessskill-awsskills
	//
	AwsSkills interface{} `field:"optional" json:"awsSkills" yaml:"awsSkills"`
	// A git repository containing the skill, cloned over HTTPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html#cfn-bedrockagentcore-harness-harnessskill-git
	//
	Git interface{} `field:"optional" json:"git" yaml:"git"`
	// The filesystem path to the skill definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html#cfn-bedrockagentcore-harness-harnessskill-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// An S3 source containing the skill.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskill.html#cfn-bedrockagentcore-harness-harnessskill-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

