package awscognito


// Constraints that can be applied to a custom attribute of string type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringAttributeConstraints := &StringAttributeConstraints{
//   	MaxLen: jsii.Number(123),
//   	MinLen: jsii.Number(123),
//   }
//
type StringAttributeConstraints struct {
	// Maximum length of this attribute.
	// Default: 2048.
	//
	MaxLen *float64 `field:"optional" json:"maxLen" yaml:"maxLen"`
	// Minimum length of this attribute.
	// Default: 0.
	//
	MinLen *float64 `field:"optional" json:"minLen" yaml:"minLen"`
}

