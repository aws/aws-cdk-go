package awsquicksight


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyle.html
//
type CfnDashboard_GeospatialMapStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyle.html#cfn-quicksight-dashboard-geospatialmapstyle-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyle.html#cfn-quicksight-dashboard-geospatialmapstyle-basemapstyle
	//
	BaseMapStyle *string `field:"optional" json:"baseMapStyle" yaml:"baseMapStyle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyle.html#cfn-quicksight-dashboard-geospatialmapstyle-basemapvisibility
	//
	BaseMapVisibility *string `field:"optional" json:"baseMapVisibility" yaml:"baseMapVisibility"`
}

