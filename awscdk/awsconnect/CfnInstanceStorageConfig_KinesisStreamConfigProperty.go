package awsconnect


// Configuration information of a Kinesis data stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamConfigProperty := &KinesisStreamConfigProperty{
//   	StreamArn: jsii.String("streamArn"),
//   }
//
type CfnInstanceStorageConfig_KinesisStreamConfigProperty struct {
	// The Amazon Resource Name (ARN) of the data stream.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

