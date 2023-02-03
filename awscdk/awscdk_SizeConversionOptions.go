// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for how to convert time to a different unit.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   awscdk.Size.mebibytes(jsii.Number(2)).toKibibytes() // yields 2048
//   awscdk.Size.kibibytes(jsii.Number(2050)).toMebibytes(&sizeConversionOptions{
//   	rounding: awscdk.SizeRoundingBehavior_FLOOR,
//   })
//
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

