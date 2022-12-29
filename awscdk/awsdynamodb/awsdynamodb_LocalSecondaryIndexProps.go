package awsdynamodb


// Properties for a local secondary index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localSecondaryIndexProps := &localSecondaryIndexProps{
//   	indexName: jsii.String("indexName"),
//   	sortKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.Aws_dynamodb.attributeType_BINARY,
//   	},
//
//   	// the properties below are optional
//   	nonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	projectionType: awscdk.*Aws_dynamodb.projectionType_KEYS_ONLY,
//   }
//
type LocalSecondaryIndexProps struct {
	// The name of the secondary index.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	ProjectionType ProjectionType `field:"optional" json:"projectionType" yaml:"projectionType"`
	// The attribute of a sort key for the local secondary index.
	SortKey *Attribute `field:"required" json:"sortKey" yaml:"sortKey"`
}

