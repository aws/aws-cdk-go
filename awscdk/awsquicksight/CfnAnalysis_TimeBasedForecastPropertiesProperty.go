package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedForecastPropertiesProperty := &TimeBasedForecastPropertiesProperty{
//   	LowerBoundary: jsii.Number(123),
//   	PeriodsBackward: jsii.Number(123),
//   	PeriodsForward: jsii.Number(123),
//   	PredictionInterval: jsii.Number(123),
//   	Seasonality: jsii.Number(123),
//   	UpperBoundary: jsii.Number(123),
//   }
//
type CfnAnalysis_TimeBasedForecastPropertiesProperty struct {
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.LowerBoundary`.
	LowerBoundary *float64 `field:"optional" json:"lowerBoundary" yaml:"lowerBoundary"`
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.PeriodsBackward`.
	PeriodsBackward *float64 `field:"optional" json:"periodsBackward" yaml:"periodsBackward"`
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.PeriodsForward`.
	PeriodsForward *float64 `field:"optional" json:"periodsForward" yaml:"periodsForward"`
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.PredictionInterval`.
	PredictionInterval *float64 `field:"optional" json:"predictionInterval" yaml:"predictionInterval"`
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.Seasonality`.
	Seasonality *float64 `field:"optional" json:"seasonality" yaml:"seasonality"`
	// `CfnAnalysis.TimeBasedForecastPropertiesProperty.UpperBoundary`.
	UpperBoundary *float64 `field:"optional" json:"upperBoundary" yaml:"upperBoundary"`
}

