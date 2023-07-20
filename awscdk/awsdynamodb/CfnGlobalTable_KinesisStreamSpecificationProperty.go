package awsdynamodb


// The Kinesis Data Streams configuration for the specified global table replica.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisStreamSpecificationProperty := &KinesisStreamSpecificationProperty{
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-kinesisstreamspecification.html
//
type CfnGlobalTable_KinesisStreamSpecificationProperty struct {
	// The ARN for a specific Kinesis data stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-kinesisstreamspecification.html#cfn-dynamodb-globaltable-kinesisstreamspecification-streamarn
	//
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

