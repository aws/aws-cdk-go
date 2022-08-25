package awsmsk


// Details of the CloudWatch Logs destination for broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsProperty := &cloudWatchLogsProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	logGroup: jsii.String("logGroup"),
//   }
//
type CfnCluster_CloudWatchLogsProperty struct {
	// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudWatch Logs group that is the destination for broker logs.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

