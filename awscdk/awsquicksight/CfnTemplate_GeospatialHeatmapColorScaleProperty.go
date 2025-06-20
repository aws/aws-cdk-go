package awsquicksight


// The color scale specification for the heatmap point style.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialHeatmapColorScaleProperty := &GeospatialHeatmapColorScaleProperty{
//   	Colors: []interface{}{
//   		&GeospatialHeatmapDataColorProperty{
//   			Color: jsii.String("color"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialheatmapcolorscale.html
//
type CfnTemplate_GeospatialHeatmapColorScaleProperty struct {
	// The list of colors to be used in heatmap point style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialheatmapcolorscale.html#cfn-quicksight-template-geospatialheatmapcolorscale-colors
	//
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
}

