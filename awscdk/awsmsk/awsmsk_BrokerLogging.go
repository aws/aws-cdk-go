package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Configuration details related to broker logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var logGroup logGroup
//
//   brokerLogging := &brokerLogging{
//   	cloudwatchLogGroup: logGroup,
//   	firehoseDeliveryStreamName: jsii.String("firehoseDeliveryStreamName"),
//   	s3: &s3LoggingConfiguration{
//   		bucket: bucket,
//
//   		// the properties below are optional
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
// Experimental.
type BrokerLogging struct {
	// The CloudWatch Logs group that is the destination for broker logs.
	// Experimental.
	CloudwatchLogGroup awslogs.ILogGroup `field:"optional" json:"cloudwatchLogGroup" yaml:"cloudwatchLogGroup"`
	// The Kinesis Data Firehose delivery stream that is the destination for broker logs.
	// Experimental.
	FirehoseDeliveryStreamName *string `field:"optional" json:"firehoseDeliveryStreamName" yaml:"firehoseDeliveryStreamName"`
	// Details of the Amazon S3 destination for broker logs.
	// Experimental.
	S3 *S3LoggingConfiguration `field:"optional" json:"s3" yaml:"s3"`
}

