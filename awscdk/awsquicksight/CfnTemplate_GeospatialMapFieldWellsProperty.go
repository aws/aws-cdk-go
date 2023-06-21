package awsquicksight


// The field wells of a `GeospatialMapVisual` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//
//
type CfnTemplate_GeospatialMapFieldWellsProperty struct {
	// The aggregated field well for a geospatial map.
	GeospatialMapAggregatedFieldWells interface{} `field:"optional" json:"geospatialMapAggregatedFieldWells" yaml:"geospatialMapAggregatedFieldWells"`
}

