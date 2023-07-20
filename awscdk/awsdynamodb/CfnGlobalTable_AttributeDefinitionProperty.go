package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDefinitionProperty := &AttributeDefinitionProperty{
//   	AttributeName: jsii.String("attributeName"),
//   	AttributeType: jsii.String("attributeType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-attributedefinition.html
//
type CfnGlobalTable_AttributeDefinitionProperty struct {
	// A name for the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-attributedefinition.html#cfn-dynamodb-globaltable-attributedefinition-attributename
	//
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The data type for the attribute, where:.
	//
	// - `S` - the attribute is of type String
	// - `N` - the attribute is of type Number
	// - `B` - the attribute is of type Binary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-attributedefinition.html#cfn-dynamodb-globaltable-attributedefinition-attributetype
	//
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
}

