package awsquicksight


// The configuration of a `BoxPlotVisual` .
//
// Example:
//
//
type CfnTemplate_BoxPlotChartConfigurationProperty struct {
	// The box plot chart options for a box plot visual.
	BoxPlotOptions interface{} `field:"optional" json:"boxPlotOptions" yaml:"boxPlotOptions"`
	// The label display options (grid line, range, scale, axis step) of a box plot category.
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The label options (label text, label visibility and sort Icon visibility) of a box plot category.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// `CfnTemplate.BoxPlotChartConfigurationProperty.Legend`.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The label display options (grid line, range, scale, axis step) of a box plot category.
	PrimaryYAxisDisplayOptions interface{} `field:"optional" json:"primaryYAxisDisplayOptions" yaml:"primaryYAxisDisplayOptions"`
	// The label options (label text, label visibility and sort icon visibility) of a box plot value.
	PrimaryYAxisLabelOptions interface{} `field:"optional" json:"primaryYAxisLabelOptions" yaml:"primaryYAxisLabelOptions"`
	// The reference line setup of the visual.
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The sort configuration of a `BoxPlotVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The palette (chart color) display setup of the visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

