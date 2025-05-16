package awsquicksight


// The field wells of a word cloud visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-wordcloudfieldwells.html
//
type CfnAnalysis_WordCloudFieldWellsProperty struct {
	// The aggregated field wells of a word cloud.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-wordcloudfieldwells.html#cfn-quicksight-analysis-wordcloudfieldwells-wordcloudaggregatedfieldwells
	//
	WordCloudAggregatedFieldWells interface{} `field:"optional" json:"wordCloudAggregatedFieldWells" yaml:"wordCloudAggregatedFieldWells"`
}

