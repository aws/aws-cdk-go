package awsquicksight


// The arc axis configuration of a `GaugeChartVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcAxisConfigurationProperty := &ArcAxisConfigurationProperty{
//   	Range: &ArcAxisDisplayRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	ReserveRange: jsii.Number(123),
//   }
//
type CfnAnalysis_ArcAxisConfigurationProperty struct {
	// The arc axis range of a `GaugeChartVisual` .
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// The reserved range of the arc axis.
	ReserveRange *float64 `field:"optional" json:"reserveRange" yaml:"reserveRange"`
}

