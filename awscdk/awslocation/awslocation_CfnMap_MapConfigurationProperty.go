package awslocation


// Specifies the map tile style selected from an available provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mapConfigurationProperty := &mapConfigurationProperty{
//   	style: jsii.String("style"),
//   }
//
type CfnMap_MapConfigurationProperty struct {
	// Specifies the map style selected from an available data provider.
	//
	// Valid [Esri map styles](https://docs.aws.amazon.com/location/latest/developerguide/esri.html) :
	//
	// - `VectorEsriDarkGrayCanvas` – The Esri Dark Gray Canvas map style. A vector basemap with a dark gray, neutral background with minimal colors, labels, and features that's designed to draw attention to your thematic content.
	// - `RasterEsriImagery` – The Esri Imagery map style. A raster basemap that provides one meter or better satellite and aerial imagery in many parts of the world and lower resolution satellite imagery worldwide.
	// - `VectorEsriLightGrayCanvas` – The Esri Light Gray Canvas map style, which provides a detailed vector basemap with a light gray, neutral background style with minimal colors, labels, and features that's designed to draw attention to your thematic content.
	// - `VectorEsriTopographic` – The Esri Light map style, which provides a detailed vector basemap with a classic Esri map style.
	// - `VectorEsriStreets` – The Esri World Streets map style, which provides a detailed vector basemap for the world symbolized with a classic Esri street map style. The vector tile layer is similar in content and style to the World Street Map raster map.
	// - `VectorEsriNavigation` – The Esri World Navigation map style, which provides a detailed basemap for the world symbolized with a custom navigation map style that's designed for use during the day in mobile devices.
	//
	// Valid [HERE Technologies map styles](https://docs.aws.amazon.com/location/latest/developerguide/HERE.html) :
	//
	// - `VectorHereBerlin` – The HERE Berlin map style is a high contrast detailed base map of the world that blends 3D and 2D rendering.
	// - `VectorHereExplore` – A default HERE map style containing a neutral, global map and its features including roads, buildings, landmarks, and water features. It also now includes a fully designed map of Japan.
	// - `VectorHereExploreTruck` – A global map containing truck restrictions and attributes (e.g. width / height / HAZMAT) symbolized with highlighted segments and icons on top of HERE Explore to support use cases within transport and logistics.
	Style *string `field:"required" json:"style" yaml:"style"`
}

