package previewawsquicksightmixins


// The aggregated field wells for a geospatial map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapaggregatedfieldwells.html
//
type CfnDashboardPropsMixin_GeospatialMapAggregatedFieldWellsProperty struct {
	// The color field wells of a geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapaggregatedfieldwells.html#cfn-quicksight-dashboard-geospatialmapaggregatedfieldwells-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The geospatial field wells of a geospatial map.
	//
	// Values are grouped by geospatial fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapaggregatedfieldwells.html#cfn-quicksight-dashboard-geospatialmapaggregatedfieldwells-geospatial
	//
	Geospatial interface{} `field:"optional" json:"geospatial" yaml:"geospatial"`
	// The size field wells of a geospatial map.
	//
	// Values are aggregated based on geospatial fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapaggregatedfieldwells.html#cfn-quicksight-dashboard-geospatialmapaggregatedfieldwells-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

