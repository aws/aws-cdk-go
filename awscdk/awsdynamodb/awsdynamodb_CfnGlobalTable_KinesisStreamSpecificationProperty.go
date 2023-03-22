package awsdynamodb


// The Kinesis Data Streams configuration for the specified global table replica.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamSpecificationProperty := &kinesisStreamSpecificationProperty{
//   	streamArn: jsii.String("streamArn"),
//   }
//
type CfnGlobalTable_KinesisStreamSpecificationProperty struct {
	// The ARN for a specific Kinesis data stream.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

