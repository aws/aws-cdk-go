package awsquicksight


// The configuration of a `GeospatialMapVisual` .
//
// Example:
//
//
type CfnTemplate_GeospatialMapConfigurationProperty struct {
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The map style options of the geospatial map.
	MapStyleOptions interface{} `field:"optional" json:"mapStyleOptions" yaml:"mapStyleOptions"`
	// The point style options of the geospatial map.
	PointStyleOptions interface{} `field:"optional" json:"pointStyleOptions" yaml:"pointStyleOptions"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// `CfnTemplate.GeospatialMapConfigurationProperty.VisualPalette`.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The window options of the geospatial map.
	WindowOptions interface{} `field:"optional" json:"windowOptions" yaml:"windowOptions"`
}

