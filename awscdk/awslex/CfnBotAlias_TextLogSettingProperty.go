package awslex


// Defines settings to enable text conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogSettingProperty := &TextLogSettingProperty{
//   	Destination: &TextLogDestinationProperty{
//   		CloudWatch: &CloudWatchLogGroupLogDestinationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			LogPrefix: jsii.String("logPrefix"),
//   		},
//   	},
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogsetting.html
//
type CfnBotAlias_TextLogSettingProperty struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogsetting.html#cfn-lex-botalias-textlogsetting-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether conversation logs should be stored for an alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-textlogsetting.html#cfn-lex-botalias-textlogsetting-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

