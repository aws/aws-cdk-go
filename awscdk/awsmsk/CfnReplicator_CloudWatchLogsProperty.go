package awsmsk


// Details about delivering logs to CloudWatch Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsProperty := &CloudWatchLogsProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-cloudwatchlogs.html
//
type CfnReplicator_CloudWatchLogsProperty struct {
	// Whether log delivery to CloudWatch Logs is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-cloudwatchlogs.html#cfn-msk-replicator-cloudwatchlogs-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group that is the destination for log delivery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-cloudwatchlogs.html#cfn-msk-replicator-cloudwatchlogs-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

