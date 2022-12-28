package awsconnect


// Configuration information of a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisFirehoseConfigProperty := &kinesisFirehoseConfigProperty{
//   	firehoseArn: jsii.String("firehoseArn"),
//   }
//
type CfnInstanceStorageConfig_KinesisFirehoseConfigProperty struct {
	// The Amazon Resource Name (ARN) of the delivery stream.
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
}

