package awslex


// Configures conversation logging that saves audio, text, and metadata for the conversations with your users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conversationLogSettingsProperty := &ConversationLogSettingsProperty{
//   	AudioLogSettings: []interface{}{
//   		&AudioLogSettingProperty{
//   			Destination: &AudioLogDestinationProperty{
//   				S3Bucket: &S3BucketLogDestinationProperty{
//   					LogPrefix: jsii.String("logPrefix"),
//   					S3BucketArn: jsii.String("s3BucketArn"),
//
//   					// the properties below are optional
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
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
type CfnBotAlias_ConversationLogSettingsProperty struct {
	// The Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// The Amazon CloudWatch Logs settings for logging text and metadata.
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

