package awscustomerprofiles


// Mathematical expression and a list of attribute items specified in that expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDetailsProperty := &AttributeDetailsProperty{
//   	Attributes: []interface{}{
//   		&AttributeItemProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Expression: jsii.String("expression"),
//   }
//
type CfnCalculatedAttributeDefinition_AttributeDetailsProperty struct {
	// Mathematical expression and a list of attribute items specified in that expression.
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// Mathematical expression that is performed on attribute items provided in the attribute list.
	//
	// Each element in the expression should follow the structure of \"{ObjectTypeName.AttributeName}\".
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

