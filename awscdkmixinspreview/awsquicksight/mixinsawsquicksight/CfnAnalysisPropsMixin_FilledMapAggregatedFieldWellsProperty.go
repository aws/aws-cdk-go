package mixinsawsquicksight


// The aggregated field well of the filled map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapaggregatedfieldwells.html
//
type CfnAnalysisPropsMixin_FilledMapAggregatedFieldWellsProperty struct {
	// The aggregated location field well of the filled map.
	//
	// Values are grouped by location fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapaggregatedfieldwells.html#cfn-quicksight-analysis-filledmapaggregatedfieldwells-geospatial
	//
	Geospatial interface{} `field:"optional" json:"geospatial" yaml:"geospatial"`
	// The aggregated color field well of a filled map.
	//
	// Values are aggregated based on location fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-filledmapaggregatedfieldwells.html#cfn-quicksight-analysis-filledmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

