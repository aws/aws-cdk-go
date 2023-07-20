package awsquicksight


// The field wells for a table visual.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldwells.html
//
type CfnTemplate_TableFieldWellsProperty struct {
	// The aggregated field well for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldwells.html#cfn-quicksight-template-tablefieldwells-tableaggregatedfieldwells
	//
	TableAggregatedFieldWells interface{} `field:"optional" json:"tableAggregatedFieldWells" yaml:"tableAggregatedFieldWells"`
	// The unaggregated field well for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablefieldwells.html#cfn-quicksight-template-tablefieldwells-tableunaggregatedfieldwells
	//
	TableUnaggregatedFieldWells interface{} `field:"optional" json:"tableUnaggregatedFieldWells" yaml:"tableUnaggregatedFieldWells"`
}

