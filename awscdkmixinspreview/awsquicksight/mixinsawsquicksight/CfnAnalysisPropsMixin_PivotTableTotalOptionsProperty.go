package mixinsawsquicksight


// The total options for a pivot table visual.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottabletotaloptions.html
//
type CfnAnalysisPropsMixin_PivotTableTotalOptionsProperty struct {
	// The column subtotal options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottabletotaloptions.html#cfn-quicksight-analysis-pivottabletotaloptions-columnsubtotaloptions
	//
	ColumnSubtotalOptions interface{} `field:"optional" json:"columnSubtotalOptions" yaml:"columnSubtotalOptions"`
	// The column total options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottabletotaloptions.html#cfn-quicksight-analysis-pivottabletotaloptions-columntotaloptions
	//
	ColumnTotalOptions interface{} `field:"optional" json:"columnTotalOptions" yaml:"columnTotalOptions"`
	// The row subtotal options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottabletotaloptions.html#cfn-quicksight-analysis-pivottabletotaloptions-rowsubtotaloptions
	//
	RowSubtotalOptions interface{} `field:"optional" json:"rowSubtotalOptions" yaml:"rowSubtotalOptions"`
	// The row total options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottabletotaloptions.html#cfn-quicksight-analysis-pivottabletotaloptions-rowtotaloptions
	//
	RowTotalOptions interface{} `field:"optional" json:"rowTotalOptions" yaml:"rowTotalOptions"`
}

