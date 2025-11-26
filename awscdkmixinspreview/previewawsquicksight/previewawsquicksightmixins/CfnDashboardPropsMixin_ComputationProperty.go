package previewawsquicksightmixins


// The computation union that is used in an insight visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html
//
type CfnDashboardPropsMixin_ComputationProperty struct {
	// The forecast computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-forecast
	//
	Forecast interface{} `field:"optional" json:"forecast" yaml:"forecast"`
	// The growth rate computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-growthrate
	//
	GrowthRate interface{} `field:"optional" json:"growthRate" yaml:"growthRate"`
	// The maximum and minimum computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-maximumminimum
	//
	MaximumMinimum interface{} `field:"optional" json:"maximumMinimum" yaml:"maximumMinimum"`
	// The metric comparison computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-metriccomparison
	//
	MetricComparison interface{} `field:"optional" json:"metricComparison" yaml:"metricComparison"`
	// The period over period computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-periodoverperiod
	//
	PeriodOverPeriod interface{} `field:"optional" json:"periodOverPeriod" yaml:"periodOverPeriod"`
	// The period to `DataSetIdentifier` computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-periodtodate
	//
	PeriodToDate interface{} `field:"optional" json:"periodToDate" yaml:"periodToDate"`
	// The top movers and bottom movers computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-topbottommovers
	//
	TopBottomMovers interface{} `field:"optional" json:"topBottomMovers" yaml:"topBottomMovers"`
	// The top ranked and bottom ranked computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-topbottomranked
	//
	TopBottomRanked interface{} `field:"optional" json:"topBottomRanked" yaml:"topBottomRanked"`
	// The total aggregation computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-totalaggregation
	//
	TotalAggregation interface{} `field:"optional" json:"totalAggregation" yaml:"totalAggregation"`
	// The unique values computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-computation.html#cfn-quicksight-dashboard-computation-uniquevalues
	//
	UniqueValues interface{} `field:"optional" json:"uniqueValues" yaml:"uniqueValues"`
}

