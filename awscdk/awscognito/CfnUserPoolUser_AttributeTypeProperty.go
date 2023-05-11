package awscognito


// Specifies whether the attribute is standard or custom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeTypeProperty := &AttributeTypeProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
type CfnUserPoolUser_AttributeTypeProperty struct {
	// The name of the attribute.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the attribute.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

