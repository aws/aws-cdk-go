package awsquicksight


// The aggregated field wells of a word cloud.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html
//
type CfnTemplate_WordCloudAggregatedFieldWellsProperty struct {
	// The group by field well of a word cloud.
	//
	// Values are grouped by group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html#cfn-quicksight-template-wordcloudaggregatedfieldwells-groupby
	//
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// The size field well of a word cloud.
	//
	// Values are aggregated based on group by fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html#cfn-quicksight-template-wordcloudaggregatedfieldwells-size
	//
	Size interface{} `field:"optional" json:"size" yaml:"size"`
}

