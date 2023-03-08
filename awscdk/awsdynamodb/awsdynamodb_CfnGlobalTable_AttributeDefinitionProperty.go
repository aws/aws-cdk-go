package awsdynamodb


// Represents an attribute for describing the key schema for the table and indexes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDefinitionProperty := &attributeDefinitionProperty{
//   	attributeName: jsii.String("attributeName"),
//   	attributeType: jsii.String("attributeType"),
//   }
//
type CfnGlobalTable_AttributeDefinitionProperty struct {
	// A name for the attribute.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The data type for the attribute, where:.
	//
	// - `S` - the attribute is of type String
	// - `N` - the attribute is of type Number
	// - `B` - the attribute is of type Binary.
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
}

