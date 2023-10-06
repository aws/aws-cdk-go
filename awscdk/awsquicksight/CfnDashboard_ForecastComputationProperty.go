package awsquicksight


// The forecast computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html
//
type CfnDashboard_ForecastComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The custom seasonality value setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-customseasonalityvalue
	//
	CustomSeasonalityValue *float64 `field:"optional" json:"customSeasonalityValue" yaml:"customSeasonalityValue"`
	// The lower boundary setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-lowerboundary
	//
	LowerBoundary *float64 `field:"optional" json:"lowerBoundary" yaml:"lowerBoundary"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The periods backward setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-periodsbackward
	//
	PeriodsBackward *float64 `field:"optional" json:"periodsBackward" yaml:"periodsBackward"`
	// The periods forward setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-periodsforward
	//
	PeriodsForward *float64 `field:"optional" json:"periodsForward" yaml:"periodsForward"`
	// The prediction interval setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-predictioninterval
	//
	PredictionInterval *float64 `field:"optional" json:"predictionInterval" yaml:"predictionInterval"`
	// The seasonality setup of a forecast computation. Choose one of the following options:.
	//
	// - `AUTOMATIC`
	// - `CUSTOM` : Checks the custom seasonality value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-seasonality
	//
	Seasonality *string `field:"optional" json:"seasonality" yaml:"seasonality"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The upper boundary setup of a forecast computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-upperboundary
	//
	UpperBoundary *float64 `field:"optional" json:"upperBoundary" yaml:"upperBoundary"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

