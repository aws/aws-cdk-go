package awsbedrock


// The Amazon S3 location of the prompt text.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textS3LocationProperty := &TextS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-texts3location.html
//
type CfnPrompt_TextS3LocationProperty struct {
	// The Amazon S3 bucket containing the prompt text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-texts3location.html#cfn-bedrock-prompt-texts3location-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The object key for the Amazon S3 location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-texts3location.html#cfn-bedrock-prompt-texts3location-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The version of the Amazon S3 location to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-texts3location.html#cfn-bedrock-prompt-texts3location-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

