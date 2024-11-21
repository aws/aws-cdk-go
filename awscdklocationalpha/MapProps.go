package awscdklocationalpha


// Properties for the Amazon Location Service Map.
//
// Example:
//   location.NewMap(this, jsii.String("Map"), &MapProps{
//   	MapName: jsii.String("my-map"),
//   	Style: location.Style_VECTOR_ESRI_NAVIGATION,
//   	CustomLayers: []pOI{
//   		location.CustomLayer_*pOI,
//   	},
//   })
//
// Experimental.
type MapProps struct {
	// Specifies the map style selected from an available data provider.
	// Experimental.
	Style Style `field:"required" json:"style" yaml:"style"`
	// Specifies the custom layers for the style.
	// Default: - no custom layes.
	//
	// Experimental.
	CustomLayers *[]CustomLayer `field:"optional" json:"customLayers" yaml:"customLayers"`
	// A description for the map.
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the map.
	//
	// Must be between 1 and 100 characters and contain only alphanumeric characters,
	// hyphens, periods and underscores.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	MapName *string `field:"optional" json:"mapName" yaml:"mapName"`
	// Specifies the map political view selected from an available data provider.
	//
	// The political view must be used in compliance with applicable laws, including those laws about mapping of the country or region where the maps,
	// images, and other data and third-party content which you access through Amazon Location Service is made available.
	// See: https://docs.aws.amazon.com/location/latest/developerguide/map-concepts.html#political-views
	//
	// Default: - no political view.
	//
	// Experimental.
	PoliticalView PoliticalView `field:"optional" json:"politicalView" yaml:"politicalView"`
}

