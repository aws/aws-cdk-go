package previewawsquicksightmixins


// The map definition that defines map state, map style, and geospatial layers.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html
//
type CfnDashboardPropsMixin_GeospatialLayerMapConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html#cfn-quicksight-dashboard-geospatiallayermapconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html#cfn-quicksight-dashboard-geospatiallayermapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The geospatial layers to visualize on the map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html#cfn-quicksight-dashboard-geospatiallayermapconfiguration-maplayers
	//
	MapLayers interface{} `field:"optional" json:"mapLayers" yaml:"mapLayers"`
	// The map state properties for the map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html#cfn-quicksight-dashboard-geospatiallayermapconfiguration-mapstate
	//
	MapState interface{} `field:"optional" json:"mapState" yaml:"mapState"`
	// The map style properties for the map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayermapconfiguration.html#cfn-quicksight-dashboard-geospatiallayermapconfiguration-mapstyle
	//
	MapStyle interface{} `field:"optional" json:"mapStyle" yaml:"mapStyle"`
}

