// The CDK Construct Library for AWS::MSK
package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configuration details related to broker logs.
//
// Example:
//   var vpc vpc
//   var bucket iBucket
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	Logging: &BrokerLogging{
//   		S3: &S3LoggingConfiguration{
//   			Bucket: *Bucket,
//   		},
//   	},
//   })
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

