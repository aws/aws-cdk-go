package mixinsawslex


// Configures conversation logging that saves audio, text, and metadata for the conversations with your users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conversationLogSettingsProperty := &ConversationLogSettingsProperty{
//   	AudioLogSettings: []interface{}{
//   		&AudioLogSettingProperty{
//   			Destination: &AudioLogDestinationProperty{
//   				S3Bucket: &S3BucketLogDestinationProperty{
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					LogPrefix: jsii.String("logPrefix"),
//   					S3BucketArn: jsii.String("s3BucketArn"),
//   				},
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	TextLogSettings: []interface{}{
//   		&TextLogSettingProperty{
//   			Destination: &TextLogDestinationProperty{
//   				CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   					CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   					LogPrefix: jsii.String("logPrefix"),
//   				},
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html
//
type CfnBotAliasPropsMixin_ConversationLogSettingsProperty struct {
	// The Amazon S3 settings for logging audio to an S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html#cfn-lex-botalias-conversationlogsettings-audiologsettings
	//
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// The Amazon CloudWatch Logs settings for logging text and metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-conversationlogsettings.html#cfn-lex-botalias-conversationlogsettings-textlogsettings
	//
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

