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
type CfnTemplate_TimeBasedForecastPropertiesProperty struct {
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.LowerBoundary`.
	LowerBoundary *float64 `field:"optional" json:"lowerBoundary" yaml:"lowerBoundary"`
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.PeriodsBackward`.
	PeriodsBackward *float64 `field:"optional" json:"periodsBackward" yaml:"periodsBackward"`
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.PeriodsForward`.
	PeriodsForward *float64 `field:"optional" json:"periodsForward" yaml:"periodsForward"`
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.PredictionInterval`.
	PredictionInterval *float64 `field:"optional" json:"predictionInterval" yaml:"predictionInterval"`
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.Seasonality`.
	Seasonality *float64 `field:"optional" json:"seasonality" yaml:"seasonality"`
	// `CfnTemplate.TimeBasedForecastPropertiesProperty.UpperBoundary`.
	UpperBoundary *float64 `field:"optional" json:"upperBoundary" yaml:"upperBoundary"`
}

