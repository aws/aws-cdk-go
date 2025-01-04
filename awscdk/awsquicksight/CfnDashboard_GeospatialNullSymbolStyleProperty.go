package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialNullSymbolStyleProperty := &GeospatialNullSymbolStyleProperty{
//   	FillColor: jsii.String("fillColor"),
//   	StrokeColor: jsii.String("strokeColor"),
//   	StrokeWidth: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html
//
type CfnDashboard_GeospatialNullSymbolStyleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-fillcolor
	//
	FillColor *string `field:"optional" json:"fillColor" yaml:"fillColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-strokecolor
	//
	StrokeColor *string `field:"optional" json:"strokeColor" yaml:"strokeColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialnullsymbolstyle.html#cfn-quicksight-dashboard-geospatialnullsymbolstyle-strokewidth
	//
	StrokeWidth *float64 `field:"optional" json:"strokeWidth" yaml:"strokeWidth"`
}

