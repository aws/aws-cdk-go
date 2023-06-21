package awsquicksight


// The configuration of a pie chart.
//
// Example:
//
//
type CfnAnalysis_PieChartConfigurationProperty struct {
	// The label options of the group/color that is displayed in a pie chart.
	CategoryLabelOptions interface{} `field:"optional" json:"categoryLabelOptions" yaml:"categoryLabelOptions"`
	// The contribution analysis (anomaly configuration) setup of the visual.
	ContributionAnalysisDefaults interface{} `field:"optional" json:"contributionAnalysisDefaults" yaml:"contributionAnalysisDefaults"`
	// The options that determine if visual data labels are displayed.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The options that determine the shape of the chart.
	//
	// This option determines whether the chart is a pie chart or a donut chart.
	DonutOptions interface{} `field:"optional" json:"donutOptions" yaml:"donutOptions"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The small multiples setup for the visual.
	SmallMultiplesOptions interface{} `field:"optional" json:"smallMultiplesOptions" yaml:"smallMultiplesOptions"`
	// The sort configuration of a pie chart.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
	// The label options for the value that is displayed in a pie chart.
	ValueLabelOptions interface{} `field:"optional" json:"valueLabelOptions" yaml:"valueLabelOptions"`
	// The palette (chart color) display setup of the visual.
	VisualPalette interface{} `field:"optional" json:"visualPalette" yaml:"visualPalette"`
}

