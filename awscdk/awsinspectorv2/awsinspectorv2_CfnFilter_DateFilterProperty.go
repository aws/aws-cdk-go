package awsinspectorv2


// Contains details on the time range used to filter findings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateFilterProperty := &dateFilterProperty{
//   	endInclusive: jsii.Number(123),
//   	startInclusive: jsii.Number(123),
//   }
//
type CfnFilter_DateFilterProperty struct {
	// A timestamp representing the end of the time period filtered on.
	EndInclusive *float64 `field:"optional" json:"endInclusive" yaml:"endInclusive"`
	// A timestamp representing the start of the time period filtered on.
	StartInclusive *float64 `field:"optional" json:"startInclusive" yaml:"startInclusive"`
}

