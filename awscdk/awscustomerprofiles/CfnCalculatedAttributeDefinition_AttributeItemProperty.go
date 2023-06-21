package awscustomerprofiles


// The details of a single attribute item specified in the mathematical expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeItemProperty := &AttributeItemProperty{
//   	Name: jsii.String("name"),
//   }
//
type CfnCalculatedAttributeDefinition_AttributeItemProperty struct {
	// The unique name of the calculated attribute.
	Name *string `field:"required" json:"name" yaml:"name"`
}

