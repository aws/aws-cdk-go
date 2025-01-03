package awsquicksight


// The field wells for a pivot table visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldwells.html
//
type CfnAnalysis_PivotTableFieldWellsProperty struct {
	// The aggregated field well for the pivot table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablefieldwells.html#cfn-quicksight-analysis-pivottablefieldwells-pivottableaggregatedfieldwells
	//
	PivotTableAggregatedFieldWells interface{} `field:"optional" json:"pivotTableAggregatedFieldWells" yaml:"pivotTableAggregatedFieldWells"`
}

