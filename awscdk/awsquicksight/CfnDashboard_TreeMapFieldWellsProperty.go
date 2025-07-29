package awsquicksight


// The field wells of a tree map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapfieldwells.html
//
type CfnDashboard_TreeMapFieldWellsProperty struct {
	// The aggregated field wells of a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-treemapfieldwells.html#cfn-quicksight-dashboard-treemapfieldwells-treemapaggregatedfieldwells
	//
	TreeMapAggregatedFieldWells interface{} `field:"optional" json:"treeMapAggregatedFieldWells" yaml:"treeMapAggregatedFieldWells"`
}

