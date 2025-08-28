package awsquicksight


// The top ranked and bottom ranked computation configuration.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html
//
type CfnTemplate_TopBottomRankedComputationProperty struct {
	// The ID for a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-computationid
	//
	ComputationId *string `field:"required" json:"computationId" yaml:"computationId"`
	// The computation type. Choose one of the following options:.
	//
	// - TOP: A top ranked computation.
	// - BOTTOM: A bottom ranked computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The category field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-category
	//
	Category interface{} `field:"optional" json:"category" yaml:"category"`
	// The name of a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The result size of a top and bottom ranked computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-resultsize
	//
	// Default: - 0.
	//
	ResultSize *float64 `field:"optional" json:"resultSize" yaml:"resultSize"`
	// The value field that is used in a computation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottomrankedcomputation.html#cfn-quicksight-template-topbottomrankedcomputation-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

