package awsquicksight


// The configuration of a heat map.
//
// Example:
//
//
type CfnTemplate_HeatMapConfigurationProperty struct {
	// The color options (gradient color, point of divergence) in a heat map.
	ColorScale interface{} `field:"optional" json:"colorScale" yaml:"colorScale"`
	// The label options of the column that is displayed in a heat map.
	ColumnLabelOptions interface{} `field:"optional" json:"columnLabelOptions" yaml:"columnLabelOptions"`
	// The options that determine if visual data labels are displayed.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field wells of the visual.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The legend display setup of the visual.
	Legend interface{} `field:"optional" json:"legend" yaml:"legend"`
	// The label options of the row that is displayed in a `heat map` .
	RowLabelOptions interface{} `field:"optional" json:"rowLabelOptions" yaml:"rowLabelOptions"`
	// The sort configuration of a heat map.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
	// The tooltip display setup of the visual.
	Tooltip interface{} `field:"optional" json:"tooltip" yaml:"tooltip"`
}

