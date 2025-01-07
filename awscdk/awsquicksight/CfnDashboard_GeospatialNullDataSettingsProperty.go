package awsquicksight


// The properties for the visualization of null data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialNullDataSettingsProperty := &GeospatialNullDataSettingsProperty{
//   	SymbolStyle: &GeospatialNullSymbolStyleProperty{
//   		FillColor: jsii.String("fillColor"),
//   		StrokeColor: jsii.String("strokeColor"),
//   		StrokeWidth: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnulldatasettings.html
//
type CfnDashboard_GeospatialNullDataSettingsProperty struct {
	// The symbol style for null data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnulldatasettings.html#cfn-quicksight-dashboard-geospatialnulldatasettings-symbolstyle
	//
	SymbolStyle interface{} `field:"required" json:"symbolStyle" yaml:"symbolStyle"`
}

