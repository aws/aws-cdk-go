package awsquicksight


// The configuration for a `FilledMapVisual` .
//
// Example:
//
//
type CfnTemplate_FilledMapConfigurationProperty struct {
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The map style options of the filled map visual.
	MapStyleOptions interface{} `field:"optional" json:"mapStyleOptions" yaml:"mapStyleOptions"`
	// The sort configuration of a `FilledMapVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The window options of the filled map visual.
	WindowOptions interface{} `field:"optional" json:"windowOptions" yaml:"windowOptions"`
}

