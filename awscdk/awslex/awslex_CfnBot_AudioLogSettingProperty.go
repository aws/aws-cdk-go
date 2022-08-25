package awslex


// Specifies settings for logging the audio of conversations between Amazon Lex and a user.
//
// You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogSettingProperty := &audioLogSettingProperty{
//   	destination: &audioLogDestinationProperty{
//   		s3Bucket: &s3BucketLogDestinationProperty{
//   			logPrefix: jsii.String("logPrefix"),
//   			s3BucketArn: jsii.String("s3BucketArn"),
//
//   			// the properties below are optional
//   			kmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBot_AudioLogSettingProperty struct {
	// Specifies the location of the audio log files collected when conversation logging is enabled for a bot.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether audio logging is enabled for the bot.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

