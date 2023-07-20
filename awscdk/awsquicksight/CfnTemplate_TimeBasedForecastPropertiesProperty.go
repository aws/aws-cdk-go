package awsquicksight


// The forecast properties setup of a forecast in the line chart.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html
//
type CfnTemplate_TimeBasedForecastPropertiesProperty struct {
	// The lower boundary setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-lowerboundary
	//
	LowerBoundary *float64 `field:"optional" json:"lowerBoundary" yaml:"lowerBoundary"`
	// The periods backward setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-periodsbackward
	//
	PeriodsBackward *float64 `field:"optional" json:"periodsBackward" yaml:"periodsBackward"`
	// The periods forward setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-periodsforward
	//
	PeriodsForward *float64 `field:"optional" json:"periodsForward" yaml:"periodsForward"`
	// The prediction interval setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-predictioninterval
	//
	PredictionInterval *float64 `field:"optional" json:"predictionInterval" yaml:"predictionInterval"`
	// The seasonality setup of a forecast computation. Choose one of the following options:.
	//
	// - `NULL` : The input is set to `NULL` .
	// - `NON_NULL` : The input is set to a custom value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-seasonality
	//
	Seasonality *float64 `field:"optional" json:"seasonality" yaml:"seasonality"`
	// The upper boundary setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-timebasedforecastproperties.html#cfn-quicksight-template-timebasedforecastproperties-upperboundary
	//
	UpperBoundary *float64 `field:"optional" json:"upperBoundary" yaml:"upperBoundary"`
}

