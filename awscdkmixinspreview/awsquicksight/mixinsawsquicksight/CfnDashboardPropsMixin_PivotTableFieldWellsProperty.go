package mixinsawsquicksight


// The field wells for a pivot table visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldwells.html
//
type CfnDashboardPropsMixin_PivotTableFieldWellsProperty struct {
	// The aggregated field well for the pivot table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablefieldwells.html#cfn-quicksight-dashboard-pivottablefieldwells-pivottableaggregatedfieldwells
	//
	PivotTableAggregatedFieldWells interface{} `field:"optional" json:"pivotTableAggregatedFieldWells" yaml:"pivotTableAggregatedFieldWells"`
}

