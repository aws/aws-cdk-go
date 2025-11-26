package previewawsquicksightmixins


// The field well configuration of a heat map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapfieldwells.html
//
type CfnDashboardPropsMixin_HeatMapFieldWellsProperty struct {
	// The aggregated field wells of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-heatmapfieldwells.html#cfn-quicksight-dashboard-heatmapfieldwells-heatmapaggregatedfieldwells
	//
	HeatMapAggregatedFieldWells interface{} `field:"optional" json:"heatMapAggregatedFieldWells" yaml:"heatMapAggregatedFieldWells"`
}

