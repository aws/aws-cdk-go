package previewawsquicksightmixins


// The computation union that is used in an insight visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html
//
type CfnAnalysisPropsMixin_ComputationProperty struct {
	// The forecast computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-forecast
	//
	Forecast interface{} `field:"optional" json:"forecast" yaml:"forecast"`
	// The growth rate computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-growthrate
	//
	GrowthRate interface{} `field:"optional" json:"growthRate" yaml:"growthRate"`
	// The maximum and minimum computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-maximumminimum
	//
	MaximumMinimum interface{} `field:"optional" json:"maximumMinimum" yaml:"maximumMinimum"`
	// The metric comparison computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-metriccomparison
	//
	MetricComparison interface{} `field:"optional" json:"metricComparison" yaml:"metricComparison"`
	// The period over period computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-periodoverperiod
	//
	PeriodOverPeriod interface{} `field:"optional" json:"periodOverPeriod" yaml:"periodOverPeriod"`
	// The period to `DataSetIdentifier` computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-periodtodate
	//
	PeriodToDate interface{} `field:"optional" json:"periodToDate" yaml:"periodToDate"`
	// The top movers and bottom movers computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-topbottommovers
	//
	TopBottomMovers interface{} `field:"optional" json:"topBottomMovers" yaml:"topBottomMovers"`
	// The top ranked and bottom ranked computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-topbottomranked
	//
	TopBottomRanked interface{} `field:"optional" json:"topBottomRanked" yaml:"topBottomRanked"`
	// The total aggregation computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-totalaggregation
	//
	TotalAggregation interface{} `field:"optional" json:"totalAggregation" yaml:"totalAggregation"`
	// The unique values computation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-computation.html#cfn-quicksight-analysis-computation-uniquevalues
	//
	UniqueValues interface{} `field:"optional" json:"uniqueValues" yaml:"uniqueValues"`
}

