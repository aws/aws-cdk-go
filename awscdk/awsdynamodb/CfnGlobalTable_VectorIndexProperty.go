package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	VectorAttribute: &VectorAttributeProperty{
//   		AttributeName: jsii.String("attributeName"),
//   	},
//
//   	// the properties below are optional
//   	SearchSchema: []interface{}{
//   		&SearchSchemaElementProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			SearchSchemaElementType: jsii.String("searchSchemaElementType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html
//
type CfnGlobalTable_VectorIndexProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-dimensions
	//
	Dimensions *float64 `field:"required" json:"dimensions" yaml:"dimensions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-distancefunction
	//
	DistanceFunction *string `field:"required" json:"distanceFunction" yaml:"distanceFunction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-indexname
	//
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-projection
	//
	Projection interface{} `field:"required" json:"projection" yaml:"projection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-vectorattribute
	//
	VectorAttribute interface{} `field:"required" json:"vectorAttribute" yaml:"vectorAttribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-vectorindex.html#cfn-dynamodb-globaltable-vectorindex-searchschema
	//
	SearchSchema interface{} `field:"optional" json:"searchSchema" yaml:"searchSchema"`
}

