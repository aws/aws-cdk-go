package awsquicksight


// The map style properties for a map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialMapStyleProperty := &GeospatialMapStyleProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BaseMapStyle: jsii.String("baseMapStyle"),
//   	BaseMapVisibility: jsii.String("baseMapVisibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyle.html
//
type CfnAnalysis_GeospatialMapStyleProperty struct {
	// The background color and opacity values for a map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyle.html#cfn-quicksight-analysis-geospatialmapstyle-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The selected base map style.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyle.html#cfn-quicksight-analysis-geospatialmapstyle-basemapstyle
	//
	BaseMapStyle *string `field:"optional" json:"baseMapStyle" yaml:"baseMapStyle"`
	// The state of visibility for the base map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyle.html#cfn-quicksight-analysis-geospatialmapstyle-basemapvisibility
	//
	BaseMapVisibility *string `field:"optional" json:"baseMapVisibility" yaml:"baseMapVisibility"`
}

