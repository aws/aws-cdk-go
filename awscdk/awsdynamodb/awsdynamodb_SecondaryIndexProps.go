package awsdynamodb


// Properties for a secondary index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secondaryIndexProps := &SecondaryIndexProps{
//   	IndexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	NonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	ProjectionType: awscdk.Aws_dynamodb.ProjectionType_KEYS_ONLY,
//   }
//
// Experimental.
type SecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `field:"optional" json:"projectionType" yaml:"projectionType"`
}

