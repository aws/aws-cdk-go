package awsbedrock


// Details about the location of the output of the batch inference job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelInvocationJobOutputDataConfigProperty := &ModelInvocationJobOutputDataConfigProperty{
//   	S3OutputDataConfig: &ModelInvocationJobS3OutputDataConfigProperty{
//   		S3Uri: jsii.String("s3Uri"),
//
//   		// the properties below are optional
//   		S3BucketOwner: jsii.String("s3BucketOwner"),
//   		S3EncryptionKeyId: jsii.String("s3EncryptionKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjoboutputdataconfig.html
//
type CfnModelInvocationJob_ModelInvocationJobOutputDataConfigProperty struct {
	// Contains the configuration of the S3 location of the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjoboutputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjoboutputdataconfig-s3outputdataconfig
	//
	S3OutputDataConfig interface{} `field:"required" json:"s3OutputDataConfig" yaml:"s3OutputDataConfig"`
}

