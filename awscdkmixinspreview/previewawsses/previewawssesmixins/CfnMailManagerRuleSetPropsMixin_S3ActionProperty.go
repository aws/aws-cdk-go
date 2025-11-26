package previewawssesmixins


// Writes the MIME content of the email to an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ActionProperty := &S3ActionProperty{
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	RoleArn: jsii.String("roleArn"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Prefix: jsii.String("s3Prefix"),
//   	S3SseKmsKeyId: jsii.String("s3SseKmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html
//
type CfnMailManagerRuleSetPropsMixin_S3ActionProperty struct {
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the specified the bucket has been deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html#cfn-ses-mailmanagerruleset-s3action-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// The Amazon Resource Name (ARN) of the IAM Role to use while writing to S3.
	//
	// This role must have access to the s3:PutObject, kms:Encrypt, and kms:GenerateDataKey APIs for the given bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html#cfn-ses-mailmanagerruleset-s3action-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The bucket name of the S3 bucket to write to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html#cfn-ses-mailmanagerruleset-s3action-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 prefix to use for the write to the s3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html#cfn-ses-mailmanagerruleset-s3action-s3prefix
	//
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// The KMS Key ID to use to encrypt the message in S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-s3action.html#cfn-ses-mailmanagerruleset-s3action-s3ssekmskeyid
	//
	S3SseKmsKeyId *string `field:"optional" json:"s3SseKmsKeyId" yaml:"s3SseKmsKeyId"`
}

