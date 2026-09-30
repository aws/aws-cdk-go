package awsdynamodb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   searchSchemaElementProperty := &SearchSchemaElementProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	SearchSchemaElementType: jsii.String("searchSchemaElementType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-searchschemaelement.html
//
type CfnGlobalTablePropsMixin_SearchSchemaElementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-searchschemaelement.html#cfn-dynamodb-globaltable-searchschemaelement-attributename
	//
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-searchschemaelement.html#cfn-dynamodb-globaltable-searchschemaelement-searchschemaelementtype
	//
	SearchSchemaElementType *string `field:"optional" json:"searchSchemaElementType" yaml:"searchSchemaElementType"`
}

