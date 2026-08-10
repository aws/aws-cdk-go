package awsbedrock


// Contains the configuration of the S3 location of the input data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInvocationJobS3InputDataConfigProperty := &ModelInvocationJobS3InputDataConfigProperty{
//   	S3Uri: jsii.String("s3Uri"),
//
//   	// the properties below are optional
//   	S3BucketOwner: jsii.String("s3BucketOwner"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3inputdataconfig.html
//
type CfnModelInvocationJob_ModelInvocationJobS3InputDataConfigProperty struct {
	// The S3 location of the input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3inputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobs3inputdataconfig-s3uri
	//
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// The ID of the AWS account that owns the S3 bucket containing the input data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3inputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobs3inputdataconfig-s3bucketowner
	//
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
}

