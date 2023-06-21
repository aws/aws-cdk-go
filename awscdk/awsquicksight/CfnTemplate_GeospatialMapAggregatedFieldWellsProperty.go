package awsquicksight


// The aggregated field wells for a geospatial map.
//
// Example:
//
//
type CfnTemplate_GeospatialMapAggregatedFieldWellsProperty struct {
	// The color field wells of a geospatial map.
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// The geospatial field wells of a geospatial map.
	//
	// Values are grouped by geospatial fields.
	Geospatial interface{} `field:"optional" json:"geospatial" yaml:"geospatial"`
	// The size field wells of a geospatial map.
	//
	// Values are aggregated based on geospatial fields.
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

