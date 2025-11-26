package previewawsquicksightmixins


// Aggregated field wells of a tree map.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapaggregatedfieldwells.html
//
type CfnTemplatePropsMixin_TreeMapAggregatedFieldWellsProperty struct {
	// The color field well of a tree map.
	//
	// Values are grouped by aggregations based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapaggregatedfieldwells.html#cfn-quicksight-template-treemapaggregatedfieldwells-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The group by field well of a tree map.
	//
	// Values are grouped based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapaggregatedfieldwells.html#cfn-quicksight-template-treemapaggregatedfieldwells-groups
	//
	Groups interface{} `field:"optional" json:"groups" yaml:"groups"`
	// The size field well of a tree map.
	//
	// Values are aggregated based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-treemapaggregatedfieldwells.html#cfn-quicksight-template-treemapaggregatedfieldwells-sizes
	//
	Sizes interface{} `field:"optional" json:"sizes" yaml:"sizes"`
}

