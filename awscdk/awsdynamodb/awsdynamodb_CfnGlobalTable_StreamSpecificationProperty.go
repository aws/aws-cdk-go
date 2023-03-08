package awsdynamodb


// Represents the DynamoDB Streams configuration for a table in DynamoDB.
//
// You can only modify this value if your `AWS::DynamoDB::GlobalTable` contains only one entry in `Replicas` . You must specify a value for this property if your `AWS::DynamoDB::GlobalTable` contains more than one replica.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSpecificationProperty := &streamSpecificationProperty{
//   	streamViewType: jsii.String("streamViewType"),
//   }
//
type CfnGlobalTable_StreamSpecificationProperty struct {
	// When an item in the table is modified, `StreamViewType` determines what information is written to the stream for this table.
	//
	// Valid values for `StreamViewType` are:
	//
	// - `KEYS_ONLY` - Only the key attributes of the modified item are written to the stream.
	// - `NEW_IMAGE` - The entire item, as it appears after it was modified, is written to the stream.
	// - `OLD_IMAGE` - The entire item, as it appeared before it was modified, is written to the stream.
	// - `NEW_AND_OLD_IMAGES` - Both the new and the old item images of the item are written to the stream.
	StreamViewType *string `field:"required" json:"streamViewType" yaml:"streamViewType"`
}

