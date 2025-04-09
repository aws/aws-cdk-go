package awscdkpipesalpha


// Log destination configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//
//   logDestinationConfig := &LogDestinationConfig{
//   	Parameters: &LogDestinationParameters{
//   		CloudwatchLogsLogDestination: &CloudwatchLogsLogDestinationProperty{
//   			LogGroupArn: jsii.String("logGroupArn"),
//   		},
//   		FirehoseLogDestination: &FirehoseLogDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   		},
//   		S3LogDestination: &S3LogDestinationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			BucketOwner: jsii.String("bucketOwner"),
//   			OutputFormat: jsii.String("outputFormat"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// Experimental.
type LogDestinationConfig struct {
	// Get the log destination configuration parameters.
	// Experimental.
	Parameters *LogDestinationParameters `field:"required" json:"parameters" yaml:"parameters"`
}

