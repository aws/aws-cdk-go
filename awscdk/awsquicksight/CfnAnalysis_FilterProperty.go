package awsquicksight


// With a `Filter` , you can remove portions of data from a particular visual or view.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html
//
type CfnAnalysis_FilterProperty struct {
	// A `CategoryFilter` filters text values.
	//
	// For more information, see [Adding text filters](https://docs.aws.amazon.com/quicksight/latest/user/add-a-text-filter-data-prep.html) in the *Amazon QuickSight User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-categoryfilter
	//
	CategoryFilter interface{} `field:"optional" json:"categoryFilter" yaml:"categoryFilter"`
	// A `NumericEqualityFilter` filters numeric values that equal or do not equal a given numeric value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-numericequalityfilter
	//
	NumericEqualityFilter interface{} `field:"optional" json:"numericEqualityFilter" yaml:"numericEqualityFilter"`
	// A `NumericRangeFilter` filters numeric values that are either inside or outside a given numeric range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-numericrangefilter
	//
	NumericRangeFilter interface{} `field:"optional" json:"numericRangeFilter" yaml:"numericRangeFilter"`
	// A `RelativeDatesFilter` filters date values that are relative to a given date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-relativedatesfilter
	//
	RelativeDatesFilter interface{} `field:"optional" json:"relativeDatesFilter" yaml:"relativeDatesFilter"`
	// A `TimeEqualityFilter` filters date-time values that equal or do not equal a given date/time value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-timeequalityfilter
	//
	TimeEqualityFilter interface{} `field:"optional" json:"timeEqualityFilter" yaml:"timeEqualityFilter"`
	// A `TimeRangeFilter` filters date-time values that are either inside or outside a given date/time range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-timerangefilter
	//
	TimeRangeFilter interface{} `field:"optional" json:"timeRangeFilter" yaml:"timeRangeFilter"`
	// A `TopBottomFilter` filters data to the top or bottom values for a given column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filter.html#cfn-quicksight-analysis-filter-topbottomfilter
	//
	TopBottomFilter interface{} `field:"optional" json:"topBottomFilter" yaml:"topBottomFilter"`
}

