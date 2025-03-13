package awselasticache


// The configuration details of the CloudWatch Logs destination.
//
// Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsDestinationDetailsProperty := &CloudWatchLogsDestinationDetailsProperty{
//   	LogGroup: jsii.String("logGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-cloudwatchlogsdestinationdetails.html
//
type CfnReplicationGroup_CloudWatchLogsDestinationDetailsProperty struct {
	// The name of the CloudWatch Logs log group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-cloudwatchlogsdestinationdetails.html#cfn-elasticache-replicationgroup-cloudwatchlogsdestinationdetails-loggroup
	//
	LogGroup *string `field:"required" json:"logGroup" yaml:"logGroup"`
}

