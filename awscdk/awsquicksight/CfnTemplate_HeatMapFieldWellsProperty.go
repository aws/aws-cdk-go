package awsquicksight


// The field well configuration of a heat map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapfieldwells.html
//
type CfnTemplate_HeatMapFieldWellsProperty struct {
	// The aggregated field wells of a heat map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-heatmapfieldwells.html#cfn-quicksight-template-heatmapfieldwells-heatmapaggregatedfieldwells
	//
	HeatMapAggregatedFieldWells interface{} `field:"optional" json:"heatMapAggregatedFieldWells" yaml:"heatMapAggregatedFieldWells"`
}

