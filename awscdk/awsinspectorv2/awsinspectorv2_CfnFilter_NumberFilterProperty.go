package awsinspectorv2


// An object that describes the details of a number filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numberFilterProperty := &numberFilterProperty{
//   	lowerInclusive: jsii.Number(123),
//   	upperInclusive: jsii.Number(123),
//   }
//
type CfnFilter_NumberFilterProperty struct {
	// The lowest number to be included in the filter.
	LowerInclusive *float64 `field:"optional" json:"lowerInclusive" yaml:"lowerInclusive"`
	// The highest number to be included in the filter.
	UpperInclusive *float64 `field:"optional" json:"upperInclusive" yaml:"upperInclusive"`
}

