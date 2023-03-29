package awslex


// Specifies the S3 bucket location where audio logs are stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogDestinationProperty := &AudioLogDestinationProperty{
//   	S3Bucket: &S3BucketLogDestinationProperty{
//   		LogPrefix: jsii.String("logPrefix"),
//   		S3BucketArn: jsii.String("s3BucketArn"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBotAlias_AudioLogDestinationProperty struct {
	// The S3 bucket location where audio logs are stored.
	S3Bucket interface{} `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

