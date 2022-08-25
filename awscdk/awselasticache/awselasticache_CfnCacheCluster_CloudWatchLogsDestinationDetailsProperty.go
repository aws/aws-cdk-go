package awselasticache


// Configuration details of a CloudWatch Logs destination.
//
// Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchLogsDestinationDetailsProperty := &cloudWatchLogsDestinationDetailsProperty{
//   	logGroup: jsii.String("logGroup"),
//   }
//
type CfnCacheCluster_CloudWatchLogsDestinationDetailsProperty struct {
	// The name of the CloudWatch Logs log group.
	LogGroup *string `field:"required" json:"logGroup" yaml:"logGroup"`
}

