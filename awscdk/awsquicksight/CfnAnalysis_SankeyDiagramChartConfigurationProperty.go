package awsquicksight


// The configuration of a sankey diagram.
//
// Example:
//
//
type CfnAnalysis_SankeyDiagramChartConfigurationProperty struct {
	// The data label configuration of a sankey diagram.
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a sankey diagram.
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The sort configuration of a sankey diagram.
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

