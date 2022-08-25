package awslex


// Specifies the location of audio log files collected when conversation logging is enabled for a bot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogDestinationProperty := &audioLogDestinationProperty{
//   	s3Bucket: &s3BucketLogDestinationProperty{
//   		logPrefix: jsii.String("logPrefix"),
//   		s3BucketArn: jsii.String("s3BucketArn"),
//
//   		// the properties below are optional
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
type CfnBot_AudioLogDestinationProperty struct {
	// Specifies the Amazon S3 bucket where the audio files are stored.
	S3Bucket interface{} `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

