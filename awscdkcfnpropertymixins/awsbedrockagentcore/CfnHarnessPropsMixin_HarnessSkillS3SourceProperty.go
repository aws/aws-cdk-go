package awsbedrockagentcore


// An S3 source containing the skill.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessSkillS3SourceProperty := &HarnessSkillS3SourceProperty{
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskills3source.html
//
type CfnHarnessPropsMixin_HarnessSkillS3SourceProperty struct {
	// The S3 URI pointing to the skill directory (e.g., s3://bucket/skills/my-skill/).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessskills3source.html#cfn-bedrockagentcore-harness-harnessskills3source-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

