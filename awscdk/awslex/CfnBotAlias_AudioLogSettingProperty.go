package awslex


// Settings for logging audio of conversations between Amazon Lex and a user.
//
// You specify whether to log audio and the Amazon S3 bucket where the audio file is stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioLogSettingProperty := &AudioLogSettingProperty{
//   	Destination: &AudioLogDestinationProperty{
//   		S3Bucket: &S3BucketLogDestinationProperty{
//   			LogPrefix: jsii.String("logPrefix"),
//   			S3BucketArn: jsii.String("s3BucketArn"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
type CfnBotAlias_AudioLogSettingProperty struct {
	// The location of audio log files collected when conversation logging is enabled for a bot.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether audio logging in enabled for the bot.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

