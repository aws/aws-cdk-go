package mixinsawsaps


// The logging destination in an Amazon Managed Service for Prometheus workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnWorkspacePropsMixin_LoggingDestinationProperty struct {
	// Configuration details for logging to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingdestination.html#cfn-aps-workspace-loggingdestination-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Filtering criteria that determine which queries are logged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-loggingdestination.html#cfn-aps-workspace-loggingdestination-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

