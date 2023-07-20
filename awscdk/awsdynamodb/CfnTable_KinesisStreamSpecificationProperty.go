package awsdynamodb


// The Kinesis Data Streams configuration for the specified table.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html
//
type CfnTable_KinesisStreamSpecificationProperty struct {
	// The ARN for a specific Kinesis data stream.
	//
	// Length Constraints: Minimum length of 37. Maximum length of 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html#cfn-dynamodb-table-kinesisstreamspecification-streamarn
	//
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

