package awsquicksight


// The color to be used in the heatmap point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialHeatmapDataColorProperty := &GeospatialHeatmapDataColorProperty{
//   	Color: jsii.String("color"),
//   }
//
type CfnTemplate_GeospatialHeatmapDataColorProperty struct {
	// The hex color to be used in the heatmap point style.
	Color *string `field:"required" json:"color" yaml:"color"`
}

