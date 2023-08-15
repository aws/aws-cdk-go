package awscdk


// Options for how to convert time to a different unit.
//
// Example:
//   awscdk.Size_Mebibytes(jsii.Number(2)).ToKibibytes() // yields 2048
//   awscdk.Size_Kibibytes(jsii.Number(2050)).ToMebibytes(&SizeConversionOptions{
//   	Rounding: awscdk.SizeRoundingBehavior_FLOOR,
//   })
//
type SizeConversionOptions struct {
	// How conversions should behave when it encounters a non-integer result.
	// Default: SizeRoundingBehavior.FAIL
	//
	Rounding SizeRoundingBehavior `field:"optional" json:"rounding" yaml:"rounding"`
}

