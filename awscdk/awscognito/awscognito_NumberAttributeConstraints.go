package awscognito


// Constraints that can be applied to a custom attribute of number type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberAttributeConstraints := &numberAttributeConstraints{
//   	max: jsii.Number(123),
//   	min: jsii.Number(123),
//   }
//
type NumberAttributeConstraints struct {
	// Maximum value of this attribute.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Minimum value of this attribute.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

