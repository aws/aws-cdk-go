package previewawsdynamodbmixins


// Represents attributes that are copied (projected) from the table into an index.
//
// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   projectionProperty := &ProjectionProperty{
//   	NonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	ProjectionType: jsii.String("projectionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-projection.html
//
type CfnTablePropsMixin_ProjectionProperty struct {
	// Represents the non-key attribute names which will be projected into the index.
	//
	// For global and local secondary indexes, the total count of `NonKeyAttributes` summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total. This limit only applies when you specify the ProjectionType of `INCLUDE` . You still can specify the ProjectionType of `ALL` to project all attributes from the source table, even if the table has more than 100 attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-projection.html#cfn-dynamodb-table-projection-nonkeyattributes
	//
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the index:.
	//
	// - `KEYS_ONLY` - Only the index and primary keys are projected into the index.
	// - `INCLUDE` - In addition to the attributes described in `KEYS_ONLY` , the secondary index will include other non-key attributes that you specify.
	// - `ALL` - All of the table attributes are projected into the index.
	//
	// When using the DynamoDB console, `ALL` is selected by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-projection.html#cfn-dynamodb-table-projection-projectiontype
	//
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
}

