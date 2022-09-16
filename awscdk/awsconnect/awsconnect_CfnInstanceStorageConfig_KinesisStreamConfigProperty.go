package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamConfigProperty := &kinesisStreamConfigProperty{
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnInstanceStorageConfig_KinesisStreamConfigProperty struct {
	// `CfnInstanceStorageConfig.KinesisStreamConfigProperty.StreamArn`.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

