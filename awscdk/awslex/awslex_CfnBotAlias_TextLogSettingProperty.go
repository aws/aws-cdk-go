package awslex


// Defines settings to enable conversation logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textLogSettingProperty := &textLogSettingProperty{
//   	destination: &textLogDestinationProperty{
//   		cloudWatch: &cloudWatchLogGroupLogDestinationProperty{
//   			cloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			logPrefix: jsii.String("logPrefix"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnBotAlias_TextLogSettingProperty struct {
	// Defines the Amazon CloudWatch Logs destination log group for conversation text logs.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether conversation logs should be stored for an alias.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

