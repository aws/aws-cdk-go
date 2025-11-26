package previewawsdynamodbmixins


// Represents the DynamoDB Streams configuration for a table in DynamoDB .
//
// You can only modify this value for a `AWS::DynamoDB::GlobalTable` resource configured for multi-Region eventual consistency (MREC, the default) if that resource contains only one entry in `Replicas` . You must specify a value for this property for a `AWS::DynamoDB::GlobalTable` resource configured for MREC with more than one entry in `Replicas` . For Multi-Region Strong Consistency (MRSC), Streams are not required and can be changed for existing tables.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   streamSpecificationProperty := &StreamSpecificationProperty{
//   	StreamViewType: jsii.String("streamViewType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-streamspecification.html
//
type CfnGlobalTablePropsMixin_StreamSpecificationProperty struct {
	// When an item in the table is modified, `StreamViewType` determines what information is written to the stream for this table.
	//
	// Valid values for `StreamViewType` are:
	//
	// - `KEYS_ONLY` - Only the key attributes of the modified item are written to the stream.
	// - `NEW_IMAGE` - The entire item, as it appears after it was modified, is written to the stream.
	// - `OLD_IMAGE` - The entire item, as it appeared before it was modified, is written to the stream.
	// - `NEW_AND_OLD_IMAGES` - Both the new and the old item images of the item are written to the stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-streamspecification.html#cfn-dynamodb-globaltable-streamspecification-streamviewtype
	//
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
}

