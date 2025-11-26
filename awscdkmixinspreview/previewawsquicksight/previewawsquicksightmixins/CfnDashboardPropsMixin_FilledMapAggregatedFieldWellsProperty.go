package previewawsquicksightmixins


// The aggregated field well of the filled map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_FilledMapAggregatedFieldWellsProperty struct {
	// The aggregated location field well of the filled map.
	//
	// Values are grouped by location fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html#cfn-quicksight-dashboard-filledmapaggregatedfieldwells-geospatial
	//
	Geospatial interface{} `field:"optional" json:"geospatial" yaml:"geospatial"`
	// The aggregated color field well of a filled map.
	//
	// Values are aggregated based on location fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapaggregatedfieldwells.html#cfn-quicksight-dashboard-filledmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

