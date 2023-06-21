package awsquicksight


// The configuration of a `BarChartVisual` .
//
// Example:
//
//
type CfnAnalysis_BarChartConfigurationProperty struct {
	// Determines the arrangement of the bars.
	//
	// The orientation and arrangement of bars determine the type of bar that is used in the visual.
	BarsArrangement *string `field:"optional" json:"barsArrangement" yaml:"barsArrangement"`
	// The label display options (grid line, range, scale, axis step) for bar chart category.
	CategoryAxis interface{} `field:"optional" json:"categoryAxis" yaml:"categoryAxis"`
	// The label options (label text, label visibility and sort icon visibility) for a bar chart.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The label options (label text, label visibility and sort icon visibility) for a color that is used in a bar chart.
	ColorLabelOptions interface{} `field:"optional" json:"colorLabelOptions" yaml:"colorLabelOptions"`
	// The contribution analysis (anomaly configuration) setup of the visual.
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The options that determine if visual data labels are displayed.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The orientation of the bars in a bar chart visual. There are two valid values in this structure:.
	//
	// - `HORIZONTAL` : Used for charts that have horizontal bars. Visuals that use this value are horizontal bar charts, horizontal stacked bar charts, and horizontal stacked 100% bar charts.
	// - `VERTICAL` : Used for charts that have vertical bars. Visuals that use this value are vertical bar charts, vertical stacked bar charts, and vertical stacked 100% bar charts.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// The reference line setup of the visual.
	ReferenceLines interface{} `field:"optional" json:"referenceLines" yaml:"referenceLines"`
	// The small multiples setup for the visual.
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a `BarChartVisual` .
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The label display options (grid line, range, scale, axis step) for a bar chart value.
	ValueAxis interface{} `field:"optional" json:"valueAxis" yaml:"valueAxis"`
	// The label options (label text, label visibility and sort icon visibility) for a bar chart value.
	ValueLabelOptions interface{} `field:"optional" json:"valueLabelOptions" yaml:"valueLabelOptions"`
	// The palette (chart color) display setup of the visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

