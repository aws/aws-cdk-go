package mixinsawsquicksight


// The field wells of a `GeospatialMapVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapfieldwells.html
//
type CfnAnalysisPropsMixin_GeospatialMapFieldWellsProperty struct {
	// The aggregated field well for a geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapfieldwells.html#cfn-quicksight-analysis-geospatialmapfieldwells-geospatialmapaggregatedfieldwells
	//
	GeospatialMapAggregatedFieldWells interface{} `field:"optional" json:"geospatialMapAggregatedFieldWells" yaml:"geospatialMapAggregatedFieldWells"`
}

