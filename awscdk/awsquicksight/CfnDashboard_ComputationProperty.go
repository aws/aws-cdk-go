package awsquicksight


// The computation union that is used in an insight visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnDashboard_ComputationProperty struct {
	// The forecast computation configuration.
	Forecast interface{} `field:"optional" json:"forecast" yaml:"forecast"`
	// The growth rate computation configuration.
	GrowthRate interface{} `field:"optional" json:"growthRate" yaml:"growthRate"`
	// The maximum and minimum computation configuration.
	MaximumMinimum interface{} `field:"optional" json:"maximumMinimum" yaml:"maximumMinimum"`
	// The metric comparison computation configuration.
	MetricComparison interface{} `field:"optional" json:"metricComparison" yaml:"metricComparison"`
	// The period over period computation configuration.
	PeriodOverPeriod interface{} `field:"optional" json:"periodOverPeriod" yaml:"periodOverPeriod"`
	// The period to `DataSetIdentifier` computation configuration.
	PeriodToDate interface{} `field:"optional" json:"periodToDate" yaml:"periodToDate"`
	// The top movers and bottom movers computation configuration.
	TopBottomMovers interface{} `field:"optional" json:"topBottomMovers" yaml:"topBottomMovers"`
	// The top ranked and bottom ranked computation configuration.
	TopBottomRanked interface{} `field:"optional" json:"topBottomRanked" yaml:"topBottomRanked"`
	// The total aggregation computation configuration.
	TotalAggregation interface{} `field:"optional" json:"totalAggregation" yaml:"totalAggregation"`
	// The unique values computation configuration.
	UniqueValues interface{} `field:"optional" json:"uniqueValues" yaml:"uniqueValues"`
}

