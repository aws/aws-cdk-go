package awsdatapipeline


// `Attribute` is a property of `ParameterObject` that defines the attributes of a parameter object as key-value pairs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterAttributeProperty := &parameterAttributeProperty{
//   	key: jsii.String("key"),
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnPipeline_ParameterAttributeProperty struct {
	// The field identifier.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The field value, expressed as a String.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

