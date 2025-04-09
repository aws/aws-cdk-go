package awscdklocationalpha


// An additional layer you can enable for a map style.
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
type CustomLayer string

const (
	// The POI custom layer adds a richer set of places, such as shops, services, restaurants, attractions, and other points of interest to your map.
	//
	// Currently only the VectorEsriNavigation map style supports the POI custom layer.
	// Experimental.
	CustomLayer_POI CustomLayer = "POI"
)

