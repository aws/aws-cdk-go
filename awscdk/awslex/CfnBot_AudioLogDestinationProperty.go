package awslex


// The location of audio log files collected when conversation logging is enabled for a bot.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologdestination.html
//
type CfnBot_AudioLogDestinationProperty struct {
	// Specifies the Amazon S3 bucket where the audio files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologdestination.html#cfn-lex-bot-audiologdestination-s3bucket
	//
	S3Bucket interface{} `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
}

