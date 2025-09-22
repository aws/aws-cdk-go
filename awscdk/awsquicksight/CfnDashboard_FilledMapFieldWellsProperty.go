package awsquicksight


// The field wells of a `FilledMapVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapfieldwells.html
//
type CfnDashboard_FilledMapFieldWellsProperty struct {
	// The aggregated field well of the filled map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filledmapfieldwells.html#cfn-quicksight-dashboard-filledmapfieldwells-filledmapaggregatedfieldwells
	//
	FilledMapAggregatedFieldWells interface{} `field:"optional" json:"filledMapAggregatedFieldWells" yaml:"filledMapAggregatedFieldWells"`
}

