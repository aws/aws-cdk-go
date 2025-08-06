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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html
//
type CfnBot_AudioLogSettingProperty struct {
	// Specifies the location of the audio log files collected when conversation logging is enabled for a bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html#cfn-lex-bot-audiologsetting-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether audio logging in enabled for the bot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-audiologsetting.html#cfn-lex-bot-audiologsetting-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

