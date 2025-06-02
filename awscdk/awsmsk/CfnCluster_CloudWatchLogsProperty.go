package awsmsk


// Details of the CloudWatch Logs destination for broker logs.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-cloudwatchlogs.html
//
type CfnCluster_CloudWatchLogsProperty struct {
	// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-cloudwatchlogs.html#cfn-msk-cluster-cloudwatchlogs-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The CloudWatch log group that is the destination for broker logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-cloudwatchlogs.html#cfn-msk-cluster-cloudwatchlogs-loggroup
	//
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
}

