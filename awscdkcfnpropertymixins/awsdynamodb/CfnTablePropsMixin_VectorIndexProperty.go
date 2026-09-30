package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vectorIndexProperty := &VectorIndexProperty{
//   	Dimensions: jsii.Number(123),
//   	DistanceFunction: jsii.String("distanceFunction"),
//   	IndexName: jsii.String("indexName"),
//   	Projection: &ProjectionProperty{
//   		NonKeyAttributes: []*string{
//   			jsii.String("nonKeyAttributes"),
//   		},
//   		ProjectionType: jsii.String("projectionType"),
//   	},
//   	SearchSchema: []interface{}{
//   		&SearchSchemaElementProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			SearchSchemaElementType: jsii.String("searchSchemaElementType"),
//   		},
//   	},
//   	VectorAttribute: &VectorAttributeProperty{
//   		AttributeName: jsii.String("attributeName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html
//
type CfnTablePropsMixin_VectorIndexProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-dimensions
	//
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-distancefunction
	//
	DistanceFunction *string `field:"optional" json:"distanceFunction" yaml:"distanceFunction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Represents attributes that are copied (projected) from the table into an index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-projection
	//
	Projection interface{} `field:"optional" json:"projection" yaml:"projection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-searchschema
	//
	SearchSchema interface{} `field:"optional" json:"searchSchema" yaml:"searchSchema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-vectorindex.html#cfn-dynamodb-table-vectorindex-vectorattribute
	//
	VectorAttribute interface{} `field:"optional" json:"vectorAttribute" yaml:"vectorAttribute"`
}

