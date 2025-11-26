package previewawsquicksightmixins


// The configuration of a sankey diagram.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramchartconfiguration.html
//
type CfnTemplatePropsMixin_SankeyDiagramChartConfigurationProperty struct {
	// The data label configuration of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramchartconfiguration.html#cfn-quicksight-template-sankeydiagramchartconfiguration-datalabels
	//
	DataLabels interface{} `field:"optional" json:"dataLabels" yaml:"dataLabels"`
	// The field well configuration of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramchartconfiguration.html#cfn-quicksight-template-sankeydiagramchartconfiguration-fieldwells
	//
	FieldWells interface{} `field:"optional" json:"fieldWells" yaml:"fieldWells"`
	// The general visual interactions setup for a visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramchartconfiguration.html#cfn-quicksight-template-sankeydiagramchartconfiguration-interactions
	//
	Interactions interface{} `field:"optional" json:"interactions" yaml:"interactions"`
	// The sort configuration of a sankey diagram.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sankeydiagramchartconfiguration.html#cfn-quicksight-template-sankeydiagramchartconfiguration-sortconfiguration
	//
	SortConfiguration interface{} `field:"optional" json:"sortConfiguration" yaml:"sortConfiguration"`
}

