package awsquicksight


// The forecast computation configuration.
//
// Example:
//
//
type CfnAnalysis_ForecastComputationProperty struct {
	// The ID for a computation.
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The time field that is used in a computation.
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// The custom seasonality value setup of a forecast computation.
	CustomSeasonalityValue *float64 `field:"optional" json:"customSeasonalityValue" yaml:"customSeasonalityValue"`
	// The lower boundary setup of a forecast computation.
	LowerBoundary *float64 `field:"optional" json:"lowerBoundary" yaml:"lowerBoundary"`
	// The name of a computation.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The periods backward setup of a forecast computation.
	PeriodsBackward *float64 `field:"optional" json:"periodsBackward" yaml:"periodsBackward"`
	// The periods forward setup of a forecast computation.
	PeriodsForward *float64 `field:"optional" json:"periodsForward" yaml:"periodsForward"`
	// The prediction interval setup of a forecast computation.
	PredictionInterval *float64 `field:"optional" json:"predictionInterval" yaml:"predictionInterval"`
	// The seasonality setup of a forecast computation. Choose one of the following options:.
	//
	// - `AUTOMATIC`
	// - `CUSTOM` : Checks the custom seasonality value.
	Seasonality *string `field:"optional" json:"seasonality" yaml:"seasonality"`
	// The upper boundary setup of a forecast computation.
	UpperBoundary *float64 `field:"optional" json:"upperBoundary" yaml:"upperBoundary"`
	// The value field that is used in a computation.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

