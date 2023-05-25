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
type CfnBot_TextLogSettingProperty struct {
	// Specifies the Amazon CloudWatch Logs destination log group for conversation text logs.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Determines whether conversation logs should be stored for an alias.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

