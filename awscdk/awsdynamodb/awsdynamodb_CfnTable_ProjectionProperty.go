package awsdynamodb


// Represents attributes that are copied (projected) from the table into an index.
//
// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectionProperty := &projectionProperty{
//   	nonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	projectionType: jsii.String("projectionType"),
//   }
//
type CfnTable_ProjectionProperty struct {
	// Represents the non-key attribute names which will be projected into the index.
	//
	// For local secondary indexes, the total count of `NonKeyAttributes` summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the index:.
	//
	// - `KEYS_ONLY` - Only the index and primary keys are projected into the index.
	// - `INCLUDE` - In addition to the attributes described in `KEYS_ONLY` , the secondary index will include other non-key attributes that you specify.
	// - `ALL` - All of the table attributes are projected into the index.
	ProjectionType *string `field:"optional" json:"projectionType" yaml:"projectionType"`
}

