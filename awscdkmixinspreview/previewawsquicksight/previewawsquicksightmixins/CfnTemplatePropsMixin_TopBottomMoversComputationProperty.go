package previewawsquicksightmixins


// The top movers and bottom movers computation setup.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html
//
type CfnTemplatePropsMixin_TopBottomMoversComputationProperty struct {
	// The category field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-computationid
	//
	ComputationId *string `field:"optional" json:"computationId" yaml:"computationId"`
	// The mover size setup of the top and bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-moversize
	//
	// Default: - 0.
	//
	MoverSize *float64 `field:"optional" json:"moverSize" yaml:"moverSize"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order setup of the top and bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-sortorder
	//
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-time
	//
	Time interface{} `field:"optional" json:"time" yaml:"time"`
	// The computation type. Choose from the following options:.
	//
	// - TOP: Top movers computation.
	// - BOTTOM: Bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

