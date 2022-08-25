package awslex


// Specifies settings that manage logging that saves audio, text, and metadata for the conversations with your users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conversationLogSettingsProperty := &conversationLogSettingsProperty{
//   	audioLogSettings: []interface{}{
//   		&audioLogSettingProperty{
//   			destination: &audioLogDestinationProperty{
//   				s3Bucket: &s3BucketLogDestinationProperty{
//   					logPrefix: jsii.String("logPrefix"),
//   					s3BucketArn: jsii.String("s3BucketArn"),
//
//   					// the properties below are optional
//   					kmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   	textLogSettings: []interface{}{
//   		&textLogSettingProperty{
//   			destination: &textLogDestinationProperty{
//   				cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   					cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   					logPrefix: jsii.String("logPrefix"),
//   				},
//   			},
//   			enabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnBot_ConversationLogSettingsProperty struct {
	// Specifies the Amazon S3 settings for logging audio to an S3 bucket.
	AudioLogSettings interface{} `field:"optional" json:"audioLogSettings" yaml:"audioLogSettings"`
	// Specifies settings to enable text conversation logs.
	//
	// You specify the Amazon CloudWatch Logs log group and whether logs should be stored for an alias.
	TextLogSettings interface{} `field:"optional" json:"textLogSettings" yaml:"textLogSettings"`
}

