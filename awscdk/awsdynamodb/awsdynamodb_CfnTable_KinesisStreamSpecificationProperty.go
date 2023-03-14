package awsdynamodb


// The Kinesis Data Streams configuration for the specified table.
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
type CfnTable_KinesisStreamSpecificationProperty struct {
	// The ARN for a specific Kinesis data stream.
	//
	// Length Constraints: Minimum length of 37. Maximum length of 1024.
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

