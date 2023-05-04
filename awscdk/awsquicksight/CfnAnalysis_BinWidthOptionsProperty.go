package awsquicksight


// The options that determine the bin width of a histogram.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   binWidthOptionsProperty := &BinWidthOptionsProperty{
//   	BinCountLimit: jsii.Number(123),
//   	Value: jsii.Number(123),
//   }
//
type CfnAnalysis_BinWidthOptionsProperty struct {
	// The options that determine the bin count limit.
	BinCountLimit *float64 `field:"optional" json:"binCountLimit" yaml:"binCountLimit"`
	// The options that determine the bin width value.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

