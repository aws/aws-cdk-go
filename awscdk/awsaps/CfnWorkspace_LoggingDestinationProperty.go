package awsaps


// Destinations for query logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingDestinationProperty := &LoggingDestinationProperty{
//   	CloudWatchLogs: &CloudWatchLogDestinationProperty{
//   		LogGroupArn: jsii.String("logGroupArn"),
//   	},
//   	Filters: &LoggingFilterProperty{
//   		QspThreshold: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingdestination.html
//
type CfnWorkspace_LoggingDestinationProperty struct {
	// Represents a cloudwatch logs destination for query logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingdestination.html#cfn-aps-workspace-loggingdestination-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"required" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Filters for logging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingdestination.html#cfn-aps-workspace-loggingdestination-filters
	//
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
}

