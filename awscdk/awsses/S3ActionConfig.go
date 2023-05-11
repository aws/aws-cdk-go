package awsses


// S3Action configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ActionConfig := &S3ActionConfig{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	ObjectKeyPrefix: jsii.String("objectKeyPrefix"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type S3ActionConfig struct {
	// The name of the Amazon S3 bucket that you want to send incoming mail to.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The customer master key that Amazon SES should use to encrypt your emails before saving them to the Amazon S3 bucket.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The key prefix of the Amazon S3 bucket.
	ObjectKeyPrefix *string `field:"optional" json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
	// The ARN of the Amazon SNS topic to notify when the message is saved to the Amazon S3 bucket.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

