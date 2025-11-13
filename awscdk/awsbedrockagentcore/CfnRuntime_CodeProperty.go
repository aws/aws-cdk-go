package awsbedrockagentcore


// Object represents source code from zip file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   codeProperty := &CodeProperty{
//   	S3: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Prefix: jsii.String("prefix"),
//
//   		// the properties below are optional
//   		VersionId: jsii.String("versionId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-code.html
//
type CfnRuntime_CodeProperty struct {
	// S3 Location Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-code.html#cfn-bedrockagentcore-runtime-code-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

