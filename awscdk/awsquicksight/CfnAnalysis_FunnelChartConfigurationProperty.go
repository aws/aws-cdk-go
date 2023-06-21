package awsquicksight


// The configuration of a `FunnelChartVisual` .
//
// Example:
//
//
type CfnAnalysis_FunnelChartConfigurationProperty struct {
	// The label options of the categories that are displayed in a `FunnelChartVisual` .
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The options that determine the presentation of the data labels.
	DataLabelOptions interface{} `field:"optional" json:"dataLabelOptions" yaml:"dataLabelOptions"`
	// The field well configuration of a `FunnelChartVisual` .
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The sort configuration of a `FunnelChartVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip configuration of a `FunnelChartVisual` .
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The label options for the values that are displayed in a `FunnelChartVisual` .
	ValueLabelOptions interface{} `field:"optional" json:"valueLabelOptions" yaml:"valueLabelOptions"`
	// The visual palette configuration of a `FunnelChartVisual` .
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

