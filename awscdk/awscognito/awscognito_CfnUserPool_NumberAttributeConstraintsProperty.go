package awscognito


// The minimum and maximum values of an attribute that is of the number data type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberAttributeConstraintsProperty := &numberAttributeConstraintsProperty{
//   	maxValue: jsii.String("maxValue"),
//   	minValue: jsii.String("minValue"),
//   }
//
type CfnUserPool_NumberAttributeConstraintsProperty struct {
	// The maximum value of an attribute that is of the number data type.
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum value of an attribute that is of the number data type.
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
}

