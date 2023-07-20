package awsquicksight


// The configuration of a `GeospatialMapVisual` .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html
//
type CfnAnalysis_GeospatialMapConfigurationProperty struct {
	// The field wells of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-legend
	//
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The map style options of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-mapstyleoptions
	//
	MapStyleOptions interface{} `field:"optional" json:"mapStyleOptions" yaml:"mapStyleOptions"`
	// The point style options of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-pointstyleoptions
	//
	PointStyleOptions interface{} `field:"optional" json:"pointStyleOptions" yaml:"pointStyleOptions"`
	// The tooltip display setup of the visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-tooltip
	//
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-visualpalette
	//
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The window options of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapconfiguration.html#cfn-quicksight-analysis-geospatialmapconfiguration-windowoptions
	//
	WindowOptions interface{} `field:"optional" json:"windowOptions" yaml:"windowOptions"`
}

