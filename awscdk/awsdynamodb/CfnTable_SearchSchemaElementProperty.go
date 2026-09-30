package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchSchemaElementProperty := &SearchSchemaElementProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	SearchSchemaElementType: jsii.String("searchSchemaElementType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-searchschemaelement.html
//
type CfnTable_SearchSchemaElementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-searchschemaelement.html#cfn-dynamodb-table-searchschemaelement-attributename
	//
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-searchschemaelement.html#cfn-dynamodb-table-searchschemaelement-searchschemaelementtype
	//
	SearchSchemaElementType *string `field:"required" json:"searchSchemaElementType" yaml:"searchSchemaElementType"`
}

