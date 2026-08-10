package awsbedrock


// Contains the configuration of the S3 location of the output data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   modelInvocationJobS3OutputDataConfigProperty := &ModelInvocationJobS3OutputDataConfigProperty{
//   	S3BucketOwner: jsii.String("s3BucketOwner"),
//   	S3EncryptionKeyId: jsii.String("s3EncryptionKeyId"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig.html
//
type CfnModelInvocationJobPropsMixin_ModelInvocationJobS3OutputDataConfigProperty struct {
	// The ID of the AWS account that owns the S3 bucket containing the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig-s3bucketowner
	//
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// The unique identifier of the key that encrypts the S3 location of the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig-s3encryptionkeyid
	//
	S3EncryptionKeyId *string `field:"optional" json:"s3EncryptionKeyId" yaml:"s3EncryptionKeyId"`
	// The S3 location of the output data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig.html#cfn-bedrock-modelinvocationjob-modelinvocationjobs3outputdataconfig-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

