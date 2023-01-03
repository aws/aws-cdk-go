package awsconnect


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
	// `CfnInstanceStorageConfig.KinesisFirehoseConfigProperty.FirehoseArn`.
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
}

