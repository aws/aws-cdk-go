package previewawsdynamodbmixins


// The Kinesis Data Streams configuration for the specified table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kinesisStreamSpecificationProperty := &KinesisStreamSpecificationProperty{
//   	ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
//   	StreamArn: jsii.String("streamArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html
//
type CfnTablePropsMixin_KinesisStreamSpecificationProperty struct {
	// The precision for the time and date that the stream was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html#cfn-dynamodb-table-kinesisstreamspecification-approximatecreationdatetimeprecision
	//
	ApproximateCreationDateTimePrecision *string `field:"optional" json:"approximateCreationDateTimePrecision" yaml:"approximateCreationDateTimePrecision"`
	// The ARN for a specific Kinesis data stream.
	//
	// Length Constraints: Minimum length of 37. Maximum length of 1024.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-kinesisstreamspecification.html#cfn-dynamodb-table-kinesisstreamspecification-streamarn
	//
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

