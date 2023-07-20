package awsquicksight


// The top movers and bottom movers computation setup.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html
//
type CfnAnalysis_TopBottomMoversComputationProperty struct {
	// The category field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-category
	//
	Category interface{} `field:"required" json:"category" yaml:"category"`
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The time field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-time
	//
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// The computation type. Choose from the following options:.
	//
	// - TOP: Top movers computation.
	// - BOTTOM: Bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The mover size setup of the top and bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-moversize
	//
	MoverSize *float64 `field:"optional" json:"moverSize" yaml:"moverSize"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The sort order setup of the top and bottom movers computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-sortorder
	//
	SortOrder *string `field:"optional" json:"sortOrder" yaml:"sortOrder"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topbottommoverscomputation.html#cfn-quicksight-analysis-topbottommoverscomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

