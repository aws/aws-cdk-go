package awsquicksight


// The arc axis range of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcAxisDisplayRangeProperty := &ArcAxisDisplayRangeProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
type CfnTemplate_ArcAxisDisplayRangeProperty struct {
	// The maximum value of the arc axis range.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum value of the arc axis range.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

