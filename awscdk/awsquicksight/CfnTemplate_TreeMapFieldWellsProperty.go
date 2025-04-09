package awsquicksight


// The field wells of a tree map.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapfieldwells.html
//
type CfnTemplate_TreeMapFieldWellsProperty struct {
	// The aggregated field wells of a tree map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapfieldwells.html#cfn-quicksight-template-treemapfieldwells-treemapaggregatedfieldwells
	//
	TreeMapAggregatedFieldWells interface{} `field:"optional" json:"treeMapAggregatedFieldWells" yaml:"treeMapAggregatedFieldWells"`
}

