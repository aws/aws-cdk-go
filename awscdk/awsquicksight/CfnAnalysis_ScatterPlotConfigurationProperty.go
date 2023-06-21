package awsquicksight


// The configuration of a scatter plot.
//
// Example:
//
//
type CfnAnalysis_ScatterPlotConfigurationProperty struct {
	// The options that determine if visual data labels are displayed.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The legend display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The palette (chart color) display setup of the visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
	// The label display options (grid line, range, scale, and axis step) of the scatter plot's x-axis.
	XAxisDisplayOptions interface{} `field:"optional" json:"xAxisDisplayOptions" yaml:"xAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of the scatter plot's x-axis.
	XAxisLabelOptions interface{} `field:"optional" json:"xAxisLabelOptions" yaml:"xAxisLabelOptions"`
	// The label display options (grid line, range, scale, and axis step) of the scatter plot's y-axis.
	YAxisDisplayOptions interface{} `field:"optional" json:"yAxisDisplayOptions" yaml:"yAxisDisplayOptions"`
	// The label options (label text, label visibility, and sort icon visibility) of the scatter plot's y-axis.
	YAxisLabelOptions interface{} `field:"optional" json:"yAxisLabelOptions" yaml:"yAxisLabelOptions"`
}

