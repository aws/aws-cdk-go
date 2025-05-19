package awscdklocationalpha


// The map style selected from an available data provider.
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
// See: https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html
//
// Experimental.
type Style string

const (
	// The Esri Navigation map style, which provides a detailed basemap for the world symbolized with a custom navigation map style that's designed for use during the day in mobile devices.
	//
	// It also includes a richer set of places, such as shops, services, restaurants, attractions,
	// and other points of interest. Enable the POI layer by setting it in CustomLayers to leverage
	// the additional places data.
	// Experimental.
	Style_VECTOR_ESRI_NAVIGATION Style = "VECTOR_ESRI_NAVIGATION"
	// The Esri Imagery map style.
	//
	// A raster basemap that provides one meter or better
	// satellite and aerial imagery in many parts of the world and lower resolution
	// satellite imagery worldwide.
	// Experimental.
	Style_RASTER_ESRI_IMAGERY Style = "RASTER_ESRI_IMAGERY"
	// The Esri Light Gray Canvas map style, which provides a detailed vector basemap with a light gray, neutral background style with minimal colors, labels, and features that's designed to draw attention to your thematic content.
	// Experimental.
	Style_VECTOR_ESRI_LIGHT_GRAY_CANVAS Style = "VECTOR_ESRI_LIGHT_GRAY_CANVAS"
	// The Esri Light map style, which provides a detailed vector basemap with a classic Esri map style.
	// Experimental.
	Style_VECTOR_ESRI_TOPOGRAPHIC Style = "VECTOR_ESRI_TOPOGRAPHIC"
	// The Esri Street Map style, which provides a detailed vector basemap for the world symbolized with a classic Esri street map style.
	//
	// The vector tile layer is similar
	// in content and style to the World Street Map raster map.
	// Experimental.
	Style_VECTOR_ESRI_STREETS Style = "VECTOR_ESRI_STREETS"
	// The Esri Dark Gray Canvas map style.
	//
	// A vector basemap with a dark gray,
	// neutral background with minimal colors, labels, and features that's designed
	// to draw attention to your thematic content.
	// Experimental.
	Style_VECTOR_ESRI_DARK_GRAY_CANVAS Style = "VECTOR_ESRI_DARK_GRAY_CANVAS"
	// A default HERE map style containing a neutral, global map and its features including roads, buildings, landmarks, and water features.
	//
	// It also now includes
	// a fully designed map of Japan.
	// Experimental.
	Style_VECTOR_HERE_EXPLORE Style = "VECTOR_HERE_EXPLORE"
	// A global map containing high resolution satellite imagery.
	// Experimental.
	Style_RASTER_HERE_EXPLORE_SATELLITE Style = "RASTER_HERE_EXPLORE_SATELLITE"
	// A global map displaying the road network, street names, and city labels over satellite imagery.
	//
	// This style will automatically retrieve both raster
	// and vector tiles, and your charges will be based on total tiles retrieved.
	// Experimental.
	Style_HYBRID_HERE_EXPLORE_SATELLITE Style = "HYBRID_HERE_EXPLORE_SATELLITE"
	// The HERE Contrast (Berlin) map style is a high contrast detailed base map of the world that blends 3D and 2D rendering.
	// Experimental.
	Style_VECTOR_HERE_CONTRAST Style = "VECTOR_HERE_CONTRAST"
	// A global map containing truck restrictions and attributes (e.g. width / height / HAZMAT) symbolized with highlighted segments and icons on top of HERE Explore to support use cases within transport and logistics.
	// Experimental.
	Style_VECTOR_HERE_EXPLORE_TRUCK Style = "VECTOR_HERE_EXPLORE_TRUCK"
	// The Grab Standard Light map style provides a basemap with detailed land use coloring, area names, roads, landmarks, and points of interest covering Southeast Asia.
	// Experimental.
	Style_VECTOR_GRAB_STANDARD_LIGHT Style = "VECTOR_GRAB_STANDARD_LIGHT"
	// The Grab Standard Dark map style provides a dark variation of the standard basemap covering Southeast Asia.
	// Experimental.
	Style_VECTOR_GRAB_STANDARD_DARK Style = "VECTOR_GRAB_STANDARD_DARK"
	// The Open Data Standard Light map style provides a detailed basemap for the world suitable for website and mobile application use.
	//
	// The map includes highways major roads,
	// minor roads, railways, water features, cities, parks, landmarks, building footprints,
	// and administrative boundaries.
	// Experimental.
	Style_VECTOR_OPEN_DATA_STANDARD_LIGHT Style = "VECTOR_OPEN_DATA_STANDARD_LIGHT"
	// Open Data Standard Dark is a dark-themed map style that provides a detailed basemap for the world suitable for website and mobile application use.
	//
	// The map includes highways
	// major roads, minor roads, railways, water features, cities, parks, landmarks,
	// building footprints, and administrative boundaries.
	// Experimental.
	Style_VECTOR_OPEN_DATA_STANDARD_DARK Style = "VECTOR_OPEN_DATA_STANDARD_DARK"
	// The Open Data Visualization Light map style is a light-themed style with muted colors and fewer features that aids in understanding overlaid data.
	// Experimental.
	Style_VECTOR_OPEN_DATA_VISUALIZATION_LIGHT Style = "VECTOR_OPEN_DATA_VISUALIZATION_LIGHT"
	// The Open Data Visualization Dark map style is a dark-themed style with muted colors and fewer features that aids in understanding overlaid data.
	// Experimental.
	Style_VECTOR_OPEN_DATA_VISUALIZATION_DARK Style = "VECTOR_OPEN_DATA_VISUALIZATION_DARK"
)

