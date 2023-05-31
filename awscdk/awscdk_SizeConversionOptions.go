// An experiment to bundle the entire CDK into a single module
package awscdk


// Options for how to convert time to a different unit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   sizeConversionOptions := &SizeConversionOptions{
//   	Rounding: monocdk.SizeRoundingBehavior_FAIL,
//   }
//
// Experimental.
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	// Experimental.
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

